package main

import (
	"fmt"
	"log"
	"net/http"

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

func main(){
	r := mux.NewRouter()
	
	movies = append(movies, Movie{ID:"1", isbn: "13324", Title: "Movie One", Director: &Director{Firstname: "Dakota", Lastname: "Johnas"}})
	movies = append(movies, Movie{ID:"2", isbn: "14452", Title: "Movie Two", Director: &Director{Firstname: "Gopal", Lastname: "Dakait"}}

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("UPDATE")

	fmt.Printf("Starting Server At Port :8000\n")
	log.Fatal(http.ListenAndServe(8000,r))
)

}