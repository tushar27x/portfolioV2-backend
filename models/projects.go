package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Title       string         `gorm:"not null;unique" json:"title"`
	Stack       pq.StringArray `gorm:"type:text[]" json:"stack"`
	Description string         `gorm:"not null" json:"description"`
	GithubLink  string         `gorm:"not null" json:"githubLink"`
	LiveLink    string         `gorm:"not null" json:"liveLink"`
	UserId      uint           `gorm:"not null" json:"userId"`
}
