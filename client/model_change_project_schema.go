/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ChangeProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeProjectSchema{}

// ChangeProjectSchema Data required to move a feature toggle to a project.
type ChangeProjectSchema struct {
	// The project to move the feature toggle to.
	NewProjectId string `json:"newProjectId"`
}

// NewChangeProjectSchema instantiates a new ChangeProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeProjectSchema(newProjectId string) *ChangeProjectSchema {
	this := ChangeProjectSchema{}
	this.NewProjectId = newProjectId
	return &this
}

// NewChangeProjectSchemaWithDefaults instantiates a new ChangeProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeProjectSchemaWithDefaults() *ChangeProjectSchema {
	this := ChangeProjectSchema{}
	return &this
}

// GetNewProjectId returns the NewProjectId field value
func (o *ChangeProjectSchema) GetNewProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewProjectId
}

// GetNewProjectIdOk returns a tuple with the NewProjectId field value
// and a boolean to check if the value has been set.
func (o *ChangeProjectSchema) GetNewProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewProjectId, true
}

// SetNewProjectId sets field value
func (o *ChangeProjectSchema) SetNewProjectId(v string) {
	o.NewProjectId = v
}

func (o ChangeProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newProjectId"] = o.NewProjectId
	return toSerialize, nil
}

type NullableChangeProjectSchema struct {
	value *ChangeProjectSchema
	isSet bool
}

func (v NullableChangeProjectSchema) Get() *ChangeProjectSchema {
	return v.value
}

func (v *NullableChangeProjectSchema) Set(val *ChangeProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeProjectSchema(val *ChangeProjectSchema) *NullableChangeProjectSchema {
	return &NullableChangeProjectSchema{value: val, isSet: true}
}

func (v NullableChangeProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
