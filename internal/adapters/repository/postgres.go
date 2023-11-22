package repository

import (
	"bythecover/backend/internal/core/domain"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository() (*userPostgresRepository, error) {
	connStr := "user=postgres_admin dbname=postgres password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return &userPostgresRepository{
		db,
	}, nil
}

func (repo userPostgresRepository) Save(u domain.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (first_name, last_name, email, is_author, created_at) VALUES($1, $2, $3, $4, $5)")
	
	if err != nil {
		return err
	}

	defer stmt.Close()
	stmt.Exec(u.FirstName, u.LastName, u.Email, u.IsAuthor, u.CreatedAt)
	
	return nil
}

func (repo userPostgresRepository) GetAll() ([]domain.User, error) {
	rows, err := repo.db.Query("SELECT * from users")

	if err != nil {
		return nil, err
	}

	var people []domain.User

	for rows.Next() {
		var person domain.User
		rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.IsAuthor, &person.CreatedAt)
		people = append(people, person)
	}

	rows.Close()

	return people, nil
}