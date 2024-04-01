package ports

import (
	"bythecover/backend/internal/core/domain"
)

type VoteRepo interface {
	SubmitVote(domain.Vote) error
}

type VoteService interface {
	SubmitVote(int) error
}
