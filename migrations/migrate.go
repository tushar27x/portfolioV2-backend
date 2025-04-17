package migrations

import (
	"log"

	"github.com/tushar27x/portfolioV2-backend/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.Experience{},
		&models.Project{},
	)

	if err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}

}
