package dataloaders

//go:generate go run github.com/vektah/dataloaden AgentLoader string *golang-mongo-graphql-003/internal/models.Agent

import (
	"context"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type Loaders struct {
	// individual loaders will be defined here
}

func newLoaders(ctx context.Context) *Loaders {
	return &Loaders{
		// individual loaders will be initialized here
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}
