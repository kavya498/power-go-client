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

// VolumesCloneRequest volumes clone request
// swagger:model VolumesCloneRequest
type VolumesCloneRequest struct {

	// Prefix to use when naming the new cloned volumes
	NamingPrefix string `json:"namingPrefix,omitempty"`

	// List of volumes to be cloned
	// Required: true
	VolumeIds []string `json:"volumeIDs"`
}

// Validate validates this volumes clone request
func (m *VolumesCloneRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumesCloneRequest) validateVolumeIds(formats strfmt.Registry) error {

	if err := validate.Required("volumeIDs", "body", m.VolumeIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumesCloneRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumesCloneRequest) UnmarshalBinary(b []byte) error {
	var res VolumesCloneRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
