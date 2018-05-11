// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteCategoriesResponse openpitrix delete categories response
// swagger:model openpitrixDeleteCategoriesResponse
type OpenpitrixDeleteCategoriesResponse struct {

	// category id
	CategoryID []string `json:"category_id"`
}

// Validate validates this openpitrix delete categories response
func (m *OpenpitrixDeleteCategoriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixDeleteCategoriesResponse) validateCategoryID(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryID) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteCategoriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteCategoriesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteCategoriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
