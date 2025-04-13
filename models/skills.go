package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name     string `gorm:"not null; unique"`
	Score    int    `gorm:"check:score >= 1 AND score <= 10"`
	ImageURL string `gorm:"not null"`
	UserId   uint   `gorm:"not null"`
}
