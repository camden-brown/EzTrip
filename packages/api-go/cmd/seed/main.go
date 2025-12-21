package main

import (
	"travel-app/api-go/db"
	"travel-app/api-go/logger"
	"travel-app/api-go/seeds"
)

func main() {
	// Initialize database connection
	dbConfig := db.GetConfigFromEnv()
	database, err := db.NewGormDB(dbConfig)
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}

	logger.Log.Info("Successfully connected to PostgreSQL for seeding")

	// Run seeds
	if err := seeds.RunSeeds(database); err != nil {
		logger.Log.Fatalf("Failed to run seeds: %v", err)
	}

	logger.Log.Info("Seeds completed successfully")
}
