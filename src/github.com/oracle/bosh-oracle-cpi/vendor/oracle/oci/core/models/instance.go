// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Instance A compute host. The image used to launch the instance determines its operating system and other
// software. The shape specified during the launch process determines the number of CPUs and memory
// allocated to the instance. For more information, see
// [Overview of the Compute Service](/Content/Compute/Concepts/computeoverview.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model Instance
type Instance struct {

	// The Availability Domain the instance is running in.
	//
	// Example: `Uocm:PHX-AD-1`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	AvailabilityDomain *string `json:"availabilityDomain"`

	// The OCID of the compartment that contains the instance.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Operations": {"CostCenter": "42"}}`
	//
	DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Example: `My bare metal instance`
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Additional metadata key/value pairs that you provide.  They serve a similar purpose and functionality from fields in the 'metadata' object.
	//
	// They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).
	//
	// If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead.
	//
	ExtendedMetadata map[string]interface{} `json:"extendedMetadata,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// [Resource Tags](/Content/General/Concepts/resourcetags.htm).
	//
	// Example: `{"Department": "Finance"}`
	//
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// The OCID of the instance.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// Deprecated. Use `sourceDetails` instead.
	//
	// Max Length: 255
	// Min Length: 1
	ImageID string `json:"imageId,omitempty"`

	// When a bare metal or virtual machine
	// instance boots, the iPXE firmware that runs on the instance is
	// configured to run an iPXE script to continue the boot process.
	//
	// If you want more control over the boot process, you can provide
	// your own custom iPXE script that will run when the instance boots;
	// however, you should be aware that the same iPXE script will run
	// every time an instance boots; not only after the initial
	// LaunchInstance call.
	//
	// The default iPXE script connects to the instance's local boot
	// volume over iSCSI and performs a network boot. If you use a custom iPXE
	// script and want to network-boot from the instance's local boot volume
	// over iSCSI the same way as the default iPXE script, you should use the
	// following iSCSI IP address: 169.254.0.2, and boot volume IQN:
	// iqn.2015-02.oracle.boot.
	//
	// For more information about the Bring Your Own Image feature of
	// Oracle Cloud Infrastructure, see
	// [Bring Your Own Image](/Content/Compute/References/bringyourownimage.htm).
	//
	// For more information about iPXE, see http://ipxe.org.
	//
	// Max Length: 10240
	// Min Length: 1
	IPXEScript string `json:"ipxeScript,omitempty"`

	// Specifies the configuration mode for launching virtual machine (VM) instances. The configuration modes are:
	// * `NATIVE` - VM instances launch with iSCSI boot and VFIO devices. The default value for Oracle-provided images.
	// * `EMULATED` - VM instances launch with emulated devices, such as the E1000 network driver and emulated SCSI disk controller.
	// * `CUSTOM` - VM instances launch with custom configuration settings specified in the `LaunchOptions` parameter.
	//
	LaunchMode string `json:"launchMode,omitempty"`

	// launch options
	LaunchOptions *LaunchOptions `json:"launchOptions,omitempty"`

	// The current state of the instance.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// Custom metadata that you provide.
	Metadata map[string]string `json:"metadata,omitempty"`

	// The region that contains the Availability Domain the instance is running in.
	//
	// Example: `phx`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Region *string `json:"region"`

	// The shape of the instance. The shape determines the number of CPUs and the amount of memory
	// allocated to the instance. You can enumerate all available shapes by calling
	// [ListShapes](#/en/iaas/20160918/Shape/ListShapes).
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Shape *string `json:"shape"`

	sourceDetailsField InstanceSourceDetails

	// The date and time the instance was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	// Required: true
	TimeCreated *strfmt.DateTime `json:"timeCreated"`
}

func (m *Instance) SourceDetails() InstanceSourceDetails {
	return m.sourceDetailsField
}
func (m *Instance) SetSourceDetails(val InstanceSourceDetails) {
	m.sourceDetailsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Instance) UnmarshalJSON(raw []byte) error {
	var data struct {
		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		ExtendedMetadata map[string]interface{} `json:"extendedMetadata,omitempty"`

		FreeformTags map[string]string `json:"freeformTags,omitempty"`

		ID *string `json:"id"`

		ImageID string `json:"imageId,omitempty"`

		IPXEScript string `json:"ipxeScript,omitempty"`

		LaunchMode string `json:"launchMode,omitempty"`

		LaunchOptions *LaunchOptions `json:"launchOptions,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		Metadata map[string]string `json:"metadata,omitempty"`

		Region *string `json:"region"`

		Shape *string `json:"shape"`

		SourceDetails json.RawMessage `json:"sourceDetails,omitempty"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	sourceDetails, err := UnmarshalInstanceSourceDetails(bytes.NewBuffer(data.SourceDetails), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result Instance

	// availabilityDomain
	result.AvailabilityDomain = data.AvailabilityDomain

	// compartmentId
	result.CompartmentID = data.CompartmentID

	// definedTags
	result.DefinedTags = data.DefinedTags

	// displayName
	result.DisplayName = data.DisplayName

	// extendedMetadata
	result.ExtendedMetadata = data.ExtendedMetadata

	// freeformTags
	result.FreeformTags = data.FreeformTags

	// id
	result.ID = data.ID

	// imageId
	result.ImageID = data.ImageID

	// ipxeScript
	result.IPXEScript = data.IPXEScript

	// launchMode
	result.LaunchMode = data.LaunchMode

	// launchOptions
	result.LaunchOptions = data.LaunchOptions

	// lifecycleState
	result.LifecycleState = data.LifecycleState

	// metadata
	result.Metadata = data.Metadata

	// region
	result.Region = data.Region

	// shape
	result.Shape = data.Shape

	// sourceDetails
	result.sourceDetailsField = sourceDetails

	// timeCreated
	result.TimeCreated = data.TimeCreated

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Instance) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		AvailabilityDomain *string `json:"availabilityDomain"`

		CompartmentID *string `json:"compartmentId"`

		DefinedTags map[string]map[string]interface{} `json:"definedTags,omitempty"`

		DisplayName string `json:"displayName,omitempty"`

		ExtendedMetadata map[string]interface{} `json:"extendedMetadata,omitempty"`

		FreeformTags map[string]string `json:"freeformTags,omitempty"`

		ID *string `json:"id"`

		ImageID string `json:"imageId,omitempty"`

		IPXEScript string `json:"ipxeScript,omitempty"`

		LaunchMode string `json:"launchMode,omitempty"`

		LaunchOptions *LaunchOptions `json:"launchOptions,omitempty"`

		LifecycleState *string `json:"lifecycleState"`

		Metadata map[string]string `json:"metadata,omitempty"`

		Region *string `json:"region"`

		Shape *string `json:"shape"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`
	}{

		AvailabilityDomain: m.AvailabilityDomain,

		CompartmentID: m.CompartmentID,

		DefinedTags: m.DefinedTags,

		DisplayName: m.DisplayName,

		ExtendedMetadata: m.ExtendedMetadata,

		FreeformTags: m.FreeformTags,

		ID: m.ID,

		ImageID: m.ImageID,

		IPXEScript: m.IPXEScript,

		LaunchMode: m.LaunchMode,

		LaunchOptions: m.LaunchOptions,

		LifecycleState: m.LifecycleState,

		Metadata: m.Metadata,

		Region: m.Region,

		Shape: m.Shape,

		TimeCreated: m.TimeCreated,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		SourceDetails InstanceSourceDetails `json:"sourceDetails,omitempty"`
	}{

		SourceDetails: m.sourceDetailsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this instance
func (m *Instance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPXEScript(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaunchMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLaunchOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShape(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instance) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateImageID(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageID) { // not required
		return nil
	}

	if err := validate.MinLength("imageId", "body", string(m.ImageID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("imageId", "body", string(m.ImageID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateIPXEScript(formats strfmt.Registry) error {

	if swag.IsZero(m.IPXEScript) { // not required
		return nil
	}

	if err := validate.MinLength("ipxeScript", "body", string(m.IPXEScript), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ipxeScript", "body", string(m.IPXEScript), 10240); err != nil {
		return err
	}

	return nil
}

var instanceTypeLaunchModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NATIVE","EMULATED","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeLaunchModePropEnum = append(instanceTypeLaunchModePropEnum, v)
	}
}

const (

	// InstanceLaunchModeNATIVE captures enum value "NATIVE"
	InstanceLaunchModeNATIVE string = "NATIVE"

	// InstanceLaunchModeEMULATED captures enum value "EMULATED"
	InstanceLaunchModeEMULATED string = "EMULATED"

	// InstanceLaunchModeCUSTOM captures enum value "CUSTOM"
	InstanceLaunchModeCUSTOM string = "CUSTOM"
)

// prop value enum
func (m *Instance) validateLaunchModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceTypeLaunchModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instance) validateLaunchMode(formats strfmt.Registry) error {

	if swag.IsZero(m.LaunchMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateLaunchModeEnum("launchMode", "body", m.LaunchMode); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateLaunchOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.LaunchOptions) { // not required
		return nil
	}

	if m.LaunchOptions != nil {

		if err := m.LaunchOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("launchOptions")
			}
			return err
		}

	}

	return nil
}

var instanceTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","RUNNING","STARTING","STOPPING","STOPPED","CREATING_IMAGE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeLifecycleStatePropEnum = append(instanceTypeLifecycleStatePropEnum, v)
	}
}

const (

	// InstanceLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	InstanceLifecycleStatePROVISIONING string = "PROVISIONING"

	// InstanceLifecycleStateRUNNING captures enum value "RUNNING"
	InstanceLifecycleStateRUNNING string = "RUNNING"

	// InstanceLifecycleStateSTARTING captures enum value "STARTING"
	InstanceLifecycleStateSTARTING string = "STARTING"

	// InstanceLifecycleStateSTOPPING captures enum value "STOPPING"
	InstanceLifecycleStateSTOPPING string = "STOPPING"

	// InstanceLifecycleStateSTOPPED captures enum value "STOPPED"
	InstanceLifecycleStateSTOPPED string = "STOPPED"

	// InstanceLifecycleStateCREATINGIMAGE captures enum value "CREATING_IMAGE"
	InstanceLifecycleStateCREATINGIMAGE string = "CREATING_IMAGE"

	// InstanceLifecycleStateTERMINATING captures enum value "TERMINATING"
	InstanceLifecycleStateTERMINATING string = "TERMINATING"

	// InstanceLifecycleStateTERMINATED captures enum value "TERMINATED"
	InstanceLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *Instance) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instance) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	if err := validate.MinLength("region", "body", string(*m.Region), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("region", "body", string(*m.Region), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateShape(formats strfmt.Registry) error {

	if err := validate.Required("shape", "body", m.Shape); err != nil {
		return err
	}

	if err := validate.MinLength("shape", "body", string(*m.Shape), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("shape", "body", string(*m.Shape), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated); err != nil {
		return err
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Instance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instance) UnmarshalBinary(b []byte) error {
	var res Instance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
