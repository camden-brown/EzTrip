package seeds

import (
	"travel-app/api-go/logger"
	"travel-app/api-go/user"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// SeedUsers populates the users table with sample data
func SeedUsers(db *gorm.DB) error {
	var count int64
	db.Model(&user.User{}).Count(&count)
	if count > 0 {
		logger.Log.Info("Users already seeded, skipping...")
		return nil
	}

	logger.Log.Info("Seeding users...")

	users := []user.User{
		{FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com"},
		{FirstName: "Bob", LastName: "Williams", Email: "bob@example.com"},
	}

	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			logger.Log.WithFields(logrus.Fields{
				"email": u.Email,
				"error": err,
			}).Warn("Failed to seed user")
		}
	}

	logger.Log.WithField("count", len(users)).Info("Successfully seeded users")
	return nil
}
