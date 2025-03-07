/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// OidcSettingsSchema - Settings for configuring OpenID Connect as a login provider for Unleash
type OidcSettingsSchema struct {
	OidcSettingsSchemaOneOf  *OidcSettingsSchemaOneOf
	OidcSettingsSchemaOneOf1 *OidcSettingsSchemaOneOf1
}

// OidcSettingsSchemaOneOfAsOidcSettingsSchema is a convenience function that returns OidcSettingsSchemaOneOf wrapped in OidcSettingsSchema
func OidcSettingsSchemaOneOfAsOidcSettingsSchema(v *OidcSettingsSchemaOneOf) OidcSettingsSchema {
	return OidcSettingsSchema{
		OidcSettingsSchemaOneOf: v,
	}
}

// OidcSettingsSchemaOneOf1AsOidcSettingsSchema is a convenience function that returns OidcSettingsSchemaOneOf1 wrapped in OidcSettingsSchema
func OidcSettingsSchemaOneOf1AsOidcSettingsSchema(v *OidcSettingsSchemaOneOf1) OidcSettingsSchema {
	return OidcSettingsSchema{
		OidcSettingsSchemaOneOf1: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OidcSettingsSchema) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into OidcSettingsSchemaOneOf
	err = newStrictDecoder(data).Decode(&dst.OidcSettingsSchemaOneOf)
	if err == nil {
		jsonOidcSettingsSchemaOneOf, _ := json.Marshal(dst.OidcSettingsSchemaOneOf)
		if string(jsonOidcSettingsSchemaOneOf) == "{}" { // empty struct
			dst.OidcSettingsSchemaOneOf = nil
		} else {
			if err = validator.Validate(dst.OidcSettingsSchemaOneOf); err != nil {
				dst.OidcSettingsSchemaOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.OidcSettingsSchemaOneOf = nil
	}

	// try to unmarshal data into OidcSettingsSchemaOneOf1
	err = newStrictDecoder(data).Decode(&dst.OidcSettingsSchemaOneOf1)
	if err == nil {
		jsonOidcSettingsSchemaOneOf1, _ := json.Marshal(dst.OidcSettingsSchemaOneOf1)
		if string(jsonOidcSettingsSchemaOneOf1) == "{}" { // empty struct
			dst.OidcSettingsSchemaOneOf1 = nil
		} else {
			if err = validator.Validate(dst.OidcSettingsSchemaOneOf1); err != nil {
				dst.OidcSettingsSchemaOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.OidcSettingsSchemaOneOf1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.OidcSettingsSchemaOneOf = nil
		dst.OidcSettingsSchemaOneOf1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OidcSettingsSchema)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OidcSettingsSchema)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OidcSettingsSchema) MarshalJSON() ([]byte, error) {
	if src.OidcSettingsSchemaOneOf != nil {
		return json.Marshal(&src.OidcSettingsSchemaOneOf)
	}

	if src.OidcSettingsSchemaOneOf1 != nil {
		return json.Marshal(&src.OidcSettingsSchemaOneOf1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OidcSettingsSchema) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OidcSettingsSchemaOneOf != nil {
		return obj.OidcSettingsSchemaOneOf
	}

	if obj.OidcSettingsSchemaOneOf1 != nil {
		return obj.OidcSettingsSchemaOneOf1
	}

	// all schemas are nil
	return nil
}

type NullableOidcSettingsSchema struct {
	value *OidcSettingsSchema
	isSet bool
}

func (v NullableOidcSettingsSchema) Get() *OidcSettingsSchema {
	return v.value
}

func (v *NullableOidcSettingsSchema) Set(val *OidcSettingsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcSettingsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcSettingsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcSettingsSchema(val *OidcSettingsSchema) *NullableOidcSettingsSchema {
	return &NullableOidcSettingsSchema{value: val, isSet: true}
}

func (v NullableOidcSettingsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcSettingsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
