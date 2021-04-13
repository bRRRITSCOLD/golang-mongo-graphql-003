package models

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Author struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	AuthorID    string    `json:"authorId" bson:"authorId"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" bson:"updatedDate,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Website     string    `json:"website" bson:"website,omitempty"`
	AgentID     string    `json:"agentId" bson:"agentId,omitempty"`
}

type CreateAuthorInput struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	AgentID string `json:"agentId"`
}

func NewAuthor(author Author) Author {
	return Author{
		ID:          author.ID,
		AuthorID:    author.AuthorID,
		CreatedDate: author.CreatedDate,
		UpdatedDate: author.UpdatedDate,
		Name:        author.Name,
		Website:     author.Website,
		AgentID:     author.AgentID,
	}
}

func PointerAuthor(author Author) *Author {
	pointerAuthor := author
	return &pointerAuthor
}

func PointerAuthors(authors []Author) []*Author {
	var pointerAuthors []*Author
	for _, author := range authors {
		pointerAuthors = append(pointerAuthors, PointerAuthor(author))
	}
	return pointerAuthors
}
