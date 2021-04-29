package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// create new token
	validToken, err := GenerateJWT()

	// if err, throw err
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	// if err, throw err
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	// if err, throw err
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	// respond with token
	fmt.Fprintf(w, "%s", string(body))
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	port := 9001

	fmt.Printf("Booting up client server on port:%d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

// password for key generation (use env variables for this in production)
// var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("password")

// Generates and returns Signed JWT
func GenerateJWT() (string, error) {
	// Declare signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// create claim
	claims := token.Claims.(jwt.MapClaims)

	// give claim config
	claims["authorized"] = true
	claims["user"] = "Abhinav Robinson"
	// valid of 30 mins only
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// generate key
	tokenString, err := token.SignedString(mySigningKey)

	// if err, throw err
	if err != nil {
		fmt.Println("Error : ", err.Error())
		return "", err
	}

	// return key
	return tokenString, nil
}

// Main Program
func main() {
	handleRequests()
}
