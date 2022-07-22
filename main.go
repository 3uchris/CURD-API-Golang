package main

import( 
	"fmt"            //name of a package
	"log"
	"encoding/json"  //path of the package
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
	r := mux.NewRouter()  //NewRouter is a library inside the Gorilla/mux

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies/", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")
	
	fmt.printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
	

}