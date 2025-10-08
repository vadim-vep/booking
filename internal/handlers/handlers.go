package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/driver"
	"github.com/vadim-vep/booking/internal/forms"
	"github.com/vadim-vep/booking/internal/helpers"
	"github.com/vadim-vep/booking/internal/models"
	"github.com/vadim-vep/booking/internal/render"
	"github.com/vadim-vep/booking/internal/repository"
	"github.com/vadim-vep/booking/internal/repository/dbrepo"
)

// TemplateData sent from handlers to templates

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is Repository type type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo createst a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.gohtml", &models.TemplateData{})
}

// About handles the "/about" route and renders the About page template.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform business logic

	//send the data to the template
	render.RenderTemplate(w, r, "about.page.gohtml", &models.TemplateData{})
}

// Reservation handles the "/make-reservation" route by rendering the reservation page.
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{
		Form: forms.NewForm(nil),
		Data: data,
	})
}

// PostReservation handles the POST request for the "/make-reservation" (reservation form)
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.NewForm(r.PostForm)
	//different validations for different fields
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data})
		return
	}
	//save the reservation to the session
	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
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

	w.Write([]byte(fmt.Sprint("Start date is: " + start + " End date is: " + end)))
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
		helpers.ServerError(w, err)
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

// ReservationSummary renders the "reservation-summary" page template and saves the client's remote IP address into the session.'
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Could not get reservation from session")
		m.App.Session.Put(r.Context(), "error", "Could not get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.gohtml", &models.TemplateData{
		Data: data,
	})
}
