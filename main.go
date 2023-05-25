package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
    router.GET("/profile/:userID", getUserProfileByID)
    router.POST("/profile", postUserProfile)
    router.POST("/register", postUserRegistration)
    
    router.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Test route is working"})
    })

    router.Run("localhost:8080")
}