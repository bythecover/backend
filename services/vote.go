package services

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"
)

type voteService struct {
	repo persistence.VoteRepo
}

func NewVoteService(repo persistence.VoteRepo) voteService {
	return voteService{
		repo,
	}
}

func (service voteService) SubmitVote(id int) error {
	// TODO - verify user is logged in and has not voted yet
	// TODO - don't use hardcoded values

	err := service.repo.SubmitVote(model.Vote{PollEventId: 1, UserId: 1, Selection: 1, Source: "web"})

	return err
}
