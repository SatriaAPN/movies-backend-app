package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET /movies\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/:id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/:id", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/:id", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
