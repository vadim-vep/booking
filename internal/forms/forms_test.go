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

	if form.Has("a") {
		t.Error("form shows field, when it shouldn't")
	}

	postedData := url.Values{}
	postedData.Add("a", "1")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	form = NewForm(postedData)

	if !form.Has("a") {
		t.Error("form does not have field, when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	if form.MinLength("non_existent_field", 3) {
		t.Error("form shows min length for field that doesn't exist")
	}

	postedData := url.Values{}
	postedData.Add("field_key", "123")

	form = NewForm(postedData)
	//field key is 3 characters long, value is "123" =>len 3
	if !form.MinLength("field_key", 3) {
		t.Error("form shows invalid when min length is met")
	}
	isError := form.Errors.Get("field_key")
	if isError != "" {
		t.Error("should not have an error, but there is one")
	}

	if form.MinLength("field_key", 4) {
		t.Error("form shows valid when min length is met")
	}

	isError = form.Errors.Get("field_key")
	if isError == "" {
		t.Error("should get error, but there is none")
	}

}

func TestForm_isEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	form.IsEmail("non_existant_field")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedData := url.Values{}
	postedData.Add("email", "me@here.com")
	form = NewForm(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("form shows invalid email for valid email")
	}

	postedData2 := url.Values{}
	postedData2.Add("email", "not an email")
	form = NewForm(postedData2)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("form shows email is valid for invalid address")
	}
}
