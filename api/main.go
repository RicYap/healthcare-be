package main

import (
	"healthcare-be/config"
	"healthcare-be/models"
	"healthcare-be/routes"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
	once   sync.Once
)

func initGin() {
	// Initialize only once
	once.Do(func() {
		// Determine environment
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

		// Initialize Gin
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
		}
		router = gin.Default()

		// CORS configuration
		frontendURL := os.Getenv("FRONTEND_URL")
		if frontendURL == "" {
			frontendURL = "*" // Allow all in dev, restrict in production
		}

		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{frontendURL},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		// Initialize database
		config.ConnectDB()
		config.DB.AutoMigrate(&models.User{}, &models.LabResult{})

		// Register routes
		routes.RegisterRoutes(router)
	})
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	initGin()
	router.ServeHTTP(w, r)
}

// Local development main function
func main() {
	initGin()
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}