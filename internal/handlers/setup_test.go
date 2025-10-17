package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/helpers"
	"github.com/vadim-vep/booking/internal/models"
	"github.com/vadim-vep/booking/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func TestMain(m *testing.M) {
	//What we put in the session
	gob.Register(models.Reservation{})

	//change this to True when in production
	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	helpers.NewHelpers(&app)

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewTestRepo(&app)
	NewHandlers(repo)
	render.NewRenderer(&app)

	os.Exit(m.Run())
}

func getRoutes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(Repo.Home))
	mux.Get("/home", http.HandlerFunc(Repo.Home))
	mux.Get("/about", http.HandlerFunc(Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(Repo.Generals))
	mux.Get("/majors-suites", http.HandlerFunc(Repo.Majors))

	mux.Get("/search-availability", http.HandlerFunc(Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(Repo.PostAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(Repo.AvailabilityJSON))

	mux.Get("/contacts", http.HandlerFunc(Repo.Contacts))

	mux.Get("/make-reservation", http.HandlerFunc(Repo.Reservation))
	mux.Post("/make-reservation", http.HandlerFunc(Repo.PostReservation))
	mux.Get("/reservation-summary", http.HandlerFunc(Repo.ReservationSummary))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Secure:   app.InProduction,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTestTemplateCache creates a map of templates for testing purposes
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	//create a map of strings myCache pointing to templates that we will need to return in the function
	myCache := map[string]*template.Template{}

	// get all of the files that start with *.page.tmpl from ./templates folder
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	//range through all files ending with *page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
