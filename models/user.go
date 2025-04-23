package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Name       string         `gorm:"not null" json:"name"`
	Email      string         `gorm:"not null; unique" json:"email"`
	Passwd     string         `gorm:"not null" json:"passwd"`
	Skills     []Skill        `gorm:"foreignKey:UserId" json:"skills"`
	Experience []Experience   `gorm:"foreignKey:UserId" json:"experience"`
	Project    []Project      `gorm:"foreignKey:UserId" json:"project"`
}
