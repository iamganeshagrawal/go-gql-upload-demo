package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/iamganeshagrawal/go-gql-upload-demo/graph/generated"
	"github.com/iamganeshagrawal/go-gql-upload-demo/models"
	"github.com/iamganeshagrawal/go-gql-upload-demo/utils"
)

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return time.Now().Local().Format(time.RFC822), nil
}

func (r *queryResolver) Test(ctx context.Context, satisfies []string) (*models.TestResponse, error) {
	preloads := utils.GetPreloads(ctx, satisfies)
	fmt.Println("---------------------------------------------------------------")
	// // opCtx := graphql.GetOperationContext(ctx)
	// fields := graphql.CollectFieldsCtx(ctx, nil)
	// for _, field := range fields {
	// 	fmt.Println([]interface{}{field.Name, field.Definition.Type.Name(), field.ObjectDefinition.Name, field.ObjectDefinition.OneOf("Collection")})
	// }

	return &models.TestResponse{
		Preloads: preloads,
		A:        nil,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
