package main

import (
	    "fmt"
		//"encoding/json"
		"log"
		"net/http"
		//"math/rand"
		//"strconv"
		"github.com/gorilla/mux"
	"encoding/json"
)

//Book Struct (Model classları) go dilinde modeller değil structlar denir
//propertyler ve metodlar içinde bulundurur
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

// Init books var as a slice(collection) Book struct
var books []Book


// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// Get a Single Book
func getBook(w http.ResponseWriter, r *http.Request){

}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request){

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request){

}

// DeleteBook
func deleteBook(w http.ResponseWriter, r *http.Request){

}


func main() {
	fmt.Println("Application starting ...")

	/*initialize router, := işareti soldaki nesneye tipi bağlamak için kullanılır
	bu durumda mux.NewRouter ın tipi route nesnesine bağlanmış olur*/
	router :=mux.NewRouter()

	//Mock Data - @todo - implement DB
	books = append(books, Book{ID:"1", Isbn:"1234" , Title:"LOTR - Fellowship of the Ring", Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})
	books = append(books, Book{ID:"2", Isbn:"54634", Title:"LOTR - The Two Towers"        , Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})
	books = append(books, Book{ID:"3", Isbn:"6473" , Title:"LOTR - Return of The King"    , Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})


	//Route handlers, api endpoinlerini belirteceğiz
	router.HandleFunc("/api/books"     , getBooks  ).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook   ).Methods("GET")
	router.HandleFunc("/api/books"     , createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))


}