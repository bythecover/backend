package user_repository

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) userPostgresRepository {
	return userPostgresRepository{
		db,
	}
}

func (repo userPostgresRepository) Save(u domain.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (first_name, last_name, email, is_author) VALUES($1, $2, $3, $4)")
	
	if err != nil {
		log.Print(err)
		return err
	}

	defer stmt.Close()
	stmt.Exec(u.FirstName, u.LastName, u.Email, u.IsAuthor)
	
	return nil
}

func (repo userPostgresRepository) GetAll() ([]ports.UserResp, error) {
	rows, err := repo.db.Query("SELECT id, first_name, last_name, email, is_author, created_at from users")

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var people []ports.UserResp

	for rows.Next() {
		var person ports.UserResp
		rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.IsAuthor, &person.CreatedAt)
		people = append(people, person)
	}

	rows.Close()

	return people, nil
}

func (repo userPostgresRepository) GetUser(id int, ctx context.Context) (ports.UserResp, error) {
	var user ports.UserResp
	err := repo.db.QueryRowContext(ctx, "SELECT id, first_name, last_name, email, is_author, created_at FROM users WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsAuthor, &user.CreatedAt)

	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			return ports.UserResp{}, ports.ErrUserNotFound
		}

		return ports.UserResp{}, err
	}

	return user, nil
}