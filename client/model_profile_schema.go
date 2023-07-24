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

// checks if the ProfileSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSchema{}

// ProfileSchema struct for ProfileSchema
type ProfileSchema struct {
	RootRole RoleSchema      `json:"rootRole"`
	Projects []string        `json:"projects"`
	Features []FeatureSchema `json:"features"`
}

// NewProfileSchema instantiates a new ProfileSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSchema(rootRole RoleSchema, projects []string, features []FeatureSchema) *ProfileSchema {
	this := ProfileSchema{}
	this.RootRole = rootRole
	this.Projects = projects
	this.Features = features
	return &this
}

// NewProfileSchemaWithDefaults instantiates a new ProfileSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSchemaWithDefaults() *ProfileSchema {
	this := ProfileSchema{}
	return &this
}

// GetRootRole returns the RootRole field value
func (o *ProfileSchema) GetRootRole() RoleSchema {
	if o == nil {
		var ret RoleSchema
		return ret
	}

	return o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value
// and a boolean to check if the value has been set.
func (o *ProfileSchema) GetRootRoleOk() (*RoleSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootRole, true
}

// SetRootRole sets field value
func (o *ProfileSchema) SetRootRole(v RoleSchema) {
	o.RootRole = v
}

// GetProjects returns the Projects field value
func (o *ProfileSchema) GetProjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ProfileSchema) GetProjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ProfileSchema) SetProjects(v []string) {
	o.Projects = v
}

// GetFeatures returns the Features field value
func (o *ProfileSchema) GetFeatures() []FeatureSchema {
	if o == nil {
		var ret []FeatureSchema
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ProfileSchema) GetFeaturesOk() ([]FeatureSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ProfileSchema) SetFeatures(v []FeatureSchema) {
	o.Features = v
}

func (o ProfileSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rootRole"] = o.RootRole
	toSerialize["projects"] = o.Projects
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

type NullableProfileSchema struct {
	value *ProfileSchema
	isSet bool
}

func (v NullableProfileSchema) Get() *ProfileSchema {
	return v.value
}

func (v *NullableProfileSchema) Set(val *ProfileSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSchema(val *ProfileSchema) *NullableProfileSchema {
	return &NullableProfileSchema{value: val, isSet: true}
}

func (v NullableProfileSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}