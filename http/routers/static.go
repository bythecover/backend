package routers

import (
	"net/http"
)

func NewStaticHttpAdapter(router *http.ServeMux) {
	registerRoutes(router)
}

func registerRoutes(router *http.ServeMux) {
	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/assets"))))
}
