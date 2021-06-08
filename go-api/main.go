package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id		string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
}

var Articles []Article

// GET all articles
func returnALlArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: All Articles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "key: "+ key)

	for _,article:= range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to page")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/articles", returnALlArticles)
	fmt.Println("listening on port : 10000")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}
