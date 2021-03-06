package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreatePolicyDetails create policy details
// swagger:model CreatePolicyDetails
type CreatePolicyDetails struct {

	// The OCID of the compartment containing the policy (either the tenancy or another compartment).
	// Required: true
	CompartmentID *string `json:"compartmentId"`

	// The description you assign to the policy during creation. Does not have to be unique, and it's changeable.
	//
	// Required: true
	// Max Length: 400
	// Min Length: 1
	Description *string `json:"description"`

	// The name you assign to the policy during creation. The name must be unique across all policies
	// in the tenancy and cannot be changed.
	//
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// An array of policy statements written in the policy language. See
	// [How Policies Work](/Content/Identity/Concepts/policies.htm) and
	// [Common Policies](/Content/Identity/Concepts/commonpolicies.htm).
	//
	// Required: true
	Statements []string `json:"statements"`

	// The version of the policy. If null or set to an empty string, when a request comes in for authorization, the
	// policy will be evaluated according to the current behavior of the services at that moment. If set to a particular
	// date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date.
	//
	VersionDate strfmt.DateTime `json:"versionDate,omitempty"`
}

// Validate validates this create policy details
func (m *CreatePolicyDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatements(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePolicyDetails) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	return nil
}

func (m *CreatePolicyDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 400); err != nil {
		return err
	}

	return nil
}

func (m *CreatePolicyDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *CreatePolicyDetails) validateStatements(formats strfmt.Registry) error {

	if err := validate.Required("statements", "body", m.Statements); err != nil {
		return err
	}

	return nil
}
