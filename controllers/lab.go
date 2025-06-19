package controllers

import (
	"log"
	"net/http"

	"healthcare-be/config"
	"healthcare-be/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddLabResult(c *gin.Context) {
	var input models.LabInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := models.LabResult{
		UserID:           input.UserID,
		Date:             input.Date,
		Glucose:          input.Results.Glucose,
		CholesterolTotal: input.Results.Cholesterol.Total,
		LDL:              input.Results.Cholesterol.LDL,
		HDL:              input.Results.Cholesterol.HDL,
		Systolic:         input.Results.BloodPressure.Systolic,
		Diastolic:        input.Results.BloodPressure.Diastolic,
	}

	config.DB.Create(&result)
	c.JSON(http.StatusCreated, result)
}

func GetLabResults(c *gin.Context) {
	userID := c.MustGet("userId").(uuid.UUID)
	var results []models.LabResult
	config.DB.Where("user_id = ?", userID).Order("date desc").Find(&results)

	c.JSON(http.StatusOK, results)
}
