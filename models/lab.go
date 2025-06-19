package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type LabResult struct {
	ID               uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID           uuid.UUID `gorm:"type:char(36);not null"`
	Date             time.Time `gorm:"type:date"`
	Glucose          float64   `gorm:"type:float"`
	CholesterolTotal float64   `gorm:"type:float"`
	LDL              float64   `gorm:"type:float"`
	HDL              float64   `gorm:"type:float"`
	Systolic         float64   `gorm:"type:int"`
	Diastolic        float64   `gorm:"type:int"`
}

type LabInput struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	Date    time.Time `json:"date"`
	Results Lab       `json:"results"`
}

type Lab struct {
	Glucose       float64         `json:"glucose"`
	Cholesterol   CholesterolInfo `json:"cholesterol"`
	BloodPressure BloodPressure   `json:"bloodPressure"`
}

type CholesterolInfo struct {
	Total float64 `json:"total"`
	LDL   float64 `json:"ldl"`
	HDL   float64 `json:"hdl"`
}

type BloodPressure struct {
	Systolic  float64 `json:"systolic"`
	Diastolic float64 `json:"diastolic"`
}

func (l *LabResult) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
