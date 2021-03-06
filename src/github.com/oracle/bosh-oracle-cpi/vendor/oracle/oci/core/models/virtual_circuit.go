// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualCircuit For use with Oracle Cloud Infrastructure FastConnect.
//
// A virtual circuit is an isolated network path that runs over one or more physical
// network connections to provide a single, logical connection between the edge router
// on the customer's existing network and Oracle Cloud Infrastructure. *Private*
// virtual circuits support private peering, and *public* virtual circuits support
// public peering. For more information, see [FastConnect Overview](/Content/Network/Concepts/fastconnect.htm).
//
// Each virtual circuit is made up of information shared between a customer, Oracle,
// and a provider (if the customer is using FastConnect via a provider). Who fills in
// a given property of a virtual circuit depends on whether the BGP session related to
// that virtual circuit goes from the customer's edge router to Oracle, or from the provider's
// edge router to Oracle. Also, in the case where the customer is using a provider, values
// for some of the properties may not be present immediately, but may get filled in as the
// provider and Oracle each do their part to provision the virtual circuit.
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model VirtualCircuit
type VirtualCircuit struct {

	// The provisioned data rate of the connection.
	BandwidthShapeName string `json:"bandwidthShapeName,omitempty"`

	// BGP management option.
	//
	BgpManagement string `json:"bgpManagement,omitempty"`

	// The state of the BGP session associated with the virtual circuit.
	BgpSessionState string `json:"bgpSessionState,omitempty"`

	// The OCID of the compartment containing the virtual circuit.
	// Max Length: 255
	// Min Length: 1
	CompartmentID string `json:"compartmentId,omitempty"`

	// An array of mappings, each containing properties for a
	// cross-connect or cross-connect group that is associated with this
	// virtual circuit.
	//
	// Maximum: 2
	// Minimum: 0
	CrossConnectMappings []*CrossConnectMapping `json:"crossConnectMappings"`

	// The BGP ASN of the network at the other end of the BGP
	// session from Oracle. If the session is between the customer's
	// edge router and Oracle, the value is the customer's ASN. If the BGP
	// session is between the provider's edge router and Oracle, the value
	// is the provider's ASN.
	//
	CustomerBgpAsn int64 `json:"customerBgpAsn,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The OCID of the customer's [Dynamic Routing Gateway (DRG)](#/en/iaas/20160918/Drg)
	// that this virtual circuit uses. Applicable only to private virtual circuits.
	//
	// Max Length: 255
	// Min Length: 1
	GatewayID string `json:"gatewayId,omitempty"`

	// The virtual circuit's Oracle ID (OCID).
	// Max Length: 255
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// The virtual circuit's current state. For information about
	// the different states, see
	// [FastConnect Overview](/Content/Network/Concepts/fastconnect.htm).
	//
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The Oracle BGP ASN.
	OracleBgpAsn int64 `json:"oracleBgpAsn,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	//
	// Max Length: 255
	// Min Length: 1
	ProviderName string `json:"providerName,omitempty"`

	// The OCID of the service offered by the provider (if the customer is connecting via a provider).
	//
	// Max Length: 255
	// Min Length: 1
	ProviderServiceID string `json:"providerServiceId,omitempty"`

	// Deprecated. Instead use `providerServiceId`.
	//
	// Max Length: 255
	// Min Length: 1
	ProviderServiceName string `json:"providerServiceName,omitempty"`

	// The provider's state in relation to this virtual circuit (if the
	// customer is connecting via a provider). ACTIVE means
	// the provider has provisioned the virtual circuit from their end.
	// INACTIVE means the provider has not yet provisioned the virtual
	// circuit, or has de-provisioned it.
	//
	ProviderState string `json:"providerState,omitempty"`

	// For a public virtual circuit. The public IP prefixes (CIDRs) the customer wants to
	// advertise across the connection. Each prefix must be /24 or less specific.
	//
	// Maximum: 50
	// Minimum: 0
	PublicPrefixes []string `json:"publicPrefixes"`

	// Provider-supplied reference information about this virtual circuit
	// (if the customer is connecting via a provider).
	//
	ReferenceComment string `json:"referenceComment,omitempty"`

	// The Oracle Cloud Infrastructure region where this virtual
	// circuit is located.
	//
	// Max Length: 255
	// Min Length: 1
	Region string `json:"region,omitempty"`

	// Provider service type.
	//
	ServiceType string `json:"serviceType,omitempty"`

	// The date and time the virtual circuit was created,
	// in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// Whether the virtual circuit supports private or public peering. For more information,
	// see [FastConnect Overview](/Content/Network/Concepts/fastconnect.htm).
	//
	Type string `json:"type,omitempty"`
}

// Validate validates this virtual circuit
func (m *VirtualCircuit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBgpManagement(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBgpSessionState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCrossConnectMappings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderServiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderServiceName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProviderState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublicPrefixes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var virtualCircuitTypeBgpManagementPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CUSTOMER_MANAGED","PROVIDER_MANAGED","ORACLE_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeBgpManagementPropEnum = append(virtualCircuitTypeBgpManagementPropEnum, v)
	}
}

const (

	// VirtualCircuitBgpManagementCUSTOMERMANAGED captures enum value "CUSTOMER_MANAGED"
	VirtualCircuitBgpManagementCUSTOMERMANAGED string = "CUSTOMER_MANAGED"

	// VirtualCircuitBgpManagementPROVIDERMANAGED captures enum value "PROVIDER_MANAGED"
	VirtualCircuitBgpManagementPROVIDERMANAGED string = "PROVIDER_MANAGED"

	// VirtualCircuitBgpManagementORACLEMANAGED captures enum value "ORACLE_MANAGED"
	VirtualCircuitBgpManagementORACLEMANAGED string = "ORACLE_MANAGED"
)

// prop value enum
func (m *VirtualCircuit) validateBgpManagementEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeBgpManagementPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateBgpManagement(formats strfmt.Registry) error {

	if swag.IsZero(m.BgpManagement) { // not required
		return nil
	}

	// value enum
	if err := m.validateBgpManagementEnum("bgpManagement", "body", m.BgpManagement); err != nil {
		return err
	}

	return nil
}

var virtualCircuitTypeBgpSessionStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UP","DOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeBgpSessionStatePropEnum = append(virtualCircuitTypeBgpSessionStatePropEnum, v)
	}
}

const (

	// VirtualCircuitBgpSessionStateUP captures enum value "UP"
	VirtualCircuitBgpSessionStateUP string = "UP"

	// VirtualCircuitBgpSessionStateDOWN captures enum value "DOWN"
	VirtualCircuitBgpSessionStateDOWN string = "DOWN"
)

// prop value enum
func (m *VirtualCircuit) validateBgpSessionStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeBgpSessionStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateBgpSessionState(formats strfmt.Registry) error {

	if swag.IsZero(m.BgpSessionState) { // not required
		return nil
	}

	// value enum
	if err := m.validateBgpSessionStateEnum("bgpSessionState", "body", m.BgpSessionState); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateCompartmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.CompartmentID) { // not required
		return nil
	}

	if err := validate.MinLength("compartmentId", "body", string(m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateCrossConnectMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.CrossConnectMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.CrossConnectMappings); i++ {

		if swag.IsZero(m.CrossConnectMappings[i]) { // not required
			continue
		}

		if m.CrossConnectMappings[i] != nil {

			if err := m.CrossConnectMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("crossConnectMappings" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *VirtualCircuit) validateDisplayName(formats strfmt.Registry) error {

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

func (m *VirtualCircuit) validateGatewayID(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayID) { // not required
		return nil
	}

	if err := validate.MinLength("gatewayId", "body", string(m.GatewayID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("gatewayId", "body", string(m.GatewayID), 255); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(m.ID), 255); err != nil {
		return err
	}

	return nil
}

var virtualCircuitTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING_PROVIDER","VERIFYING","PROVISIONING","PROVISIONED","FAILED","INACTIVE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeLifecycleStatePropEnum = append(virtualCircuitTypeLifecycleStatePropEnum, v)
	}
}

const (

	// VirtualCircuitLifecycleStatePENDINGPROVIDER captures enum value "PENDING_PROVIDER"
	VirtualCircuitLifecycleStatePENDINGPROVIDER string = "PENDING_PROVIDER"

	// VirtualCircuitLifecycleStateVERIFYING captures enum value "VERIFYING"
	VirtualCircuitLifecycleStateVERIFYING string = "VERIFYING"

	// VirtualCircuitLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	VirtualCircuitLifecycleStatePROVISIONING string = "PROVISIONING"

	// VirtualCircuitLifecycleStatePROVISIONED captures enum value "PROVISIONED"
	VirtualCircuitLifecycleStatePROVISIONED string = "PROVISIONED"

	// VirtualCircuitLifecycleStateFAILED captures enum value "FAILED"
	VirtualCircuitLifecycleStateFAILED string = "FAILED"

	// VirtualCircuitLifecycleStateINACTIVE captures enum value "INACTIVE"
	VirtualCircuitLifecycleStateINACTIVE string = "INACTIVE"

	// VirtualCircuitLifecycleStateTERMINATING captures enum value "TERMINATING"
	VirtualCircuitLifecycleStateTERMINATING string = "TERMINATING"

	// VirtualCircuitLifecycleStateTERMINATED captures enum value "TERMINATED"
	VirtualCircuitLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *VirtualCircuit) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateLifecycleState(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleState) { // not required
		return nil
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateProviderName(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderName) { // not required
		return nil
	}

	if err := validate.MinLength("providerName", "body", string(m.ProviderName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("providerName", "body", string(m.ProviderName), 255); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateProviderServiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderServiceID) { // not required
		return nil
	}

	if err := validate.MinLength("providerServiceId", "body", string(m.ProviderServiceID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("providerServiceId", "body", string(m.ProviderServiceID), 255); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateProviderServiceName(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderServiceName) { // not required
		return nil
	}

	if err := validate.MinLength("providerServiceName", "body", string(m.ProviderServiceName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("providerServiceName", "body", string(m.ProviderServiceName), 255); err != nil {
		return err
	}

	return nil
}

var virtualCircuitTypeProviderStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeProviderStatePropEnum = append(virtualCircuitTypeProviderStatePropEnum, v)
	}
}

const (

	// VirtualCircuitProviderStateACTIVE captures enum value "ACTIVE"
	VirtualCircuitProviderStateACTIVE string = "ACTIVE"

	// VirtualCircuitProviderStateINACTIVE captures enum value "INACTIVE"
	VirtualCircuitProviderStateINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *VirtualCircuit) validateProviderStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeProviderStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateProviderState(formats strfmt.Registry) error {

	if swag.IsZero(m.ProviderState) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderStateEnum("providerState", "body", m.ProviderState); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validatePublicPrefixes(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicPrefixes) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicPrefixes); i++ {

		if err := validate.MinLength("publicPrefixes"+"."+strconv.Itoa(i), "body", string(m.PublicPrefixes[i]), 9); err != nil {
			return err
		}

		if err := validate.MaxLength("publicPrefixes"+"."+strconv.Itoa(i), "body", string(m.PublicPrefixes[i]), 50); err != nil {
			return err
		}

	}

	return nil
}

func (m *VirtualCircuit) validateRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.Region) { // not required
		return nil
	}

	if err := validate.MinLength("region", "body", string(m.Region), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("region", "body", string(m.Region), 255); err != nil {
		return err
	}

	return nil
}

var virtualCircuitTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COLOCATED","LAYER2","LAYER3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeServiceTypePropEnum = append(virtualCircuitTypeServiceTypePropEnum, v)
	}
}

const (

	// VirtualCircuitServiceTypeCOLOCATED captures enum value "COLOCATED"
	VirtualCircuitServiceTypeCOLOCATED string = "COLOCATED"

	// VirtualCircuitServiceTypeLAYER2 captures enum value "LAYER2"
	VirtualCircuitServiceTypeLAYER2 string = "LAYER2"

	// VirtualCircuitServiceTypeLAYER3 captures enum value "LAYER3"
	VirtualCircuitServiceTypeLAYER3 string = "LAYER3"
)

// prop value enum
func (m *VirtualCircuit) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeServiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateServiceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceTypeEnum("serviceType", "body", m.ServiceType); err != nil {
		return err
	}

	return nil
}

func (m *VirtualCircuit) validateTimeCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("timeCreated", "body", "date-time", m.TimeCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

var virtualCircuitTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualCircuitTypeTypePropEnum = append(virtualCircuitTypeTypePropEnum, v)
	}
}

const (

	// VirtualCircuitTypePUBLIC captures enum value "PUBLIC"
	VirtualCircuitTypePUBLIC string = "PUBLIC"

	// VirtualCircuitTypePRIVATE captures enum value "PRIVATE"
	VirtualCircuitTypePRIVATE string = "PRIVATE"
)

// prop value enum
func (m *VirtualCircuit) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, virtualCircuitTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VirtualCircuit) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualCircuit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualCircuit) UnmarshalBinary(b []byte) error {
	var res VirtualCircuit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
