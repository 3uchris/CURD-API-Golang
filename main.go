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

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //encode movies into json and send
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //the id we pass from my postman which will go as a param 
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...) //delete an item using append
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().set("Content-Type", "application/json")
	params := mux.Vars(r)
	for , item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)//blank identifier
	//get a new movie id
	movie.ID = strconv.Itoa(rand.Intn(100000000)) //format it into string
	movies = append(movies, movie)
	json.NewEncode(w).Encode(movie)
}

func main(){
	r := mux.NewRouter()  //NewRouter is a library inside the Gorilla/mux

	//instances
	movies = append(movies, Movie{ID:"1", Isbn:"Movie One", Director:&Director{Firstname:"John", Lastname:"Doe"}})
	movies = append(movies, Movie{ID:"1", Isbn:"Movie Two", Director:&Director{Firstname:"Steve", Lastname:"Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies/", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Method("DELETE")
	
	fmt.printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))


}