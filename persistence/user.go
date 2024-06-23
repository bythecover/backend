package persistence

import (
	"database/sql"
	"github.com/bythecover/backend/model"
	"log"

	_ "github.com/lib/pq"
)

type userPostgresRepository struct {
	db *sql.DB
}

type UserRepo interface {
	Save(model.User) error
	GetAll() ([]model.UserResp, error)
	GetUser(int) (model.UserResp, error)
}

func NewUserPostgresRepository(db *sql.DB) userPostgresRepository {
	return userPostgresRepository{
		db,
	}
}

func (repo userPostgresRepository) Save(u model.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (id, first_name, last_name, email, is_author) VALUES($1, $2, $3, $4, $5)")

	if err != nil {
		log.Print(err)
		return err
	}

	defer stmt.Close()
	stmt.Exec(u.Id, u.FirstName, u.LastName, u.Email, u.IsAuthor)

	return nil
}

func (repo userPostgresRepository) GetAll() ([]model.UserResp, error) {
	rows, err := repo.db.Query("SELECT id, first_name, last_name, email, is_author, created_at from users")

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var people []model.UserResp

	for rows.Next() {
		var person model.UserResp
		rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.IsAuthor, &person.CreatedAt)
		people = append(people, person)
	}

	rows.Close()

	return people, nil
}

func (repo userPostgresRepository) GetUser(id int) (model.UserResp, error) {
	var user model.UserResp
	err := repo.db.QueryRow("SELECT id, first_name, last_name, email, is_author, created_at FROM users WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.IsAuthor, &user.CreatedAt)

	if err != nil {
		log.Print(err)
		if err == sql.ErrNoRows {
			return model.UserResp{}, model.ErrUserNotFound
		}

		return model.UserResp{}, err
	}

	return user, nil
}
