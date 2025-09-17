package render

import (
	"booking/models"
	"booking/pkg/config"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for a template function
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Renders tempaltes using HTML
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		//get the template cache from the App config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	//get the requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//create a map of strings myCache pointing to templates that we will need to return in the function
	myCache := map[string]*template.Template{}

	// get all of the files that start with *.page.tmpl from ./templates folder
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return myCache, err
	}
	//range through all files ending with *page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil

}
