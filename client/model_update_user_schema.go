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

// checks if the UpdateUserSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserSchema{}

// UpdateUserSchema All fields that can be directly changed for the user
type UpdateUserSchema struct {
	// The user's email address. Must be provided if username is not provided.
	Email *string `json:"email,omitempty"`
	// The user's name (not the user's username).
	Name     *string                   `json:"name,omitempty"`
	RootRole *CreateUserSchemaRootRole `json:"rootRole,omitempty"`
}

// NewUpdateUserSchema instantiates a new UpdateUserSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserSchema() *UpdateUserSchema {
	this := UpdateUserSchema{}
	return &this
}

// NewUpdateUserSchemaWithDefaults instantiates a new UpdateUserSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserSchemaWithDefaults() *UpdateUserSchema {
	this := UpdateUserSchema{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateUserSchema) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSchema) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserSchema) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateUserSchema) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateUserSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateUserSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateUserSchema) SetName(v string) {
	o.Name = &v
}

// GetRootRole returns the RootRole field value if set, zero value otherwise.
func (o *UpdateUserSchema) GetRootRole() CreateUserSchemaRootRole {
	if o == nil || IsNil(o.RootRole) {
		var ret CreateUserSchemaRootRole
		return ret
	}
	return *o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSchema) GetRootRoleOk() (*CreateUserSchemaRootRole, bool) {
	if o == nil || IsNil(o.RootRole) {
		return nil, false
	}
	return o.RootRole, true
}

// HasRootRole returns a boolean if a field has been set.
func (o *UpdateUserSchema) HasRootRole() bool {
	if o != nil && !IsNil(o.RootRole) {
		return true
	}

	return false
}

// SetRootRole gets a reference to the given CreateUserSchemaRootRole and assigns it to the RootRole field.
func (o *UpdateUserSchema) SetRootRole(v CreateUserSchemaRootRole) {
	o.RootRole = &v
}

func (o UpdateUserSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RootRole) {
		toSerialize["rootRole"] = o.RootRole
	}
	return toSerialize, nil
}

type NullableUpdateUserSchema struct {
	value *UpdateUserSchema
	isSet bool
}

func (v NullableUpdateUserSchema) Get() *UpdateUserSchema {
	return v.value
}

func (v *NullableUpdateUserSchema) Set(val *UpdateUserSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserSchema(val *UpdateUserSchema) *NullableUpdateUserSchema {
	return &NullableUpdateUserSchema{value: val, isSet: true}
}

func (v NullableUpdateUserSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
