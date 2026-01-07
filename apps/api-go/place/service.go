package place

import (
	"context"
	"fmt"
	"time"

	"eztrip/api-go/logger"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Service provides place-related business logic
type Service struct {
	db *gorm.DB
}

// NewService creates a new place service
func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// GetByGooglePlaceID retrieves a place by Google Place ID
func (s *Service) GetByGooglePlaceID(ctx context.Context, googlePlaceID string) (*Place, error) {
	var place Place

	err := s.db.WithContext(ctx).
		Where("google_place_id = ?", googlePlaceID).
		First(&place).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Log.WithFields(logrus.Fields{
			"google_place_id": googlePlaceID,
			"error":           err.Error(),
		}).Error("Failed to fetch place by Google Place ID")
		return nil, fmt.Errorf("failed to fetch place: %w", err)
	}

	return &place, nil
}

// GetByID retrieves a place by internal ID
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Place, error) {
	var place Place

	err := s.db.WithContext(ctx).First(&place, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Log.WithFields(logrus.Fields{
			"place_id": id,
			"error":    err.Error(),
		}).Error("Failed to fetch place by ID")
		return nil, fmt.Errorf("failed to fetch place: %w", err)
	}

	return &place, nil
}

// Create creates a new place record
func (s *Service) Create(ctx context.Context, place *Place) error {
	place.LastFetchedAt = time.Now()

	err := s.db.WithContext(ctx).Create(place).Error
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"google_place_id": place.GooglePlaceID,
			"name":            place.Name,
			"error":           err.Error(),
		}).Error("Failed to create place")
		return fmt.Errorf("failed to create place: %w", err)
	}

	logger.Log.WithFields(logrus.Fields{
		"place_id":        place.ID,
		"google_place_id": place.GooglePlaceID,
		"name":            place.Name,
	}).Info("Place created successfully")

	return nil
}

// Update updates an existing place record
func (s *Service) Update(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	updates["last_fetched_at"] = time.Now()

	result := s.db.WithContext(ctx).
		Model(&Place{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		logger.Log.WithFields(logrus.Fields{
			"place_id": id,
			"error":    result.Error.Error(),
		}).Error("Failed to update place")
		return fmt.Errorf("failed to update place: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("place not found")
	}

	logger.Log.WithFields(logrus.Fields{
		"place_id": id,
	}).Info("Place updated successfully")

	return nil
}

// RefreshIfStale checks if a place is stale and needs refreshing
func (s *Service) RefreshIfStale(ctx context.Context, place *Place) (bool, error) {
	if !place.IsStale() {
		return false, nil
	}

	logger.Log.WithFields(logrus.Fields{
		"place_id":        place.ID,
		"google_place_id": place.GooglePlaceID,
		"last_fetched_at": place.LastFetchedAt,
	}).Info("Place data is stale, refresh needed")

	// TODO: Implement Google Places API refresh in step 7
	// For now, just return true to indicate refresh is needed
	return true, nil
}

// GetOrCreate retrieves a place by Google Place ID or creates it if it doesn't exist
func (s *Service) GetOrCreate(ctx context.Context, googlePlaceID string) (*Place, error) {
	// Try to get existing place
	place, err := s.GetByGooglePlaceID(ctx, googlePlaceID)
	if err != nil {
		return nil, err
	}

	// If found, check if stale and return
	if place != nil {
		_, _ = s.RefreshIfStale(ctx, place)
		return place, nil
	}

	// TODO: If not found, fetch from Google Places API in step 7
	// For now, return nil to indicate place doesn't exist
	logger.Log.WithFields(logrus.Fields{
		"google_place_id": googlePlaceID,
	}).Warn("Place not found in database, Google API integration needed")

	return nil, fmt.Errorf("place not found and Google API not yet integrated")
}
