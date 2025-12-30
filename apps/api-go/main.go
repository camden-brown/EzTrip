package main

import (
	"os"

	"eztrip/api-go/app"
	"eztrip/api-go/db"
	"eztrip/api-go/logger"
	"eztrip/api-go/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConfig := db.GetConfigFromEnv()
	database, err := db.NewGormDB(dbConfig)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Log.WithField("component", "database").Info("Successfully connected to PostgreSQL with GORM")

	if os.Getenv("GIN_MODE") != "release" {
		if err := migrations.AutoMigrate(database); err != nil {
			logger.Log.Fatalf("AutoMigrate failed: %v", err)
		}
	}

	enforcer, err := app.InitializeRBAC(database)
	if err != nil {
		logger.Log.Fatalf("Failed to initialize RBAC: %v", err)
	}

	router := gin.New()

	if err := app.SetupMiddleware(router, database, enforcer); err != nil {
		logger.Log.Fatalf("Failed to configure middleware: %v", err)
	}

	app.SetupRoutes(router, database)

	logger.Log.WithFields(map[string]interface{}{
		"component": "server",
		"port":      "8080",
	}).Info("Starting server")
	router.Run(":8080")
}
