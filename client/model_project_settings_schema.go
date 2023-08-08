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

// checks if the ProjectSettingsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSettingsSchema{}

// ProjectSettingsSchema struct for ProjectSettingsSchema
type ProjectSettingsSchema struct {
	// Default stickiness for project
	DefaultStickiness NullableString `json:"defaultStickiness"`
	Mode              NullableString `json:"mode"`
}

// NewProjectSettingsSchema instantiates a new ProjectSettingsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSettingsSchema(defaultStickiness NullableString, mode NullableString) *ProjectSettingsSchema {
	this := ProjectSettingsSchema{}
	this.DefaultStickiness = defaultStickiness
	this.Mode = mode
	return &this
}

// NewProjectSettingsSchemaWithDefaults instantiates a new ProjectSettingsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSettingsSchemaWithDefaults() *ProjectSettingsSchema {
	this := ProjectSettingsSchema{}
	return &this
}

// GetDefaultStickiness returns the DefaultStickiness field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectSettingsSchema) GetDefaultStickiness() string {
	if o == nil || o.DefaultStickiness.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultStickiness.Get()
}

// GetDefaultStickinessOk returns a tuple with the DefaultStickiness field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSettingsSchema) GetDefaultStickinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultStickiness.Get(), o.DefaultStickiness.IsSet()
}

// SetDefaultStickiness sets field value
func (o *ProjectSettingsSchema) SetDefaultStickiness(v string) {
	o.DefaultStickiness.Set(&v)
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectSettingsSchema) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSettingsSchema) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// SetMode sets field value
func (o *ProjectSettingsSchema) SetMode(v string) {
	o.Mode.Set(&v)
}

func (o ProjectSettingsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSettingsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultStickiness"] = o.DefaultStickiness.Get()
	toSerialize["mode"] = o.Mode.Get()
	return toSerialize, nil
}

type NullableProjectSettingsSchema struct {
	value *ProjectSettingsSchema
	isSet bool
}

func (v NullableProjectSettingsSchema) Get() *ProjectSettingsSchema {
	return v.value
}

func (v *NullableProjectSettingsSchema) Set(val *ProjectSettingsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSettingsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSettingsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSettingsSchema(val *ProjectSettingsSchema) *NullableProjectSettingsSchema {
	return &NullableProjectSettingsSchema{value: val, isSet: true}
}

func (v NullableProjectSettingsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSettingsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
