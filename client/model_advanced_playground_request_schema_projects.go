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

// AdvancedPlaygroundRequestSchemaProjects - A list of projects to check for toggles in.
type AdvancedPlaygroundRequestSchemaProjects struct {
	ArrayOfString *[]string
	String        *string
}

// []stringAsAdvancedPlaygroundRequestSchemaProjects is a convenience function that returns []string wrapped in AdvancedPlaygroundRequestSchemaProjects
func ArrayOfStringAsAdvancedPlaygroundRequestSchemaProjects(v *[]string) AdvancedPlaygroundRequestSchemaProjects {
	return AdvancedPlaygroundRequestSchemaProjects{
		ArrayOfString: v,
	}
}

// stringAsAdvancedPlaygroundRequestSchemaProjects is a convenience function that returns string wrapped in AdvancedPlaygroundRequestSchemaProjects
func StringAsAdvancedPlaygroundRequestSchemaProjects(v *string) AdvancedPlaygroundRequestSchemaProjects {
	return AdvancedPlaygroundRequestSchemaProjects{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AdvancedPlaygroundRequestSchemaProjects) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfString
	err = newStrictDecoder(data).Decode(&dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfString = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfString = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AdvancedPlaygroundRequestSchemaProjects)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AdvancedPlaygroundRequestSchemaProjects)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AdvancedPlaygroundRequestSchemaProjects) MarshalJSON() ([]byte, error) {
	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AdvancedPlaygroundRequestSchemaProjects) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfString != nil {
		return obj.ArrayOfString
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableAdvancedPlaygroundRequestSchemaProjects struct {
	value *AdvancedPlaygroundRequestSchemaProjects
	isSet bool
}

func (v NullableAdvancedPlaygroundRequestSchemaProjects) Get() *AdvancedPlaygroundRequestSchemaProjects {
	return v.value
}

func (v *NullableAdvancedPlaygroundRequestSchemaProjects) Set(val *AdvancedPlaygroundRequestSchemaProjects) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedPlaygroundRequestSchemaProjects) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedPlaygroundRequestSchemaProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedPlaygroundRequestSchemaProjects(val *AdvancedPlaygroundRequestSchemaProjects) *NullableAdvancedPlaygroundRequestSchemaProjects {
	return &NullableAdvancedPlaygroundRequestSchemaProjects{value: val, isSet: true}
}

func (v NullableAdvancedPlaygroundRequestSchemaProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedPlaygroundRequestSchemaProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
