package user

import (
	"context"
	"fmt"

	"eztrip/api-go/rbac"
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
	userUUID := GetUserUUID(ctx)
	return r.Service.GetByID(ctx, userUUID)
}

func (r *Resolver) Users(ctx context.Context) ([]*User, error) {
	userUUID := GetUserUUID(ctx)

	if err := rbac.RequireAdminRole(ctx, userUUID); err != nil {
		return nil, err
	}

	return r.Service.GetAll(ctx)
}

func (r *Resolver) User(ctx context.Context, id string) (*User, error) {
	userUUID := GetUserUUID(ctx)

	if err := rbac.RequireAdminRole(ctx, userUUID); err != nil {
		return nil, err
	}

	return r.Service.GetByID(ctx, id)
}

func (r *Resolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return r.Service.Create(ctx, input)
}

func (r *Resolver) UpdateUser(ctx context.Context, input UpdateUserInput) (*User, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	userUUID := GetUserUUID(ctx)
	return r.Service.Update(ctx, userUUID, input)
}
