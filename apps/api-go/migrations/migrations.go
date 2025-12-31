package migrations

import (
	"fmt"
	"os"

	"eztrip/api-go/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// RunMigrations applies all database schema migrations using golang-migrate
func RunMigrations(db *gorm.DB) error {
	logger.Log.Info("Running database migrations...")

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnvOrDefault("POSTGRES_USER", "postgres"),
		getEnvOrDefault("POSTGRES_PASSWORD", "postgres"),
		getEnvOrDefault("POSTGRES_HOST", "localhost"),
		getEnvOrDefault("POSTGRES_PORT", "5432"),
		getEnvOrDefault("POSTGRES_DB", "eztrip"),
	)

	m, err := migrate.New(
		"file://migrations/sql",
		connString,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	if err == migrate.ErrNoChange {
		logger.Log.Info("No new migrations to apply")
	} else {
		logger.Log.Info("Migrations completed successfully")
	}

	return nil
}

// RollbackMigration rolls back the last migration
func RollbackMigration(db *gorm.DB) error {
	logger.Log.Info("Rolling back last migration...")

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnvOrDefault("POSTGRES_USER", "postgres"),
		getEnvOrDefault("POSTGRES_PASSWORD", "postgres"),
		getEnvOrDefault("POSTGRES_HOST", "localhost"),
		getEnvOrDefault("POSTGRES_PORT", "5432"),
		getEnvOrDefault("POSTGRES_DB", "eztrip"),
	)

	m, err := migrate.New(
		"file://migrations/sql",
		connString,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("rollback failed: %w", err)
	}

	logger.Log.Info("Rollback completed successfully")
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
