package handler

import (
	"net/http"

	"github.com/zdev147/hello_world_web_with_go/pkg/config"
	"github.com/zdev147/hello_world_web_with_go/pkg/models"
	"github.com/zdev147/hello_world_web_with_go/pkg/render"
)

// Home is the home page handler
// using repositoy as reciever
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
// using repositoy as reciever
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	data := models.TemplateData{
		StringData: map[string]string{
			"title": "About",
			"name":  "Ahmed Imran",
		},
	}

	render.RenderTemplates(w, "about.page.tmpl", &data)
}

// Repository pattern for app config

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func CreateNewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
