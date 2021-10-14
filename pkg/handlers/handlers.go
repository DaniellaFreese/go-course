package handlers

import (
	"net/http"

	"github.com/DaniellaFreese/go-course/pkg/render"
)

//Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.html")
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.go.html")
}
