package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/iamganeshagrawal/go-gql-upload-demo/graph/generated"
	"github.com/iamganeshagrawal/go-gql-upload-demo/models"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*models.File, error) {
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	newFileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	uploadFilePath := fmt.Sprintf("./uploads/%s", newFileName)
	if err = ioutil.WriteFile(uploadFilePath, content, 0644); err != nil {
		return nil, err
	}

	return &models.File{
		ID:          int(file.Size),
		Name:        file.Filename,
		Content:     fmt.Sprintf("http://localhost:8080/upload/%s", newFileName),
		ContentType: file.ContentType,
	}, nil
}

func (r *mutationResolver) SingleUploadWithPayload(ctx context.Context, req models.UploadFile) (*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MultipleUpload(ctx context.Context, files []*graphql.Upload) ([]*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*models.UploadFile) ([]*models.File, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
