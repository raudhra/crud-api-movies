package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:director`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["ID"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["ID"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func updateMovie(w http.ResponseWrite, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["ID"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID:"1", isbn: "13324", Title: "Movie One", Director: &Director{Firstname: "Dakota", Lastname: "Johnas"}})
	movies = append(movies, Movie{ID:"2", isbn: "14452", Title: "Movie Two", Director: &Director{Firstname: "Gopal", Lastname: "Dakait"}}

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("UPDATE")

	fmt.Printf("Starting Server At Port :8000\n")
	log.Fatal(http.ListenAndServe(8000,r))
)

}