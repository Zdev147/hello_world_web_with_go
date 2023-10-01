package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	// Fprintln is used to print on web
	// _ is used to ignore a value, as go require a variable to use if it is declared
	_, err := fmt.Fprintln(w, "This home page")

	if err != nil {
		// Println is used to print on console
		fmt.Println(err)
	}
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "This about page")

	if err != nil {
		// Println is used to print on console
		fmt.Println(err)
	}
}

// main is the entry point of go program
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
