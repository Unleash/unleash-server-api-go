/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ContextFieldStrategiesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextFieldStrategiesSchema{}

// ContextFieldStrategiesSchema A wrapper object containing all strategies that use a specific context field
type ContextFieldStrategiesSchema struct {
	// List of strategies using the context field
	Strategies []ContextFieldStrategiesSchemaStrategiesInner `json:"strategies"`
}

// NewContextFieldStrategiesSchema instantiates a new ContextFieldStrategiesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextFieldStrategiesSchema(strategies []ContextFieldStrategiesSchemaStrategiesInner) *ContextFieldStrategiesSchema {
	this := ContextFieldStrategiesSchema{}
	this.Strategies = strategies
	return &this
}

// NewContextFieldStrategiesSchemaWithDefaults instantiates a new ContextFieldStrategiesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextFieldStrategiesSchemaWithDefaults() *ContextFieldStrategiesSchema {
	this := ContextFieldStrategiesSchema{}
	return &this
}

// GetStrategies returns the Strategies field value
func (o *ContextFieldStrategiesSchema) GetStrategies() []ContextFieldStrategiesSchemaStrategiesInner {
	if o == nil {
		var ret []ContextFieldStrategiesSchemaStrategiesInner
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *ContextFieldStrategiesSchema) GetStrategiesOk() ([]ContextFieldStrategiesSchemaStrategiesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategies, true
}

// SetStrategies sets field value
func (o *ContextFieldStrategiesSchema) SetStrategies(v []ContextFieldStrategiesSchemaStrategiesInner) {
	o.Strategies = v
}

func (o ContextFieldStrategiesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextFieldStrategiesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["strategies"] = o.Strategies
	return toSerialize, nil
}

type NullableContextFieldStrategiesSchema struct {
	value *ContextFieldStrategiesSchema
	isSet bool
}

func (v NullableContextFieldStrategiesSchema) Get() *ContextFieldStrategiesSchema {
	return v.value
}

func (v *NullableContextFieldStrategiesSchema) Set(val *ContextFieldStrategiesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableContextFieldStrategiesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableContextFieldStrategiesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextFieldStrategiesSchema(val *ContextFieldStrategiesSchema) *NullableContextFieldStrategiesSchema {
	return &NullableContextFieldStrategiesSchema{value: val, isSet: true}
}

func (v NullableContextFieldStrategiesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextFieldStrategiesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
