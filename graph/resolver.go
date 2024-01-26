package graph

import "github.com/antch57/goose/src/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbConn *db.DB
}
