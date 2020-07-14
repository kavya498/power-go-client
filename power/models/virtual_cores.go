// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualCores virtual cores
// swagger:model VirtualCores
type VirtualCores struct {

	// The active virtual Cores
	// Required: true
	Assigned *int64 `json:"assigned"`

	// The maximum DLPAR range for virtual Cores
	// Required: true
	Max *int64 `json:"max"`

	// The minimum DLPAR range for virtual Cores
	// Required: true
	Min *int64 `json:"min"`
}

// Validate validates this virtual cores
func (m *VirtualCores) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssigned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualCores) validateAssigned(formats strfmt.Registry) error {

	if err := validate.Required("assigned", "body", m.Assigned); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCores) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCores) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualCores) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualCores) UnmarshalBinary(b []byte) error {
	var res VirtualCores
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}