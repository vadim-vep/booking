package forms

type errors map[string][]string

// AddError adds an error to the errors map
func (e errors) AddError(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
