package main

import (
	"fmt"
	"net/http"

	"github.com/Johan-Liebert1/gowebapp/packages/handlers"
)

var port string = ":7000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Application started on port %s", port))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}
