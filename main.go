package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Movie struct{
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json::"director"' //this is a pointer
}

type Director struct{
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}

var movie []Movie

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	

}