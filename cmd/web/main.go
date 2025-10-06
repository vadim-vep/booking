package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vadim-vep/booking/internal/config"
	"github.com/vadim-vep/booking/internal/handlers"
	"github.com/vadim-vep/booking/internal/helpers"
	"github.com/vadim-vep/booking/internal/models"
	"github.com/vadim-vep/booking/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

type AnyWithTags[T any] struct {
	A string
	B struct {
		C T `asdasd`
	}
}

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server started on port: %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() error {
	//What we put in the session
	gob.Register(models.Reservation{})

	//change this to True when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)
	return nil
}
