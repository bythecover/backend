package persistence

import (
	"bythecover/backend/internal/core/domain"
	"context"
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

func (repo votePostgresRepository) SubmitVote(ctx context.Context, submission domain.Vote) error {
	stmt, err := repo.db.PrepareContext(ctx, "INSERT INTO votes (selection, poll_event_id, source, user_id) VALUES($1, $2, $3, $4)")
	defer stmt.Close()

	if err != nil {
		log.Print(err)
		return err
	}

	res, err2 := stmt.ExecContext(ctx, submission.Selection, submission.PollEventId, submission.Source, submission.UserId)

	log.Print(res.LastInsertId())

	if err2 != nil {
		return err
	}

	return nil
}
