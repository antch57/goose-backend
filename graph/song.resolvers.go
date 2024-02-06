package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/antch57/goose/graph/model"
)

// Band is the resolver for the band field.
func (r *songResolver) Band(ctx context.Context, obj *model.Song) (*model.Band, error) {
	res, err := r.BandRepo.GetBandBySongId(obj.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Song returns SongResolver implementation.
func (r *Resolver) Song() SongResolver { return &songResolver{r} }

type songResolver struct{ *Resolver }
