# graphql with go, gin, gqlgen


end goal is to make a website with this... for now its my playground to learn some graphql.

run: `go server.go` to start the graphql server
run: `go run github.com/99designs/gqlgen generate` to regenerate resolvers if you make changes to your graphql schema.


flow goes like this:

- `server.go` is your graphql entry point
- `src/` holds packages that get called from resolvers
- `graph/resolver.go` a spot where you can inject dependencies of graphql. DB for example.
- `schema.graphqls` where you define all your data shapes
- `schema.resolvers.go` where you call your packages you wrote in `src/`
- `graph/model/models_gen.go` import this file when you want to use your data structures as types. for example import into your packages that live in `src/`.

