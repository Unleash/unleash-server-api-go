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

// checks if the PlaygroundStrategySchemaResultAnyOf1Variant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchemaResultAnyOf1Variant{}

// PlaygroundStrategySchemaResultAnyOf1Variant The feature variant you receive based on the provided context or the _disabled                           variant_. If a feature is disabled or doesn't have any                           variants, you would get the _disabled variant_.                           Otherwise, you'll get one of the feature's defined variants.
type PlaygroundStrategySchemaResultAnyOf1Variant struct {
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool                                                `json:"enabled"`
	Payload *PlaygroundStrategySchemaResultAnyOf1VariantPayload `json:"payload,omitempty"`
}

// NewPlaygroundStrategySchemaResultAnyOf1Variant instantiates a new PlaygroundStrategySchemaResultAnyOf1Variant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchemaResultAnyOf1Variant(name string, enabled bool) *PlaygroundStrategySchemaResultAnyOf1Variant {
	this := PlaygroundStrategySchemaResultAnyOf1Variant{}
	this.Name = name
	this.Enabled = enabled
	return &this
}

// NewPlaygroundStrategySchemaResultAnyOf1VariantWithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf1Variant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaResultAnyOf1VariantWithDefaults() *PlaygroundStrategySchemaResultAnyOf1Variant {
	this := PlaygroundStrategySchemaResultAnyOf1Variant{}
	return &this
}

// GetName returns the Name field value
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetPayload() PlaygroundStrategySchemaResultAnyOf1VariantPayload {
	if o == nil || IsNil(o.Payload) {
		var ret PlaygroundStrategySchemaResultAnyOf1VariantPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) GetPayloadOk() (*PlaygroundStrategySchemaResultAnyOf1VariantPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PlaygroundStrategySchemaResultAnyOf1VariantPayload and assigns it to the Payload field.
func (o *PlaygroundStrategySchemaResultAnyOf1Variant) SetPayload(v PlaygroundStrategySchemaResultAnyOf1VariantPayload) {
	o.Payload = &v
}

func (o PlaygroundStrategySchemaResultAnyOf1Variant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchemaResultAnyOf1Variant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullablePlaygroundStrategySchemaResultAnyOf1Variant struct {
	value *PlaygroundStrategySchemaResultAnyOf1Variant
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1Variant) Get() *PlaygroundStrategySchemaResultAnyOf1Variant {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1Variant) Set(val *PlaygroundStrategySchemaResultAnyOf1Variant) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1Variant) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1Variant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResultAnyOf1Variant(val *PlaygroundStrategySchemaResultAnyOf1Variant) *NullablePlaygroundStrategySchemaResultAnyOf1Variant {
	return &NullablePlaygroundStrategySchemaResultAnyOf1Variant{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1Variant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1Variant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
