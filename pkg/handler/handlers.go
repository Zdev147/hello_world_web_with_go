package handler

import (
	"net/http"

	"github.com/zdev147/hello_world_web_with_go/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplates(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl")
}
