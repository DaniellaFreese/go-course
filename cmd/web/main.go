package main

import (
	"fmt"
	"net/http"

	"github.com/DaniellaFreese/go-course/pkg/handlers"
)

const portNumber = ":8090"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
