package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var messageTemplate *template.Template

func init() {
	var err error
	messageTemplate, err = template.New("Message").
		Parse(`<div hx-swap-oob="afterbegin:#chatbox"><span class="chat-message">{{.}}</span></div>`)
	if err != nil {
		log.Fatalln("FATAL", err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if input := r.FormValue("chatinput"); len(input) > 300 {
		w.Write(
			[]byte(
				`<div hx-swap-oob="afterbegin:#chatbox"><span class="chat-message error">Message was too long</span></div>
				<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`,
			),
		)
		return
	}

	newInput := `<input hx-post="/message" autofocus hx-swap="outerHTML" class="chat-input" type="text" name="chatinput" value="">`
	w.Write([]byte(newInput))

	buf := &bytes.Buffer{}
	err := messageTemplate.ExecuteTemplate(buf, "Message", r.FormValue("chatinput"))
	if err != nil {
		panic(err)
	}
	Hub.Broadcast <- buf.Bytes()
}
