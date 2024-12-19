package routers

import (
	"net/http"

	"github.com/bythecover/backend/sessions"
	"github.com/bythecover/backend/templates/pages"
)

func Home(router *http.ServeMux) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.WithSession(r.Context())
		pages.Home(session).Render(r.Context(), w)
	})
}
