package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// tc is the variable that stores the cache
var tc = map[string]*template.Template{}

// RenderTemplates is used to parse the template, store it in cache and template on web
func RenderTemplates(w http.ResponseWriter, tmplName string) {
	var template *template.Template

	_, isInMap := tc[tmplName]
	if isInMap {
		fmt.Println("reading from cache")
	} else {
		fmt.Println("creating template")
		err := CreateTemplateCache(tmplName)
		if err != nil {
			fmt.Println("error in creating cache:", err)
			return
		}
	}

	template = tc[tmplName]
	err := template.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsing template:", err)
	}
}

// CreateTemplateCache is creating the cache for templates
func CreateTemplateCache(tmplName string) error {
	templatesToParse := []string{
		"./templates/" + tmplName,
		"./templates/base.layout.tmpl",
	}

	parsedTemplate, err := template.ParseFiles(templatesToParse...)
	if err != nil {
		return err
	} else {
		tc[tmplName] = parsedTemplate
		return nil
	}
}
