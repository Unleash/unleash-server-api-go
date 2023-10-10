/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.6.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectAccessSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectAccessSchema{}

// ProjectAccessSchema An object describing access permissions for a given project.
type ProjectAccessSchema struct {
	// A list of groups that have access to this project
	Groups []GroupWithProjectRoleSchema `json:"groups"`
	// A list of users and their roles within this project
	Users []UserWithProjectRoleSchema `json:"users"`
	// A list of roles that are available within this project.
	Roles []RoleSchema `json:"roles"`
}

// NewProjectAccessSchema instantiates a new ProjectAccessSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAccessSchema(groups []GroupWithProjectRoleSchema, users []UserWithProjectRoleSchema, roles []RoleSchema) *ProjectAccessSchema {
	this := ProjectAccessSchema{}
	this.Groups = groups
	this.Users = users
	this.Roles = roles
	return &this
}

// NewProjectAccessSchemaWithDefaults instantiates a new ProjectAccessSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAccessSchemaWithDefaults() *ProjectAccessSchema {
	this := ProjectAccessSchema{}
	return &this
}

// GetGroups returns the Groups field value
func (o *ProjectAccessSchema) GetGroups() []GroupWithProjectRoleSchema {
	if o == nil {
		var ret []GroupWithProjectRoleSchema
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *ProjectAccessSchema) GetGroupsOk() ([]GroupWithProjectRoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *ProjectAccessSchema) SetGroups(v []GroupWithProjectRoleSchema) {
	o.Groups = v
}

// GetUsers returns the Users field value
func (o *ProjectAccessSchema) GetUsers() []UserWithProjectRoleSchema {
	if o == nil {
		var ret []UserWithProjectRoleSchema
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ProjectAccessSchema) GetUsersOk() ([]UserWithProjectRoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ProjectAccessSchema) SetUsers(v []UserWithProjectRoleSchema) {
	o.Users = v
}

// GetRoles returns the Roles field value
func (o *ProjectAccessSchema) GetRoles() []RoleSchema {
	if o == nil {
		var ret []RoleSchema
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ProjectAccessSchema) GetRolesOk() ([]RoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ProjectAccessSchema) SetRoles(v []RoleSchema) {
	o.Roles = v
}

func (o ProjectAccessSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAccessSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	toSerialize["users"] = o.Users
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableProjectAccessSchema struct {
	value *ProjectAccessSchema
	isSet bool
}

func (v NullableProjectAccessSchema) Get() *ProjectAccessSchema {
	return v.value
}

func (v *NullableProjectAccessSchema) Set(val *ProjectAccessSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAccessSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAccessSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAccessSchema(val *ProjectAccessSchema) *NullableProjectAccessSchema {
	return &NullableProjectAccessSchema{value: val, isSet: true}
}

func (v NullableProjectAccessSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAccessSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
