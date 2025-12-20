package migrations

import (
	"travel-app/api-go/logger"
	"travel-app/api-go/user"

	"gorm.io/gorm"
)

// RunMigrations applies all database schema migrations
func RunMigrations(db *gorm.DB) error {
	logger.Log.Info("Running database migrations...")

	// AutoMigrate will create tables, missing columns and indexes
	// It will NOT delete columns or change column types
	err := db.AutoMigrate(
		&user.User{},
		// Add other models here as you create them
	)

	if err != nil {
		return err
	}

	logger.Log.Info("Migrations completed successfully")
	return nil
}
