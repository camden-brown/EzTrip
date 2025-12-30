package main

import (
	"eztrip/api-go/db"
	"eztrip/api-go/logger"
	"eztrip/api-go/migrations"
)

func main() {
	logger.Log.Info("Starting database migration...")

	dbConfig := db.GetConfigFromEnv()
	database, err := db.NewGormDB(dbConfig)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := migrations.RunMigrations(database); err != nil {
		logger.Log.Fatalf("Migration failed: %v", err)
	}

	logger.Log.Info("Migration completed successfully")
}
