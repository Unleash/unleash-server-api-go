/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ApiTokensSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTokensSchema{}

// ApiTokensSchema An object with [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys)
type ApiTokensSchema struct {
	// A list of Unleash API tokens.
	Tokens []ApiTokenSchema `json:"tokens"`
}

type _ApiTokensSchema ApiTokensSchema

// NewApiTokensSchema instantiates a new ApiTokensSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokensSchema(tokens []ApiTokenSchema) *ApiTokensSchema {
	this := ApiTokensSchema{}
	this.Tokens = tokens
	return &this
}

// NewApiTokensSchemaWithDefaults instantiates a new ApiTokensSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokensSchemaWithDefaults() *ApiTokensSchema {
	this := ApiTokensSchema{}
	return &this
}

// GetTokens returns the Tokens field value
func (o *ApiTokensSchema) GetTokens() []ApiTokenSchema {
	if o == nil {
		var ret []ApiTokenSchema
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *ApiTokensSchema) GetTokensOk() ([]ApiTokenSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *ApiTokensSchema) SetTokens(v []ApiTokenSchema) {
	o.Tokens = v
}

func (o ApiTokensSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTokensSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tokens"] = o.Tokens
	return toSerialize, nil
}

func (o *ApiTokensSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tokens",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiTokensSchema := _ApiTokensSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTokensSchema)

	if err != nil {
		return err
	}

	*o = ApiTokensSchema(varApiTokensSchema)

	return err
}

type NullableApiTokensSchema struct {
	value *ApiTokensSchema
	isSet bool
}

func (v NullableApiTokensSchema) Get() *ApiTokensSchema {
	return v.value
}

func (v *NullableApiTokensSchema) Set(val *ApiTokensSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokensSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokensSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokensSchema(val *ApiTokensSchema) *NullableApiTokensSchema {
	return &NullableApiTokensSchema{value: val, isSet: true}
}

func (v NullableApiTokensSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokensSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
