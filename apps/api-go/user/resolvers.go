package user

import (
	"context"

	"eztrip/api-go/middleware"
)

type Resolver struct {
	Service *Service
}

func NewResolver(service *Service) *Resolver {
	return &Resolver{
		Service: service,
	}
}

// CurrentUser returns the currently authenticated user based on the JWT token.
// The middleware ensures this endpoint is only called with a valid JWT.
func (r *Resolver) CurrentUser(ctx context.Context) (*User, error) {
	auth0UserID := middleware.GetUserIDFromContext(ctx)

	return r.Service.GetByAuth0ID(ctx, auth0UserID)
}

// Users returns all users in the system.
func (r *Resolver) Users(ctx context.Context) ([]*User, error) {
	return r.Service.GetAll(ctx)
}

// User returns a single user by ID.
func (r *Resolver) User(ctx context.Context, id string) (*User, error) {
	return r.Service.GetByID(ctx, id)
}

// CreateUser creates a new user with the provided input.
func (r *Resolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	return r.Service.Create(ctx, input)
}
