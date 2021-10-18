package handlers

import (
	"net/http"

	"github.com/DaniellaFreese/go-course/pkg/config"
	"github.com/DaniellaFreese/go-course/pkg/models"
	"github.com/DaniellaFreese/go-course/pkg/render"
)

//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.go.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"

	render.RenderTemplate(w, "about.page.go.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
