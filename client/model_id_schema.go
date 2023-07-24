/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IdSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdSchema{}

// IdSchema struct for IdSchema
type IdSchema struct {
	Id string `json:"id"`
}

// NewIdSchema instantiates a new IdSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdSchema(id string) *IdSchema {
	this := IdSchema{}
	this.Id = id
	return &this
}

// NewIdSchemaWithDefaults instantiates a new IdSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdSchemaWithDefaults() *IdSchema {
	this := IdSchema{}
	return &this
}

// GetId returns the Id field value
func (o *IdSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdSchema) SetId(v string) {
	o.Id = v
}

func (o IdSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableIdSchema struct {
	value *IdSchema
	isSet bool
}

func (v NullableIdSchema) Get() *IdSchema {
	return v.value
}

func (v *NullableIdSchema) Set(val *IdSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableIdSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableIdSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdSchema(val *IdSchema) *NullableIdSchema {
	return &NullableIdSchema{value: val, isSet: true}
}

func (v NullableIdSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
