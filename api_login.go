package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"
)

func postUserAuthorization(c *gin.Context) {
	var userAuthorizationBody model.LasagnaLoveAuthRequest
	var emailAddress string
	var errors []string

	// First check for a username/password combination
	if err := c.BindJSON(&userAuthorizationBody); err == nil {
		if userAuthorizationBody.RefreshToken != "" {
			emailAddress, err = internal.VerifyRefreshJWT(userAuthorizationBody.RefreshToken)
			if err != nil {
				errors = append(errors, "Refresh token invalid or expired")
			}
		} else {
			if userAuthorizationBody.Email == "" {
				errors = append(errors, "Required parameter email not supplied or empty")
			}
			if userAuthorizationBody.Password == "" {
				errors = append(errors, "Required parameter password not supplied or empty")
			}
			var lasagnaLoveUser model.LasagnaLoveUser
			lasagnaLoveUser, err = internal.AuthorizeUser(userAuthorizationBody.Email, userAuthorizationBody.Password)
			if err != nil {
				c.JSON(http.StatusUnauthorized,
					gin.H{"errors": []string{"Supplied user could not be authorized with supplied password"}})
				return
			}
			emailAddress = lasagnaLoveUser.Email
		}
	} else {
		if err.Error() == "EOF" {
			errors = append(errors, "Missing or unparsable JSON body")
		} else {
			errors = append(errors, "Error parsing supplied message body")
		}
	}
	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	accessJWTToken, accessErr := internal.GenerateAccessJWT(emailAddress)
	if accessErr != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"errors": []string{"Error generating JWT token for supplied userName and password"}})
		return
	}
	refreshJWTToken, refreshErr := internal.GenerateRefreshJWT(emailAddress)
	if refreshErr != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"errors": []string{"Error generating JWT refresh token"}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessJWTToken,
		"refresh_token": refreshJWTToken})
}
