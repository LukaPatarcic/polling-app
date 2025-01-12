package templates

import (
	"html/template"
	"net/http"
)

func handleTemplate(path string, w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles(path)
	template.Must(html, err)
	err = html.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	return
}

func HandleHomeTemplate(w http.ResponseWriter, r *http.Request) {
	handleTemplate("templates/home.go.html", w, r)
}

func HandleSuccessTemplate(w http.ResponseWriter, r *http.Request) {
	handleTemplate("templates/success.go.html", w, r)
}

func HandlePollTemplate(w http.ResponseWriter, r *http.Request) {
	handleTemplate("templates/poll.go.html", w, r)
}
