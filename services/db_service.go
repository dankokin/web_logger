package services

import (
	"database/sql"
	"fmt"
	"github.com/web_logger/models"

	_ "github.com/lib/pq"
)

// structure for functions that access the database
type DB struct {
	*sql.DB
}

// Creating new database
func NewDB(conf models.Config) (*DB, error) {
	db, err := sql.Open("postgres", conf.ConnectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return &DB{db}, nil
}
