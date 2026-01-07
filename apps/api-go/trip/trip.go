package trip

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Trip represents a travel itinerary with multiple days and activities
type Trip struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     uuid.UUID      `gorm:"type:uuid;not null;index"`
	Title       string         `gorm:"column:title;not null"`
	Destination string         `gorm:"column:destination;not null"`
	StartDate   time.Time      `gorm:"column:start_date;not null"`
	EndDate     time.Time      `gorm:"column:end_date;not null"`
	Travelers   int            `gorm:"column:travelers;default:1"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Relationships
	Itinerary     []ItineraryDay     `gorm:"foreignKey:TripID;constraint:OnDelete:CASCADE"`
	Collaborators []TripCollaborator `gorm:"foreignKey:TripID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for the Trip model
func (Trip) TableName() string {
	return "trips"
}

