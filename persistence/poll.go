package persistence

import (
	"database/sql"

	"github.com/bythecover/backend/logger"
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

// Gets a poll by its id
func (repo pollPostgresRepository) GetById(bookId int, authorName string) (model.Poll, error) {
	var poll model.Poll
	err := repo.db.QueryRow("SELECT id, title, created_by, expired FROM poll_events WHERE id = $1 AND created_by = $2", bookId, authorName).Scan(&poll.Id, &poll.Title, &poll.CreatedBy, &poll.Expired)

	if err != nil {
		return model.Poll{}, err
	}

	rows, err := repo.db.Query("SELECT name, image, id FROM option WHERE poll_event_id = $1", bookId)

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

// Creates a poll
func (repo pollPostgresRepository) CreatePoll(poll model.Poll) error {
	stmt, err := repo.db.Prepare("INSERT INTO poll_events (title, created_by) VALUES ($1, $2)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(poll.Title, poll.CreatedBy)

	if err != nil {
		return err
	}

	var pollId int
	rows := repo.db.QueryRow("SELECT id FROM poll_events WHERE created_by = $1 ORDER BY poll_events.created_at DESC", poll.CreatedBy)
	rows.Scan(&pollId)

	for _, item := range poll.Options {
		stmt, err := repo.db.Prepare("INSERT INTO option (poll_event_id, name, image) VALUES ($1, $2, $3)")

		if err != nil {
			return err
		}

		_, err = stmt.Exec(pollId, item.Name, item.Image)

		if err != nil {
			return err
		}

	}

	return nil
}

// Gets a list of polls by the author id
func (repo pollPostgresRepository) GetCreatedBy(userId string) ([]model.Poll, error) {
	rows, err := repo.db.Query("SELECT id, title, expired FROM poll_events WHERE created_by = $1", userId)

	if err != nil {
		return nil, err
	}

	var polls []model.Poll
	for rows.Next() {
		var poll model.Poll
		err := rows.Scan(&poll.Id, &poll.Title, &poll.Expired)

		if err != nil {
			logger.Error.Println(err)
		}

		polls = append(polls, poll)
	}

	return polls, nil
}

func (repo pollPostgresRepository) ExpirePoll(pollId int) error {
	row := repo.db.QueryRow("UPDATE poll_events SET expired = true WHERE id = $1", pollId)

	return row.Err()
}
