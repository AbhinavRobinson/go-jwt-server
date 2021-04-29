package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

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
	// generate key
	tokenString, _ := GenerateJWT()

	// print key (debugging)
	fmt.Println(tokenString)
}
