// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product Product defines the structure of a sample product
//
// swagger:model Product
type Product struct {

	// the description for this poduct
	// Max Length: 1000
	Description string `json:"desc,omitempty"`

	// the id for this product
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// the name for this product
	// Required: true
	// Max Length: 255
	Name *string `json:"name"`

	// the price for the product
	// Required: true
	// Minimum: 1
	Price *int64 `json:"price"`

	// the SKU for the product
	// Required: true
	// Pattern: [0-9]+-[0-9]+-[0-9]+
	SKU *string `json:"sku"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("desc", "body", m.Description, 1000); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", m.ID, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *Product) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if err := validate.MinimumInt("price", "body", *m.Price, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateSKU(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.SKU); err != nil {
		return err
	}

	if err := validate.Pattern("sku", "body", *m.SKU, `[0-9]+-[0-9]+-[0-9]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this product based on context it is used
func (m *Product) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
