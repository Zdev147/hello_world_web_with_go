package main

import (
	"fmt"
	"net/http"
)

func TestMidleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Midleware called")
		next.ServeHTTP(w, r)
	})
}