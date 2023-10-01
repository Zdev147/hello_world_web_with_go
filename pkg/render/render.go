package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplates is used to render parse and render html templates on web
func RenderTemplates(w http.ResponseWriter, tmplName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmplName, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsing template:", err)
	}
}
