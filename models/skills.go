package models

import (
	"time"

	"gorm.io/gorm"
)

type Skill struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Name      string         `gorm:"not null; unique" json:"name"`
	Score     int            `gorm:"check:score >= 1 AND score <= 10" json:"score"`
	ImageURL  string         `gorm:"not null" json:"imageUrl"`
	UserId    uint           `gorm:"not null" json:"userId"`
}
