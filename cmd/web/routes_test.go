package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/vadim-vep/booking/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("type is not *chi.Mux %t", v)
	}
}
