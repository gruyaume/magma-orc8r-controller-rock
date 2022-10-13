// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriberStatus Status of a subscriber device
// swagger:model subscriber_status
type SubscriberStatus struct {

	// icmp
	Icmp *IcmpStatus `json:"icmp,omitempty"`
}

// Validate validates this subscriber status
func (m *SubscriberStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIcmp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriberStatus) validateIcmp(formats strfmt.Registry) error {

	if swag.IsZero(m.Icmp) { // not required
		return nil
	}

	if m.Icmp != nil {
		if err := m.Icmp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icmp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriberStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriberStatus) UnmarshalBinary(b []byte) error {
	var res SubscriberStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
