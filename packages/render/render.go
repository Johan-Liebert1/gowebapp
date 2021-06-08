package render

import (
	"fmt"
	"net/http"
	"text/template"
)

/*
Renders templates using HTML templates
*/
func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + templateName)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}
