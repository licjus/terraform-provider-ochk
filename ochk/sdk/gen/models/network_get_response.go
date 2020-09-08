// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkGetResponse NetworkGetResponse
//
// swagger:model NetworkGetResponse
type NetworkGetResponse struct {

	// messages
	Messages string `json:"messages,omitempty"`

	// network instance
	NetworkInstance *VCSNetworkInstance `json:"networkInstance,omitempty"`

	// success
	Success bool `json:"success,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this network get response
func (m *NetworkGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkInstance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkGetResponse) validateNetworkInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInstance) { // not required
		return nil
	}

	if m.NetworkInstance != nil {
		if err := m.NetworkInstance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkInstance")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkGetResponse) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkGetResponse) UnmarshalBinary(b []byte) error {
	var res NetworkGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
