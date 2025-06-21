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
	userID := c.MustGet("userId").(uuid.UUID)

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result := models.LabResult{
		UserID:           userID,
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
	config.DB.Where("user_id = ?", userID).Order("date asc").Find(&results)

	var resultsJSON []models.LabInput
	for _, result := range results {
		var resultJSON models.LabInput
		resultJSON.ID = result.ID
		resultJSON.Date = result.Date
		resultJSON.Results.Glucose = result.Glucose
		resultJSON.Results.Cholesterol.Total = result.CholesterolTotal
		resultJSON.Results.Cholesterol.HDL = result.HDL
		resultJSON.Results.Cholesterol.LDL = result.LDL
		resultJSON.Results.BloodPressure.Diastolic = result.Diastolic
		resultJSON.Results.BloodPressure.Systolic = result.Systolic
		resultJSON.UserID = result.UserID

		resultsJSON = append(resultsJSON, resultJSON)
	}

	c.JSON(http.StatusOK, resultsJSON)
}
