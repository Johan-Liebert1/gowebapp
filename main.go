package main

import (
	"fmt"
	"net/http"
)

var port string = ":7000"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Application started on port %s", port))

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}
