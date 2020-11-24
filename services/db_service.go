package services

import "database/sql"

type DB struct {
	*sql.DB
}

func NewDBConnection() *DB {
	return new(DB)
}
