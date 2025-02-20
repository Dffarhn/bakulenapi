package main

import (
	"log"
	"os"

	v1 "github.com/Dffarhn/bakulenapi/api/v1"
	"github.com/Dffarhn/bakulenapi/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize Firebase
	config.InitFirebase()

	// Setup Gin router
	router := gin.Default()
// Create an instance of AuthHandler
	authHandler := v1.NewAuthHandler()
	userHandler := v1.NewUserHandler()

	// Register the routes
	v1Routes := router.Group("/v1")
	{
		v1.RegisterAuthRoutes(v1Routes, authHandler)
		v1.RegisterUserRoutes(v1Routes, userHandler)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	router.Run(":" + port)
}
