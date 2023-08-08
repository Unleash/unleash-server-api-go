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

// checks if the ValidatePasswordSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatePasswordSchema{}

// ValidatePasswordSchema Used to validate passwords obeying [Unleash password guidelines](https://docs.getunleash.io/reference/deploy/securing-unleash#password-requirements)
type ValidatePasswordSchema struct {
	// The password to validate
	Password string `json:"password"`
}

// NewValidatePasswordSchema instantiates a new ValidatePasswordSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatePasswordSchema(password string) *ValidatePasswordSchema {
	this := ValidatePasswordSchema{}
	this.Password = password
	return &this
}

// NewValidatePasswordSchemaWithDefaults instantiates a new ValidatePasswordSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatePasswordSchemaWithDefaults() *ValidatePasswordSchema {
	this := ValidatePasswordSchema{}
	return &this
}

// GetPassword returns the Password field value
func (o *ValidatePasswordSchema) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ValidatePasswordSchema) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ValidatePasswordSchema) SetPassword(v string) {
	o.Password = v
}

func (o ValidatePasswordSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatePasswordSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableValidatePasswordSchema struct {
	value *ValidatePasswordSchema
	isSet bool
}

func (v NullableValidatePasswordSchema) Get() *ValidatePasswordSchema {
	return v.value
}

func (v *NullableValidatePasswordSchema) Set(val *ValidatePasswordSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatePasswordSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatePasswordSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatePasswordSchema(val *ValidatePasswordSchema) *NullableValidatePasswordSchema {
	return &NullableValidatePasswordSchema{value: val, isSet: true}
}

func (v NullableValidatePasswordSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatePasswordSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
