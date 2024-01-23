package services

import (
	"bythecover/backend/internal/core/domain"
	"bythecover/backend/internal/core/ports"
	"context"
)

type pollService struct{
	repo ports.PollRepo
}

func NewPollService(repo ports.PollRepo) pollService {
	return pollService{
		repo,
	}
}

func (service pollService) GetById(ctx context.Context, id int) (domain.Poll, error) {
	poll, err := service.repo.GetById(ctx, id)

	return poll, err
}
