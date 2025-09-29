package forms

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Form creates a custom form struct and embeds an url.Values object
type Form struct {
	url.Values
	Errors errors
}

// NewForm initializes a new "form" struct
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a field exists in the form and is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required checks if a field exists in the form and is not empty
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(field, "This field is required")
		}
	}
}

// Valid checks if the form has any errors (true if no errors)
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// MinLength checks if a field is at least a certain length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.AddError(field, "This field must be at least "+strconv.Itoa(length)+" characters long")
		return false
	}
	return true
}
