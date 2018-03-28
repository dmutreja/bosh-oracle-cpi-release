// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// AttachIScsiVolumeDetailsAllOf1 attach i scsi volume details all of1
// swagger:discriminator attachIScsiVolumeDetailsAllOf1 iscsi
type AttachIScsiVolumeDetailsAllOf1 interface {

	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap() bool
	SetUseChap(bool)
}

type attachIScsiVolumeDetailsAllOf1 struct {
	useChapField bool
}

func (m *attachIScsiVolumeDetailsAllOf1) UseChap() bool {
	return m.useChapField
}
func (m *attachIScsiVolumeDetailsAllOf1) SetUseChap(val bool) {
	m.useChapField = val
}

// UnmarshalAttachIScsiVolumeDetailsAllOf1Slice unmarshals polymorphic slices of AttachIScsiVolumeDetailsAllOf1
func UnmarshalAttachIScsiVolumeDetailsAllOf1Slice(reader io.Reader, consumer runtime.Consumer) ([]AttachIScsiVolumeDetailsAllOf1, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AttachIScsiVolumeDetailsAllOf1
	for _, element := range elements {
		obj, err := unmarshalAttachIScsiVolumeDetailsAllOf1(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAttachIScsiVolumeDetailsAllOf1 unmarshals polymorphic AttachIScsiVolumeDetailsAllOf1
func UnmarshalAttachIScsiVolumeDetailsAllOf1(reader io.Reader, consumer runtime.Consumer) (AttachIScsiVolumeDetailsAllOf1, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAttachIScsiVolumeDetailsAllOf1(data, consumer)
}

func unmarshalAttachIScsiVolumeDetailsAllOf1(data []byte, consumer runtime.Consumer) (AttachIScsiVolumeDetailsAllOf1, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the iscsi property.
	var getType struct {
		Iscsi string `json:"iscsi"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("iscsi", "body", getType.Iscsi); err != nil {
		return nil, err
	}

	// The value of iscsi is used to determine which type to create and unmarshal the data into
	switch getType.Iscsi {
	case DiscriminatorTypeValues["AttachIScsiVolumeDetails"]:
		var result AttachIScsiVolumeDetails
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "attachIScsiVolumeDetailsAllOf1":
		var result attachIScsiVolumeDetailsAllOf1
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid iscsi value: %q", getType.Iscsi)

}
