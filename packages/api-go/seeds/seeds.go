package seeds

import (
	"log"

	"gorm.io/gorm"
)

// RunSeeds populates the database with sample data (development only).
// Add new seed functions here as you create more models.
func RunSeeds(db *gorm.DB) error {
	log.Println("Running database seeds...")

	// Seed users
	if err := SeedUsers(db); err != nil {
		return err
	}

	// Add more seed functions here as needed:
	// if err := SeedPosts(db); err != nil {
	//     return err
	// }

	log.Println("All seeds completed successfully")
	return nil
}
