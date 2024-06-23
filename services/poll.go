package services

import (
	"errors"

	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/persistence"
)

type pollRepo interface {
	GetById(int) (model.Poll, error)
}

type PollService struct {
	pollRepo pollRepo
	voteRepo persistence.VoteRepo
}

func NewPollService(pollRepo pollRepo, voteRepo persistence.VoteRepo) PollService {
	return PollService{
		pollRepo,
		voteRepo,
	}
}

func (service PollService) GetById(id int) (model.Poll, error) {
	return service.pollRepo.GetById(id)
}

func (service PollService) SubmitVote(vote model.Vote) error {
	if service.voteRepo.HasUserVoted(vote.UserId, vote.PollEventId) {
		return errors.New("User has already voted")
	}

	return service.voteRepo.SubmitVote(vote)
}
