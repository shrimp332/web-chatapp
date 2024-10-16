package main

import (
	"log"
	"net/http"

	"github.com/shrimp332/web-chatapp/internal/handlers"
	"github.com/shrimp332/web-chatapp/internal/middleware"
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
