package persistence

import (
	"database/sql"
	"github.com/bythecover/backend/model"
	"log"

	_ "github.com/lib/pq"
)

type votePostgresRepository struct {
	db *sql.DB
}

type VoteRepo interface {
	SubmitVote(model.Vote) error
	HasUserVoted(string, int) bool
}

func NewVotePostgresRepository(db *sql.DB) votePostgresRepository {
	if db == nil {
		log.Fatalln("SQL Client was not passed to Vote Repo Constructor")
	}

	return votePostgresRepository{
		db,
	}
}

func (repo votePostgresRepository) SubmitVote(submission model.Vote) error {
	stmt, err := repo.db.Prepare("INSERT INTO votes (selection, poll_event_id, source, user_id) VALUES($1, $2, $3, $4)")
	defer stmt.Close()

	if err != nil {
		log.Print(err)
		return err
	}

	log.Println(submission)
	_, err2 := stmt.Exec(submission.Selection, submission.PollEventId, submission.Source, submission.UserId)

	if err2 != nil {
		return err
	}

	return nil
}

func (repo votePostgresRepository) HasUserVoted(userId string, pollId int) bool {
	var foundId string
	err := repo.db.QueryRow("SELECT id FROM votes WHERE user_id = $1 AND poll_event_id = $2", userId, pollId).Scan(&foundId)

	if err != nil {
		log.Println(err)
		return false
	}

	return foundId != ""
}
