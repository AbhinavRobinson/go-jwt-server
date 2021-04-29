package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// hidden information
	fmt.Fprintf(w, "secret information")
}

// key, same as client
var mySigningKey = []byte("password")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if Token in header
		if r.Header["Token"] != nil {
			// parse token
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, "Error : %s", err.Error())
			}

			// if valid, return endpoint
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			// if no token, reject request
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	// create route
	http.Handle("/", isAuthorized(homePage))
	port := 9000
	fmt.Printf("Booting up server on port:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func main() {
	handleRequests()
}
