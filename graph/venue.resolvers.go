package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/antch57/jam-statz/graph/model"
)

// Performances is the resolver for the performances field.
func (r *venueResolver) Performances(ctx context.Context, obj *model.Venue) ([]*model.Performance, error) {
	res, err := r.PerformanceRepo.GetPerformancesByVenueID(obj.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Venue returns VenueResolver implementation.
func (r *Resolver) Venue() VenueResolver { return &venueResolver{r} }

type venueResolver struct{ *Resolver }