package main

import (
	"fmt"
	"net/http"
)

func index_hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func api(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(&r)
	fmt.Fprintf(w, "This is Api Path!")
}

func main() {
	http.HandleFunc("/", index_hander)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8000", nil)
}
