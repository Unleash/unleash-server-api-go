/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProjectUsersSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectUsersSchema{}

// ProjectUsersSchema struct for ProjectUsersSchema
type ProjectUsersSchema struct {
	Users []UserWithProjectRoleSchema `json:"users"`
	Roles []RoleSchema `json:"roles"`
}

// NewProjectUsersSchema instantiates a new ProjectUsersSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUsersSchema(users []UserWithProjectRoleSchema, roles []RoleSchema) *ProjectUsersSchema {
	this := ProjectUsersSchema{}
	this.Users = users
	this.Roles = roles
	return &this
}

// NewProjectUsersSchemaWithDefaults instantiates a new ProjectUsersSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUsersSchemaWithDefaults() *ProjectUsersSchema {
	this := ProjectUsersSchema{}
	return &this
}

// GetUsers returns the Users field value
func (o *ProjectUsersSchema) GetUsers() []UserWithProjectRoleSchema {
	if o == nil {
		var ret []UserWithProjectRoleSchema
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ProjectUsersSchema) GetUsersOk() ([]UserWithProjectRoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ProjectUsersSchema) SetUsers(v []UserWithProjectRoleSchema) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *ProjectUsersSchema) GetRoles() []RoleSchema {
	if o == nil {
		var ret []RoleSchema
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ProjectUsersSchema) GetRolesOk() ([]RoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ProjectUsersSchema) SetRoles(v []RoleSchema) {
	o.Roles = v
}

func (o ProjectUsersSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectUsersSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableProjectUsersSchema struct {
	value *ProjectUsersSchema
	isSet bool
}

func (v NullableProjectUsersSchema) Get() *ProjectUsersSchema {
	return v.value
}

func (v *NullableProjectUsersSchema) Set(val *ProjectUsersSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUsersSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUsersSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUsersSchema(val *ProjectUsersSchema) *NullableProjectUsersSchema {
	return &NullableProjectUsersSchema{value: val, isSet: true}
}

func (v NullableProjectUsersSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUsersSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


