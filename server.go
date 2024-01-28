package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/antch57/goose/graph"
	"github.com/antch57/goose/internal/db"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler(resolver *graph.Resolver) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

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

	// Open up DB connection to use throughout the app
	db.Open()
	defer db.Close()

	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler(&graph.Resolver{}))
	r.GET("/", playgroundHandler())
	r.Run()
}
