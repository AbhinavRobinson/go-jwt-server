package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "secret information")
}

var mySigningKey = []byte("password")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	port := 9000
	fmt.Printf("Booting up server on port:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func main() {
	handleRequests()
}
