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

// checks if the StrategiesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StrategiesSchema{}

// StrategiesSchema List of strategies
type StrategiesSchema struct {
	// Version of the strategies schema
	Version int32 `json:"version"`
	// List of strategies
	Strategies []StrategySchema `json:"strategies"`
}

// NewStrategiesSchema instantiates a new StrategiesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategiesSchema(version int32, strategies []StrategySchema) *StrategiesSchema {
	this := StrategiesSchema{}
	this.Version = version
	this.Strategies = strategies
	return &this
}

// NewStrategiesSchemaWithDefaults instantiates a new StrategiesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategiesSchemaWithDefaults() *StrategiesSchema {
	this := StrategiesSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *StrategiesSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *StrategiesSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *StrategiesSchema) SetVersion(v int32) {
	o.Version = v
}

// GetStrategies returns the Strategies field value
func (o *StrategiesSchema) GetStrategies() []StrategySchema {
	if o == nil {
		var ret []StrategySchema
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *StrategiesSchema) GetStrategiesOk() ([]StrategySchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategies, true
}

// SetStrategies sets field value
func (o *StrategiesSchema) SetStrategies(v []StrategySchema) {
	o.Strategies = v
}

func (o StrategiesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StrategiesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["strategies"] = o.Strategies
	return toSerialize, nil
}

type NullableStrategiesSchema struct {
	value *StrategiesSchema
	isSet bool
}

func (v NullableStrategiesSchema) Get() *StrategiesSchema {
	return v.value
}

func (v *NullableStrategiesSchema) Set(val *StrategiesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategiesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategiesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategiesSchema(val *StrategiesSchema) *NullableStrategiesSchema {
	return &NullableStrategiesSchema{value: val, isSet: true}
}

func (v NullableStrategiesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategiesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
