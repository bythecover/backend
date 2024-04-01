package ports

import (
	"bythecover/backend/internal/core/domain"
)

type HtmxService interface {
	// Pages
	VotePage(domain.Poll) error
	SubmitVote() error
}
