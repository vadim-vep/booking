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

// renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About handles the "/about" route and renders the About page template with dynamic data from the session.
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

// Reservation handles the "/make-reservation" route by saving the client's remote IP to the session and rendering the reservation page.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

// Generals renders the "generals" page template and saves the client's remote IP address into the session.
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{})
}

// Majors renders the "majors" page template and saves the client's remote IP address into the session.'
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "majors.page.gohtml", &models.TemplateData{})
}

// Availability renders the "search-availability" page template and saves the client's remote IP address into the session.'
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "search-availability.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Contacts(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "contacts.page.gohtml", &models.TemplateData{})
}
