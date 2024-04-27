package http

import (
	"bythecover/backend/internal/core/ports"
	"bythecover/backend/internal/core/templates/components"
	"bythecover/backend/internal/core/templates/pages"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

type pollHttpHandler struct {
	poll ports.PollService
}

func NewPollHttpHandler(poll ports.PollService) pollHttpHandler {
	return pollHttpHandler{
		poll,
	}
}

func (adapter pollHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		// convert string to number
		id, _ := strconv.Atoi(r.PathValue("id"))

		poll, err := adapter.poll.GetById(id)
		if err != nil {
			log.Fatalln(err)
		}
		templ.Handler(pages.VotePage(poll)).ServeHTTP(w, r)
	})

	// vote submission
	router.HandleFunc("POST /polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		selectedId, _ := strconv.Atoi(r.PostFormValue("selection"))
		pollId, _ := strconv.Atoi(r.PathValue("id"))

		err := adapter.poll.SubmitVote(selectedId, pollId)
		dialog := components.Dialog(nil)

		if err != nil {
			dialog = components.Dialog(err)
		}

		templ.Handler(dialog).ServeHTTP(w, r)
	})
}
