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

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

// createdComments, createCommentssErr := comment.CreateComments([]comment.Comment{
// 	{
// 		IssueID: input.IssueID,
// 		Body:    input.Body,
// 	},
// })
// if createCommentssErr != nil {
// 	return nil, createCommentssErr
// }
// return comment.PointerComment(createdComments[0]), nil

// "github.com/99designs/gqlgen/graphql/handler"
// "github.com/99designs/gqlgen/graphql/playground"

// Defining the Graphql handler
func graphqlHandler(dl dataloaders.Retriever) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{
		DataLoaders: dl,
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
	// initialize the dataloaders
	dl := dataloaders.NewRetriever() // <- here we initialize the dataloader.Retriever
	// Setting up Gin
	r := gin.Default()
	r.Use(dataloaders.Middleware())
	r.Use(middleware.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler(dl))
	r.GET("/", playgroundHandler())
	r.Run()
}
