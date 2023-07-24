/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VariantFlagSchemaPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariantFlagSchemaPayload{}

// VariantFlagSchemaPayload struct for VariantFlagSchemaPayload
type VariantFlagSchemaPayload struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewVariantFlagSchemaPayload instantiates a new VariantFlagSchemaPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariantFlagSchemaPayload() *VariantFlagSchemaPayload {
	this := VariantFlagSchemaPayload{}
	return &this
}

// NewVariantFlagSchemaPayloadWithDefaults instantiates a new VariantFlagSchemaPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariantFlagSchemaPayloadWithDefaults() *VariantFlagSchemaPayload {
	this := VariantFlagSchemaPayload{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VariantFlagSchemaPayload) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantFlagSchemaPayload) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VariantFlagSchemaPayload) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VariantFlagSchemaPayload) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VariantFlagSchemaPayload) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariantFlagSchemaPayload) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VariantFlagSchemaPayload) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VariantFlagSchemaPayload) SetValue(v string) {
	o.Value = &v
}

func (o VariantFlagSchemaPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariantFlagSchemaPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableVariantFlagSchemaPayload struct {
	value *VariantFlagSchemaPayload
	isSet bool
}

func (v NullableVariantFlagSchemaPayload) Get() *VariantFlagSchemaPayload {
	return v.value
}

func (v *NullableVariantFlagSchemaPayload) Set(val *VariantFlagSchemaPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableVariantFlagSchemaPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableVariantFlagSchemaPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariantFlagSchemaPayload(val *VariantFlagSchemaPayload) *NullableVariantFlagSchemaPayload {
	return &NullableVariantFlagSchemaPayload{value: val, isSet: true}
}

func (v NullableVariantFlagSchemaPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariantFlagSchemaPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
