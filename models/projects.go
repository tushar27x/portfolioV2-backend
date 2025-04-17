package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string         `gorm:"not null"`
	Stack       pq.StringArray `gorm:"type:text[]"`
	Description string         `gorm:"not null"`
	GithubLink  string         `gorm:"not null"`
	LiveLink    string         `gorm:"not null"`
	UserId      uint           `gorm:"not null"`
}
