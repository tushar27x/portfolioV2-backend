package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/controllers"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", func(c *gin.Context) {
		controllers.RegisterUser(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		controllers.LoginUser(c, db)
	})
}
