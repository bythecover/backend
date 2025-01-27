package routers

import (
	"net/http"
	"strconv"

	"github.com/bythecover/backend/logger"
	"github.com/bythecover/backend/model"
	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/components"
	"github.com/bythecover/backend/templates/pages"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/a-h/templ"
)

type pollHttpAdapter struct {
	pollRepo   model.PollRepo
	voteRepo   model.VoteRepo
	cloudinary *cloudinary.Cloudinary
}

func RegisterPollRoutes(router *http.ServeMux, pollRepo model.PollRepo, voteRepo model.VoteRepo, cloudinary *cloudinary.Cloudinary) {
	adapter := pollHttpAdapter{pollRepo, voteRepo, cloudinary}

	isAuthorizedAsAuthorOrUser := CreateAuthorizedHandler([]string{"author", "user"})
	router.Handle("GET /a/{authorName}/{bookid}", isAuthorizedAsAuthorOrUser(http.HandlerFunc(adapter.getPollPage)))
	router.Handle("POST /polls/{id}", isAuthorizedAsAuthorOrUser(http.HandlerFunc(adapter.submitVote)))

	isAuthorizedAsAuthor := CreateAuthorizedHandler([]string{"author"})
	router.Handle("GET /create", isAuthorizedAsAuthor(http.HandlerFunc(adapter.getCreatePollPage)))
	router.Handle("POST /polls/admin", isAuthorizedAsAuthor(http.HandlerFunc(adapter.createNewPoll)))
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

	err = adapter.voteRepo.SubmitVote(submission)
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

	// convert string to number
	authorId := r.PathValue("authorName")
	bookId, err := strconv.Atoi(r.PathValue("bookid"))

	poll, err := adapter.pollRepo.GetById(bookId, authorId)

	if session.Profile.UserId == poll.CreatedBy {
		// TODO: Show a live results page to the author that created the poll
		adapter.getResultPage(w, r)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err)
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
	err = adapter.pollRepo.CreatePoll(poll)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
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

	pollId, err := strconv.Atoi(r.PathValue("bookid"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error.Println(err.Error())
		return
	}

	results := adapter.voteRepo.GetResults(pollId)
	pages.Results(session, results).Render(r.Context(), w)
}
