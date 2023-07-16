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

// checks if the FeaturesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturesSchema{}

// FeaturesSchema A list of features
type FeaturesSchema struct {
	// The version of the feature's schema
	Version int32 `json:"version"`
	// A list of features
	Features []FeatureSchema `json:"features"`
}

// NewFeaturesSchema instantiates a new FeaturesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesSchema(version int32, features []FeatureSchema) *FeaturesSchema {
	this := FeaturesSchema{}
	this.Version = version
	this.Features = features
	return &this
}

// NewFeaturesSchemaWithDefaults instantiates a new FeaturesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesSchemaWithDefaults() *FeaturesSchema {
	this := FeaturesSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *FeaturesSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *FeaturesSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *FeaturesSchema) SetVersion(v int32) {
	o.Version = v
}

// GetFeatures returns the Features field value
func (o *FeaturesSchema) GetFeatures() []FeatureSchema {
	if o == nil {
		var ret []FeatureSchema
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *FeaturesSchema) GetFeaturesOk() ([]FeatureSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *FeaturesSchema) SetFeatures(v []FeatureSchema) {
	o.Features = v
}

func (o FeaturesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

type NullableFeaturesSchema struct {
	value *FeaturesSchema
	isSet bool
}

func (v NullableFeaturesSchema) Get() *FeaturesSchema {
	return v.value
}

func (v *NullableFeaturesSchema) Set(val *FeaturesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesSchema(val *FeaturesSchema) *NullableFeaturesSchema {
	return &NullableFeaturesSchema{value: val, isSet: true}
}

func (v NullableFeaturesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
