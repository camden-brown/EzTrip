package place

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Place represents cached data from Google Places API
type Place struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	GooglePlaceID    string         `gorm:"column:google_place_id;uniqueIndex;not null"` // Google's place_id
	Name             string         `gorm:"column:name;not null"`
	Rating           float64        `gorm:"column:rating"`
	ReviewCount      int            `gorm:"column:review_count"`
	PrimaryPhotoURL  string         `gorm:"column:primary_photo_url;type:text"` // Store one primary photo
	Address          string         `gorm:"column:address;type:text"`
	FormattedAddress string         `gorm:"column:formatted_address;type:text"`
	Website          string         `gorm:"column:website;type:text"`
	PhoneNumber      string         `gorm:"column:phone_number"`
	PriceLevel       int            `gorm:"column:price_level"` // 0-4 scale from Google
	LastFetchedAt    time.Time      `gorm:"column:last_fetched_at;not null"`
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// TableName specifies the table name for the Place model
func (Place) TableName() string {
	return "places"
}

// IsStale checks if the place data is older than 30 days
func (p *Place) IsStale() bool {
	return time.Since(p.LastFetchedAt) > 30*24*time.Hour
}
