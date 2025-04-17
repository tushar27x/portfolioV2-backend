package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/controllers"
	"github.com/tushar27x/portfolioV2-backend/middlewares"
	"gorm.io/gorm"
)

func RegisterProjectRoutes(r *gin.Engine, db *gorm.DB) {
	projectGroup := r.Group("/project")
	projectGroup.Use(middlewares.AuthMiddleWare())

	projectGroup.GET("/", func(ctx *gin.Context) {
		controllers.GetProjectsForUser(ctx, db)
	})

	projectGroup.POST("/", func(ctx *gin.Context) {
		controllers.AddProject(ctx, db)
	})

	projectGroup.PUT("/:id", func(ctx *gin.Context) {
		controllers.UpdateProject(ctx, db)
	})

	projectGroup.DELETE("/:id", func(ctx *gin.Context) {
		controllers.DeleteProject(ctx, db)
	})
}
