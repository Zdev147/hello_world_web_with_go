package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zdev147/hello_world_web_with_go/pkg/config"
	"github.com/zdev147/hello_world_web_with_go/pkg/models"
)

const templatesPath = "./templates/"

var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

// RenderTemplates is used to parse and render html templates on web
func RenderTemplates(w http.ResponseWriter, tmplName string, tmplData *models.TemplateData) {
	var templateCache map[string]*template.Template
	var err error

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	parsedTemplate := templateCache[tmplName]

	err = parsedTemplate.Execute(w, tmplData)
	if err != nil {
		fmt.Println("error in executing template:", err)
	}
}

// CreateTemplateCache is used to create template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprint(templatesPath, "*.page.tmpl"))
	if err != nil {
		return templateCache, err
	}

	// _ here is index
	for _, page := range pages {
		name := filepath.Base(page)

		parsedTemplatesSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		layoutMatches, err := filepath.Glob(fmt.Sprint(templatesPath, "*.layout.tmpl"))
		if err != nil {
			return templateCache, err
		}

		if len(layoutMatches) > 0 {
			parsedTemplatesSet, err = parsedTemplatesSet.ParseGlob(fmt.Sprint(templatesPath, "*.layout.tmpl"))
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = parsedTemplatesSet
	}

	return templateCache, nil
}
