package trip

import (
	"context"
	"fmt"

	"eztrip/api-go/llm"

	"github.com/google/uuid"
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

// GetByID retrieves a trip by ID
func (s *Service) GetByID(ctx context.Context, id string) (*Trip, error) {
	tripID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid trip ID: %w", err)
	}

	var trip Trip
	if err := s.db.WithContext(ctx).First(&trip, "id = ?", tripID).Error; err != nil {
		return nil, fmt.Errorf("trip not found: %w", err)
	}

	return &trip, nil
}

// GetSuggestion generates an AI-powered travel suggestion
func (s *Service) GetSuggestion(ctx context.Context, prompt string) (string, error) {
	if s.llm == nil {
		return "", fmt.Errorf("AI features are not available")
	}

	return s.llm.Complete(ctx, systemPrompt, prompt)
}
