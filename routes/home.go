package routes

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("public/index.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fs.ServeHTTP(w, r)
}
