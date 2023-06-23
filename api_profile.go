package main

import (
	"fmt"
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
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

func patchCurrentUserProfile(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	_, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}
}

type patchUpdateStruct map[string]interface{}

func (inStruct patchUpdateStruct) pascalCase() patchUpdateStruct {
	retStruct := make(patchUpdateStruct)
	for k, v := range inStruct {
		// TODO: iterate here
		switch vt := v.(type) {
		case map[string]interface{}:
			retStruct[strcase.ToCamel(k)] = (patchUpdateStruct)(vt).pascalCase()
		default:
			retStruct[strcase.ToCamel(k)] = vt
		}
	}
	return retStruct
}

func patchUserProfileByID(c *gin.Context) {
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

	_, err = internal.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"errors": []string{"Could not retrieve user profile for specified user id"}})
		return
	}

	// NOTE: This will get the JSON as a general struct so we can pass it along
	// Using the filled update struct type doesn't work because it defaults unfilled values,
	// so we can't tell the difference between ones that are being set to default/blank
	// values and those that weren't provided by the caller
	var generalStruct patchUpdateStruct
	if err = c.BindJSON(&generalStruct); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Missing or unparsable JSON body"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing supplied message body"}})
		}
		return
	}

	var pascalEncasedUpdates patchUpdateStruct = generalStruct.pascalCase()
	_, err = internal.UpdateUser(userID, pascalEncasedUpdates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{err.Error()}})
		return
	}
	// TODO: curently the API documentation specifies a successful PATCH returns
	// an HTTP 204 NO CONTENT. We do have the updated Profile returned if we want
	// to return it. Should we offer that option? Or let the caller
	// re-GET the Profile if they want it?
	c.Status(http.StatusNoContent)
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
