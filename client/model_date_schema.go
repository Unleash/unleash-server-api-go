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
	"time"
)

// DateSchema - struct for DateSchema
type DateSchema struct {
	Float32  *float32
	TimeTime *time.Time
}

// float32AsDateSchema is a convenience function that returns float32 wrapped in DateSchema
func Float32AsDateSchema(v *float32) DateSchema {
	return DateSchema{
		Float32: v,
	}
}

// time.TimeAsDateSchema is a convenience function that returns time.Time wrapped in DateSchema
func TimeTimeAsDateSchema(v *time.Time) DateSchema {
	return DateSchema{
		TimeTime: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DateSchema) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into TimeTime
	err = newStrictDecoder(data).Decode(&dst.TimeTime)
	if err == nil {
		jsonTimeTime, _ := json.Marshal(dst.TimeTime)
		if string(jsonTimeTime) == "{}" { // empty struct
			dst.TimeTime = nil
		} else {
			match++
		}
	} else {
		dst.TimeTime = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.TimeTime = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DateSchema)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DateSchema)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DateSchema) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DateSchema) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.TimeTime != nil {
		return obj.TimeTime
	}

	// all schemas are nil
	return nil
}

type NullableDateSchema struct {
	value *DateSchema
	isSet bool
}

func (v NullableDateSchema) Get() *DateSchema {
	return v.value
}

func (v *NullableDateSchema) Set(val *DateSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableDateSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableDateSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateSchema(val *DateSchema) *NullableDateSchema {
	return &NullableDateSchema{value: val, isSet: true}
}

func (v NullableDateSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
