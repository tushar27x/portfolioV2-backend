package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAllRoutes(r *gin.Engine, db *gorm.DB) {
	RegisterSkillRoutes(r, db)
	RegisterExperienceRoutes(r, db)
	RegisterProjectRoutes(r, db)
	RegisterAuthRoutes(r, db)
}
