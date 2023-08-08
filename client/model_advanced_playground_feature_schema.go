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

// checks if the AdvancedPlaygroundFeatureSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedPlaygroundFeatureSchema{}

// AdvancedPlaygroundFeatureSchema A simplified feature toggle model intended for the Unleash playground.
type AdvancedPlaygroundFeatureSchema struct {
	// The feature's name.
	Name string `json:"name"`
	// The ID of the project that contains this feature.
	ProjectId string `json:"projectId"`
	// The lists of features that have been evaluated grouped by environment.
	Environments map[string][]AdvancedPlaygroundEnvironmentFeatureSchema `json:"environments"`
}

// NewAdvancedPlaygroundFeatureSchema instantiates a new AdvancedPlaygroundFeatureSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedPlaygroundFeatureSchema(name string, projectId string, environments map[string][]AdvancedPlaygroundEnvironmentFeatureSchema) *AdvancedPlaygroundFeatureSchema {
	this := AdvancedPlaygroundFeatureSchema{}
	this.Name = name
	this.ProjectId = projectId
	this.Environments = environments
	return &this
}

// NewAdvancedPlaygroundFeatureSchemaWithDefaults instantiates a new AdvancedPlaygroundFeatureSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedPlaygroundFeatureSchemaWithDefaults() *AdvancedPlaygroundFeatureSchema {
	this := AdvancedPlaygroundFeatureSchema{}
	return &this
}

// GetName returns the Name field value
func (o *AdvancedPlaygroundFeatureSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundFeatureSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdvancedPlaygroundFeatureSchema) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *AdvancedPlaygroundFeatureSchema) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundFeatureSchema) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AdvancedPlaygroundFeatureSchema) SetProjectId(v string) {
	o.ProjectId = v
}

// GetEnvironments returns the Environments field value
func (o *AdvancedPlaygroundFeatureSchema) GetEnvironments() map[string][]AdvancedPlaygroundEnvironmentFeatureSchema {
	if o == nil {
		var ret map[string][]AdvancedPlaygroundEnvironmentFeatureSchema
		return ret
	}

	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *AdvancedPlaygroundFeatureSchema) GetEnvironmentsOk() (*map[string][]AdvancedPlaygroundEnvironmentFeatureSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value
func (o *AdvancedPlaygroundFeatureSchema) SetEnvironments(v map[string][]AdvancedPlaygroundEnvironmentFeatureSchema) {
	o.Environments = v
}

func (o AdvancedPlaygroundFeatureSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedPlaygroundFeatureSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	toSerialize["environments"] = o.Environments
	return toSerialize, nil
}

type NullableAdvancedPlaygroundFeatureSchema struct {
	value *AdvancedPlaygroundFeatureSchema
	isSet bool
}

func (v NullableAdvancedPlaygroundFeatureSchema) Get() *AdvancedPlaygroundFeatureSchema {
	return v.value
}

func (v *NullableAdvancedPlaygroundFeatureSchema) Set(val *AdvancedPlaygroundFeatureSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedPlaygroundFeatureSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedPlaygroundFeatureSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedPlaygroundFeatureSchema(val *AdvancedPlaygroundFeatureSchema) *NullableAdvancedPlaygroundFeatureSchema {
	return &NullableAdvancedPlaygroundFeatureSchema{value: val, isSet: true}
}

func (v NullableAdvancedPlaygroundFeatureSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedPlaygroundFeatureSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
