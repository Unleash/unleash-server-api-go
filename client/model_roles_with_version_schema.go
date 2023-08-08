/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the RolesWithVersionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesWithVersionSchema{}

// RolesWithVersionSchema struct for RolesWithVersionSchema
type RolesWithVersionSchema struct {
	Version float32      `json:"version"`
	Roles   []RoleSchema `json:"roles"`
}

// NewRolesWithVersionSchema instantiates a new RolesWithVersionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesWithVersionSchema(version float32, roles []RoleSchema) *RolesWithVersionSchema {
	this := RolesWithVersionSchema{}
	this.Version = version
	this.Roles = roles
	return &this
}

// NewRolesWithVersionSchemaWithDefaults instantiates a new RolesWithVersionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesWithVersionSchemaWithDefaults() *RolesWithVersionSchema {
	this := RolesWithVersionSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *RolesWithVersionSchema) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RolesWithVersionSchema) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RolesWithVersionSchema) SetVersion(v float32) {
	o.Version = v
}

// GetRoles returns the Roles field value
func (o *RolesWithVersionSchema) GetRoles() []RoleSchema {
	if o == nil {
		var ret []RoleSchema
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RolesWithVersionSchema) GetRolesOk() ([]RoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *RolesWithVersionSchema) SetRoles(v []RoleSchema) {
	o.Roles = v
}

func (o RolesWithVersionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesWithVersionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableRolesWithVersionSchema struct {
	value *RolesWithVersionSchema
	isSet bool
}

func (v NullableRolesWithVersionSchema) Get() *RolesWithVersionSchema {
	return v.value
}

func (v *NullableRolesWithVersionSchema) Set(val *RolesWithVersionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesWithVersionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesWithVersionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesWithVersionSchema(val *RolesWithVersionSchema) *NullableRolesWithVersionSchema {
	return &NullableRolesWithVersionSchema{value: val, isSet: true}
}

func (v NullableRolesWithVersionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesWithVersionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}