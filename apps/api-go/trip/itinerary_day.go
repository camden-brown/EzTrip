package trip

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ItineraryDay represents a single day in a trip's itinerary
type ItineraryDay struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TripID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	Date      time.Time      `gorm:"column:date;not null"`
	DayNumber int            `gorm:"column:day_number;not null"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`

	// Relationships
	Activities []Activity `gorm:"foreignKey:ItineraryDayID;constraint:OnDelete:CASCADE"`
}

// TableName specifies the table name for the ItineraryDay model
func (ItineraryDay) TableName() string {
	return "itinerary_days"
}
