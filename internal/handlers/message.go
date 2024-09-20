package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var (
	subs map[string]string
	reg  *regexp.Regexp
)

func init() {
	// regex init
	subs = map[string]string{
		"&":  "&amp;",
		"<":  "&lt;",
		">":  "&gt;",
		"\"": "&quot;",
		"'":  "&#x27;",
		"/":  "&#x2F;",
	}
	pattern := "/|&|<|>|\"|'"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalln("FATAL", err)
	}
	reg = regex
}

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

	message := santizeInput(r.FormValue("chatinput"))
	resp := fmt.Sprintf(
		`<div hx-swap-oob="afterbegin:#chatbox"><span class="chat-message">%s</span></div>`,
		message,
	)

	newInput := `<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`
	w.Write([]byte(newInput))
	Hub.Broadcast <- []byte(resp)
}

func santizeInput(s string) string {
	return reg.ReplaceAllStringFunc(s, func(match string) string {
		return subs[match]
	})
}
