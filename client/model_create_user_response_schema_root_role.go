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

// CreateUserResponseSchemaRootRole - Which [root role](https://docs.getunleash.io/reference/rbac#predefined-roles) this user is assigned. Usually a numeric role ID, but can be a string when returning newly created user with an explicit string role.
type CreateUserResponseSchemaRootRole struct {
	Int32  *int32
	String *string
}

// int32AsCreateUserResponseSchemaRootRole is a convenience function that returns int32 wrapped in CreateUserResponseSchemaRootRole
func Int32AsCreateUserResponseSchemaRootRole(v *int32) CreateUserResponseSchemaRootRole {
	return CreateUserResponseSchemaRootRole{
		Int32: v,
	}
}

// stringAsCreateUserResponseSchemaRootRole is a convenience function that returns string wrapped in CreateUserResponseSchemaRootRole
func StringAsCreateUserResponseSchemaRootRole(v *string) CreateUserResponseSchemaRootRole {
	return CreateUserResponseSchemaRootRole{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateUserResponseSchemaRootRole) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(CreateUserResponseSchemaRootRole)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateUserResponseSchemaRootRole)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateUserResponseSchemaRootRole) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateUserResponseSchemaRootRole) GetActualInstance() interface{} {
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

type NullableCreateUserResponseSchemaRootRole struct {
	value *CreateUserResponseSchemaRootRole
	isSet bool
}

func (v NullableCreateUserResponseSchemaRootRole) Get() *CreateUserResponseSchemaRootRole {
	return v.value
}

func (v *NullableCreateUserResponseSchemaRootRole) Set(val *CreateUserResponseSchemaRootRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponseSchemaRootRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponseSchemaRootRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponseSchemaRootRole(val *CreateUserResponseSchemaRootRole) *NullableCreateUserResponseSchemaRootRole {
	return &NullableCreateUserResponseSchemaRootRole{value: val, isSet: true}
}

func (v NullableCreateUserResponseSchemaRootRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponseSchemaRootRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
