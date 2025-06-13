package main

import (
	"healthcare-be/config"
	"healthcare-be/models"
	"healthcare-be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.LabResult{})

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
