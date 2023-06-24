package main

import (
	"log"
	"project-ricotta/bechamel-api/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.RuntimeConfig = config.NewLocalhostDevConfig()

	router := gin.Default()
	var corsConfig = cors.DefaultConfig()
	// TODO: for the time being, we allow requests from all origins.
	// We can consider narrowing this if/when the Project Ricotta front-end has
	// specific known origin domains that we want to restrict requests to originating from.
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "authorization")
	router.Use(cors.New(corsConfig))

	dev := router.Group("/dev")
	{
		dev.POST("/login", postUserAuthorization)
		dev.GET("/profile", getCurrentUserProfile)
		dev.PATCH("/profile", patchCurrentUserProfile)
		dev.POST("/profile", postUserProfile)
		dev.GET("/profile/:userID", getUserProfileByID)
		dev.PATCH("/profile/:userID", patchUserProfileByID)
		dev.GET("/request/:requestID", getRequestByID)
		dev.POST("/request", postRequest)
	}
	v0 := router.Group("/v0")
	{
		v0.POST("/login", postUserAuthorization)
		v0.GET("/profile", getCurrentUserProfile)
		v0.PATCH("/profile", patchCurrentUserProfile)
		v0.POST("/profile", postUserProfile)
		v0.GET("/profile/:userID", getUserProfileByID)
		v0.PATCH("/profile/:userID", patchUserProfileByID)
		v0.GET("/request/:requestID", getRequestByID)
		v0.POST("/request", postRequest)
	}

	log.Fatal(router.Run("0.0.0.0:8080"))
}
