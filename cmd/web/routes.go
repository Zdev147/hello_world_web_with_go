package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/zdev147/hello_world_web_with_go/pkg/config"
	"github.com/zdev147/hello_world_web_with_go/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux
}