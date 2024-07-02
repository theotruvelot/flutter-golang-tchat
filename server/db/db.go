package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@postgres:5432/flutter_golang_tchat_db?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	if err := d.db.Close(); err != nil {
		panic(err)
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
