package main

import (
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"

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
	if userAuthorizationBody.Username == "" {
		errors = append(errors, "Required parameter userName not supplied or empty")
	}
	if userAuthorizationBody.Password == "" {
		errors = append(errors, "Required parameter password not supplied or empty")
	}
	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	lasagnaLoveUser, err := internal.AuthorizeUser(userAuthorizationBody.Username,
		userAuthorizationBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			gin.H{"errors": []string{"Supplied user could not be authorized with supplied password"}})
		return
	}

	jwtToken, err := internal.GenerateJWT(lasagnaLoveUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"errors": []string{"Error generating JWT token for supplied userName and password"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
