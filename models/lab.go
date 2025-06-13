package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type LabResult struct {
	ID               uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID           uuid.UUID `gorm:"type:char(36);not null"`
	Date             time.Time `gorm:"type:datetime"`
	Glucose          float64   `gorm:"type:float"`
	CholesterolTotal float64   `gorm:"type:float"`
	LDL              float64   `gorm:"type:float"`
	HDL              float64   `gorm:"type:float"`
	Systolic         int       `gorm:"type:int"`
	Diastolic        int       `gorm:"type:int"`
}

func (l *LabResult) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New()
	return
}
