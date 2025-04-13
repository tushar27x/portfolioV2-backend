package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/controllers"
	"github.com/tushar27x/portfolioV2-backend/middlewares"
	"gorm.io/gorm"
)

func RegisterSkillRoutes(r *gin.Engine, db *gorm.DB) {
	skills := r.Group("/skills")
	skills.Use(middlewares.AuthMiddleWare())
	{
		skills.GET("/", func(c *gin.Context) {
			controllers.GetSkillsForUser(c, db)
		})
		skills.POST("/", func(c *gin.Context) {
			controllers.AddSkill(c, db)
		})
		skills.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateSkill(c, db)
		})
		skills.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteSkill(c, db)
		})
	}
}
