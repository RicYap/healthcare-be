package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Email    string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string    `gorm:"type:text;not null" json:"-"`
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
