package handlers

import (
	"myapp/config"
	"myapp/models"
	"myapp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello from handler"
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
