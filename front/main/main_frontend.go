package main

import (
	"log"
	"net/http"
)

func main() {
	fsStatic := http.FileServer(http.Dir("/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	fsWeb := http.FileServer(http.Dir("/web"))
	http.Handle("/web/", http.StripPrefix("/web/", fsWeb))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/web/home.html")
	})

	log.Println("Starting frontend server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
