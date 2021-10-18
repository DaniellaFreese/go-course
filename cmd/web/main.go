package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DaniellaFreese/go-course/pkg/config"
	"github.com/DaniellaFreese/go-course/pkg/handlers"
	"github.com/DaniellaFreese/go-course/pkg/render"
)

const portNumber = ":8090"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	// app.UseCache = false
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("starting app on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
