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
	app.UseCache = false

	render.SetConfig(&app)

	repo := handler.CreateNewRepo(&app)
	handler.NewHandlers(repo)

	// initializing server routes and server related stuff
	srv := http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting server on port %s\n", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
