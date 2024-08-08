/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ProjectsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsSchema{}

// ProjectsSchema An overview of all the projects in the Unleash instance
type ProjectsSchema struct {
	// The schema version used to represent the project data.
	Version int32 `json:"version"`
	// A list of projects in the Unleash instance
	Projects []ProjectSchema `json:"projects"`
}

type _ProjectsSchema ProjectsSchema

// NewProjectsSchema instantiates a new ProjectsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsSchema(version int32, projects []ProjectSchema) *ProjectsSchema {
	this := ProjectsSchema{}
	this.Version = version
	this.Projects = projects
	return &this
}

// NewProjectsSchemaWithDefaults instantiates a new ProjectsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsSchemaWithDefaults() *ProjectsSchema {
	this := ProjectsSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *ProjectsSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ProjectsSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ProjectsSchema) SetVersion(v int32) {
	o.Version = v
}

// GetProjects returns the Projects field value
func (o *ProjectsSchema) GetProjects() []ProjectSchema {
	if o == nil {
		var ret []ProjectSchema
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *ProjectsSchema) GetProjectsOk() ([]ProjectSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *ProjectsSchema) SetProjects(v []ProjectSchema) {
	o.Projects = v
}

func (o ProjectsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["projects"] = o.Projects
	return toSerialize, nil
}

func (o *ProjectsSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"projects",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectsSchema := _ProjectsSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectsSchema)

	if err != nil {
		return err
	}

	*o = ProjectsSchema(varProjectsSchema)

	return err
}

type NullableProjectsSchema struct {
	value *ProjectsSchema
	isSet bool
}

func (v NullableProjectsSchema) Get() *ProjectsSchema {
	return v.value
}

func (v *NullableProjectsSchema) Set(val *ProjectsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsSchema(val *ProjectsSchema) *NullableProjectsSchema {
	return &NullableProjectsSchema{value: val, isSet: true}
}

func (v NullableProjectsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
