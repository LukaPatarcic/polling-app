package routes

import (
	"net/http"
	"polling-app.com/m/database"
	"polling-app.com/m/models"
	"strconv"
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
	poll := models.Poll{
		Title:        title,
		Option1:      option1,
		Option2:      option2,
		Option3:      option3,
		Option1Votes: 0,
		Option2Votes: 0,
		Option3Votes: 0,
	}
	result := db.Create(&poll)
	println(result.Error)
	if result.Error != nil {
		http.Error(w, "Unable to create poll", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/success?id="+strconv.Itoa(int(poll.ID)), http.StatusSeeOther)
}

func HandlePollVote(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Read form values
	id := r.FormValue("id")
	answer := r.FormValue("answer")

	cookieExists, _ := r.Cookie("voted_" + id)
	if cookieExists != nil {
		http.Error(w, "You have already voted for this poll", http.StatusInternalServerError)
		return
	}

	db := database.GetInstance()

	var poll models.Poll
	db.Model(models.Poll{}).First(&poll, id)

	// Update vote counts
	if answer == poll.Option1 {
		poll.Option1Votes++
	} else if answer == poll.Option2 {
		poll.Option2Votes++
	} else if answer == poll.Option3 {
		poll.Option3Votes++
	} else {
		http.Error(w, "Invalid poll option", http.StatusBadRequest)
		return
	}

	result := db.Save(&poll)
	if result.Error != nil {
		http.Error(w, "Unable to create vote", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "voted_" + id,
		Value:    "true",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/poll/"+id, http.StatusSeeOther)
}
