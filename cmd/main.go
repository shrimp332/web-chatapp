package main

import (
	"log"
	"net/http"

	"htmx-site/internal/handlers"
	"htmx-site/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	handlers.Load(mux)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8000",
		Handler: stack(mux),
	}

	log.Printf("INFO serving on http://localhost%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
