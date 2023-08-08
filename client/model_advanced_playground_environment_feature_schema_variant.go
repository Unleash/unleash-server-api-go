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

// checks if the AdvancedPlaygroundEnvironmentFeatureSchemaVariant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedPlaygroundEnvironmentFeatureSchemaVariant{}

// AdvancedPlaygroundEnvironmentFeatureSchemaVariant The feature variant you receive based on the provided context or the _disabled                           variant_. If a feature is disabled or doesn't have any                           variants, you would get the _disabled variant_.                           Otherwise, you'll get one of the feature's defined variants.
type AdvancedPlaygroundEnvironmentFeatureSchemaVariant struct {
	// The variant's name. If there is no variant or if the toggle is disabled, this will be `disabled`
	Name string `json:"name"`
	// Whether the variant is enabled or not. If the feature is disabled or if it doesn't have variants, this property will be `false`
	Enabled bool                                                      `json:"enabled"`
	Payload *AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload `json:"payload,omitempty"`
}

// NewAdvancedPlaygroundEnvironmentFeatureSchemaVariant instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedPlaygroundEnvironmentFeatureSchemaVariant(name string, enabled bool) *AdvancedPlaygroundEnvironmentFeatureSchemaVariant {
	this := AdvancedPlaygroundEnvironmentFeatureSchemaVariant{}
	this.Name = name
	this.Enabled = enabled
	return &this
}

// NewAdvancedPlaygroundEnvironmentFeatureSchemaVariantWithDefaults instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedPlaygroundEnvironmentFeatureSchemaVariantWithDefaults() *AdvancedPlaygroundEnvironmentFeatureSchemaVariant {
	this := AdvancedPlaygroundEnvironmentFeatureSchemaVariant{}
	return &this
}

// GetName returns the Name field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetPayload() AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload {
	if o == nil || IsNil(o.Payload) {
		var ret AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) GetPayloadOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload and assigns it to the Payload field.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) SetPayload(v AdvancedPlaygroundEnvironmentFeatureSchemaVariantPayload) {
	o.Payload = &v
}

func (o AdvancedPlaygroundEnvironmentFeatureSchemaVariant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedPlaygroundEnvironmentFeatureSchemaVariant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant struct {
	value *AdvancedPlaygroundEnvironmentFeatureSchemaVariant
	isSet bool
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) Get() *AdvancedPlaygroundEnvironmentFeatureSchemaVariant {
	return v.value
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) Set(val *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant(val *AdvancedPlaygroundEnvironmentFeatureSchemaVariant) *NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant {
	return &NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant{value: val, isSet: true}
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
