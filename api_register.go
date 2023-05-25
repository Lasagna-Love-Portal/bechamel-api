package main

import (
	"net/http"
	"project-ricotta/bechamel-api/internal"
	"project-ricotta/bechamel-api/model"

	"github.com/gin-gonic/gin"
)

func postUserRegistration(c *gin.Context) {
	var newUserProfile model.LasagnaLoveUser
	if err := c.BindJSON(&newUserProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Error parsing request body"}})
		return
	}

	// Attempt to add new user
	user, err := internal.AddNewUser(newUserProfile)
	if err != nil {
		// Return an error message if registration fails
		c.JSON(http.StatusInternalServerError, gin.H{"errors": []string{err.Error()}})
		return
	}

	// If user registration is successful, return the user object
	c.JSON(http.StatusOK, user)
}
