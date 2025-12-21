package graph

import (
	"travel-app/api-go/user"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	UserService *user.Service
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver{
		UserService: user.NewService(db),
	}
}
