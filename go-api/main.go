package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to page")
	fmt.Println("Endpoint hit: homepage")
}

func handleRequests() {
	http.handleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	handleRequests()
}