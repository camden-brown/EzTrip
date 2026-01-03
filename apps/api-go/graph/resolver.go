package graph

import (
	"eztrip/api-go/trip"
	"eztrip/api-go/user"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	UserResolver *user.Resolver
	TripResolver *trip.Resolver
}

func NewResolver(db *gorm.DB) *Resolver {
	userService := user.NewService(db)
	tripService := trip.NewService(db)

	return &Resolver{
		UserResolver: user.NewResolver(userService),
		TripResolver: trip.NewResolver(tripService),
	}
}
