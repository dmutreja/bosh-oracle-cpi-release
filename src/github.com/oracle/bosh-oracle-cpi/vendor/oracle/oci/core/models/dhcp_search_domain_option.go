// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DhcpSearchDomainOption DHCP option for specifying a search domain name for DNS queries. For more information, see
// [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).
//
// swagger:model DhcpSearchDomainOption
type DhcpSearchDomainOption struct {
	typeField string

	searchDomainNamesField []string
}

func (m *DhcpSearchDomainOption) Type() string {
	//return m.typeField
	return DiscriminatorTypeValues["DhcpSearchDomainOption"]
}
func (m *DhcpSearchDomainOption) SetType(val string) {
	m.typeField = val
}

func (m *DhcpSearchDomainOption) SearchDomainNames() []string {
	return m.searchDomainNamesField
}
func (m *DhcpSearchDomainOption) SetSearchDomainNames(val []string) {
	m.searchDomainNamesField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DhcpSearchDomainOption) UnmarshalJSON(raw []byte) error {
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

		Type string `json:"type"`

		SearchDomainNames []string `json:"searchDomainNames"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DhcpSearchDomainOption

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.searchDomainNamesField = base.SearchDomainNames

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DhcpSearchDomainOption) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Type string `json:"type"`

		SearchDomainNames []string `json:"searchDomainNames"`
	}{

		Type: m.Type(),

		SearchDomainNames: m.SearchDomainNames(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this dhcp search domain option
func (m *DhcpSearchDomainOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchDomainNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpSearchDomainOption) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type())); err != nil {
		return err
	}

	return nil
}

func (m *DhcpSearchDomainOption) validateSearchDomainNames(formats strfmt.Registry) error {

	if err := validate.Required("searchDomainNames", "body", m.SearchDomainNames()); err != nil {
		return err
	}

	iSearchDomainNamesSize := int64(len(m.SearchDomainNames()))

	if err := validate.MaxItems("searchDomainNames", "body", iSearchDomainNamesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.SearchDomainNames()); i++ {

		if err := validate.MinLength("searchDomainNames"+"."+strconv.Itoa(i), "body", string(m.searchDomainNamesField[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("searchDomainNames"+"."+strconv.Itoa(i), "body", string(m.searchDomainNamesField[i]), 251); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DhcpSearchDomainOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DhcpSearchDomainOption) UnmarshalBinary(b []byte) error {
	var res DhcpSearchDomainOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
