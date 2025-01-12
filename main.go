package main

import (
	"net/http"
	"polling-app.com/m/database"
	"polling-app.com/m/routes"
	"polling-app.com/m/templates"
)

func main() {
	database.CheckIfDatabaseFileExists()
	database.AutoMigrate()
	server := http.NewServeMux()

	server.Handle("/public", http.FileServer(http.Dir("./public")))

	server.HandleFunc("/", templates.HandleHomeTemplate)
	server.HandleFunc("/success", templates.HandleSuccessTemplate)
	server.HandleFunc("/poll", templates.HandlePollTemplate)
	server.HandleFunc("/poll/create", routes.HandlePollCreate)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		println("Failed to start server", err)
	}

	println("Server started on port 8080")

}
