package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string       `gorm:"not null"`
	Email      string       `gorm:"not null; unique"`
	Passwd     string       `gorm:"not null"`
	Skills     []Skill      `gorm:"foreignKey:UserId"`
	Experience []Experience `gorm:"foreignKey:UserId"`
}
