package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bythecover/backend/http/middleware"
	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/services"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/components"
	"github.com/bythecover/backend/templates/pages"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

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

	isAuthorizedAsAuthorOrUser := middleware.CreateAuthorizedHandler([]string{"author", "user"})
	router.Handle("GET /a/{authorName}/{bookid}", isAuthorizedAsAuthorOrUser(http.HandlerFunc(adapter.getPollPage)))
	router.Handle("POST /a/{authorName}/{bookid}", isAuthorizedAsAuthorOrUser(http.HandlerFunc(adapter.submitVote)))

	isAuthorizedAsAuthor := middleware.CreateAuthorizedHandler([]string{"author"})
	router.Handle("GET /a/{authorName}/create", isAuthorizedAsAuthor(http.HandlerFunc(adapter.getCreatePollPage)))
	router.Handle("POST /a/{authorName}", isAuthorizedAsAuthor(http.HandlerFunc(adapter.createNewPoll)))
}

func (adapter pollHttpAdapter) submitVote(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.ParseForm()
	selectedId, _ := strconv.Atoi(r.PostFormValue("selection"))
	pollId, _ := strconv.Atoi(r.PathValue("bookid"))

	submission := model.Vote{
		Selection:   selectedId,
		PollEventId: pollId,
		UserId:      session.Profile.UserId,
		Source:      "web",
	}

	err = adapter.poll.SubmitVote(submission)
	dialog := components.Dialog(nil)

	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusForbidden)
		dialog = components.Dialog(err)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	dialog.Render(r.Context(), w)
	return
}

func (adapter pollHttpAdapter) getPollPage(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())
	authorId := r.PathValue("authorName")

	if session.Profile.UserId == authorId {
		adapter.getResultPage(w, r)
		return
	}

	// convert string to number
	bookId, err := strconv.Atoi(r.PathValue("bookid"))

	log.Println(authorId)

	poll, err := adapter.poll.GetByIdAndAuthorName(bookId, authorId)

	if err != nil {
		logger.Error.Fatalln(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	templ.Handler(pages.VotePage(poll, session)).ServeHTTP(w, r)
}

func (adapter pollHttpAdapter) getCreatePollPage(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.WithSession(r.Context())
	pages.CreatePage(session).Render(r.Context(), w)
}

func (adapter pollHttpAdapter) createNewPoll(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		logger.Error.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(1 << 20)

	if err != nil {
		logger.Error.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["image"]
	values := r.MultipartForm.Value
	options := []model.Option{}

	for i, item := range files {
		res, _ := adapter.cloudinary.Upload.Upload(r.Context(), item, uploader.UploadParams{})
		options = append(options, model.Option{
			Image: res.PublicID,
			Name:  values["name"][i],
		})
	}

	poll := model.Poll{
		CreatedBy: session.Profile.UserId,
		Title:     values["title"][0],
		Options:   options,
	}

	if err = adapter.poll.CreatePoll(poll); err != nil {
		logger.Error.Println(err)
	}

	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("<p>Success</p>"))
}

func (adapter pollHttpAdapter) getResultPage(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	pollId, _ := strconv.Atoi(r.PathValue("bookid"))
	results := adapter.poll.GetResults(pollId)

	pages.Results(session, results).Render(r.Context(), w)
}
