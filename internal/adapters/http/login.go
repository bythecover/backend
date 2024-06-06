package http

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"

	"bythecover/backend/internal/core/services/authenticator"
	"bythecover/backend/internal/core/services/sessions"
)

type loginHttpAdapter struct {
	authenticator *authenticator.Authenticator
}

func NewLoginHttpAdapter(authenticator *authenticator.Authenticator, router *http.ServeMux) loginHttpAdapter {
	adapter := loginHttpAdapter{
		authenticator,
	}
	adapter.RegisterRoutes(router)
	return adapter
}

func (adapter loginHttpAdapter) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /test", adapter.loginHandler)
}

func (adapter loginHttpAdapter) loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.WithSession(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	if session.State != "" {
		log.Println("session already has state")
		log.Println(session)
		w.WriteHeader(http.StatusOK)
		return
	}

	state, _ := generateRandomState()

	session.State = state
	session.Save()
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
