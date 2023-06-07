package main

import (
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: all requests for user, see api_profile for inspiration
// TODO: this functionality should only be permitted for users who have a role with permissions
func getRequestByID(c *gin.Context) {
	var err error
	var requestID int
	var request model.LasagnaLoveRequest

	authHeader := c.GetHeader("Authorization")
	_, err = internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	requestID, err = strconv.Atoi(c.Params.ByName("requestID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Could not parse id for request to retrieve"}})
		return
	}

	request, err = internal.GetRequestByID(requestID)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"errors": []string{"Could not retrieve request for specified request id"}})
		return
	}

	c.JSON(http.StatusOK, request)
}

// TODO: should all authorized users be permitted to make new requests?
func postRequest(c *gin.Context) {
	var newRequestProfile model.LasagnaLoveRequest

	authHeader := c.GetHeader("Authorization")
	_, err := internal.GetUserFromAuthHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": []string{"Invalid access token provided"}})
		return
	}

	if err := c.BindJSON(&newRequestProfile); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Missing or unparsable JSON body"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing supplied message body"}})
		}
		return
	}

	newRequestProfile, err = internal.AddNewRequest(newRequestProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": []string{"Could not add new request"}})
		return
	}

	c.JSON(http.StatusOK, newRequestProfile)
}
