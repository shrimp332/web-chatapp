package handlers

import (
	"log"
	"net/http"

	ws "github.com/shrimp332/web-chatapp/internal/websocket"
)

var Hub *ws.Hub

func init() {
	log.Println("INFO Starting websocket")
	Hub = ws.NewHub()
	go Hub.Run()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWs(Hub, w, r)
}
