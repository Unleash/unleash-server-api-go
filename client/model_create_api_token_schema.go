/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// CreateApiTokenSchema - The data required to create an [Unleash API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys).
type CreateApiTokenSchema struct {
	CreateApiTokenSchemaOneOf  *CreateApiTokenSchemaOneOf
	CreateApiTokenSchemaOneOf1 *CreateApiTokenSchemaOneOf1
	CreateApiTokenSchemaOneOf2 *CreateApiTokenSchemaOneOf2
	CreateApiTokenSchemaOneOf3 *CreateApiTokenSchemaOneOf3
}

// CreateApiTokenSchemaOneOfAsCreateApiTokenSchema is a convenience function that returns CreateApiTokenSchemaOneOf wrapped in CreateApiTokenSchema
func CreateApiTokenSchemaOneOfAsCreateApiTokenSchema(v *CreateApiTokenSchemaOneOf) CreateApiTokenSchema {
	return CreateApiTokenSchema{
		CreateApiTokenSchemaOneOf: v,
	}
}

// CreateApiTokenSchemaOneOf1AsCreateApiTokenSchema is a convenience function that returns CreateApiTokenSchemaOneOf1 wrapped in CreateApiTokenSchema
func CreateApiTokenSchemaOneOf1AsCreateApiTokenSchema(v *CreateApiTokenSchemaOneOf1) CreateApiTokenSchema {
	return CreateApiTokenSchema{
		CreateApiTokenSchemaOneOf1: v,
	}
}

// CreateApiTokenSchemaOneOf2AsCreateApiTokenSchema is a convenience function that returns CreateApiTokenSchemaOneOf2 wrapped in CreateApiTokenSchema
func CreateApiTokenSchemaOneOf2AsCreateApiTokenSchema(v *CreateApiTokenSchemaOneOf2) CreateApiTokenSchema {
	return CreateApiTokenSchema{
		CreateApiTokenSchemaOneOf2: v,
	}
}

// CreateApiTokenSchemaOneOf3AsCreateApiTokenSchema is a convenience function that returns CreateApiTokenSchemaOneOf3 wrapped in CreateApiTokenSchema
func CreateApiTokenSchemaOneOf3AsCreateApiTokenSchema(v *CreateApiTokenSchemaOneOf3) CreateApiTokenSchema {
	return CreateApiTokenSchema{
		CreateApiTokenSchemaOneOf3: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateApiTokenSchema) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateApiTokenSchemaOneOf
	err = newStrictDecoder(data).Decode(&dst.CreateApiTokenSchemaOneOf)
	if err == nil {
		jsonCreateApiTokenSchemaOneOf, _ := json.Marshal(dst.CreateApiTokenSchemaOneOf)
		if string(jsonCreateApiTokenSchemaOneOf) == "{}" { // empty struct
			dst.CreateApiTokenSchemaOneOf = nil
		} else {
			match++
		}
	} else {
		dst.CreateApiTokenSchemaOneOf = nil
	}

	// try to unmarshal data into CreateApiTokenSchemaOneOf1
	err = newStrictDecoder(data).Decode(&dst.CreateApiTokenSchemaOneOf1)
	if err == nil {
		jsonCreateApiTokenSchemaOneOf1, _ := json.Marshal(dst.CreateApiTokenSchemaOneOf1)
		if string(jsonCreateApiTokenSchemaOneOf1) == "{}" { // empty struct
			dst.CreateApiTokenSchemaOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.CreateApiTokenSchemaOneOf1 = nil
	}

	// try to unmarshal data into CreateApiTokenSchemaOneOf2
	err = newStrictDecoder(data).Decode(&dst.CreateApiTokenSchemaOneOf2)
	if err == nil {
		jsonCreateApiTokenSchemaOneOf2, _ := json.Marshal(dst.CreateApiTokenSchemaOneOf2)
		if string(jsonCreateApiTokenSchemaOneOf2) == "{}" { // empty struct
			dst.CreateApiTokenSchemaOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.CreateApiTokenSchemaOneOf2 = nil
	}

	// try to unmarshal data into CreateApiTokenSchemaOneOf3
	err = newStrictDecoder(data).Decode(&dst.CreateApiTokenSchemaOneOf3)
	if err == nil {
		jsonCreateApiTokenSchemaOneOf3, _ := json.Marshal(dst.CreateApiTokenSchemaOneOf3)
		if string(jsonCreateApiTokenSchemaOneOf3) == "{}" { // empty struct
			dst.CreateApiTokenSchemaOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.CreateApiTokenSchemaOneOf3 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateApiTokenSchemaOneOf = nil
		dst.CreateApiTokenSchemaOneOf1 = nil
		dst.CreateApiTokenSchemaOneOf2 = nil
		dst.CreateApiTokenSchemaOneOf3 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateApiTokenSchema)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateApiTokenSchema)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateApiTokenSchema) MarshalJSON() ([]byte, error) {
	if src.CreateApiTokenSchemaOneOf != nil {
		return json.Marshal(&src.CreateApiTokenSchemaOneOf)
	}

	if src.CreateApiTokenSchemaOneOf1 != nil {
		return json.Marshal(&src.CreateApiTokenSchemaOneOf1)
	}

	if src.CreateApiTokenSchemaOneOf2 != nil {
		return json.Marshal(&src.CreateApiTokenSchemaOneOf2)
	}

	if src.CreateApiTokenSchemaOneOf3 != nil {
		return json.Marshal(&src.CreateApiTokenSchemaOneOf3)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateApiTokenSchema) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateApiTokenSchemaOneOf != nil {
		return obj.CreateApiTokenSchemaOneOf
	}

	if obj.CreateApiTokenSchemaOneOf1 != nil {
		return obj.CreateApiTokenSchemaOneOf1
	}

	if obj.CreateApiTokenSchemaOneOf2 != nil {
		return obj.CreateApiTokenSchemaOneOf2
	}

	if obj.CreateApiTokenSchemaOneOf3 != nil {
		return obj.CreateApiTokenSchemaOneOf3
	}

	// all schemas are nil
	return nil
}

type NullableCreateApiTokenSchema struct {
	value *CreateApiTokenSchema
	isSet bool
}

func (v NullableCreateApiTokenSchema) Get() *CreateApiTokenSchema {
	return v.value
}

func (v *NullableCreateApiTokenSchema) Set(val *CreateApiTokenSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiTokenSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiTokenSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiTokenSchema(val *CreateApiTokenSchema) *NullableCreateApiTokenSchema {
	return &NullableCreateApiTokenSchema{value: val, isSet: true}
}

func (v NullableCreateApiTokenSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiTokenSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
