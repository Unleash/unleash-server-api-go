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

// checks if the ResetPasswordSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordSchema{}

// ResetPasswordSchema struct for ResetPasswordSchema
type ResetPasswordSchema struct {
	ResetPasswordUrl string `json:"resetPasswordUrl"`
}

// NewResetPasswordSchema instantiates a new ResetPasswordSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordSchema(resetPasswordUrl string) *ResetPasswordSchema {
	this := ResetPasswordSchema{}
	this.ResetPasswordUrl = resetPasswordUrl
	return &this
}

// NewResetPasswordSchemaWithDefaults instantiates a new ResetPasswordSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordSchemaWithDefaults() *ResetPasswordSchema {
	this := ResetPasswordSchema{}
	return &this
}

// GetResetPasswordUrl returns the ResetPasswordUrl field value
func (o *ResetPasswordSchema) GetResetPasswordUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetPasswordUrl
}

// GetResetPasswordUrlOk returns a tuple with the ResetPasswordUrl field value
// and a boolean to check if the value has been set.
func (o *ResetPasswordSchema) GetResetPasswordUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetPasswordUrl, true
}

// SetResetPasswordUrl sets field value
func (o *ResetPasswordSchema) SetResetPasswordUrl(v string) {
	o.ResetPasswordUrl = v
}

func (o ResetPasswordSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resetPasswordUrl"] = o.ResetPasswordUrl
	return toSerialize, nil
}

type NullableResetPasswordSchema struct {
	value *ResetPasswordSchema
	isSet bool
}

func (v NullableResetPasswordSchema) Get() *ResetPasswordSchema {
	return v.value
}

func (v *NullableResetPasswordSchema) Set(val *ResetPasswordSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordSchema(val *ResetPasswordSchema) *NullableResetPasswordSchema {
	return &NullableResetPasswordSchema{value: val, isSet: true}
}

func (v NullableResetPasswordSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
