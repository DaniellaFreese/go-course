package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8090"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
