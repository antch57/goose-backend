package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/antch57/goose/graph"
	"github.com/antch57/goose/src/db"
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

// func helloHandler(resolver *graph.Resolver) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		data, err := resolver.
// 			c.JSON(200, gin.H{
// 			"message": "Hello world!",
// 		})
// 	}
// }

func main() {
	// DB connection
	creds := db.DBCredentials{
		Host:     "localhost",
		Port:     "3306",
		Database: "goose",
		Username: "admin",
		Password: "my-secret-pw",
	}

	conn, err := db.Open(creds)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Resolvers
	resolver := &graph.Resolver{}

	// Setting up Gin
	r := gin.Default()
	// r.GET("/hello", helloHandler())
	r.POST("/query", graphqlHandler(resolver))
	r.GET("/", playgroundHandler())
	r.Run()
}
