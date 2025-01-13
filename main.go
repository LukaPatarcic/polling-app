package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"polling-app.com/m/database"
	"polling-app.com/m/routes"
	"polling-app.com/m/templates"
)

func main() {
	database.CheckIfDatabaseFileExists()
	database.AutoMigrate()
	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	router.HandleFunc("/", templates.HandleHomeTemplate)
	router.HandleFunc("/success", templates.HandleSuccessTemplate)
	router.HandleFunc("/poll/create", routes.HandlePollCreate)
	router.HandleFunc("/poll/vote/create", routes.HandlePollVote)
	router.HandleFunc("/poll/{id}", templates.HandlePollTemplate)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		println("Failed to start server", err)
	}

	println("Server started on port 8080")

}
