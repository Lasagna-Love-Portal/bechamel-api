package main

import (
	"fmt"
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: return something more useful for error messages
func getCurrentUserProfile(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	userProfile, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

// TODO: this functionality should only be permitted for users who have a role with permissions
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

// TODO: should an API key or similar be required to POST a request to create a new user?
// Possibly API key or a JWT token with a role with suitable permissions?
// For demonstration and debugging purposes, any authenticated user can make new user profiles
// at the present time. This is most certainly not what we want to deploy.

func postUserProfile(c *gin.Context) {
	var newUserProfile model.LasagnaLoveUser

	authHeader := c.GetHeader("Authorization")
	_, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	if err = c.BindJSON(&newUserProfile); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Missing or unparsable JSON body"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing supplied message body"}})
		}
		return
	}

	newUserProfile, err = internal.AddNewUser(newUserProfile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []string{fmt.Sprintf("error adding profile: %v", err)}})
		return
	}

	c.JSON(http.StatusOK, newUserProfile)
}
