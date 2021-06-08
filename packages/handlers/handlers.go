package handlers

import (
	"net/http"

	"github.com/Johan-Liebert1/gowebapp/packages/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// numBytesWritten, err := fmt.Fprintf(w, "This is the home page")
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
