package model

import (
	"time"

	"gorm.io/gorm"
)

type UserOTP struct {
	gorm.Model
	UserID     uint
	Code       string `json:"kode" gorm:"type:varchar(8);not null;"`
	Expire 		time.Time `json:"Expire" gorm:"type:timestamp;not null"`
	Used 	bool `json:"used" gorm:"default:false"`
}