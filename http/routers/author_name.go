package routers

import (
	"net/http"
	"strconv"

	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/pages"
)

type authorHttpAdapter struct {
	pollService services.PollService
}

func NewAuthorHttpAdapter(router *http.ServeMux, pollService services.PollService) authorHttpAdapter {
	adapter := authorHttpAdapter{pollService}
	adapter.registerRoutes(router)
	return adapter
}

func (adapter *authorHttpAdapter) registerRoutes(router *http.ServeMux) {
	isAuthorized := middleware.IsAuthorizedAsAuthor()
	router.Handle("GET /{author_name}", isAuthorized(http.HandlerFunc(adapter.getAuthorPage)))
	router.Handle("POST /{author_name}", isAuthorized(http.HandlerFunc(adapter.finalizePoll)))
}

func (adapter *authorHttpAdapter) getAuthorPage(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		logger.Warn.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authorId := r.PathValue("author_name")

	if authorId != session.Profile.UserId {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	polls, err := adapter.pollService.GetCreatedBy(authorId)

	if err != nil {
		logger.Error.Println(err)
	}

	pages.AuthorPage(session, polls).Render(r.Context(), w)
}

func (adapter *authorHttpAdapter) finalizePoll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pollId, err := strconv.Atoi(r.Form["pollId"][0])

	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	adapter.pollService.ExpirePoll(pollId)

	w.WriteHeader(http.StatusOK)
}
