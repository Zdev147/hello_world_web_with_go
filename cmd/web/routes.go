package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zdev147/hello_world_web_with_go/pkg/config"
	"github.com/zdev147/hello_world_web_with_go/pkg/handler"
)

// // routes with pat
// func routes(app *config.AppConfig) http.Handler {
// 	router := pat.New()

// 	router.Get("/", http.HandlerFunc(handler.Repo.Home))
// 	router.Get("/about", http.HandlerFunc(handler.Repo.About))

// 	return router
// }

// routes with chi
func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(TestMidleware)

	router.Get("/", handler.Repo.Home)
	router.Get("/about", handler.Repo.About)

	return router
}
