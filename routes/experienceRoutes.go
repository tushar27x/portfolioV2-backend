// routes/experienceRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/controllers"
	"github.com/tushar27x/portfolioV2-backend/middlewares"
	"gorm.io/gorm"
)

func RegisterExperienceRoutes(r *gin.Engine, db *gorm.DB) {
	experienceGroup := r.Group("/experience")
	experienceGroup.Use(middlewares.AuthMiddleWare())

	experienceGroup.GET("/", func(ctx *gin.Context) {
		controllers.GetExperiencesForUser(ctx, db)
	})

	experienceGroup.POST("/", func(ctx *gin.Context) {
		controllers.AddExperience(ctx, db)
	})

	experienceGroup.PUT("/:id", func(ctx *gin.Context) {
		controllers.UpdateExperience(ctx, db)
	})

	experienceGroup.DELETE("/:id", func(ctx *gin.Context) {
		controllers.DeleteExperience(ctx, db)
	})
}
