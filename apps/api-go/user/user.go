package user

import (
	"time"

	"gorm.io/gorm"
)

type CreateUserInput struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=100"`
	LastName  string `json:"lastName" validate:"required,min=1,max=100"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=72,password_complexity"`
}

type UpdateUserInput struct {
	FirstName *string `json:"firstName" validate:"omitempty,min=1,max=100"`
	LastName  *string `json:"lastName" validate:"omitempty,min=1,max=100"`
}

type User struct {
	ID          string         `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Auth0UserID *string        `json:"-" gorm:"column:auth0_user_id;uniqueIndex"` // Internal only - not exposed to client
	FirstName   string         `json:"firstName" gorm:"column:first_name;not null"`
	LastName    string         `json:"lastName" gorm:"column:last_name;not null"`
	Email       string         `json:"email" gorm:"uniqueIndex;not null"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete support
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}
