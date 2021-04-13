package main

import (
	"golang-mongo-graphql-003/internal/api"
	"golang-mongo-graphql-003/internal/api/generated"
	"golang-mongo-graphql-003/internal/dataloaders"
	"golang-mongo-graphql-003/internal/middleware"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Graphql handler
func graphqlHandler(dataloadersRetriever dataloaders.Retriever) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{
		DataLoaders: dataloadersRetriever,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	r.Use(dataloaders.Middleware())
	r.Use(middleware.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler(dataloaders.NewRetriever() /* <- here we initialize the dataloader.Retriever */))
	r.GET("/", playgroundHandler())
	r.Run()
}
