package graph

import (
	"eztrip/api-go/user"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	UserService  *user.Service
	UserResolver *user.Resolver
}

func NewResolver(db *gorm.DB) *Resolver {
	userService := user.NewService(db)
	return &Resolver{
		UserService:  userService,
		UserResolver: user.NewResolver(userService),
	}
}
