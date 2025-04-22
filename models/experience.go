package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	CompanyName string         `gorm:"not null" json:"companyName"`
	Designation string         `gorm:"not null" json:"designation"`
	Description string         `gorm:"not null" json:"description"`
	UserId      uint           `gorm:"not null" json:"userId"`
}
