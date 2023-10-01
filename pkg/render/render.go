package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const templatesPath = "./templates/"

// RenderTemplates is used to parse and render html templates on web
func RenderTemplates(w http.ResponseWriter, tmplName string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	parsedTemplate := templateCache[tmplName]

	err = parsedTemplate.Execute(w, nil)
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
