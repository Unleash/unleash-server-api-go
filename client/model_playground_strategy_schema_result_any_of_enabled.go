/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// PlaygroundStrategySchemaResultAnyOfEnabled Whether this strategy resolves to `false` or if it might resolve to `true`. Because Unleash can't evaluate the strategy, it can't say for certain whether it will be `true`, but if you have failing constraints or segments, it _can_ determine that your strategy would be `false`.
type PlaygroundStrategySchemaResultAnyOfEnabled struct {
	bool   *bool
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PlaygroundStrategySchemaResultAnyOfEnabled) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into bool
	err = json.Unmarshal(data, &dst.bool)
	if err == nil {
		jsonbool, _ := json.Marshal(dst.bool)
		if string(jsonbool) == "{}" { // empty struct
			dst.bool = nil
		} else {
			return nil // data stored in dst.bool, return on the first match
		}
	} else {
		dst.bool = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PlaygroundStrategySchemaResultAnyOfEnabled)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PlaygroundStrategySchemaResultAnyOfEnabled) MarshalJSON() ([]byte, error) {
	if src.bool != nil {
		return json.Marshal(&src.bool)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePlaygroundStrategySchemaResultAnyOfEnabled struct {
	value *PlaygroundStrategySchemaResultAnyOfEnabled
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResultAnyOfEnabled) Get() *PlaygroundStrategySchemaResultAnyOfEnabled {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResultAnyOfEnabled) Set(val *PlaygroundStrategySchemaResultAnyOfEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResultAnyOfEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResultAnyOfEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResultAnyOfEnabled(val *PlaygroundStrategySchemaResultAnyOfEnabled) *NullablePlaygroundStrategySchemaResultAnyOfEnabled {
	return &NullablePlaygroundStrategySchemaResultAnyOfEnabled{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResultAnyOfEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResultAnyOfEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
