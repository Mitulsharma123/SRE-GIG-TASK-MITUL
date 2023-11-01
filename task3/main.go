package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Movie struct{
	Name string `json:"name"`
	Genre string `json:"genre"`
	Director string `json:"director"`
	Rating float64 `json:"rating"`
}


func main(){
	fmt.Println("Initiate Listen and Serve....")
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/movie", GetMovie)
	
	log.Fatal(http.ListenAndServe(":8085", nil))
	
}

func Hello(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style = 'color:blue'> Hello Gophers !</h1>"))
}

func GetMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	movie := Movie{
		Name: "Mission Impossible Dead Reckoning",
		Genre: "Action",
		Director: "Christopher McQuarrie",
		Rating: 7.8,
	}

	json.NewEncoder(w).Encode(movie)
}