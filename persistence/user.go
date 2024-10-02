package persistence

import (
	"database/sql"

	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"

	_ "github.com/lib/pq"
)

type userPostgresRepository struct {
	db *sql.DB
}

type UserRepo interface {
	Save(model.User) error
	GetAll() ([]model.UserResp, error)
	GetUser(string) (model.UserResp, error)
}

func NewUserPostgresRepository(db *sql.DB) userPostgresRepository {
	return userPostgresRepository{
		db,
	}
}

func (repo userPostgresRepository) Save(u model.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (id, user_role) VALUES($1, $2)")
	defer stmt.Close()

	if err != nil {
		logger.Error.Println(err)
		return err
	}

	_, err = stmt.Exec(u.Id, u.Role)

	if err != nil {
		logger.Error.Println(err)
		return err
	}

	return err
}

func (repo userPostgresRepository) GetAll() ([]model.UserResp, error) {
	rows, err := repo.db.Query("SELECT id, first_name, last_name, email, user_role, created_at from users")

	if err != nil {
		logger.Error.Println(err)
		return nil, err
	}

	var people []model.UserResp

	for rows.Next() {
		var person model.UserResp
		rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Email, &person.Role, &person.CreatedAt)
		people = append(people, person)
	}

	rows.Close()

	return people, nil
}

func (repo userPostgresRepository) GetUser(id string) (model.UserResp, error) {
	var user model.UserResp
	err := repo.db.QueryRow("SELECT id, user_role FROM users WHERE id = $1", id).Scan(&user.Id, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserResp{}, model.ErrUserNotFound
		}

		logger.Error.Println(err)
		return model.UserResp{}, err
	}

	return user, nil
}
