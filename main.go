package main

import (
	"fmt"
	"os"

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
	routes.RegisterAllRoutes(r, db)

	r.Run(":" + port)

	fmt.Println("Started Server")
}
