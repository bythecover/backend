package persistence

import (
	"bythecover/backend/internal/core/domain"
	"database/sql"

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

func (repo pollPostgresRepository) GetById(id int) (domain.Poll, error) {
	var poll domain.Poll
	err := repo.db.QueryRow("SELECT id, title, created_by, created_at, expiration_date, expired FROM poll_events WHERE id = $1", id).Scan(&poll.Id, &poll.Title, &poll.CreatedBy, &poll.CreatedAt, &poll.ExpirationDate, &poll.Expired)

	if err != nil {
		return domain.Poll{}, err
	}

	rows, err := repo.db.Query("SELECT name, image FROM option WHERE poll_event_id = $1", id)

	if err != nil {
		return domain.Poll{}, err
	}

	var options []domain.Option
	for rows.Next() {
		var option domain.Option
		rows.Scan(&option.Name, &option.Image)
		options = append(options, option)
	}

	poll.Options = options

	return poll, nil
}
