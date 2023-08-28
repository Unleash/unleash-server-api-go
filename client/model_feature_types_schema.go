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

// checks if the FeatureTypesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureTypesSchema{}

// FeatureTypesSchema A list of [feature toggle types](https://docs.getunleash.io/reference/feature-toggle-types) and the schema version used to represent those feature types.
type FeatureTypesSchema struct {
	// The schema version used to describe the feature toggle types listed in the `types` property.
	Version int32 `json:"version"`
	// The list of feature toggle types.
	Types []FeatureTypeSchema `json:"types"`
}

// NewFeatureTypesSchema instantiates a new FeatureTypesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureTypesSchema(version int32, types []FeatureTypeSchema) *FeatureTypesSchema {
	this := FeatureTypesSchema{}
	this.Version = version
	this.Types = types
	return &this
}

// NewFeatureTypesSchemaWithDefaults instantiates a new FeatureTypesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureTypesSchemaWithDefaults() *FeatureTypesSchema {
	this := FeatureTypesSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *FeatureTypesSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FeatureTypesSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FeatureTypesSchema) SetVersion(v int32) {
	o.Version = v
}

// GetTypes returns the Types field value
func (o *FeatureTypesSchema) GetTypes() []FeatureTypeSchema {
	if o == nil {
		var ret []FeatureTypeSchema
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *FeatureTypesSchema) GetTypesOk() ([]FeatureTypeSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *FeatureTypesSchema) SetTypes(v []FeatureTypeSchema) {
	o.Types = v
}

func (o FeatureTypesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureTypesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["types"] = o.Types
	return toSerialize, nil
}

type NullableFeatureTypesSchema struct {
	value *FeatureTypesSchema
	isSet bool
}

func (v NullableFeatureTypesSchema) Get() *FeatureTypesSchema {
	return v.value
}

func (v *NullableFeatureTypesSchema) Set(val *FeatureTypesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureTypesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureTypesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureTypesSchema(val *FeatureTypesSchema) *NullableFeatureTypesSchema {
	return &NullableFeatureTypesSchema{value: val, isSet: true}
}

func (v NullableFeatureTypesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureTypesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
