package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
)

type pollService struct {
	pollRepo ports.PollRepo
	voteRepo ports.VoteRepo
}

func NewPollService(pollRepo ports.PollRepo, voteRepo ports.VoteRepo) pollService {
	return pollService{
		pollRepo,
		voteRepo,
	}
}

func (service pollService) GetById(id int) (domain.Poll, error) {
	poll, err := service.pollRepo.GetById(id)

	return poll, err
}

func (service pollService) SubmitVote(vote domain.Vote) error {
	return service.voteRepo.SubmitVote(vote)
}
