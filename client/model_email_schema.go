/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the EmailSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSchema{}

// EmailSchema Represents the email of a user. Used to send email communication (reset password, welcome mail etc)
type EmailSchema struct {
	// The email address
	Email string `json:"email"`
}

// NewEmailSchema instantiates a new EmailSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSchema(email string) *EmailSchema {
	this := EmailSchema{}
	this.Email = email
	return &this
}

// NewEmailSchemaWithDefaults instantiates a new EmailSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSchemaWithDefaults() *EmailSchema {
	this := EmailSchema{}
	return &this
}

// GetEmail returns the Email field value
func (o *EmailSchema) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *EmailSchema) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *EmailSchema) SetEmail(v string) {
	o.Email = v
}

func (o EmailSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableEmailSchema struct {
	value *EmailSchema
	isSet bool
}

func (v NullableEmailSchema) Get() *EmailSchema {
	return v.value
}

func (v *NullableEmailSchema) Set(val *EmailSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSchema(val *EmailSchema) *NullableEmailSchema {
	return &NullableEmailSchema{value: val, isSet: true}
}

func (v NullableEmailSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
