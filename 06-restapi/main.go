package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct (Modal)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

var count int

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count++
	// fmt.Println(count)
	t := time.Now()
	fmt.Printf("%s: %d\n", t.Format("2006-01-02 15:04:05"), count)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(params)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		Isbn:  "4040214",
		Title: "Book 1",
		Author: &Author{
			Firstname: "rayman",
			Lastname:  "ruan"}})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "4040211",
		Title: "Book 2",
		Author: &Author{
			Firstname: "ray",
			Lastname:  "ruan"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
