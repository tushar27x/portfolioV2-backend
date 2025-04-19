package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tushar27x/portfolioV2-backend/config"
	"github.com/tushar27x/portfolioV2-backend/migrations"
	"github.com/tushar27x/portfolioV2-backend/routes"
)

func main() {
	config.LoadEnv()
	db := config.ConnectToDB()

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		migrations.Migrate(db)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-production-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAllRoutes(r, db)

	r.Run(":" + port)

	fmt.Println("Started Server")
}
