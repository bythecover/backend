package http

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"bythecover/backend/internal/core/services"
)

type loginHttpAdapter struct {
	authenticator *services.Authenticator
}

func NewLoginHttpAdapter(authenticator *services.Authenticator) loginHttpAdapter {
	return loginHttpAdapter{
		authenticator,
	}
}

func (adapter loginHttpAdapter) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /test", adapter.loginHandler)
}

// Handler for our login.
func (adapter loginHttpAdapter) loginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := generateRandomState()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	session := r.Context().Value("session").(*Session)

	if session == nil {
		log.Fatalln("session is nil")
	}

	if session.State != "" {
		log.Println("session already has state")
		log.Println(session.State)
		w.WriteHeader(http.StatusOK)
		return
	}

	session.State = state
	log.Println(session.State)

	http.Redirect(w, r, adapter.authenticator.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
