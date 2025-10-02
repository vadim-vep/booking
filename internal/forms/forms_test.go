package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("form is not valid")
	}
}

func TestForm_required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}

	postedData.Add("a", "1")
	postedData.Add("b", "2")
	postedData.Add("c", "3")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows form doesn't have required fields when it does")
	}
}

func TestFormHas_field(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	if form.Has("a", r) {
		t.Error("form shows field, when it shouldn't")
	}

	postedData := url.Values{}
	postedData.Add("a", "1")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	form = NewForm(postedData)

	if !form.Has("a", r) {
		t.Error("form does not have field, when it should")
	}
}
