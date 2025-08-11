package handlers

import (
	"booking/models"
	"booking/pkg/config"
	"booking/pkg/render"
	"net/http"
)

// TemplateData sent from handlers to templates

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is Repository type type
type Repository struct {
	App *config.AppConfig
}

// NewRepo createst a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
