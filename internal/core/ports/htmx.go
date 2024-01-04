package ports

import (
	"bythecover/backend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type HtmxService interface {
	// Pages
	VotePage(domain.Poll, *gin.Context) error
}
