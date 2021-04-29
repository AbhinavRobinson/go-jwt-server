package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// add code
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	port := 9000
	fmt.Printf("Booting up server on port:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func main() {
	handleRequests()
}
