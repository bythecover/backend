package persistence

import (
	"bythecover/backend/internal/core/domain"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type votePostgresRepository struct {
	db *sql.DB
}

func NewVotePostgresRepository(db *sql.DB) votePostgresRepository {
	return votePostgresRepository{
		db,
	}
}

func (repo votePostgresRepository) SubmitVote(submission domain.Vote) error {
	stmt, err := repo.db.Prepare("INSERT INTO votes (selection, poll_event_id, source, user_id) VALUES($1, $2, $3, $4)")
	defer stmt.Close()

	if err != nil {
		log.Print(err)
		return err
	}

	_, err2 := stmt.Exec(submission.Selection, submission.PollEventId, submission.Source, submission.UserId)

	if err2 != nil {
		return err
	}

	return nil
}
