package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/web_logger/models"

)

// structure for functions that access the database
type DB struct {
	*sql.DB
}

// Preparing an expression for connecting to the database
func ReadDatabaseSettings(conf models.Config) string {
	DbDriver := conf.DataBase.Driver
	DbUsername := conf.DataBase.Username
	DbPassword := conf.DataBase.Password
	DbHost := conf.DataBase.Host
	DbPort := conf.DataBase.Port
	DbName := conf.DataBase.Name
	DbSslMode := conf.DataBase.SslMode

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		DbDriver, DbUsername, DbPassword, DbHost, DbPort, DbName, DbSslMode)
}

// Creating new database
func NewDB(conf models.Config) (*DB, error) {
	dbSourceName := ReadDatabaseSettings(conf)
	db, err := sql.Open("postgres", dbSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return &DB{db}, nil
}
