package api

import "golang-mongo-graphql-003/internal/dataloaders"

// import "golang-mongo-graphql-003/internal/dataloaders"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DataLoaders dataloaders.Retriever
}
