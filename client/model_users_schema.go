/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UsersSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersSchema{}

// UsersSchema Users and root roles
type UsersSchema struct {
	// A list of users in the Unleash instance.
	Users []UserSchema `json:"users"`
	// A list of [root roles](https://docs.getunleash.io/reference/rbac#standard-roles) in the Unleash instance.
	RootRoles []RoleSchema `json:"rootRoles,omitempty"`
}

// NewUsersSchema instantiates a new UsersSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersSchema(users []UserSchema) *UsersSchema {
	this := UsersSchema{}
	this.Users = users
	return &this
}

// NewUsersSchemaWithDefaults instantiates a new UsersSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersSchemaWithDefaults() *UsersSchema {
	this := UsersSchema{}
	return &this
}

// GetUsers returns the Users field value
func (o *UsersSchema) GetUsers() []UserSchema {
	if o == nil {
		var ret []UserSchema
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *UsersSchema) GetUsersOk() ([]UserSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *UsersSchema) SetUsers(v []UserSchema) {
	o.Users = v
}

// GetRootRoles returns the RootRoles field value if set, zero value otherwise.
func (o *UsersSchema) GetRootRoles() []RoleSchema {
	if o == nil || IsNil(o.RootRoles) {
		var ret []RoleSchema
		return ret
	}
	return o.RootRoles
}

// GetRootRolesOk returns a tuple with the RootRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersSchema) GetRootRolesOk() ([]RoleSchema, bool) {
	if o == nil || IsNil(o.RootRoles) {
		return nil, false
	}
	return o.RootRoles, true
}

// HasRootRoles returns a boolean if a field has been set.
func (o *UsersSchema) HasRootRoles() bool {
	if o != nil && !IsNil(o.RootRoles) {
		return true
	}

	return false
}

// SetRootRoles gets a reference to the given []RoleSchema and assigns it to the RootRoles field.
func (o *UsersSchema) SetRootRoles(v []RoleSchema) {
	o.RootRoles = v
}

func (o UsersSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	if !IsNil(o.RootRoles) {
		toSerialize["rootRoles"] = o.RootRoles
	}
	return toSerialize, nil
}

type NullableUsersSchema struct {
	value *UsersSchema
	isSet bool
}

func (v NullableUsersSchema) Get() *UsersSchema {
	return v.value
}

func (v *NullableUsersSchema) Set(val *UsersSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersSchema(val *UsersSchema) *NullableUsersSchema {
	return &NullableUsersSchema{value: val, isSet: true}
}

func (v NullableUsersSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
