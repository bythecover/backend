package ports

import (
	"bythecover/backend/internal/core/domain"
)

type PollRepo interface {
	GetById(int) (domain.Poll, error)
}

type PollService interface {
	GetById(int) (domain.Poll, error)
}

