// Entry point for the service
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kashyapkrlucky/ff-go-src/db"
	"github.com/kashyapkrlucky/ff-go-src/internal/handlers"
	"github.com/kashyapkrlucky/ff-go-src/internal/repositories"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db.InitDB()
	defer db.DB.Close()

	// Initialize repository and handler
	repo := repositories.NewFeatureFlagRepo(db.DB, db.RedisClient)
	handler := handlers.NewFeatureFlagHandler(repo)

	// Setup Gin router
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/flags", handler.GetAllFlags)
		api.POST("/flags", handler.CreateFlag)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port", port)
	router.Run(":" + port)
}
