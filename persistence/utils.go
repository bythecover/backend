package persistence

import (
	"database/sql"
	"os"

	"github.com/bythecover/backend/logger"
	_ "github.com/lib/pq"
)

func NewPostgresDatabase() *sql.DB {
	connStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		logger.Error.Fatalln(err)
	}

	return db
}
