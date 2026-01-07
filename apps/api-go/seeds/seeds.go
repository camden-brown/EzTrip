package seeds

import (
	"eztrip/api-go/logger"
	"eztrip/api-go/rbac"

	"gorm.io/gorm"
)

// RunSeeds populates the database with sample data (development only).
// Add new seed functions here as you create more models.
func RunSeeds(db *gorm.DB) error {
	logger.Log.Info("Running database seeds...")

	enforcer, err := rbac.NewEnforcer(db)
	if err != nil {
		logger.Log.WithField("error", err.Error()).Error("Failed to create RBAC enforcer for seeding")
		return err
	}

	if err := SeedUsers(db, enforcer); err != nil {
		return err
	}

	if err := SeedTrips(db); err != nil {
		return err
	}

	logger.Log.Info("All seeds completed successfully")
	return nil
}
