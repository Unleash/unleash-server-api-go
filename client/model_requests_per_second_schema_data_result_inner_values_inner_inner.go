/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// RequestsPerSecondSchemaDataResultInnerValuesInnerInner struct for RequestsPerSecondSchemaDataResultInnerValuesInnerInner
type RequestsPerSecondSchemaDataResultInnerValuesInnerInner struct {
	float32 *float32
	string  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestsPerSecondSchemaDataResultInnerValuesInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.float32)
	if err == nil {
		jsonfloat32, _ := json.Marshal(dst.float32)
		if string(jsonfloat32) == "{}" { // empty struct
			dst.float32 = nil
		} else {
			return nil // data stored in dst.float32, return on the first match
		}
	} else {
		dst.float32 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestsPerSecondSchemaDataResultInnerValuesInnerInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestsPerSecondSchemaDataResultInnerValuesInnerInner) MarshalJSON() ([]byte, error) {
	if src.float32 != nil {
		return json.Marshal(&src.float32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner struct {
	value *RequestsPerSecondSchemaDataResultInnerValuesInnerInner
	isSet bool
}

func (v NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) Get() *RequestsPerSecondSchemaDataResultInnerValuesInnerInner {
	return v.value
}

func (v *NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) Set(val *RequestsPerSecondSchemaDataResultInnerValuesInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner(val *RequestsPerSecondSchemaDataResultInnerValuesInnerInner) *NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner {
	return &NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner{value: val, isSet: true}
}

func (v NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPerSecondSchemaDataResultInnerValuesInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
