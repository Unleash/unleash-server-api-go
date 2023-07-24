/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PlaygroundStrategySchemaResultAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchemaResultAnyOf1{}

// PlaygroundStrategySchemaResultAnyOf1 struct for PlaygroundStrategySchemaResultAnyOf1
type PlaygroundStrategySchemaResultAnyOf1 struct {
	// Signals that this strategy was evaluated successfully.
	EvaluationStatus string `json:"evaluationStatus"`
	// Whether this strategy evaluates to true or not.
	Enabled bool `json:"enabled"`
}

// NewPlaygroundStrategySchemaResultAnyOf1 instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchemaResultAnyOf1(evaluationStatus string, enabled bool) *PlaygroundStrategySchemaResultAnyOf1 {
	this := PlaygroundStrategySchemaResultAnyOf1{}
	this.EvaluationStatus = evaluationStatus
	this.Enabled = enabled
	return &this
}

// NewPlaygroundStrategySchemaResultAnyOf1WithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaResultAnyOf1WithDefaults() *PlaygroundStrategySchemaResultAnyOf1 {
	this := PlaygroundStrategySchemaResultAnyOf1{}
	return &this
}

// GetEvaluationStatus returns the EvaluationStatus field value
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationStatus
}

// GetEvaluationStatusOk returns a tuple with the EvaluationStatus field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationStatus, true
}

// SetEvaluationStatus sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1) SetEvaluationStatus(v string) {
	o.EvaluationStatus = v
}

// GetEnabled returns the Enabled field value
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1) SetEnabled(v bool) {
	o.Enabled = v
}

func (o PlaygroundStrategySchemaResultAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchemaResultAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluationStatus"] = o.EvaluationStatus
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullablePlaygroundStrategySchemaResultAnyOf1 struct {
	value *PlaygroundStrategySchemaResultAnyOf1
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) Get() *PlaygroundStrategySchemaResultAnyOf1 {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) Set(val *PlaygroundStrategySchemaResultAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResultAnyOf1(val *PlaygroundStrategySchemaResultAnyOf1) *NullablePlaygroundStrategySchemaResultAnyOf1 {
	return &NullablePlaygroundStrategySchemaResultAnyOf1{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}