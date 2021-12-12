package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/iamganeshagrawal/go-gql-upload-demo/graph/generated"
)

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return time.Now().Local().Format(time.RFC822), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
