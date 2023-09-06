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

// checks if the ProjectAddRoleAccessSchemaUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAddRoleAccessSchemaUsersInner{}

// ProjectAddRoleAccessSchemaUsersInner struct for ProjectAddRoleAccessSchemaUsersInner
type ProjectAddRoleAccessSchemaUsersInner struct {
	// A user ID
	Id int32 `json:"id"`
}

// NewProjectAddRoleAccessSchemaUsersInner instantiates a new ProjectAddRoleAccessSchemaUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAddRoleAccessSchemaUsersInner(id int32) *ProjectAddRoleAccessSchemaUsersInner {
	this := ProjectAddRoleAccessSchemaUsersInner{}
	this.Id = id
	return &this
}

// NewProjectAddRoleAccessSchemaUsersInnerWithDefaults instantiates a new ProjectAddRoleAccessSchemaUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAddRoleAccessSchemaUsersInnerWithDefaults() *ProjectAddRoleAccessSchemaUsersInner {
	this := ProjectAddRoleAccessSchemaUsersInner{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectAddRoleAccessSchemaUsersInner) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectAddRoleAccessSchemaUsersInner) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectAddRoleAccessSchemaUsersInner) SetId(v int32) {
	o.Id = v
}

func (o ProjectAddRoleAccessSchemaUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAddRoleAccessSchemaUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableProjectAddRoleAccessSchemaUsersInner struct {
	value *ProjectAddRoleAccessSchemaUsersInner
	isSet bool
}

func (v NullableProjectAddRoleAccessSchemaUsersInner) Get() *ProjectAddRoleAccessSchemaUsersInner {
	return v.value
}

func (v *NullableProjectAddRoleAccessSchemaUsersInner) Set(val *ProjectAddRoleAccessSchemaUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAddRoleAccessSchemaUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAddRoleAccessSchemaUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAddRoleAccessSchemaUsersInner(val *ProjectAddRoleAccessSchemaUsersInner) *NullableProjectAddRoleAccessSchemaUsersInner {
	return &NullableProjectAddRoleAccessSchemaUsersInner{value: val, isSet: true}
}

func (v NullableProjectAddRoleAccessSchemaUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAddRoleAccessSchemaUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
