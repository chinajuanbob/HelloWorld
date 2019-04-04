// Code generated by go-swagger; DO NOT EDIT.

package modules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PbTodo pb todo
// swagger:model pbTodo
type PbTodo struct {

	// content
	Content string `json:"content,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last updated
	LastUpdated *TimestampTimestamp `json:"last_updated,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this pb todo
func (m *PbTodo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PbTodo) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if m.LastUpdated != nil {
		if err := m.LastUpdated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_updated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PbTodo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbTodo) UnmarshalBinary(b []byte) error {
	var res PbTodo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
