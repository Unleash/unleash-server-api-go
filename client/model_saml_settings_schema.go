/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.12+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// SamlSettingsSchema - Settings used to authenticate via SAML
type SamlSettingsSchema struct {
	SamlSettingsSchemaOneOf  *SamlSettingsSchemaOneOf
	SamlSettingsSchemaOneOf1 *SamlSettingsSchemaOneOf1
}

// SamlSettingsSchemaOneOfAsSamlSettingsSchema is a convenience function that returns SamlSettingsSchemaOneOf wrapped in SamlSettingsSchema
func SamlSettingsSchemaOneOfAsSamlSettingsSchema(v *SamlSettingsSchemaOneOf) SamlSettingsSchema {
	return SamlSettingsSchema{
		SamlSettingsSchemaOneOf: v,
	}
}

// SamlSettingsSchemaOneOf1AsSamlSettingsSchema is a convenience function that returns SamlSettingsSchemaOneOf1 wrapped in SamlSettingsSchema
func SamlSettingsSchemaOneOf1AsSamlSettingsSchema(v *SamlSettingsSchemaOneOf1) SamlSettingsSchema {
	return SamlSettingsSchema{
		SamlSettingsSchemaOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SamlSettingsSchema) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SamlSettingsSchemaOneOf
	err = newStrictDecoder(data).Decode(&dst.SamlSettingsSchemaOneOf)
	if err == nil {
		jsonSamlSettingsSchemaOneOf, _ := json.Marshal(dst.SamlSettingsSchemaOneOf)
		if string(jsonSamlSettingsSchemaOneOf) == "{}" { // empty struct
			dst.SamlSettingsSchemaOneOf = nil
		} else {
			if err = validator.Validate(dst.SamlSettingsSchemaOneOf); err != nil {
				dst.SamlSettingsSchemaOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.SamlSettingsSchemaOneOf = nil
	}

	// try to unmarshal data into SamlSettingsSchemaOneOf1
	err = newStrictDecoder(data).Decode(&dst.SamlSettingsSchemaOneOf1)
	if err == nil {
		jsonSamlSettingsSchemaOneOf1, _ := json.Marshal(dst.SamlSettingsSchemaOneOf1)
		if string(jsonSamlSettingsSchemaOneOf1) == "{}" { // empty struct
			dst.SamlSettingsSchemaOneOf1 = nil
		} else {
			if err = validator.Validate(dst.SamlSettingsSchemaOneOf1); err != nil {
				dst.SamlSettingsSchemaOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.SamlSettingsSchemaOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SamlSettingsSchemaOneOf = nil
		dst.SamlSettingsSchemaOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SamlSettingsSchema)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SamlSettingsSchema)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SamlSettingsSchema) MarshalJSON() ([]byte, error) {
	if src.SamlSettingsSchemaOneOf != nil {
		return json.Marshal(&src.SamlSettingsSchemaOneOf)
	}

	if src.SamlSettingsSchemaOneOf1 != nil {
		return json.Marshal(&src.SamlSettingsSchemaOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SamlSettingsSchema) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SamlSettingsSchemaOneOf != nil {
		return obj.SamlSettingsSchemaOneOf
	}

	if obj.SamlSettingsSchemaOneOf1 != nil {
		return obj.SamlSettingsSchemaOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableSamlSettingsSchema struct {
	value *SamlSettingsSchema
	isSet bool
}

func (v NullableSamlSettingsSchema) Get() *SamlSettingsSchema {
	return v.value
}

func (v *NullableSamlSettingsSchema) Set(val *SamlSettingsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSettingsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSettingsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSettingsSchema(val *SamlSettingsSchema) *NullableSamlSettingsSchema {
	return &NullableSamlSettingsSchema{value: val, isSet: true}
}

func (v NullableSamlSettingsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSettingsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
