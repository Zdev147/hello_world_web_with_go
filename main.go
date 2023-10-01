package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8000"

// main is the entry point of go program
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting server on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about.page.tmpl")
}

// renderTemplates is used to render parse and render html templates on web
func renderTemplates(w http.ResponseWriter, tmplName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmplName)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsing template:", err)
	}
}
