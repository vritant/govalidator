package govalidator

import (
	"errors"
)

var (
	errStringToInt          = errors.New("govalidator: unable to parse string to integer")
	errStringToFloat        = errors.New("govalidator: unable to parse string to float")
	errValidateArgsMismatch = errors.New("govalidator: provide at least *http.Request and rules for Validate method")
	errInvalidArgument      = errors.New("govalidator: invalid number of argument")
	errRequirePtr           = errors.New("govalidator: provide pointer to the data structure")
)

// ValidationError respresents validation error
type ValidationError map[string][]string

// Add add an error to validation error
func (v ValidationError) Add(key, value string) {
	v[key] = append(v[key], value)
}

// Get the error by key
func (v ValidationError) Get(key string) string {
	vs := v[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// HasError check if any error exist
func (v ValidationError) HasError() bool {
	return len(v) != 0
}
