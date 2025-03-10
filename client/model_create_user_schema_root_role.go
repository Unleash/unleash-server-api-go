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

// CreateUserSchemaRootRole - The role to assign to the user. Can be either the role's ID or its unique name.
type CreateUserSchemaRootRole struct {
	Int32  *int32
	String *string
}

// int32AsCreateUserSchemaRootRole is a convenience function that returns int32 wrapped in CreateUserSchemaRootRole
func Int32AsCreateUserSchemaRootRole(v *int32) CreateUserSchemaRootRole {
	return CreateUserSchemaRootRole{
		Int32: v,
	}
}

// stringAsCreateUserSchemaRootRole is a convenience function that returns string wrapped in CreateUserSchemaRootRole
func StringAsCreateUserSchemaRootRole(v *string) CreateUserSchemaRootRole {
	return CreateUserSchemaRootRole{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateUserSchemaRootRole) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			if err = validator.Validate(dst.Int32); err != nil {
				dst.Int32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateUserSchemaRootRole)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateUserSchemaRootRole)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateUserSchemaRootRole) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateUserSchemaRootRole) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateUserSchemaRootRole struct {
	value *CreateUserSchemaRootRole
	isSet bool
}

func (v NullableCreateUserSchemaRootRole) Get() *CreateUserSchemaRootRole {
	return v.value
}

func (v *NullableCreateUserSchemaRootRole) Set(val *CreateUserSchemaRootRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserSchemaRootRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserSchemaRootRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserSchemaRootRole(val *CreateUserSchemaRootRole) *NullableCreateUserSchemaRootRole {
	return &NullableCreateUserSchemaRootRole{value: val, isSet: true}
}

func (v NullableCreateUserSchemaRootRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserSchemaRootRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
