package handlers

import (
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(
		[]byte(
			`<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`,
		),
	)
}
