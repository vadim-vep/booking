package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/models"
	"github.com/vadim-vep/booking/internal/render"
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
	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// About handles the "/about" route and renders the About page template with dynamic data from the session.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation handles the "/make-reservation" route by saving the client's remote IP to the session and rendering the reservation page.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{})
}

// Generals renders the "generals" page template and saves the client's remote IP address into the session.
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, r, "generals.page.gohtml", &models.TemplateData{})
}

// Majors renders the "majors" page template and saves the client's remote IP address into the session.'
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, r, "majors.page.gohtml", &models.TemplateData{})
}

// Availability renders the "search-availability" page template and saves the client's remote IP address into the session.'
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, r, "search-availability.page.gohtml", &models.TemplateData{})
}

// PostAvailability renders the "search-availability" page template and saves the client's remote IP address into the session.'
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is: " + start + " End date is: " + end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles the request for availability and sends JSON request
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	output, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// Contacts renders the "contacts" page template and saves the client's remote IP address into the session.'
func (m *Repository) Contacts(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, r, "contacts.page.gohtml", &models.TemplateData{})
}
