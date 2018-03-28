// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ExportImageViaObjectStorageURIDetails export image via object storage Uri details
// swagger:model ExportImageViaObjectStorageUriDetails
type ExportImageViaObjectStorageURIDetails struct {
	destinationTypeField string

	destinationUriField *string
}

func (m *ExportImageViaObjectStorageURIDetails) DestinationType() string {
	// return m.destinationTypeField
	return DiscriminatorTypeValues["ExportImageViaObjectStorageUriDetails"]
}
func (m *ExportImageViaObjectStorageURIDetails) SetDestinationType(val string) {
	m.destinationTypeField = val
}

func (m *ExportImageViaObjectStorageURIDetails) DestinationURI() *string {
	return m.destinationUriField
}
func (m *ExportImageViaObjectStorageURIDetails) SetDestinationURI(val *string) {
	m.destinationUriField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ExportImageViaObjectStorageURIDetails) UnmarshalJSON(raw []byte) error {
	var data struct {
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		DestinationType string `json:"destinationType"`

		DestinationURI *string `json:"destinationUri"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result ExportImageViaObjectStorageURIDetails

	if base.DestinationType != result.DestinationType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid destinationType value: %q", base.DestinationType)
	}

	result.destinationUriField = base.DestinationURI

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ExportImageViaObjectStorageURIDetails) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		DestinationType string `json:"destinationType"`

		DestinationURI *string `json:"destinationUri"`
	}{

		DestinationType: m.DestinationType(),

		DestinationURI: m.DestinationURI(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this export image via object storage Uri details
func (m *ExportImageViaObjectStorageURIDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportImageViaObjectStorageURIDetails) validateDestinationType(formats strfmt.Registry) error {

	if err := validate.RequiredString("destinationType", "body", string(m.DestinationType())); err != nil {
		return err
	}

	return nil
}

func (m *ExportImageViaObjectStorageURIDetails) validateDestinationURI(formats strfmt.Registry) error {

	if err := validate.Required("destinationUri", "body", m.DestinationURI()); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportImageViaObjectStorageURIDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportImageViaObjectStorageURIDetails) UnmarshalBinary(b []byte) error {
	var res ExportImageViaObjectStorageURIDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
