package ports

import (
	"bythecover/backend/internal/core/domain"
	"context"
)

type VoteRepo interface {
	SubmitVote(context.Context, domain.Vote) error
}

type VoteService interface {
	SubmitVote(context.Context, int) error
}
