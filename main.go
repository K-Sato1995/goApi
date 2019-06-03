package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRecuests() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRecuests()
}
