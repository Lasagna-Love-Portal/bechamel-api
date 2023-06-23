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
	router.POST("/login", postUserAuthorization)
	router.GET("/profile", getCurrentUserProfile)
	router.PATCH("/profile", patchCurrentUserProfile)
	router.POST("/profile", postUserProfile)
	router.GET("/profile/:userID", getUserProfileByID)
	router.PATCH("/profile/:userID", patchUserProfileByID)
	router.GET("/request/:requestID", getRequestByID)
	router.POST("/request", postRequest)

	log.Fatal(router.Run("0.0.0.0:8080"))
}
