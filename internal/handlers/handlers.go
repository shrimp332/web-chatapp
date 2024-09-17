package handlers

import (
	"net/http"
	"strings"
)

func Load(mux *http.ServeMux) {
	fs := noDirList(http.FileServer(http.Dir("static")))
	mux.Handle("GET /", http.StripPrefix("/", fs))
	mux.HandleFunc("POST /message", messageHandler)
}

func noDirList(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
