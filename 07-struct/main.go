package main

import (
	"encoding/json"
	"fmt"
)

// Book object
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author object
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// var books []Book

func main() {

	author := Author{
		Firstname: "rayman",
		Lastname:  "ruan",
	}

	book := Book{
		ID:     "1001",
		Isbn:   "4040214",
		Title:  "Taitanic",
		Author: &author,
	}

	// JSON 序列化
	jsonBytes, _ := json.Marshal(book)
	str := string(jsonBytes)
	fmt.Println(str)

	// JSON 反序列化
	var bookObject Book
	err := json.Unmarshal([]byte(str), &bookObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bookObject)
}
