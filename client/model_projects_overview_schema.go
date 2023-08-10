/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ProjectsOverviewSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsOverviewSchema{}

// ProjectsOverviewSchema struct for ProjectsOverviewSchema
type ProjectsOverviewSchema struct {
	FeatureCount float32         `json:"featureCount"`
	MemberCount  float32         `json:"memberCount"`
	Projects     []ProjectSchema `json:"projects"`
}

// NewProjectsOverviewSchema instantiates a new ProjectsOverviewSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsOverviewSchema(featureCount float32, memberCount float32, projects []ProjectSchema) *ProjectsOverviewSchema {
	this := ProjectsOverviewSchema{}
	this.FeatureCount = featureCount
	this.MemberCount = memberCount
	this.Projects = projects
	return &this
}

// NewProjectsOverviewSchemaWithDefaults instantiates a new ProjectsOverviewSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsOverviewSchemaWithDefaults() *ProjectsOverviewSchema {
	this := ProjectsOverviewSchema{}
	return &this
}

// GetFeatureCount returns the FeatureCount field value
func (o *ProjectsOverviewSchema) GetFeatureCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FeatureCount
}

// GetFeatureCountOk returns a tuple with the FeatureCount field value
// and a boolean to check if the value has been set.
func (o *ProjectsOverviewSchema) GetFeatureCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureCount, true
}

// SetFeatureCount sets field value
func (o *ProjectsOverviewSchema) SetFeatureCount(v float32) {
	o.FeatureCount = v
}

// GetMemberCount returns the MemberCount field value
func (o *ProjectsOverviewSchema) GetMemberCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value
// and a boolean to check if the value has been set.
func (o *ProjectsOverviewSchema) GetMemberCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemberCount, true
}

// SetMemberCount sets field value
func (o *ProjectsOverviewSchema) SetMemberCount(v float32) {
	o.MemberCount = v
}

// GetProjects returns the Projects field value
func (o *ProjectsOverviewSchema) GetProjects() []ProjectSchema {
	if o == nil {
		var ret []ProjectSchema
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ProjectsOverviewSchema) GetProjectsOk() ([]ProjectSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ProjectsOverviewSchema) SetProjects(v []ProjectSchema) {
	o.Projects = v
}

func (o ProjectsOverviewSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsOverviewSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["featureCount"] = o.FeatureCount
	toSerialize["memberCount"] = o.MemberCount
	toSerialize["projects"] = o.Projects
	return toSerialize, nil
}

type NullableProjectsOverviewSchema struct {
	value *ProjectsOverviewSchema
	isSet bool
}

func (v NullableProjectsOverviewSchema) Get() *ProjectsOverviewSchema {
	return v.value
}

func (v *NullableProjectsOverviewSchema) Set(val *ProjectsOverviewSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsOverviewSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsOverviewSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsOverviewSchema(val *ProjectsOverviewSchema) *NullableProjectsOverviewSchema {
	return &NullableProjectsOverviewSchema{value: val, isSet: true}
}

func (v NullableProjectsOverviewSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsOverviewSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
