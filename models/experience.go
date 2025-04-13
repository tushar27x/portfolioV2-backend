package models

import "gorm.io/gorm"

type Experience struct {
	gorm.Model
	CompanyName string `gorm:"not null"`
	Designation string `gorm:"not null"`
	Description string `gorm:"not null"`
	UserId      uint   `gorm:"not null"`
}
