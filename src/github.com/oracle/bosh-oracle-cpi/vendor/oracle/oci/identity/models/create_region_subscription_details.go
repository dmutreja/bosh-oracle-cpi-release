package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CreateRegionSubscriptionDetails create region subscription details
// swagger:model CreateRegionSubscriptionDetails
type CreateRegionSubscriptionDetails struct {

	// The key of the region such as PHX, IAD.
	// Required: true
	// Max Length: 16
	// Min Length: 1
	RegionKey *string `json:"regionKey"`
}

// Validate validates this create region subscription details
func (m *CreateRegionSubscriptionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegionKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRegionSubscriptionDetails) validateRegionKey(formats strfmt.Registry) error {

	if err := validate.Required("regionKey", "body", m.RegionKey); err != nil {
		return err
	}

	if err := validate.MinLength("regionKey", "body", string(*m.RegionKey), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("regionKey", "body", string(*m.RegionKey), 16); err != nil {
		return err
	}

	return nil
}
