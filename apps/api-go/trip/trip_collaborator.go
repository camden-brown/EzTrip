package trip

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TripCollaborator represents a user who has access to manage a trip
type TripCollaborator struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TripID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

// TableName specifies the table name for the TripCollaborator model
func (TripCollaborator) TableName() string {
	return "trip_collaborators"
}
