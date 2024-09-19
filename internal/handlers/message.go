package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if input := r.FormValue("chatinput"); len(input) > 300 {
		w.Write(
			[]byte(
				`<div hx-swap-oob="afterbegin:#chatbox"><span style="color:darkred" class="chat-message">Message was too long</span></div>
				<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`,
			),
		)
		return
	}
	message, err := santizeInput(r.FormValue("chatinput"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf(
		`<div hx-swap-oob="afterbegin:#chatbox"><span class="chat-message">%s</span></div>`,
		message,
	)

	newInput := `<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`
	w.Write([]byte(newInput))
	Hub.Broadcast <- []byte(resp)
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
