# Jam Statz Backend

## Overview
This is the backend for a website that will track a band's performances and albums. The project is actively being worked on and is not yet fully built out. Right now jam-statz-backend is written in go using the gin web framework. It utilizes graphql as the query language and MariaDB for storage.

This project is linked directly with the [jam-statz-frontend](https://github.com/antch57/jam-statz-frontend) repo. If you run both repos locally they will integrate out of the box.

## Quick Start:

### Dev Database setup

Build dockerfile:
- `docker build -t jam-statz-db pkg/db/`

Run docker image:
- `docker run --detach --name jam-statz-db --env MARIADB_ROOT_PASSWORD=my-secret-pw -p 3306:3306 jam-statz-db --general_log_file=/var/lib/mysql/mysql.log --general_log=1`

you now have your db setup with logging enabled 👍

if you want to access your db just run:
- `docker exec -it test-jam-statz-mariadb mariadb -u root -p`

just be sure to swap out for your creds.

### Install Dependencies and Start Backend

Installing dependencies:
- `go mod download` 
- `go mod tidy`

Start backend:
- `go run server.go`

  
### Handy Commands

if you update `graph/schemas/schema.graphqls` you have to regenerate by running the following command:
- `go run github.com/99designs/gqlgen generate`

to start your graphql server:
- `go run server.go`

to test out your grahpql:
- start graphql server
- go to localhost:8080

to access db:
- `docker exec -it test-jam-statz-mariadb mariadb -u root -p`

to tail db logs:
- `docker exec -it test-jam-statz-mariadb tail -f /var/lib/mysql/mysql.log`
 
