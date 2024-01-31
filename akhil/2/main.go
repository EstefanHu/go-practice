package main

import (
  "fmt"
  "log"
  "encoding/json"
  "math/random"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
  Firstname string `json:"firstname"`
  LastName string `json:"lastname"`
}

var movies []Movie

func main() {
  r := mux.NewRouter()

  movies = append(movies, Movie{ID: "1", Isbn: "43877", Title: "Movie One", Director: &Director{Firstname: "Gary", Lastname: "Host"} })
  movies = append(movies, Movie{ID: "2", Isbn: "43878", Title: "Movie Two", Director: &Director{Firstname: "Gladys", Lastname: "Host"} })
  r.handleFunc("/movies", getMovies).Methods("GET")
  r.handleFunc("/movies/{id}", getMovie).Methods("GET")
  r.handleFunc("/movies", createMovie).Methods("POST")
  r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
  r.handleFunc("/movies/{id}", deletMovie).Methods("DELETE")

  fmt.Printf("Starting server at port 8000\n")
  log.Fatal(http.ListenAndServer(":8000", r)
}
