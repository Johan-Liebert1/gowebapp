package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Johan-Liebert1/gowebapp/packages/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// sets the config for templates package
func NewTemplates(a *config.AppConfig) {
	app = a
}

/*
Renders templates using HTML templates
*/
func RenderTemplate(w http.ResponseWriter, templateName string) {
	// get the template cache from app config

	tc := app.TemplateCache

	t, ok := tc[templateName]

	if !ok {
		log.Fatal("cannnot find template")
	}

	buffer := new(bytes.Buffer)

	_ = t.Execute(buffer, nil)

	_, err := buffer.WriteTo(w)

	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./html/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./html/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./html/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
