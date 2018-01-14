package main

import (
	    "fmt"
		"log"
		"net/http"
		"strconv"
		"github.com/gorilla/mux"
		"encoding/json"
)

// Init books var as a slice(collection) Book struct
var books []Book

//Book Struct (Model class)
type Book struct {
	ID int `json:"id"`
	Isbn int `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`

}

//Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}


func getBooks(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}


func getBook(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params:=mux.Vars(request) //Get parameters from request

	// Loop through books and find with id
	for _, item := range books{
		if strconv.Itoa(item.ID) == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}


func createBook(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)

	book.ID = books[len(books)-1].ID +1 // get last id and add 1
	book.Isbn = books[len(books)-1].Isbn+1 // get last isbn no and add 1

	books = append(books, book)

	json.NewEncoder(w).Encode(book)


}


func updateBook(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(request)


	for index, item:= range books {

		if strconv.Itoa(item.ID) == params["id"]{
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = books[index].ID // do not change id
			book.Isbn = books[index].Isbn // do not change isbn
			books = append(books[:index], books[index+1:]...)//delete the book
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})



}


func deleteBook(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(request)

	for index, item:= range books {
		if strconv.Itoa(item.ID) == params["id"]{
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})


}


func main() {
	fmt.Println("Application starting ...")

	/*initialize router, := type inference*/
	router :=mux.NewRouter()

	//Mock Data - @todo - implement DB
	books = append(books, Book{ID:1, Isbn:1 , Title:"LOTR - Fellowship of the Ring", Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})
	books = append(books, Book{ID:2, Isbn:2, Title:"LOTR - The Two Towers"        , Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})
	books = append(books, Book{ID:3, Isbn:3 , Title:"LOTR - Return of The King"    , Author: &Author{Firstname: "J.R.R", Lastname:"Tolkien" }})


	//Route handlers, api endpoinlerini belirteceğiz
	router.HandleFunc("/api/books"     , getBooks  ).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook   ).Methods("GET")
	router.HandleFunc("/api/books"     , createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000", router))




}