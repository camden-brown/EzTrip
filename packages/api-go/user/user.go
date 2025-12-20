package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FirstName string         `json:"firstName" gorm:"column:first_name;not null"`
	LastName  string         `json:"lastName" gorm:"column:last_name;not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete support
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}
