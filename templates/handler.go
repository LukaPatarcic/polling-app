package templates

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"polling-app.com/m/database"
	"polling-app.com/m/models"
	"strconv"
)

func handleTemplate(path string, w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	html, err := template.ParseFiles(path)
	template.Must(html, err)
	err = html.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	return
}

func HandleHomeTemplate(w http.ResponseWriter, r *http.Request) {
	handleTemplate("templates/home.go.html", w, r, nil)
}

func HandleSuccessTemplate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	data := map[string]interface{}{
		"ID": id,
	}
	handleTemplate("templates/success.go.html", w, r, data)
}

func HandlePollTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
	}
	db := database.GetInstance()
	var poll models.Poll
	result := map[string]interface{}{}
	db.Model(&poll).First(&result, id)
	handleTemplate("templates/poll.go.html", w, r, result)
}
