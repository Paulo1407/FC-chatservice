// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package ondeck

import (
	"context"
)

type Querier interface {
	// Create a new city. The slug must be unique.
	// This is the second line of the comment
	// This is the third line
	CreateCity(ctx context.Context, arg CreateCityParams) (City, error)
	CreateVenue(ctx context.Context, arg CreateVenueParams) (int32, error)
	DeleteVenue(ctx context.Context, slug string) error
	GetCity(ctx context.Context, slug string) (City, error)
	GetVenue(ctx context.Context, arg GetVenueParams) (Venue, error)
	ListCities(ctx context.Context) ([]City, error)
	ListVenues(ctx context.Context, city string) ([]Venue, error)
	UpdateCityName(ctx context.Context, arg UpdateCityNameParams) error
	UpdateVenueName(ctx context.Context, arg UpdateVenueNameParams) (int32, error)
	VenueCountByCity(ctx context.Context) ([]VenueCountByCityRow, error)
}

var _ Querier = (*Queries)(nil)