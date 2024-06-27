package persistence

import (
	"database/sql"

	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"

	_ "github.com/lib/pq"
)

type votePostgresRepository struct {
	db *sql.DB
}

type Result struct {
	model.Option
	Total int
}

type VoteRepo interface {
	SubmitVote(model.Vote) error
	HasUserVoted(string, int) bool
	GetResults(int) []Result
}

func NewVotePostgresRepository(db *sql.DB) votePostgresRepository {
	if db == nil {
		logger.Error.Fatalln("SQL Client was not passed to Vote Repo Constructor")
	}

	return votePostgresRepository{
		db,
	}
}

func (repo votePostgresRepository) SubmitVote(submission model.Vote) error {
	stmt, err := repo.db.Prepare("INSERT INTO votes (selection, poll_event_id, source, user_id) VALUES($1, $2, $3, $4)")
	defer stmt.Close()

	if err != nil {
		logger.Error.Println(err)
		return err
	}

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
		return false
	}

	return foundId != ""
}

func (repo votePostgresRepository) GetResults(pollId int) []Result {
	stmt, err := repo.db.Prepare("SELECT option.name, option.image, option.id, COUNT(*) as total_votes FROM votes INNER JOIN option ON option.id = votes.selection WHERE votes.poll_event_id = $1 GROUP BY option.id ORDER BY total_votes DESC;")

	if err != nil {
		logger.Error.Println(err.Error())
	}

	var results []Result
	rows, err := stmt.Query(pollId)

	if err != nil {
		logger.Error.Println(err.Error())
	}

	for rows.Next() {
		var result Result
		err := rows.Scan(&result.Name, &result.Image, &result.Id, &result.Total)

		if err != nil {
			logger.Error.Println(err.Error())
		}

		results = append(results, result)
	}

	return results
}
