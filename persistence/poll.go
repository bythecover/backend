package persistence

import (
	"database/sql"
	"github.com/bythecover/backend/model"

	_ "github.com/lib/pq"
)

type pollPostgresRepository struct {
	db *sql.DB
}

func NewPollPostgresRepository(db *sql.DB) pollPostgresRepository {
	return pollPostgresRepository{
		db,
	}
}

func (repo pollPostgresRepository) GetById(id int) (model.Poll, error) {
	var poll model.Poll
	err := repo.db.QueryRow("SELECT id, title, created_by, created_at, expiration_date, expired FROM poll_events WHERE id = $1", id).Scan(&poll.Id, &poll.Title, &poll.CreatedBy, &poll.CreatedAt, &poll.ExpirationDate, &poll.Expired)

	if err != nil {
		return model.Poll{}, err
	}

	rows, err := repo.db.Query("SELECT name, image, id FROM option WHERE poll_event_id = $1", id)

	if err != nil {
		return model.Poll{}, err
	}

	var options []model.Option
	for rows.Next() {
		var option model.Option
		rows.Scan(&option.Name, &option.Image, &option.Id)
		options = append(options, option)
	}

	poll.Options = options

	return poll, nil
}
