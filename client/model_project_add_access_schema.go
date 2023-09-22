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

// checks if the ProjectAddAccessSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAddAccessSchema{}

// ProjectAddAccessSchema An object containing a collection of roles, a collection of groups and a collection of users.
type ProjectAddAccessSchema struct {
	// A list of role IDs
	Roles []int32 `json:"roles"`
	// A list of group IDs
	Groups []int32 `json:"groups"`
	// A list of user IDs
	Users []int32 `json:"users"`
}

// NewProjectAddAccessSchema instantiates a new ProjectAddAccessSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAddAccessSchema(roles []int32, groups []int32, users []int32) *ProjectAddAccessSchema {
	this := ProjectAddAccessSchema{}
	this.Roles = roles
	this.Groups = groups
	this.Users = users
	return &this
}

// NewProjectAddAccessSchemaWithDefaults instantiates a new ProjectAddAccessSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAddAccessSchemaWithDefaults() *ProjectAddAccessSchema {
	this := ProjectAddAccessSchema{}
	return &this
}

// GetRoles returns the Roles field value
func (o *ProjectAddAccessSchema) GetRoles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ProjectAddAccessSchema) GetRolesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ProjectAddAccessSchema) SetRoles(v []int32) {
	o.Roles = v
}

// GetGroups returns the Groups field value
func (o *ProjectAddAccessSchema) GetGroups() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *ProjectAddAccessSchema) GetGroupsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *ProjectAddAccessSchema) SetGroups(v []int32) {
	o.Groups = v
}

// GetUsers returns the Users field value
func (o *ProjectAddAccessSchema) GetUsers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ProjectAddAccessSchema) GetUsersOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ProjectAddAccessSchema) SetUsers(v []int32) {
	o.Users = v
}

func (o ProjectAddAccessSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAddAccessSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	toSerialize["groups"] = o.Groups
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

type NullableProjectAddAccessSchema struct {
	value *ProjectAddAccessSchema
	isSet bool
}

func (v NullableProjectAddAccessSchema) Get() *ProjectAddAccessSchema {
	return v.value
}

func (v *NullableProjectAddAccessSchema) Set(val *ProjectAddAccessSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAddAccessSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAddAccessSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAddAccessSchema(val *ProjectAddAccessSchema) *NullableProjectAddAccessSchema {
	return &NullableProjectAddAccessSchema{value: val, isSet: true}
}

func (v NullableProjectAddAccessSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAddAccessSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}