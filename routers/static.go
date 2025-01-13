package routers

import (
	"embed"
	"net/http"
)

//go:embed static/assets/*
var content embed.FS

func RegisterStaticRoutes(router *http.ServeMux) {
	router.Handle("GET /static/", http.FileServer(http.FS(content)))
}
