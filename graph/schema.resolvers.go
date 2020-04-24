package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ipreferwater/netflikss-golang/di"
	"github.com/ipreferwater/netflikss-golang/graph/generated"
	"github.com/ipreferwater/netflikss-golang/graph/model"
	"github.com/ipreferwater/netflikss-golang/organizer"
)

func (r *mutationResolver) BuildSeriesFromInfo(ctx context.Context, input *bool) (bool, error) {
	series := organizer.ReadAllInside()
	for idx := range series {
		r.series = append(r.series, &series[idx])
	}

	return true, nil
}

func (r *mutationResolver) CreateInfoJSON(ctx context.Context, input *bool) (bool, error) {
	organizer.BuildInfoJSONFile()

	return true, nil
}

func (r *queryResolver) Series(ctx context.Context) ([]*model.Serie, error) {
	return r.series, nil
}

func (r *queryResolver) Netflikss(ctx context.Context) (*model.Data, error) {
	configuration := &model.Configuration{
		FileServerPath: di.Configuration.FileServerPath,
	}
	data := &model.Data{
		Series:        r.series,
		Configuration: configuration,
	}
	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
