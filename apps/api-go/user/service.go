package user

import (
	"context"
	"strings"

	"eztrip/api-go/auth0"
	appErrors "eztrip/api-go/errors"
	"eztrip/api-go/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db          *gorm.DB
	auth0Client *auth0.Client
}

func NewService(db *gorm.DB) *Service {
	auth0Client, err := auth0.NewClient()
	if err != nil {
		logger.Log.WithError(err).Error("Failed to initialize Auth0 client")
		// Continue without Auth0 client - will fail on user creation attempts
	}

	return &Service{
		db:          db,
		auth0Client: auth0Client,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	result := s.db.WithContext(ctx).Order("created_at DESC").Find(&users)
	if result.Error != nil {
		logger.Log.WithError(result.Error).Error("Failed to fetch all users")
		return nil, appErrors.Internal("Failed to fetch users")
	}
	logger.Log.WithField("count", len(users)).Debug("Fetched users")
	return users, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	var user User
	result := s.db.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			logger.Log.WithField("id", id).Warn("User not found")
			return nil, appErrors.NotFound("User")
		}
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Failed to fetch user by ID")
		return nil, appErrors.Internal("Failed to fetch user")
	}
	logger.Log.WithField("id", id).Debug("Fetched user by ID")
	return &user, nil
}

func (s *Service) GetByAuth0ID(ctx context.Context, auth0UserID string) (*User, error) {
	var user User
	result := s.db.WithContext(ctx).Where("auth0_user_id = ?", auth0UserID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			logger.Log.WithField("auth0_user_id", auth0UserID).Warn("User not found")
			return nil, appErrors.NotFound("User")
		}
		logger.Log.WithFields(logrus.Fields{
			"auth0_user_id": auth0UserID,
			"error":         result.Error,
		}).Error("Failed to fetch user by Auth0 ID")
		return nil, appErrors.Internal("Failed to fetch user")
	}
	logger.Log.WithField("auth0_user_id", auth0UserID).Debug("Fetched user by Auth0 ID")
	return &user, nil
}

func (s *Service) GetCurrent(ctx context.Context) (*User, error) {
	var current User
	result := s.db.WithContext(ctx).Order("created_at DESC").Limit(1).First(&current)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			logger.Log.Warn("No current user found")
			return nil, appErrors.NotFound("User")
		}
		logger.Log.WithError(result.Error).Error("Failed to fetch current user")
		return nil, appErrors.Internal("Failed to fetch current user")
	}
	return &current, nil
}

func (s *Service) Create(ctx context.Context, input CreateUserInput) (*User, error) {
	if s.auth0Client == nil {
		return nil, appErrors.Internal("Auth0 client not initialized")
	}

	var createdUser *User

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		user, err := s.createUserInDatabase(tx, input)
		if err != nil {
			return err
		}

		auth0UserID, err := s.createUserInAuth0(input)
		if err != nil {
			return err
		}

		if err := s.linkAuth0User(tx, user, auth0UserID); err != nil {
			s.cleanupAuth0User(auth0UserID)
			return err
		}

		user.Auth0UserID = &auth0UserID
		createdUser = user
		return nil
	})

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": input.Email,
			"error": err,
		}).Error("User creation transaction failed")
		return nil, err
	}

	logger.Log.WithFields(logrus.Fields{
		"id":            createdUser.ID,
		"email":         createdUser.Email,
		"auth0_user_id": *createdUser.Auth0UserID,
	}).Info("User created successfully")

	return createdUser, nil
}

func (s *Service) createUserInDatabase(tx *gorm.DB, input CreateUserInput) (*User, error) {
	user := User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	if err := tx.Create(&user).Error; err != nil {
		if dupErr := s.checkDuplicateEmail(err, input.Email); dupErr != nil {
			return nil, dupErr
		}

		logger.Log.WithFields(logrus.Fields{
			"email": input.Email,
			"error": err,
		}).Error("Failed to create user in database")
		return nil, appErrors.Internal("Failed to create user")
	}

	logger.Log.WithFields(logrus.Fields{
		"id":    user.ID,
		"email": user.Email,
	}).Debug("User created in database")

	return &user, nil
}

func (s *Service) checkDuplicateEmail(err error, email string) error {
	if strings.Contains(err.Error(), "duplicate key") ||
		strings.Contains(err.Error(), "unique constraint") {
		logger.Log.WithField("email", email).Warn("Duplicate email attempted")
		return DuplicateEmailError()
	}
	return nil
}

func (s *Service) createUserInAuth0(input CreateUserInput) (string, error) {
	auth0User, err := s.auth0Client.CreateUser(input.Email, input.Password, input.FirstName, input.LastName)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": input.Email,
			"error": err,
		}).Error("Failed to create user in Auth0")
		return "", appErrors.Internal("Failed to create user account")
	}

	logger.Log.WithField("auth0_user_id", auth0User.UserID).Debug("User created in Auth0")
	return auth0User.UserID, nil
}

func (s *Service) linkAuth0User(tx *gorm.DB, user *User, auth0UserID string) error {
	if err := tx.Model(user).Update("auth0_user_id", auth0UserID).Error; err != nil {
		logger.Log.WithFields(logrus.Fields{
			"user_id":       user.ID,
			"auth0_user_id": auth0UserID,
			"error":         err,
		}).Error("Failed to update user with Auth0 ID")
		return appErrors.Internal("Failed to link user account")
	}

	logger.Log.WithFields(logrus.Fields{
		"user_id":       user.ID,
		"auth0_user_id": auth0UserID,
	}).Debug("Linked Auth0 user to database record")

	return nil
}

func (s *Service) cleanupAuth0User(auth0UserID string) {
	if err := s.auth0Client.DeleteUser(auth0UserID); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"auth0_user_id": auth0UserID,
			"error":         err,
		}).Error("Failed to cleanup Auth0 user after transaction failure")
	}
}

func (s *Service) Update(ctx context.Context, id string, input UpdateUserInput) (*User, error) {
	var user User

	result := s.db.WithContext(ctx).Model(&user).Where("id = ?", id).Updates(input)
	if result.Error != nil {
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Failed to update user")
		return nil, appErrors.Internal("Failed to update user")
	}

	if result.RowsAffected == 0 {
		logger.Log.WithField("id", id).Warn("User not found for update")
		return nil, appErrors.NotFound("User")
	}

	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": err,
		}).Error("Failed to fetch updated user")
		return nil, appErrors.Internal("Failed to fetch updated user")
	}

	logger.Log.WithField("id", id).Info("User updated successfully")
	return &user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	result := s.db.WithContext(ctx).Delete(&User{}, "id = ?", id)
	if result.Error != nil {
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Failed to delete user")
		return appErrors.Internal("Failed to delete user")
	}
	if result.RowsAffected == 0 {
		logger.Log.WithField("id", id).Warn("User not found for deletion")
		return appErrors.NotFound("User")
	}
	logger.Log.WithField("id", id).Info("User deleted successfully")
	return nil
}
