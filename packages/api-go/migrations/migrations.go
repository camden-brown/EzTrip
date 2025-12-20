package migrations

import (
	"log"

	"travel-app/api-go/user"

	"gorm.io/gorm"
)

// RunMigrations applies all database schema migrations
func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// AutoMigrate will create tables, missing columns and indexes
	// It will NOT delete columns or change column types
	err := db.AutoMigrate(
		&user.User{},
		// Add other models here as you create them
	)

	if err != nil {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
