package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/dankokin/web_logger/models"
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

func Setup(filename string, db *DB) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Setupfile opening error: ", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Error after opening setupfile: ", err)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Error after opening setupfile: ", err)
		panic(err)
	}

	command := string(bs)
	_, err = db.Exec(command)
	if err != nil {
		fmt.Println("Command error")
	}
}
