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

// checks if the ProjectEnvironmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEnvironmentSchema{}

// ProjectEnvironmentSchema Add an environment to a project, optionally also sets if change requests are enabled for this environment on the project
type ProjectEnvironmentSchema struct {
	// The environment to add to the project
	Environment string `json:"environment"`
	// Whether change requests should be enabled or for this environment on the project or not
	ChangeRequestsEnabled *bool                        `json:"changeRequestsEnabled,omitempty"`
	DefaultStrategy       *CreateFeatureStrategySchema `json:"defaultStrategy,omitempty"`
}

// NewProjectEnvironmentSchema instantiates a new ProjectEnvironmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEnvironmentSchema(environment string) *ProjectEnvironmentSchema {
	this := ProjectEnvironmentSchema{}
	this.Environment = environment
	return &this
}

// NewProjectEnvironmentSchemaWithDefaults instantiates a new ProjectEnvironmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEnvironmentSchemaWithDefaults() *ProjectEnvironmentSchema {
	this := ProjectEnvironmentSchema{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *ProjectEnvironmentSchema) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ProjectEnvironmentSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ProjectEnvironmentSchema) SetEnvironment(v string) {
	o.Environment = v
}

// GetChangeRequestsEnabled returns the ChangeRequestsEnabled field value if set, zero value otherwise.
func (o *ProjectEnvironmentSchema) GetChangeRequestsEnabled() bool {
	if o == nil || IsNil(o.ChangeRequestsEnabled) {
		var ret bool
		return ret
	}
	return *o.ChangeRequestsEnabled
}

// GetChangeRequestsEnabledOk returns a tuple with the ChangeRequestsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEnvironmentSchema) GetChangeRequestsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ChangeRequestsEnabled) {
		return nil, false
	}
	return o.ChangeRequestsEnabled, true
}

// HasChangeRequestsEnabled returns a boolean if a field has been set.
func (o *ProjectEnvironmentSchema) HasChangeRequestsEnabled() bool {
	if o != nil && !IsNil(o.ChangeRequestsEnabled) {
		return true
	}

	return false
}

// SetChangeRequestsEnabled gets a reference to the given bool and assigns it to the ChangeRequestsEnabled field.
func (o *ProjectEnvironmentSchema) SetChangeRequestsEnabled(v bool) {
	o.ChangeRequestsEnabled = &v
}

// GetDefaultStrategy returns the DefaultStrategy field value if set, zero value otherwise.
func (o *ProjectEnvironmentSchema) GetDefaultStrategy() CreateFeatureStrategySchema {
	if o == nil || IsNil(o.DefaultStrategy) {
		var ret CreateFeatureStrategySchema
		return ret
	}
	return *o.DefaultStrategy
}

// GetDefaultStrategyOk returns a tuple with the DefaultStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEnvironmentSchema) GetDefaultStrategyOk() (*CreateFeatureStrategySchema, bool) {
	if o == nil || IsNil(o.DefaultStrategy) {
		return nil, false
	}
	return o.DefaultStrategy, true
}

// HasDefaultStrategy returns a boolean if a field has been set.
func (o *ProjectEnvironmentSchema) HasDefaultStrategy() bool {
	if o != nil && !IsNil(o.DefaultStrategy) {
		return true
	}

	return false
}

// SetDefaultStrategy gets a reference to the given CreateFeatureStrategySchema and assigns it to the DefaultStrategy field.
func (o *ProjectEnvironmentSchema) SetDefaultStrategy(v CreateFeatureStrategySchema) {
	o.DefaultStrategy = &v
}

func (o ProjectEnvironmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEnvironmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	if !IsNil(o.ChangeRequestsEnabled) {
		toSerialize["changeRequestsEnabled"] = o.ChangeRequestsEnabled
	}
	if !IsNil(o.DefaultStrategy) {
		toSerialize["defaultStrategy"] = o.DefaultStrategy
	}
	return toSerialize, nil
}

type NullableProjectEnvironmentSchema struct {
	value *ProjectEnvironmentSchema
	isSet bool
}

func (v NullableProjectEnvironmentSchema) Get() *ProjectEnvironmentSchema {
	return v.value
}

func (v *NullableProjectEnvironmentSchema) Set(val *ProjectEnvironmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvironmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvironmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvironmentSchema(val *ProjectEnvironmentSchema) *NullableProjectEnvironmentSchema {
	return &NullableProjectEnvironmentSchema{value: val, isSet: true}
}

func (v NullableProjectEnvironmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvironmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
