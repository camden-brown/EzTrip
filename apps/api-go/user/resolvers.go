package user

import (
	"context"

	"eztrip/api-go/middleware"
	"eztrip/api-go/validation"
)

type Resolver struct {
	Service *Service
}

func NewResolver(service *Service) *Resolver {
	return &Resolver{
		Service: service,
	}
}

func (r *Resolver) CurrentUser(ctx context.Context) (*User, error) {
	auth0UserID := middleware.GetUserIDFromContext(ctx)

	return r.Service.GetByAuth0ID(ctx, auth0UserID)
}

func (r *Resolver) Users(ctx context.Context) ([]*User, error) {
	return r.Service.GetAll(ctx)
}

func (r *Resolver) User(ctx context.Context, id string) (*User, error) {
	return r.Service.GetByID(ctx, id)
}

func (r *Resolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, err
	}

	return r.Service.Create(ctx, input)
}
