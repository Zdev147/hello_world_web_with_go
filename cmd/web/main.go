package main

import (
	"fmt"
	"net/http"

	"github.com/zdev147/hello_world_web_with_go/pkg/handler"
)

const portNumber = ":8000"

// main is the entry point of go program
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Starting server on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
