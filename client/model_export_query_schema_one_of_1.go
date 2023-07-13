/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ExportQuerySchemaOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportQuerySchemaOneOf1{}

// ExportQuerySchemaOneOf1 struct for ExportQuerySchemaOneOf1
type ExportQuerySchemaOneOf1 struct {
	// Selects features to export by tag. Takes precedence over the features field.
	Tag string `json:"tag"`
}

// NewExportQuerySchemaOneOf1 instantiates a new ExportQuerySchemaOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportQuerySchemaOneOf1(tag string) *ExportQuerySchemaOneOf1 {
	this := ExportQuerySchemaOneOf1{}
	this.Tag = tag
	return &this
}

// NewExportQuerySchemaOneOf1WithDefaults instantiates a new ExportQuerySchemaOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportQuerySchemaOneOf1WithDefaults() *ExportQuerySchemaOneOf1 {
	this := ExportQuerySchemaOneOf1{}
	return &this
}

// GetTag returns the Tag field value
func (o *ExportQuerySchemaOneOf1) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *ExportQuerySchemaOneOf1) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *ExportQuerySchemaOneOf1) SetTag(v string) {
	o.Tag = v
}

func (o ExportQuerySchemaOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportQuerySchemaOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tag"] = o.Tag
	return toSerialize, nil
}

type NullableExportQuerySchemaOneOf1 struct {
	value *ExportQuerySchemaOneOf1
	isSet bool
}

func (v NullableExportQuerySchemaOneOf1) Get() *ExportQuerySchemaOneOf1 {
	return v.value
}

func (v *NullableExportQuerySchemaOneOf1) Set(val *ExportQuerySchemaOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableExportQuerySchemaOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableExportQuerySchemaOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportQuerySchemaOneOf1(val *ExportQuerySchemaOneOf1) *NullableExportQuerySchemaOneOf1 {
	return &NullableExportQuerySchemaOneOf1{value: val, isSet: true}
}

func (v NullableExportQuerySchemaOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportQuerySchemaOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


