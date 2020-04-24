package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/ipreferwater/netflikss-golang/configuration"
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

func (r *mutationResolver) UpdateConfig(ctx context.Context, input *model.InputConfiguration) (bool, error) {
	var errors []string
	var copyDiConf = di.Configuration

	if input.FileServerPath != nil {
		copyDiConf.FileServerPath = *input.FileServerPath
	}

	if input.StockPath != nil {
		copyDiConf.StockPath = *input.StockPath
	}

	if input.Port != nil {
		if organizer.IsPortNumber(*input.Port) {
			copyDiConf.ServerConfiguration.Port = *input.Port
		} else {
			errors = append(errors, "the port number is invalid format")
		}
	}

	if input.AllowedOrigin != nil {
		if organizer.IsURL(*input.AllowedOrigin) {
			copyDiConf.ServerConfiguration.AllowedOrigin = *input.AllowedOrigin
		} else {
			errors = append(errors, "the allowedOrigin is invalid format")
		}
	}

	//TODO if no value has changed, dont update it
	/*if copyDiConf == di.Configuration {
		return false, fmt.Errorf("UpdateConfig won't update because the config received is the same")
	}*/

	if len(errors) > 0 {
		return false, fmt.Errorf(strings.Join(errors, ","))
	}

	configuration.SetConfiguration(copyDiConf)
	//TODO if we changed the server configuration, we might need to reboot it ?

	return true, nil
}

func (r *queryResolver) Netflikss(ctx context.Context) (*model.Data, error) {
	configuration := di.Configuration

	data := &model.Data{
		Series:        r.series,
		Configuration: &configuration,
	}
	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
