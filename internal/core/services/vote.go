package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"context"
)

type voteService struct {
	repo ports.VoteRepo
}

func NewVoteService(repo ports.VoteRepo) voteService {
	return voteService{
		repo,
	}
}

func (service voteService) SubmitVote(ctx context.Context, id int) error {
	// TODO - verify user is logged in and has not voted yet
	// TODO - don't use hardcoded values

	err := service.repo.SubmitVote(ctx, domain.Vote{PollEventId: 1, UserId: 1, Selection: 1, Source: "web"})

	return err
}
