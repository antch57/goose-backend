package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBCredentials struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func InitDB() (*gorm.DB, error) {

	// FIXME: don't hardcode credentials
	creds := DBCredentials{
		Host:     "localhost",
		Port:     "3306",
		Database: "music_library",
		Username: "root",
		Password: "my-secret-pw",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO: proper error handling
		return nil, err
	}

	return db, nil
}
