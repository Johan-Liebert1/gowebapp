package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

/*
Renders templates using HTML templates
*/
func RenderTemplate(w http.ResponseWriter, templateName string) {
	tc, err := CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[templateName]

	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	_ = t.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)

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
