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

// CreateInstanceConsoleConnectionDetails The details for creating a instance console connection.
// The instance console connection is created in the same compartment as the instance.
//
// swagger:model CreateInstanceConsoleConnectionDetails
type CreateInstanceConsoleConnectionDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Operations": {"CostCenter": "42"}}`
	//
	DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Department": "Finance"}`
	//
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// The OCID of the instance to create the console connection to.
	// Required: true
	InstanceID *string `json:"instanceId"`

	// The SSH public key used to authenticate the console connection.
	// Required: true
	PublicKey *string `json:"publicKey"`
}

// Validate validates this create instance console connection details
func (m *CreateInstanceConsoleConnectionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateInstanceConsoleConnectionDetails) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instanceId", "body", m.InstanceID); err != nil {
		return err
	}

	return nil
}

func (m *CreateInstanceConsoleConnectionDetails) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateInstanceConsoleConnectionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateInstanceConsoleConnectionDetails) UnmarshalBinary(b []byte) error {
	var res CreateInstanceConsoleConnectionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
