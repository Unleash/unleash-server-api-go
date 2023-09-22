/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the AdminPermissionsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminPermissionsSchema{}

// AdminPermissionsSchema What kind of permissions are available
type AdminPermissionsSchema struct {
	Permissions AdminPermissionsSchemaPermissions `json:"permissions"`
	// The api version of this response. A natural increasing number. Only increases if format changes
	Version int32 `json:"version"`
}

// NewAdminPermissionsSchema instantiates a new AdminPermissionsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminPermissionsSchema(permissions AdminPermissionsSchemaPermissions, version int32) *AdminPermissionsSchema {
	this := AdminPermissionsSchema{}
	this.Permissions = permissions
	this.Version = version
	return &this
}

// NewAdminPermissionsSchemaWithDefaults instantiates a new AdminPermissionsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminPermissionsSchemaWithDefaults() *AdminPermissionsSchema {
	this := AdminPermissionsSchema{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *AdminPermissionsSchema) GetPermissions() AdminPermissionsSchemaPermissions {
	if o == nil {
		var ret AdminPermissionsSchemaPermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionsSchema) GetPermissionsOk() (*AdminPermissionsSchemaPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *AdminPermissionsSchema) SetPermissions(v AdminPermissionsSchemaPermissions) {
	o.Permissions = v
}

// GetVersion returns the Version field value
func (o *AdminPermissionsSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AdminPermissionsSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AdminPermissionsSchema) SetVersion(v int32) {
	o.Version = v
}

func (o AdminPermissionsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminPermissionsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAdminPermissionsSchema struct {
	value *AdminPermissionsSchema
	isSet bool
}

func (v NullableAdminPermissionsSchema) Get() *AdminPermissionsSchema {
	return v.value
}

func (v *NullableAdminPermissionsSchema) Set(val *AdminPermissionsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminPermissionsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminPermissionsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminPermissionsSchema(val *AdminPermissionsSchema) *NullableAdminPermissionsSchema {
	return &NullableAdminPermissionsSchema{value: val, isSet: true}
}

func (v NullableAdminPermissionsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminPermissionsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}