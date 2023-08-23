package main

import (
	"html/template"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albumList = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	http.HandleFunc("/albums", CORS(albums))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func albums(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "albums", albumList)
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, albums []album) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	for i := 0; i < len(albums); i++ {
		err = t.Execute(w, albums[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	log.Print("Successfully rendered template: " + "../templates/" + tmpl + ".html")
}
