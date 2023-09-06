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

// checks if the CreateGroupSchemaUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupSchemaUsersInner{}

// CreateGroupSchemaUsersInner A minimal user object
type CreateGroupSchemaUsersInner struct {
	User CreateGroupSchemaUsersInnerUser `json:"user"`
}

// NewCreateGroupSchemaUsersInner instantiates a new CreateGroupSchemaUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupSchemaUsersInner(user CreateGroupSchemaUsersInnerUser) *CreateGroupSchemaUsersInner {
	this := CreateGroupSchemaUsersInner{}
	this.User = user
	return &this
}

// NewCreateGroupSchemaUsersInnerWithDefaults instantiates a new CreateGroupSchemaUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupSchemaUsersInnerWithDefaults() *CreateGroupSchemaUsersInner {
	this := CreateGroupSchemaUsersInner{}
	return &this
}

// GetUser returns the User field value
func (o *CreateGroupSchemaUsersInner) GetUser() CreateGroupSchemaUsersInnerUser {
	if o == nil {
		var ret CreateGroupSchemaUsersInnerUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *CreateGroupSchemaUsersInner) GetUserOk() (*CreateGroupSchemaUsersInnerUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *CreateGroupSchemaUsersInner) SetUser(v CreateGroupSchemaUsersInnerUser) {
	o.User = v
}

func (o CreateGroupSchemaUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupSchemaUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableCreateGroupSchemaUsersInner struct {
	value *CreateGroupSchemaUsersInner
	isSet bool
}

func (v NullableCreateGroupSchemaUsersInner) Get() *CreateGroupSchemaUsersInner {
	return v.value
}

func (v *NullableCreateGroupSchemaUsersInner) Set(val *CreateGroupSchemaUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupSchemaUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupSchemaUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupSchemaUsersInner(val *CreateGroupSchemaUsersInner) *NullableCreateGroupSchemaUsersInner {
	return &NullableCreateGroupSchemaUsersInner{value: val, isSet: true}
}

func (v NullableCreateGroupSchemaUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupSchemaUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
