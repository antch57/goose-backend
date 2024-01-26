package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBCredentials struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type DB struct {
	*sql.DB
}

func Open(creds DBCredentials) (*DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	fmt.Printf("dsn: %s\n", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (conn *DB) Close() error {
	err := conn.DB.Close()
	if err != nil {
		return fmt.Errorf("close failed: %w", err)
	}
	return nil
}

func (conn *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, args...) // Call Query method on db.DB instead of db
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

func (conn *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := conn.DB.Exec(query, args...) // Call Exec method on db.DB instead of db
	if err != nil {
		return nil, fmt.Errorf("exec failed: %w", err)
	}
	return result, nil
}
