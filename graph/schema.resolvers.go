package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ipreferwater/netflikss-golang/graph/generated"
	"github.com/ipreferwater/netflikss-golang/graph/model"
)

func (r *mutationResolver) CreateSerie(ctx context.Context, input model.InputSerie) (*model.Serie, error) {
	/*serie := &model.Serie{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Label: input.Label,
	}
	r.series = append(r.series, serie)
	return serie, nil*/
	//we should trigger creation from files
	serie := &model.Serie{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Label: "fake",
	}
	return serie, nil
}

func (r *queryResolver) Series(ctx context.Context) ([]*model.Serie, error) {
	return r.series, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
