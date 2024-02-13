# graphql with go, gin, gqlgen


end goal is to make a website with this...

run: `go server.go` to start the graphql server

run: `go run github.com/99designs/gqlgen generate` to regenerate resolvers if you make changes to your graphql schema.


flow goes like this:

- `server.go` is your graphql entry point
- `internal/` holds packages that get called from resolvers
- `schema.graphqls` where you define all your data shapes
- `schema.resolvers.go` where you call your packages you wrote in `internal/`
- `graph/model/models_gen.go` import this file when you want to use your data structures as types. for example import into your packages that live in `internal/`.

## to get started:
### setup db for deving

step 1 build dockerfile:
- `docker build -t jam-statz-mariadb pkg/db/`

step 2 run dockerfile:
- `docker run --detach --name test-jam-statz-mariadb --env MARIADB_ROOT_PASSWORD=my-secret-pw -p 3306:3306 jam-statz-mariadb --general_log_file=/var/lib/mysql/mysql.log --general_log=1`

you now have your db setup.

if you want to access your db just run:
- `docker exec -it test-jam-statz-mariadb mariadb -u root -p`

just be sure to swap out for your creds.

### handy commands

- if you update schema.graphqls you have to regenerate by running the following command: `go run github.com/99designs/gqlgen generate`

- to start your graphql server: `go run server.go`

- to test out your grahpql:
    - start graphql server
    - go to localhost:8080

- tail db logs:
    - `docker exec -it test-jam-statz-mariadb tail -f /var/lib/mysql/mysql.log`

## TODO:

major stuffs:
- impliment dataloader for graphql
- create auth workflow
