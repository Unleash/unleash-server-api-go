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

// checks if the ValidateTagTypeSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateTagTypeSchema{}

// ValidateTagTypeSchema The result of validating a tag type.
type ValidateTagTypeSchema struct {
	// Whether or not the tag type is valid.
	Valid   bool          `json:"valid"`
	TagType TagTypeSchema `json:"tagType"`
}

// NewValidateTagTypeSchema instantiates a new ValidateTagTypeSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateTagTypeSchema(valid bool, tagType TagTypeSchema) *ValidateTagTypeSchema {
	this := ValidateTagTypeSchema{}
	this.Valid = valid
	this.TagType = tagType
	return &this
}

// NewValidateTagTypeSchemaWithDefaults instantiates a new ValidateTagTypeSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateTagTypeSchemaWithDefaults() *ValidateTagTypeSchema {
	this := ValidateTagTypeSchema{}
	return &this
}

// GetValid returns the Valid field value
func (o *ValidateTagTypeSchema) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *ValidateTagTypeSchema) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *ValidateTagTypeSchema) SetValid(v bool) {
	o.Valid = v
}

// GetTagType returns the TagType field value
func (o *ValidateTagTypeSchema) GetTagType() TagTypeSchema {
	if o == nil {
		var ret TagTypeSchema
		return ret
	}

	return o.TagType
}

// GetTagTypeOk returns a tuple with the TagType field value
// and a boolean to check if the value has been set.
func (o *ValidateTagTypeSchema) GetTagTypeOk() (*TagTypeSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagType, true
}

// SetTagType sets field value
func (o *ValidateTagTypeSchema) SetTagType(v TagTypeSchema) {
	o.TagType = v
}

func (o ValidateTagTypeSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateTagTypeSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valid"] = o.Valid
	toSerialize["tagType"] = o.TagType
	return toSerialize, nil
}

type NullableValidateTagTypeSchema struct {
	value *ValidateTagTypeSchema
	isSet bool
}

func (v NullableValidateTagTypeSchema) Get() *ValidateTagTypeSchema {
	return v.value
}

func (v *NullableValidateTagTypeSchema) Set(val *ValidateTagTypeSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateTagTypeSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateTagTypeSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateTagTypeSchema(val *ValidateTagTypeSchema) *NullableValidateTagTypeSchema {
	return &NullableValidateTagTypeSchema{value: val, isSet: true}
}

func (v NullableValidateTagTypeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateTagTypeSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
