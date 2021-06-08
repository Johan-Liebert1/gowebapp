package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// numBytesWritten, err := fmt.Fprintf(w, "This is the home page")
	renderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./html/" + templateName)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}
