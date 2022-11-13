// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProviderConfiguration provider configuration
//
// swagger:model ProviderConfiguration
type ProviderConfiguration struct {

	// open stack
	OpenStack *OpenStack `json:"openStack,omitempty"`
}

// Validate validates this provider configuration
func (m *ProviderConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpenStack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderConfiguration) validateOpenStack(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenStack) { // not required
		return nil
	}

	if m.OpenStack != nil {
		if err := m.OpenStack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openStack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openStack")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this provider configuration based on the context it is used
func (m *ProviderConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOpenStack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProviderConfiguration) contextValidateOpenStack(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenStack != nil {
		if err := m.OpenStack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openStack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openStack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProviderConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProviderConfiguration) UnmarshalBinary(b []byte) error {
	var res ProviderConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}