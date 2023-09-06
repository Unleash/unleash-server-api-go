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

// checks if the ProxyFeaturesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyFeaturesSchema{}

// ProxyFeaturesSchema Frontend SDK features list
type ProxyFeaturesSchema struct {
	// The actual features returned to the Frontend SDK
	Toggles []ProxyFeatureSchema `json:"toggles"`
}

// NewProxyFeaturesSchema instantiates a new ProxyFeaturesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyFeaturesSchema(toggles []ProxyFeatureSchema) *ProxyFeaturesSchema {
	this := ProxyFeaturesSchema{}
	this.Toggles = toggles
	return &this
}

// NewProxyFeaturesSchemaWithDefaults instantiates a new ProxyFeaturesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyFeaturesSchemaWithDefaults() *ProxyFeaturesSchema {
	this := ProxyFeaturesSchema{}
	return &this
}

// GetToggles returns the Toggles field value
func (o *ProxyFeaturesSchema) GetToggles() []ProxyFeatureSchema {
	if o == nil {
		var ret []ProxyFeatureSchema
		return ret
	}

	return o.Toggles
}

// GetTogglesOk returns a tuple with the Toggles field value
// and a boolean to check if the value has been set.
func (o *ProxyFeaturesSchema) GetTogglesOk() ([]ProxyFeatureSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Toggles, true
}

// SetToggles sets field value
func (o *ProxyFeaturesSchema) SetToggles(v []ProxyFeatureSchema) {
	o.Toggles = v
}

func (o ProxyFeaturesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyFeaturesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toggles"] = o.Toggles
	return toSerialize, nil
}

type NullableProxyFeaturesSchema struct {
	value *ProxyFeaturesSchema
	isSet bool
}

func (v NullableProxyFeaturesSchema) Get() *ProxyFeaturesSchema {
	return v.value
}

func (v *NullableProxyFeaturesSchema) Set(val *ProxyFeaturesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyFeaturesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyFeaturesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyFeaturesSchema(val *ProxyFeaturesSchema) *NullableProxyFeaturesSchema {
	return &NullableProxyFeaturesSchema{value: val, isSet: true}
}

func (v NullableProxyFeaturesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyFeaturesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
