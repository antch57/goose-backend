package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBCredentials struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var Db *sql.DB

func Open() {
	var err error
	// FIXME: don't hardcode credentials
	// DB connection
	creds := DBCredentials{
		Host:     "localhost",
		Port:     "3306",
		Database: "music_library",
		Username: "root",
		Password: "my-secret-pw",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	fmt.Printf("dsn: %s\n", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	Db = db
}

func Close() error {
	return Db.Close()
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := Db.Query(query, args...) // Call Query method on db.DB instead of db
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := Db.Exec(query, args...) // Call Exec method on db.DB instead of db
	if err != nil {
		return nil, fmt.Errorf("exec failed: %w", err)
	}
	return result, nil
}

func Transacntion() (*sql.Tx, error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}
	return tx, nil
}
