package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zdev147/hello_world_web_with_go/pkg/config"
	"github.com/zdev147/hello_world_web_with_go/pkg/handler"
	"github.com/zdev147/hello_world_web_with_go/pkg/render"
)

const portNumber = ":8000"

// main is the entry point of go program
func main() {
	// initializing app config
	var app config.AppConfig

	cache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = cache

	// true for live server and false for debug
	app.UseCache = true

	render.SetConfig(&app)

	repo := handler.CreateNewRepo(&app)
	handler.NewHandlers(repo)

	// initializing server routes and server related stuff
	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Printf("Starting server on port %s\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
