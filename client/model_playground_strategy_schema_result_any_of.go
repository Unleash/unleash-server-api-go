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

// checks if the PlaygroundStrategySchemaResultAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchemaResultAnyOf{}

// PlaygroundStrategySchemaResultAnyOf struct for PlaygroundStrategySchemaResultAnyOf
type PlaygroundStrategySchemaResultAnyOf struct {
	// Signals that this strategy could not be evaluated. This is most likely because you're using a custom strategy that Unleash doesn't know about.
	EvaluationStatus string                                     `json:"evaluationStatus"`
	Enabled          PlaygroundStrategySchemaResultAnyOfEnabled `json:"enabled"`
}

// NewPlaygroundStrategySchemaResultAnyOf instantiates a new PlaygroundStrategySchemaResultAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchemaResultAnyOf(evaluationStatus string, enabled PlaygroundStrategySchemaResultAnyOfEnabled) *PlaygroundStrategySchemaResultAnyOf {
	this := PlaygroundStrategySchemaResultAnyOf{}
	this.EvaluationStatus = evaluationStatus
	this.Enabled = enabled
	return &this
}

// NewPlaygroundStrategySchemaResultAnyOfWithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaResultAnyOfWithDefaults() *PlaygroundStrategySchemaResultAnyOf {
	this := PlaygroundStrategySchemaResultAnyOf{}
	return &this
}

// GetEvaluationStatus returns the EvaluationStatus field value
func (o *PlaygroundStrategySchemaResultAnyOf) GetEvaluationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationStatus
}

// GetEvaluationStatusOk returns a tuple with the EvaluationStatus field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf) GetEvaluationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationStatus, true
}

// SetEvaluationStatus sets field value
func (o *PlaygroundStrategySchemaResultAnyOf) SetEvaluationStatus(v string) {
	o.EvaluationStatus = v
}

// GetEnabled returns the Enabled field value
func (o *PlaygroundStrategySchemaResultAnyOf) GetEnabled() PlaygroundStrategySchemaResultAnyOfEnabled {
	if o == nil {
		var ret PlaygroundStrategySchemaResultAnyOfEnabled
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf) GetEnabledOk() (*PlaygroundStrategySchemaResultAnyOfEnabled, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PlaygroundStrategySchemaResultAnyOf) SetEnabled(v PlaygroundStrategySchemaResultAnyOfEnabled) {
	o.Enabled = v
}

func (o PlaygroundStrategySchemaResultAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchemaResultAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluationStatus"] = o.EvaluationStatus
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullablePlaygroundStrategySchemaResultAnyOf struct {
	value *PlaygroundStrategySchemaResultAnyOf
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResultAnyOf) Get() *PlaygroundStrategySchemaResultAnyOf {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf) Set(val *PlaygroundStrategySchemaResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResultAnyOf(val *PlaygroundStrategySchemaResultAnyOf) *NullablePlaygroundStrategySchemaResultAnyOf {
	return &NullablePlaygroundStrategySchemaResultAnyOf{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
