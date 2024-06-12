package routers

import (
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/templates/components"
	"github.com/bythecover/backend/templates/pages"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

type pollHttpAdapter struct {
	poll services.PollService
}

func NewPollHttpAdapter(poll services.PollService, router *http.ServeMux) pollHttpAdapter {
	adapter := pollHttpAdapter{
		poll,
	}
	adapter.registerRoutes(router)
	return adapter
}

func (adapter pollHttpAdapter) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /polls/{id}", adapter.getPollPage)
	router.HandleFunc("POST /polls/{id}", adapter.submitVote)
}

func (adapter pollHttpAdapter) submitVote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	selectedId, _ := strconv.Atoi(r.PostFormValue("selection"))
	pollId, _ := strconv.Atoi(r.PathValue("id"))

	submission := model.Vote{
		Selection:   selectedId,
		PollEventId: pollId,
	}

	err := adapter.poll.SubmitVote(submission)
	dialog := components.Dialog(nil)

	if err != nil {
		dialog = components.Dialog(err)
	}

	templ.Handler(dialog).ServeHTTP(w, r)
}

func (adapter pollHttpAdapter) getPollPage(w http.ResponseWriter, r *http.Request) {
	// convert string to number
	id, _ := strconv.Atoi(r.PathValue("id"))
	poll, err := adapter.poll.GetById(id)
	if err != nil {
		log.Fatalln(err)
	}
	templ.Handler(pages.VotePage(poll)).ServeHTTP(w, r)
}
