package user

import (
	"context"

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
		return nil, result.Error
	}
	return users, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*User, error) {
	var user User
	result := s.db.WithContext(ctx).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (s *Service) Create(ctx context.Context, user *User) (*User, error) {
	result := s.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *Service) Update(ctx context.Context, id string, updates map[string]interface{}) (*User, error) {
	var user User
	result := s.db.WithContext(ctx).Model(&user).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	// Fetch updated user
	s.db.WithContext(ctx).Where("id = ?", id).First(&user)
	return &user, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	result := s.db.WithContext(ctx).Delete(&User{}, "id = ?", id)
	return result.Error
}
