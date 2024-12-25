package routers

import (
	"embed"
	"net/http"
)

func NewStaticHttpAdapter(router *http.ServeMux) {
	registerRoutes(router)
}

//go:embed static/assets/*
var content embed.FS

func registerRoutes(router *http.ServeMux) {
	router.Handle("GET /static/", http.FileServer(http.FS(content)))
}
