package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", postUserAuthorization)
	router.GET("/profile", getCurrentUserProfile)
	router.GET("/profile/:userID", getUserProfileByID)
	router.POST("/profile", postUserProfile)

	router.Run("localhost:8080")
}
