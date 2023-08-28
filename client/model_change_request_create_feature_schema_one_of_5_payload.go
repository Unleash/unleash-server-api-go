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

// checks if the ChangeRequestCreateFeatureSchemaOneOf5Payload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestCreateFeatureSchemaOneOf5Payload{}

// ChangeRequestCreateFeatureSchemaOneOf5Payload struct for ChangeRequestCreateFeatureSchemaOneOf5Payload
type ChangeRequestCreateFeatureSchemaOneOf5Payload struct {
	Variants []VariantSchema `json:"variants"`
}

// NewChangeRequestCreateFeatureSchemaOneOf5Payload instantiates a new ChangeRequestCreateFeatureSchemaOneOf5Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestCreateFeatureSchemaOneOf5Payload(variants []VariantSchema) *ChangeRequestCreateFeatureSchemaOneOf5Payload {
	this := ChangeRequestCreateFeatureSchemaOneOf5Payload{}
	this.Variants = variants
	return &this
}

// NewChangeRequestCreateFeatureSchemaOneOf5PayloadWithDefaults instantiates a new ChangeRequestCreateFeatureSchemaOneOf5Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestCreateFeatureSchemaOneOf5PayloadWithDefaults() *ChangeRequestCreateFeatureSchemaOneOf5Payload {
	this := ChangeRequestCreateFeatureSchemaOneOf5Payload{}
	return &this
}

// GetVariants returns the Variants field value
func (o *ChangeRequestCreateFeatureSchemaOneOf5Payload) GetVariants() []VariantSchema {
	if o == nil {
		var ret []VariantSchema
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestCreateFeatureSchemaOneOf5Payload) GetVariantsOk() ([]VariantSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *ChangeRequestCreateFeatureSchemaOneOf5Payload) SetVariants(v []VariantSchema) {
	o.Variants = v
}

func (o ChangeRequestCreateFeatureSchemaOneOf5Payload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestCreateFeatureSchemaOneOf5Payload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variants"] = o.Variants
	return toSerialize, nil
}

type NullableChangeRequestCreateFeatureSchemaOneOf5Payload struct {
	value *ChangeRequestCreateFeatureSchemaOneOf5Payload
	isSet bool
}

func (v NullableChangeRequestCreateFeatureSchemaOneOf5Payload) Get() *ChangeRequestCreateFeatureSchemaOneOf5Payload {
	return v.value
}

func (v *NullableChangeRequestCreateFeatureSchemaOneOf5Payload) Set(val *ChangeRequestCreateFeatureSchemaOneOf5Payload) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestCreateFeatureSchemaOneOf5Payload) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestCreateFeatureSchemaOneOf5Payload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestCreateFeatureSchemaOneOf5Payload(val *ChangeRequestCreateFeatureSchemaOneOf5Payload) *NullableChangeRequestCreateFeatureSchemaOneOf5Payload {
	return &NullableChangeRequestCreateFeatureSchemaOneOf5Payload{value: val, isSet: true}
}

func (v NullableChangeRequestCreateFeatureSchemaOneOf5Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestCreateFeatureSchemaOneOf5Payload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
