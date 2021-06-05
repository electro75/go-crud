package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to page")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
