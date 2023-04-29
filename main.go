package main

import (
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func postUserAuthorization(c *gin.Context) {
	var userAuthorizationBody model.LasagnaLoveUserAuthRequest
	var errors []string

	if err := c.BindJSON(&userAuthorizationBody); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Missing or unparsable JSON body"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing supplied message body"}})
		}
		return
	}
	if userAuthorizationBody.UserName == "" {
		errors = append(errors, "Required parameter userName not supplied or empty")
	}
	if userAuthorizationBody.Password == "" {
		errors = append(errors, "Required parameter password not supplied or empty")
	}
	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	lasagnaLoveUser, err := internal.AuthorizeUser(userAuthorizationBody.UserName, userAuthorizationBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"errors": []string{"Supplied user could not be authorized with supplied password"}})
		return
	}

	jwtToken, err := internal.GenerateJWT(lasagnaLoveUser.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"errors": []string{"Error generating JWT token for supplied userName and password"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

/* TODO: return something more useful for error messages */
func getCurrentUserProfile(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	userProfile, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

/* TODO: this functionality should only be permitted for users who have a role with permissions */
func getUserProfileByID(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	_, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	userID, err := strconv.Atoi(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Could not parse id for user profile to retrieve"}})
		return
	}

	userProfile, err := internal.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"errors": []string{"Could not retrieve user profile for specified user id"}})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

/*
TODO: should an API key or similar be required to POST a request to create a new user?

Possibly API key or a JWT token with a role with suitable permissions?

For demonstration and debugging purposes, any authenticated user can make new user profiles
at the present time. This is most certainly not what we want to deploy.
*/
func postUserProfile(c *gin.Context) {
	var newUserProfile model.LasagnaLoveUser

	authHeader := c.GetHeader("Authorization")
	_, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	if err := c.BindJSON(&newUserProfile); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Missing or unparsable JSON body"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing supplied message body"}})
		}
		return
	}

	newUserProfile, err = internal.AddNewUser(newUserProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": []string{"Could not add new user profile"}})
		return
	}

	c.JSON(http.StatusOK, newUserProfile)
}

func main() {
	router := gin.Default()
	router.POST("/login", postUserAuthorization)
	router.GET("/profile", getCurrentUserProfile)
	router.GET("/profile/:userID", getUserProfileByID)
	router.POST("/profile", postUserProfile)

	router.Run("localhost:8080")
}
