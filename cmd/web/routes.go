package main

import (
	"net/http"

	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/handlers"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurf)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(handlers.Repo.Generals))
	mux.Get("/majors-suites", http.HandlerFunc(handlers.Repo.Majors))

	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))

	mux.Get("/contacts", http.HandlerFunc(handlers.Repo.Contacts))

	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Post("/make-reservation", http.HandlerFunc(handlers.Repo.PostReservation))
	mux.Get("/reservation-summary", http.HandlerFunc(handlers.Repo.ReservationSummary))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux

}

type test struct {
	App              *config.AppConfig
	a, b, c, Session *scs.SessionManager
}
