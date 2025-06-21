package main

import (
	"healthcare-be/config"
	"healthcare-be/models"
	"healthcare-be/routes"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Determine the environment: development or production
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Load .env file only in development
	if env != "production" {
		dotenvFile := ".env.development"
		if err := godotenv.Load(dotenvFile); err != nil {
			log.Printf("Warning: failed to load %s file\n", dotenvFile)
		}
	}

	// Connect to the database and run migrations
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.LabResult{})

	// Get FRONTEND_URL for CORS
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		log.Fatal("FRONTEND_URL environment variable not set")
	}

	// Initialize Gin router
	r := gin.Default()

	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// CORS middleware setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register routes
	routes.RegisterRoutes(r)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server in %s mode on port %s...\n", env, port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
