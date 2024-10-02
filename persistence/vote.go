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

func (repo votePostgresRepository) GetResults(pollId int) []model.PollResult {
	stmt, err := repo.db.Prepare("SELECT option.name, option.image, option.id, COUNT(votes.id) as total_votes FROM option LEFT JOIN votes ON option.id = votes.selection WHERE option.poll_event_id = $1 GROUP BY option.id ORDER BY total_votes DESC;")

	if err != nil {
		logger.Error.Println(err.Error())
	}

	var results []model.PollResult
	rows, err := stmt.Query(pollId)

	if err != nil {
		logger.Error.Println(err.Error())
	}

	for rows.Next() {
		var result model.PollResult
		err := rows.Scan(&result.Name, &result.Image, &result.Id, &result.Total)

		if err != nil {
			logger.Error.Println(err.Error())
		}

		results = append(results, result)
	}

	return results
}
