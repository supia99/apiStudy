package main

import (
	"fmt"
	"log"
	"net/http"
	scrape "supia99/scrapeero/cmd"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func requestURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Success!")
	fmt.Println("Endpoint Hit: requestURL " + r.Method)

	url := r.URL.Query().Get("url")
	fmt.Printf("url:%s\n", url)
	scrape.Accept(url)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/url", requestURL)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
