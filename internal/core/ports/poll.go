package ports

import (
	"bythecover/backend/internal/core/domain"
	"context"
)

type PollRepo interface {
	GetById(context.Context, int) (domain.Poll, error)
}

type PollService interface {
	GetById(context.Context, int) (domain.Poll, error)
}