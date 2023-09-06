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

// checks if the DoraFeaturesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DoraFeaturesSchema{}

// DoraFeaturesSchema The representation of a dora time to production feature metric
type DoraFeaturesSchema struct {
	// The name of a feature toggle
	Name string `json:"name"`
	// The average number of days it takes a feature toggle to get into production
	TimeToProduction float32 `json:"timeToProduction"`
}

// NewDoraFeaturesSchema instantiates a new DoraFeaturesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDoraFeaturesSchema(name string, timeToProduction float32) *DoraFeaturesSchema {
	this := DoraFeaturesSchema{}
	this.Name = name
	this.TimeToProduction = timeToProduction
	return &this
}

// NewDoraFeaturesSchemaWithDefaults instantiates a new DoraFeaturesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDoraFeaturesSchemaWithDefaults() *DoraFeaturesSchema {
	this := DoraFeaturesSchema{}
	return &this
}

// GetName returns the Name field value
func (o *DoraFeaturesSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DoraFeaturesSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DoraFeaturesSchema) SetName(v string) {
	o.Name = v
}

// GetTimeToProduction returns the TimeToProduction field value
func (o *DoraFeaturesSchema) GetTimeToProduction() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TimeToProduction
}

// GetTimeToProductionOk returns a tuple with the TimeToProduction field value
// and a boolean to check if the value has been set.
func (o *DoraFeaturesSchema) GetTimeToProductionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeToProduction, true
}

// SetTimeToProduction sets field value
func (o *DoraFeaturesSchema) SetTimeToProduction(v float32) {
	o.TimeToProduction = v
}

func (o DoraFeaturesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DoraFeaturesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["timeToProduction"] = o.TimeToProduction
	return toSerialize, nil
}

type NullableDoraFeaturesSchema struct {
	value *DoraFeaturesSchema
	isSet bool
}

func (v NullableDoraFeaturesSchema) Get() *DoraFeaturesSchema {
	return v.value
}

func (v *NullableDoraFeaturesSchema) Set(val *DoraFeaturesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableDoraFeaturesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableDoraFeaturesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDoraFeaturesSchema(val *DoraFeaturesSchema) *NullableDoraFeaturesSchema {
	return &NullableDoraFeaturesSchema{value: val, isSet: true}
}

func (v NullableDoraFeaturesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDoraFeaturesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
