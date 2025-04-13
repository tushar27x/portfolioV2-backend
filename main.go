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
	db := config.ConnectToDB()

	if os.Getenv("RUN_MIGRATIONS") == "true" {
		migrations.Migrate(db)
	}
	r := gin.Default()

	routes.RegisterUserRoutes(r, db)
	routes.RegisterSkillRoutes(r, db)
	routes.RegisterExperienceRoutes(r, db)
	routes.RegisterAuthRoutes(r, db)

	r.Run(":8080")
	fmt.Println("Started Server")
}
