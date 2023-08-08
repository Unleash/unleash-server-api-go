/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ValidateProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateProjectSchema{}

// ValidateProjectSchema struct for ValidateProjectSchema
type ValidateProjectSchema struct {
	Id string `json:"id"`
}

// NewValidateProjectSchema instantiates a new ValidateProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateProjectSchema(id string) *ValidateProjectSchema {
	this := ValidateProjectSchema{}
	this.Id = id
	return &this
}

// NewValidateProjectSchemaWithDefaults instantiates a new ValidateProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateProjectSchemaWithDefaults() *ValidateProjectSchema {
	this := ValidateProjectSchema{}
	return &this
}

// GetId returns the Id field value
func (o *ValidateProjectSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValidateProjectSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValidateProjectSchema) SetId(v string) {
	o.Id = v
}

func (o ValidateProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableValidateProjectSchema struct {
	value *ValidateProjectSchema
	isSet bool
}

func (v NullableValidateProjectSchema) Get() *ValidateProjectSchema {
	return v.value
}

func (v *NullableValidateProjectSchema) Set(val *ValidateProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateProjectSchema(val *ValidateProjectSchema) *NullableValidateProjectSchema {
	return &NullableValidateProjectSchema{value: val, isSet: true}
}

func (v NullableValidateProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
