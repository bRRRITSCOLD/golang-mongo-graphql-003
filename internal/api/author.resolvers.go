package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-mongo-graphql-003/internal/api/generated"
	"golang-mongo-graphql-003/internal/data_management"
	"golang-mongo-graphql-003/internal/middleware"
	"golang-mongo-graphql-003/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *authorResolver) Agent(ctx context.Context, obj *models.Author) (*models.Agent, error) {
	gc, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.DataLoaders.Retrieve(gc.Request.Context()).AgentsByAgentIds.Load(obj.AgentID)
}

func (r *authorResolver) Books(ctx context.Context, obj *models.Author) ([]*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input *models.CreateAuthorInput) (*models.Author, error) {
	createdAuthors, createAuthorsErr := data_management.CreateAuthors([]models.Author{
		{
			Name:    input.Name,
			Website: input.Website,
			AgentID: input.AgentID,
			// BookIDs: input.BookIDs,
		},
	})
	if createAuthorsErr != nil {
		return nil, createAuthorsErr
	}
	return models.PointerAuthor(createdAuthors[0]), nil
}

func (r *queryResolver) Author(ctx context.Context, authorID string) (*models.Author, error) {
	foundAuthors, findAuthorsErr := data_management.FindAuthors(models.Author{
		AuthorID: authorID,
	})
	if findAuthorsErr != nil {
		return nil, findAuthorsErr
	}
	return models.PointerAuthor(foundAuthors[0]), nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*models.Author, error) {
	foundAuthors, findAuthorsErr := data_management.FindAuthors(bson.D{})
	if findAuthorsErr != nil {
		return nil, findAuthorsErr
	}
	return models.PointerAuthors(foundAuthors), nil
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

type authorResolver struct{ *Resolver }
