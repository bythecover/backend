package http

import (
	"bythecover/backend/internal/core/templates/pages"
	"net/http"

	"github.com/a-h/templ"
)

type TestHttpHandler struct{}

func (adapter TestHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(pages.NewPage()).ServeHTTP(w, r)
	})
}
