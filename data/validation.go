package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// ValidationError wraps the validators FieldError
type ValidationError struct {
	validator.FieldError
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Error returns a custom error message for Validation errors
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validate checks an interface for the validations specified on each field
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs, ok := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 || !ok {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU returns true if the sku tag is valid on the specified field
func validateSKU(flvl validator.FieldLevel) bool {
	// SKU must be of the format 000-000-000
	regex := regexp.MustCompile(`[0-9]+-[0-9]+-[0-9]+`)
	matches := regex.FindAllString(flvl.Field().String(), -1)

	return len(matches) == 1
}
