package govalidator

import (
	"testing"
)

var ve = &ValidationError{}

func TestValidationErrorAdd(t *testing.T) {
	ve.Add("username", "The username field is required")
	ve.Add("username", "The username field must be between 3-40 chars")
	ve.Add("password", "The password field must between 8-15 chars")
	if len(*ve) != 2 {
		t.Error("failed to add error")
	}
}

func TestValidationErrorGet(t *testing.T) {
	if ve.Get("username") == "" {
		t.Error("failed to get value from ValidationError")
	}

	if ve.Get("invalidKey") != "" {
		t.Error("get must return empty string from ValidationError")
	}
}
