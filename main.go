package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Could not parse template: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	})
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "blog", nil)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
