package http

import (
	"log"
	"net/http"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &wrappedWriter{
			w,
			http.StatusOK,
		}
		next.ServeHTTP(writer, r)
		log.Println(writer.statusCode, r.Method, r.URL.Path)
	})
}

type Middleware func(http.Handler) http.Handler

func CreateStack(handlers ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(handlers) - 1; i >= 0; i-- {
			next = handlers[i](next)
		}

		return next
	}
}
