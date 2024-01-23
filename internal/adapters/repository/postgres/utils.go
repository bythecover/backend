package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDatabase() *sql.DB {
	connStr := "user=postgres_admin dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}