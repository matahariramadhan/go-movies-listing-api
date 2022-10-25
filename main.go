package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var Movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/:id", getMovie).Methods("GET")
	r.HandleFunc("/:id", createMovie).Methods("POST")
	r.HandleFunc("/:id", updateMovie).Methods("PATCH")
	r.HandleFunc("/:id", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
