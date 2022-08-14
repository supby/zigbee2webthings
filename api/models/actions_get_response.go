// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionsGetResponse actions get response
//
// swagger:model ActionsGetResponse
type ActionsGetResponse []map[string]ActionsGetResponseAnon

// Validate validates this actions get response
func (m ActionsGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		for k := range m[i] {

			if swag.IsZero(m[i][k]) { // not required
				continue
			}
			if val, ok := m[i][k]; ok {
				if err := val.Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName(strconv.Itoa(i) + "." + k)
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName(strconv.Itoa(i) + "." + k)
					}
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this actions get response based on the context it is used
func (m ActionsGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		for k := range m[i] {

			if val, ok := m[i][k]; ok {
				if err := val.ContextValidate(ctx, formats); err != nil {
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ActionsGetResponseAnon actions get response anon
//
// swagger:model ActionsGetResponseAnon
type ActionsGetResponseAnon struct {

	// href
	Href string `json:"href,omitempty"`

	// input
	Input map[string]string `json:"input,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// time requested
	TimeRequested string `json:"timeRequested,omitempty"`
}

// Validate validates this actions get response anon
func (m *ActionsGetResponseAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this actions get response anon based on context it is used
func (m *ActionsGetResponseAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActionsGetResponseAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionsGetResponseAnon) UnmarshalBinary(b []byte) error {
	var res ActionsGetResponseAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
