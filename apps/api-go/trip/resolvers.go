package trip

import (
	"context"
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

// TripSuggestion generates an AI-powered travel suggestion
func (r *Resolver) TripSuggestion(ctx context.Context, prompt string) (string, error) {
	return r.Service.GetSuggestion(ctx, prompt)
}
