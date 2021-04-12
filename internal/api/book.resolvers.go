package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang-mongo-graphql-003/internal/api/generated"
	"golang-mongo-graphql-003/internal/data_management"
	"golang-mongo-graphql-003/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *bookResolver) Authors(ctx context.Context, obj *models.Book) ([]*models.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBook(ctx context.Context, input models.CreateBookInput) (*models.Book, error) {
	createdBooks, createBooksErr := data_management.CreateBooks([]models.Book{
		{
			Title:       input.Title,
			Description: input.Description,
			Cover:       input.Cover,
			AuthorIDs:   input.AuthorIDs,
		},
	})
	if createBooksErr != nil {
		return nil, createBooksErr
	}
	return models.PointerBook(createdBooks[0]), nil
}

func (r *queryResolver) Book(ctx context.Context, bookID string) (*models.Book, error) {
	foundBooks, findBooksErr := data_management.FindBooks(models.Book{
		BookID: bookID,
	})
	if findBooksErr != nil {
		return nil, findBooksErr
	}
	return models.PointerBook(foundBooks[0]), nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
	foundBooks, findBooksErr := data_management.FindBooks(bson.D{})
	if findBooksErr != nil {
		return nil, findBooksErr
	}
	return models.PointerBooks(foundBooks), nil
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
