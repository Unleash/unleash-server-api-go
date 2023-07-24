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

// checks if the TagTypesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagTypesSchema{}

// TagTypesSchema struct for TagTypesSchema
type TagTypesSchema struct {
	Version  int32           `json:"version"`
	TagTypes []TagTypeSchema `json:"tagTypes"`
}

// NewTagTypesSchema instantiates a new TagTypesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagTypesSchema(version int32, tagTypes []TagTypeSchema) *TagTypesSchema {
	this := TagTypesSchema{}
	this.Version = version
	this.TagTypes = tagTypes
	return &this
}

// NewTagTypesSchemaWithDefaults instantiates a new TagTypesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagTypesSchemaWithDefaults() *TagTypesSchema {
	this := TagTypesSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *TagTypesSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagTypesSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TagTypesSchema) SetVersion(v int32) {
	o.Version = v
}

// GetTagTypes returns the TagTypes field value
func (o *TagTypesSchema) GetTagTypes() []TagTypeSchema {
	if o == nil {
		var ret []TagTypeSchema
		return ret
	}

	return o.TagTypes
}

// GetTagTypesOk returns a tuple with the TagTypes field value
// and a boolean to check if the value has been set.
func (o *TagTypesSchema) GetTagTypesOk() ([]TagTypeSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagTypes, true
}

// SetTagTypes sets field value
func (o *TagTypesSchema) SetTagTypes(v []TagTypeSchema) {
	o.TagTypes = v
}

func (o TagTypesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagTypesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["tagTypes"] = o.TagTypes
	return toSerialize, nil
}

type NullableTagTypesSchema struct {
	value *TagTypesSchema
	isSet bool
}

func (v NullableTagTypesSchema) Get() *TagTypesSchema {
	return v.value
}

func (v *NullableTagTypesSchema) Set(val *TagTypesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTagTypesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTagTypesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagTypesSchema(val *TagTypesSchema) *NullableTagTypesSchema {
	return &NullableTagTypesSchema{value: val, isSet: true}
}

func (v NullableTagTypesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagTypesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
