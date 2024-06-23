package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/components"
	"github.com/bythecover/backend/templates/pages"
	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/a-h/templ"
)

type pollHttpAdapter struct {
	poll       services.PollService
	cloudinary *cloudinary.Cloudinary
}

func NewPollHttpAdapter(poll services.PollService, cloudinary *cloudinary.Cloudinary, router *http.ServeMux) pollHttpAdapter {
	adapter := pollHttpAdapter{
		poll,
		cloudinary,
	}
	adapter.registerRoutes(router)
	return adapter
}

func (adapter pollHttpAdapter) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /polls/{id}", adapter.getPollPage)
	router.HandleFunc("POST /polls/{id}", adapter.submitVote)

	router.HandleFunc("GET /polls/admin", adapter.getCreatePollPage)
	router.HandleFunc("POST /polls/admin", adapter.createNewPoll)
}

func (adapter pollHttpAdapter) submitVote(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.ParseForm()
	selectedId, _ := strconv.Atoi(r.PostFormValue("selection"))
	pollId, _ := strconv.Atoi(r.PathValue("id"))

	submission := model.Vote{
		Selection:   selectedId,
		PollEventId: pollId,
		UserId:      session.Profile.UserId,
		Source:      "web",
	}

	err = adapter.poll.SubmitVote(submission)
	dialog := components.Dialog(nil)

	if err != nil {
		log.Println("Error: ", err)
		dialog = components.Dialog(err)
		return
	}

	log.Println(w)

	w.Header().Add("Content-Type", "text/plain")
	dialog.Render(r.Context(), w)
	return
}

func (adapter pollHttpAdapter) getPollPage(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.WithSession(r.Context())
	// convert string to number
	id, _ := strconv.Atoi(r.PathValue("id"))
	poll, err := adapter.poll.GetById(id)
	if err != nil {
		log.Fatalln(err)
	}
	templ.Handler(pages.VotePage(poll, session)).ServeHTTP(w, r)
}

func (adapter pollHttpAdapter) getCreatePollPage(w http.ResponseWriter, r *http.Request) {
	pages.CreatePage(nil).Render(r.Context(), w)
}

func (adapter pollHttpAdapter) createNewPoll(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["image"]

	for _, item := range files {
		log.Println(item.Filename)
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("<p>Success</p>"))
}
