package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message, err := santizeInput(r.FormValue("chatinput"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf(`
		<div class="chat-box" id="chatbox" hx-swap-oob="afterbegin:#chatbox"><span class="chat-message">%s</span></div>
		<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">
	`, message)

	w.Write([]byte(resp))
}

func santizeInput(s string) (string, error) {
	subs := map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&#x27;",
		"/":  "&#x2F;",
	}
	pattern := "/|&|<|>|\"|'"

	reg, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllStringFunc(s, func(match string) string {
		return subs[match]
	}), nil
}
