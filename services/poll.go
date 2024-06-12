package services

import (
	"github.com/bythecover/backend/model"
)

type pollRepo interface {
	GetById(int) (model.Poll, error)
}

type voteRepo interface {
	SubmitVote(model.Vote) error
}

type PollService struct {
	pollRepo pollRepo
	voteRepo voteRepo
}

func NewPollService(pollRepo pollRepo, voteRepo voteRepo) PollService {
	return PollService{
		pollRepo,
		voteRepo,
	}
}

func (service PollService) GetById(id int) (model.Poll, error) {
	return service.pollRepo.GetById(id)
}

func (service PollService) SubmitVote(vote model.Vote) error {
	return service.voteRepo.SubmitVote(vote)
}
