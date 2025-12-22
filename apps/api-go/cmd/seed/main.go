package main

import (
	"eztrip/api-go/db"
	"eztrip/api-go/logger"
	"eztrip/api-go/seeds"
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
