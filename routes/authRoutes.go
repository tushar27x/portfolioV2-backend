package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/controllers"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", func(c *gin.Context) {
			controllers.RegisterUser(c, db) // Calls the controller for user registration
		})
		auth.POST("/login", func(c *gin.Context) {
			controllers.LoginUser(c, db) // Calls the controller for user login
		})
	}
}
