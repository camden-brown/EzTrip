package trip

import (
	"context"
	"fmt"

	appErrors "eztrip/api-go/errors"
	"eztrip/api-go/llm"
	"eztrip/api-go/logger"
	"eztrip/api-go/user"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	systemPrompt = "You are a helpful travel assistant. Provide personalized travel suggestions, recommendations, and advice. Be concise and friendly."
)

// Service handles trip operations
type Service struct {
	db  *gorm.DB
	llm *llm.Service
}

// NewService creates a new trip service
func NewService(db *gorm.DB) *Service {
	var llmService *llm.Service
	if svc, err := llm.NewDefaultService(); err == nil {
		llmService = svc
	}

	return &Service{
		db:  db,
		llm: llmService,
	}
}

// GetAll retrieves all trips for the authenticated user (owner or collaborator)
func (s *Service) GetAll(ctx context.Context) ([]Trip, error) {
	_, userID, err := user.GetAuthenticatedUser(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var trips []Trip
	err = s.db.WithContext(ctx).
		Preload("Itinerary.Activities").
		Preload("Collaborators").
		Where("owner_id = ?", userID).
		Or("id IN (?)",
			s.db.Table("trip_collaborators").
				Select("trip_id").
				Where("user_id = ? AND deleted_at IS NULL", userID),
		).
		Find(&trips).Error

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err.Error(),
		}).Error("Failed to fetch trips for user")
		return nil, appErrors.Internal("Failed to fetch trips")
	}

	return trips, nil
}

// GetByID retrieves a trip by ID with all related data
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*Trip, error) {
	_, userID, err := user.GetAuthenticatedUser(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var trip Trip
	err = s.db.WithContext(ctx).
		Preload("Itinerary", func(db *gorm.DB) *gorm.DB {
			return db.Order("date ASC")
		}).
		Preload("Itinerary.Activities", func(db *gorm.DB) *gorm.DB {
			return db.Order("time ASC")
		}).
		Preload("Collaborators").
		First(&trip, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appErrors.NotFound("Trip")
		}
		logger.Log.WithFields(logrus.Fields{
			"trip_id": id,
			"error":   err.Error(),
		}).Error("Failed to fetch trip by ID")
		return nil, appErrors.Internal("Failed to fetch trip")
	}

	// Check if user is owner or collaborator
	isOwner := trip.OwnerID == userID
	isCollaborator := false
	for _, collaborator := range trip.Collaborators {
		if collaborator.UserID == userID {
			isCollaborator = true
			break
		}
	}

	if !isOwner && !isCollaborator {
		logger.Log.WithFields(logrus.Fields{
			"trip_id": id,
			"user_id": userID,
		}).Warn("User attempted to access trip without permission")
		return nil, appErrors.Forbidden("You don't have permission to access this trip")
	}

	return &trip, nil
}

// GetActivityByID retrieves an activity by ID and verifies user authorization
func (s *Service) GetActivityByID(ctx context.Context, id uuid.UUID) (*Activity, error) {
	_, userID, err := user.GetAuthenticatedUser(ctx, s.db)
	if err != nil {
		return nil, err
	}

	var activity Activity
	err = s.db.WithContext(ctx).First(&activity, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, appErrors.NotFound("Activity")
		}
		logger.Log.WithFields(logrus.Fields{
			"activity_id": id,
			"error":       err.Error(),
		}).Error("Failed to fetch activity by ID")
		return nil, appErrors.Internal("Failed to fetch activity")
	}

	// Load the itinerary day to get the trip ID
	var itineraryDay ItineraryDay
	err = s.db.WithContext(ctx).First(&itineraryDay, "id = ?", activity.ItineraryDayID).Error
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"itinerary_day_id": activity.ItineraryDayID,
			"error":            err.Error(),
		}).Error("Failed to fetch itinerary day")
		return nil, appErrors.Internal("Failed to fetch itinerary day")
	}

	// Load the trip with collaborators to check authorization
	var trip Trip
	err = s.db.WithContext(ctx).
		Preload("Collaborators").
		First(&trip, "id = ?", itineraryDay.TripID).Error

	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"trip_id": itineraryDay.TripID,
			"error":   err.Error(),
		}).Error("Failed to fetch trip")
		return nil, appErrors.Internal("Failed to fetch trip")
	}

	// Check if user is owner or collaborator
	isOwner := trip.OwnerID == userID
	isCollaborator := false
	for _, collaborator := range trip.Collaborators {
		if collaborator.UserID == userID {
			isCollaborator = true
			break
		}
	}

	if !isOwner && !isCollaborator {
		logger.Log.WithFields(logrus.Fields{
			"activity_id": id,
			"trip_id":     itineraryDay.TripID,
			"user_id":     userID,
		}).Warn("User attempted to access activity without permission")
		return nil, appErrors.Forbidden("You don't have permission to access this activity")
	}

	return &activity, nil
}

// GetSuggestion generates an AI-powered travel suggestion
func (s *Service) GetSuggestion(ctx context.Context, prompt string) (string, error) {
	if s.llm == nil {
		return "", fmt.Errorf("AI features are not available")
	}

	return s.llm.Complete(ctx, systemPrompt, prompt)
}
