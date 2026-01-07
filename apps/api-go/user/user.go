package user

import (
	"context"
	"time"

	appErrors "eztrip/api-go/errors"
	"eztrip/api-go/logger"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userUUIDContextKey struct{}

// SetUserUUID stores the authenticated user's UUID in the request context.
func SetUserAuth0ID(ctx context.Context, auth0ID string) context.Context {
	return context.WithValue(ctx, userUUIDContextKey{}, auth0ID)
}

// GetUserAuth0ID retrieves the authenticated user's Auth0 ID from the request context.
// Returns empty string if not authenticated or user not found.
func GetUserAuth0ID(ctx context.Context) string {
	if auth0ID, ok := ctx.Value(userUUIDContextKey{}).(string); ok {
		return auth0ID
	}
	return ""
}

// GetAuthenticatedUser retrieves the authenticated user from the database and returns both the user and their UUID.
// This is a common utility function for services that need to verify user authentication.
func GetAuthenticatedUser(ctx context.Context, db *gorm.DB) (*User, uuid.UUID, error) {
	auth0ID := GetUserAuth0ID(ctx)
	if auth0ID == "" {
		return nil, uuid.Nil, appErrors.Unauthorized("User not authenticated")
	}

	var currentUser User
	if err := db.WithContext(ctx).Where("auth0_user_id = ?", auth0ID).First(&currentUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, uuid.Nil, appErrors.NotFound("User")
		}
		logger.Log.WithFields(logrus.Fields{
			"auth0_id": auth0ID,
			"error":    err.Error(),
		}).Error("Failed to fetch user by Auth0 ID")
		return nil, uuid.Nil, appErrors.Internal("Failed to fetch user")
	}

	return &currentUser, currentUser.ID, nil
}

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
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
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
