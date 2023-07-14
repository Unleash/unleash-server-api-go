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

// checks if the FeatureVariantsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureVariantsSchema{}

// FeatureVariantsSchema struct for FeatureVariantsSchema
type FeatureVariantsSchema struct {
	Version  int32           `json:"version"`
	Variants []VariantSchema `json:"variants"`
}

// NewFeatureVariantsSchema instantiates a new FeatureVariantsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureVariantsSchema(version int32, variants []VariantSchema) *FeatureVariantsSchema {
	this := FeatureVariantsSchema{}
	this.Version = version
	this.Variants = variants
	return &this
}

// NewFeatureVariantsSchemaWithDefaults instantiates a new FeatureVariantsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureVariantsSchemaWithDefaults() *FeatureVariantsSchema {
	this := FeatureVariantsSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *FeatureVariantsSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FeatureVariantsSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FeatureVariantsSchema) SetVersion(v int32) {
	o.Version = v
}

// GetVariants returns the Variants field value
func (o *FeatureVariantsSchema) GetVariants() []VariantSchema {
	if o == nil {
		var ret []VariantSchema
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *FeatureVariantsSchema) GetVariantsOk() ([]VariantSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *FeatureVariantsSchema) SetVariants(v []VariantSchema) {
	o.Variants = v
}

func (o FeatureVariantsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureVariantsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["variants"] = o.Variants
	return toSerialize, nil
}

type NullableFeatureVariantsSchema struct {
	value *FeatureVariantsSchema
	isSet bool
}

func (v NullableFeatureVariantsSchema) Get() *FeatureVariantsSchema {
	return v.value
}

func (v *NullableFeatureVariantsSchema) Set(val *FeatureVariantsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureVariantsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureVariantsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureVariantsSchema(val *FeatureVariantsSchema) *NullableFeatureVariantsSchema {
	return &NullableFeatureVariantsSchema{value: val, isSet: true}
}

func (v NullableFeatureVariantsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureVariantsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
