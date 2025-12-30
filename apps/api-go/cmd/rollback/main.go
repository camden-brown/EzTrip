package main

import (
	"eztrip/api-go/db"
	"eztrip/api-go/logger"
	"eztrip/api-go/migrations"
)

func main() {
	logger.Log.Info("Rolling back last migration...")

	dbConfig := db.GetConfigFromEnv()
	database, err := db.NewGormDB(dbConfig)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := migrations.RollbackMigration(database); err != nil {
		logger.Log.Fatalf("Rollback failed: %v", err)
	}

	logger.Log.Info("Rollback completed successfully")
}
