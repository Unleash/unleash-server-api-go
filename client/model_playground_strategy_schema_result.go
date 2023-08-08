/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// PlaygroundStrategySchemaResult The strategy's evaluation result. If the strategy is a custom strategy that Unleash can't evaluate, `evaluationStatus` will be `unknown`. Otherwise, it will be `true` or `false`
type PlaygroundStrategySchemaResult struct {
	PlaygroundStrategySchemaResultAnyOf  *PlaygroundStrategySchemaResultAnyOf
	PlaygroundStrategySchemaResultAnyOf1 *PlaygroundStrategySchemaResultAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PlaygroundStrategySchemaResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PlaygroundStrategySchemaResultAnyOf
	err = json.Unmarshal(data, &dst.PlaygroundStrategySchemaResultAnyOf)
	if err == nil {
		jsonPlaygroundStrategySchemaResultAnyOf, _ := json.Marshal(dst.PlaygroundStrategySchemaResultAnyOf)
		if string(jsonPlaygroundStrategySchemaResultAnyOf) == "{}" { // empty struct
			dst.PlaygroundStrategySchemaResultAnyOf = nil
		} else {
			return nil // data stored in dst.PlaygroundStrategySchemaResultAnyOf, return on the first match
		}
	} else {
		dst.PlaygroundStrategySchemaResultAnyOf = nil
	}

	// try to unmarshal JSON data into PlaygroundStrategySchemaResultAnyOf1
	err = json.Unmarshal(data, &dst.PlaygroundStrategySchemaResultAnyOf1)
	if err == nil {
		jsonPlaygroundStrategySchemaResultAnyOf1, _ := json.Marshal(dst.PlaygroundStrategySchemaResultAnyOf1)
		if string(jsonPlaygroundStrategySchemaResultAnyOf1) == "{}" { // empty struct
			dst.PlaygroundStrategySchemaResultAnyOf1 = nil
		} else {
			return nil // data stored in dst.PlaygroundStrategySchemaResultAnyOf1, return on the first match
		}
	} else {
		dst.PlaygroundStrategySchemaResultAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PlaygroundStrategySchemaResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PlaygroundStrategySchemaResult) MarshalJSON() ([]byte, error) {
	if src.PlaygroundStrategySchemaResultAnyOf != nil {
		return json.Marshal(&src.PlaygroundStrategySchemaResultAnyOf)
	}

	if src.PlaygroundStrategySchemaResultAnyOf1 != nil {
		return json.Marshal(&src.PlaygroundStrategySchemaResultAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePlaygroundStrategySchemaResult struct {
	value *PlaygroundStrategySchemaResult
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResult) Get() *PlaygroundStrategySchemaResult {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResult) Set(val *PlaygroundStrategySchemaResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResult(val *PlaygroundStrategySchemaResult) *NullablePlaygroundStrategySchemaResult {
	return &NullablePlaygroundStrategySchemaResult{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
