package routes

import (
	"net/http"
	"polling-app.com/m/database"
	"polling-app.com/m/models"
)

func HandlePollCreate(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Read form values
	title := r.FormValue("title")
	option1 := r.FormValue("option1")
	option2 := r.FormValue("option2")
	option3 := r.FormValue("option3")

	db := database.GetInstance()
	poll := models.Poll{Title: title, Option1: option1, Option2: option2, Option3: option3}
	result := db.Create(&poll)
	if result.Error != nil {
		http.Error(w, "Unable to create poll", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/success", http.StatusSeeOther)
}
