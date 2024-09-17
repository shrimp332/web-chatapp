package handlers

import (
	"fmt"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("chatinput")
	resp := fmt.Sprintf(`
		<div class="chat-box" id="chatbox" hx-swap-oob="afterbegin:#chatbox"><span class="chat-message">%s</span></div>
		<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">
	`, message)

	w.Write([]byte(resp))
}
