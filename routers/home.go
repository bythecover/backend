package routers

import (
	"net/http"

	"github.com/bythecover/backend/templates/pages"
)

func Home(router *http.ServeMux) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home().Render(r.Context(), w)
	})
}
