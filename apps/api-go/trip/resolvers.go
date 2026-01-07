package trip

import (
	"context"

	"github.com/google/uuid"
)

// Resolver handles GraphQL resolver operations for trips
type Resolver struct {
	Service *Service
}

// NewResolver creates a new trip resolver
func NewResolver(service *Service) *Resolver {
	return &Resolver{
		Service: service,
	}
}

// Trips returns all trips for the authenticated user
func (r *Resolver) Trips(ctx context.Context) ([]Trip, error) {
	return r.Service.GetAll(ctx)
}

// Trip returns a single trip by ID
func (r *Resolver) Trip(ctx context.Context, id string) (*Trip, error) {
	tripID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.Service.GetByID(ctx, tripID)
}

// Activity returns a single activity by ID
func (r *Resolver) Activity(ctx context.Context, id string) (*Activity, error) {
	activityID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return r.Service.GetActivityByID(ctx, activityID)
}

// TripSuggestion generates an AI-powered travel suggestion
func (r *Resolver) TripSuggestion(ctx context.Context, prompt string) (string, error) {
	return r.Service.GetSuggestion(ctx, prompt)
}
