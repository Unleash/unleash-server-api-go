/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UsersGroupsBaseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersGroupsBaseSchema{}

// UsersGroupsBaseSchema struct for UsersGroupsBaseSchema
type UsersGroupsBaseSchema struct {
	Groups []GroupSchema `json:"groups,omitempty"`
	Users  []UserSchema  `json:"users,omitempty"`
}

// NewUsersGroupsBaseSchema instantiates a new UsersGroupsBaseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersGroupsBaseSchema() *UsersGroupsBaseSchema {
	this := UsersGroupsBaseSchema{}
	return &this
}

// NewUsersGroupsBaseSchemaWithDefaults instantiates a new UsersGroupsBaseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersGroupsBaseSchemaWithDefaults() *UsersGroupsBaseSchema {
	this := UsersGroupsBaseSchema{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UsersGroupsBaseSchema) GetGroups() []GroupSchema {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupSchema
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersGroupsBaseSchema) GetGroupsOk() ([]GroupSchema, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UsersGroupsBaseSchema) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSchema and assigns it to the Groups field.
func (o *UsersGroupsBaseSchema) SetGroups(v []GroupSchema) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UsersGroupsBaseSchema) GetUsers() []UserSchema {
	if o == nil || IsNil(o.Users) {
		var ret []UserSchema
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersGroupsBaseSchema) GetUsersOk() ([]UserSchema, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersGroupsBaseSchema) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserSchema and assigns it to the Users field.
func (o *UsersGroupsBaseSchema) SetUsers(v []UserSchema) {
	o.Users = v
}

func (o UsersGroupsBaseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersGroupsBaseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUsersGroupsBaseSchema struct {
	value *UsersGroupsBaseSchema
	isSet bool
}

func (v NullableUsersGroupsBaseSchema) Get() *UsersGroupsBaseSchema {
	return v.value
}

func (v *NullableUsersGroupsBaseSchema) Set(val *UsersGroupsBaseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersGroupsBaseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersGroupsBaseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersGroupsBaseSchema(val *UsersGroupsBaseSchema) *NullableUsersGroupsBaseSchema {
	return &NullableUsersGroupsBaseSchema{value: val, isSet: true}
}

func (v NullableUsersGroupsBaseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersGroupsBaseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
