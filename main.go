package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := fmt.Fprintf(w, "Hello world!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of bytes printed: %d", bytes)
	})

	http.ListenAndServe(":8000", nil)
}
