package controllers

import (
	"net/http"
	"time"

	"healthcare-be/config"
	"healthcare-be/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LabInput struct {
	Date             time.Time `json:"date"`
	Glucose          float64   `json:"glucose"`
	CholesterolTotal float64   `json:"cholesterol_total"`
	LDL              float64   `json:"ldl"`
	HDL              float64   `json:"hdl"`
	Systolic         int       `json:"systolic"`
	Diastolic        int       `json:"diastolic"`
}

func AddLabResult(c *gin.Context) {
	var input LabInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userId").(uuid.UUID)
	result := models.LabResult{
		UserID:           userID,
		Date:             input.Date,
		Glucose:          input.Glucose,
		CholesterolTotal: input.CholesterolTotal,
		LDL:              input.LDL,
		HDL:              input.HDL,
		Systolic:         input.Systolic,
		Diastolic:        input.Diastolic,
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
