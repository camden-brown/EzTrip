package seeds

import (
	"log"

	"travel-app/api-go/user"

	"gorm.io/gorm"
)

// SeedUsers populates the users table with sample data
func SeedUsers(db *gorm.DB) error {
	var count int64
	db.Model(&user.User{}).Count(&count)
	if count > 0 {
		log.Println("Users already seeded, skipping...")
		return nil
	}

	log.Println("Seeding users...")

	users := []user.User{
		{FirstName: "John", LastName: "Doe", Email: "john@example.com"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice@example.com"},
		{FirstName: "Bob", LastName: "Williams", Email: "bob@example.com"},
	}

	for _, u := range users {
		if err := db.Create(&u).Error; err != nil {
			log.Printf("Warning: Failed to seed user %s: %v", u.Email, err)
		}
	}

	log.Printf("Successfully seeded %d users", len(users))
	return nil
}
