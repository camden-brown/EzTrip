package user

import (
	"context"
	"fmt"

	"travel-app/api-go/logger"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	result := s.db.WithContext(ctx).Order("created_at DESC").Find(&users)
	if result.Error != nil {
		logger.Log.WithError(result.Error).Error("Failed to fetch all users")
		return nil, result.Error
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
			return nil, fmt.Errorf("user not found")
		}
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Failed to fetch user by ID")
		return nil, result.Error
	}
	logger.Log.WithField("id", id).Debug("Fetched user by ID")
	return &user, nil
}

func (s *Service) Create(ctx context.Context, user *User) (*User, error) {
	result := s.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		logger.Log.WithFields(logrus.Fields{
			"email": user.Email,
			"error": result.Error,
		}).Error("Failed to create user")
		return nil, result.Error
	}
	logger.Log.WithFields(logrus.Fields{
		"id":    user.ID,
		"email": user.Email,
	}).Info("User created successfully")
	return user, nil
}

func (s *Service) Update(ctx context.Context, id string, updates map[string]interface{}) (*User, error) {
	var user User
	result := s.db.WithContext(ctx).Model(&user).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		logger.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Failed to update user")
		return nil, result.Error
	}

	// Fetch updated user
	s.db.WithContext(ctx).Where("id = ?", id).First(&user)
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
		return result.Error
	}
	logger.Log.WithField("id", id).Info("User deleted successfully")
	return nil
}
