package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAllRoutes(r *gin.Engine, db *gorm.DB) {
	RegisterUserRoutes(r, db)
	RegisterSkillRoutes(r, db)
	RegisterExperienceRoutes(r, db)
}
