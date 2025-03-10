/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the RoleWithVersionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleWithVersionSchema{}

// RoleWithVersionSchema A single user role received after creation or update of a role
type RoleWithVersionSchema struct {
	// The version of this schema
	Version              int32      `json:"version"`
	Roles                RoleSchema `json:"roles"`
	AdditionalProperties map[string]interface{}
}

type _RoleWithVersionSchema RoleWithVersionSchema

// NewRoleWithVersionSchema instantiates a new RoleWithVersionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleWithVersionSchema(version int32, roles RoleSchema) *RoleWithVersionSchema {
	this := RoleWithVersionSchema{}
	this.Version = version
	this.Roles = roles
	return &this
}

// NewRoleWithVersionSchemaWithDefaults instantiates a new RoleWithVersionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithVersionSchemaWithDefaults() *RoleWithVersionSchema {
	this := RoleWithVersionSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *RoleWithVersionSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RoleWithVersionSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *RoleWithVersionSchema) SetVersion(v int32) {
	o.Version = v
}

// GetRoles returns the Roles field value
func (o *RoleWithVersionSchema) GetRoles() RoleSchema {
	if o == nil {
		var ret RoleSchema
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *RoleWithVersionSchema) GetRolesOk() (*RoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *RoleWithVersionSchema) SetRoles(v RoleSchema) {
	o.Roles = v
}

func (o RoleWithVersionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleWithVersionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["roles"] = o.Roles

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleWithVersionSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"roles",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoleWithVersionSchema := _RoleWithVersionSchema{}

	err = json.Unmarshal(data, &varRoleWithVersionSchema)

	if err != nil {
		return err
	}

	*o = RoleWithVersionSchema(varRoleWithVersionSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "roles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleWithVersionSchema struct {
	value *RoleWithVersionSchema
	isSet bool
}

func (v NullableRoleWithVersionSchema) Get() *RoleWithVersionSchema {
	return v.value
}

func (v *NullableRoleWithVersionSchema) Set(val *RoleWithVersionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleWithVersionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleWithVersionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleWithVersionSchema(val *RoleWithVersionSchema) *NullableRoleWithVersionSchema {
	return &NullableRoleWithVersionSchema{value: val, isSet: true}
}

func (v NullableRoleWithVersionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleWithVersionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
