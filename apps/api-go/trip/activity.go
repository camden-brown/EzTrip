package trip

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ActivityType represents the type of activity
type ActivityType string

const (
	ActivityTypePlaceBased ActivityType = "place_based" // Activity linked to a Google Place
	ActivityTypeCustom     ActivityType = "custom"      // User-created activity without Place
	ActivityTypeTransport  ActivityType = "transport"   // Travel/transportation activity
)

// ActivityCategory represents the category of an activity
type ActivityCategory string

const (
	ActivityCategoryBeach         ActivityCategory = "beach"
	ActivityCategoryHike          ActivityCategory = "hike"
	ActivityCategoryFood          ActivityCategory = "food"
	ActivityCategoryHotel         ActivityCategory = "hotel"
	ActivityCategoryActivity      ActivityCategory = "activity"
	ActivityCategoryTransport     ActivityCategory = "transport"
	ActivityCategoryShopping      ActivityCategory = "shopping"
	ActivityCategoryEntertainment ActivityCategory = "entertainment"
)

// Activity represents a single activity in an itinerary day
type Activity struct {
	ID             uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ItineraryDayID uuid.UUID        `gorm:"type:uuid;not null;index"`
	PlaceID        *uuid.UUID       `gorm:"type:uuid;index"` // Nullable - references places table
	Type           ActivityType     `gorm:"column:type;not null;default:'place_based'"`
	Time           time.Time        `gorm:"column:time;not null"`
	Title          string           `gorm:"column:title;not null"`
	Location       string           `gorm:"column:location"`
	Category       ActivityCategory `gorm:"column:category;not null"`
	Description    string           `gorm:"column:description;type:text"`
	Notes          string           `gorm:"column:notes;type:text"`
	CreatedAt      time.Time        `gorm:"column:created_at"`
	UpdatedAt      time.Time        `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt   `gorm:"column:deleted_at;index"`
}

// TableName specifies the table name for the Activity model
func (Activity) TableName() string {
	return "activities"
}
