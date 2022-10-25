package main

import (
	"encoding/json"
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

var Directors = []Director{
	{FirstName: "Matahari", LastName: "Ramadhan"},
	{FirstName: "Iqbal", LastName: "Adudu"},
}

var Movies = []Movie{
	{ID: "1", Isbn: "43210", Title: "Movie 1", Director: Directors[0]},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/{id}", getMovie).Methods("GET")
	r.HandleFunc("/{id}", createMovie).Methods("POST")
	r.HandleFunc("/{id}", updateMovie).Methods("PATCH")
	r.HandleFunc("/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
