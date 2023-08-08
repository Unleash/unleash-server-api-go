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

// checks if the AdvancedPlaygroundEnvironmentFeatureSchemaStrategies type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedPlaygroundEnvironmentFeatureSchemaStrategies{}

// AdvancedPlaygroundEnvironmentFeatureSchemaStrategies Feature's applicable strategies and cumulative results of the strategies
type AdvancedPlaygroundEnvironmentFeatureSchemaStrategies struct {
	Result AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult `json:"result"`
	// The strategies that apply to this feature.
	Data []PlaygroundStrategySchema `json:"data"`
}

// NewAdvancedPlaygroundEnvironmentFeatureSchemaStrategies instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaStrategies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedPlaygroundEnvironmentFeatureSchemaStrategies(result AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult, data []PlaygroundStrategySchema) *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies {
	this := AdvancedPlaygroundEnvironmentFeatureSchemaStrategies{}
	this.Result = result
	this.Data = data
	return &this
}

// NewAdvancedPlaygroundEnvironmentFeatureSchemaStrategiesWithDefaults instantiates a new AdvancedPlaygroundEnvironmentFeatureSchemaStrategies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedPlaygroundEnvironmentFeatureSchemaStrategiesWithDefaults() *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies {
	this := AdvancedPlaygroundEnvironmentFeatureSchemaStrategies{}
	return &this
}

// GetResult returns the Result field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetResult() AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult {
	if o == nil {
		var ret AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetResultOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) SetResult(v AdvancedPlaygroundEnvironmentFeatureSchemaStrategiesResult) {
	o.Result = v
}

// GetData returns the Data field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetData() []PlaygroundStrategySchema {
	if o == nil {
		var ret []PlaygroundStrategySchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) GetDataOk() ([]PlaygroundStrategySchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) SetData(v []PlaygroundStrategySchema) {
	o.Data = v
}

func (o AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies struct {
	value *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies
	isSet bool
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) Get() *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies {
	return v.value
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) Set(val *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies(val *AdvancedPlaygroundEnvironmentFeatureSchemaStrategies) *NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies {
	return &NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies{value: val, isSet: true}
}

func (v NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedPlaygroundEnvironmentFeatureSchemaStrategies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
