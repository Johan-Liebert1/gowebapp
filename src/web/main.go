package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Johan-Liebert1/gowebapp/packages/config"
	"github.com/Johan-Liebert1/gowebapp/packages/handlers"
	"github.com/Johan-Liebert1/gowebapp/packages/render"
)

var port string = ":7000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot CreateTemplateCache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Application started on port %s\n", port)

	err = http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}
