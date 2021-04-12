package models

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Book struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	BookID      string    `json:"bookId" bson:"bookId,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" bson:"updatedDate,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Cover       string    `json:"cover" bson:"cover,omitempty"`
	AuthorIDs   []string  `json:"authorIds" bson:"authorIds,omitempty"`
}

type CreateBookInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	AuthorIDs   []string `json:"authorIds"`
}

func NewBook(author Book) Book {
	return Book{
		ID:          author.ID,
		BookID:      author.BookID,
		CreatedDate: author.CreatedDate,
		UpdatedDate: author.UpdatedDate,
		Title:       author.Title,
		Description: author.Description,
		Cover:       author.Cover,
		AuthorIDs:   author.AuthorIDs,
	}
}

func PointerBook(book Book) *Book {
	pointerBook := book
	return &pointerBook
}

func PointerBooks(books []Book) []*Book {
	var pointerBooks []*Book
	for _, book := range books {
		pointerBooks = append(pointerBooks, PointerBook(book))
	}
	return pointerBooks
}
