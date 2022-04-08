package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}
	
	m := make(map[string]string)

	m["message"] = "Movie not found"
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(m)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()

	// append some movies
	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "The Go Programming Language", Director: &Director{FirstName: "Robert", LastName: "Griesemer"}})
	movies = append(movies, Movie{ID: "2", Isbn: "456", Title: "The Go Programming Language", Director: &Director{FirstName: "Robert", LastName: "Griesemer"}})
	movies = append(movies, Movie{ID: "3", Isbn: "789", Title: "The Go Programming Language", Director: &Director{FirstName: "Robert", LastName: "Griesemer"}})
	movies = append(movies, Movie{ID: "4", Isbn: "101112", Title: "The Go Programming Language", Director: &Director{FirstName: "Robert", LastName: "Griesemer"}})

	r.HandleFunc("/movies", getMovies)
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
