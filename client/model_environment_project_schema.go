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

// checks if the EnvironmentProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentProjectSchema{}

// EnvironmentProjectSchema Describes a project's configuration in a given environment.
type EnvironmentProjectSchema struct {
	// The name of the environment
	Name string `json:"name"`
	// The [type of environment](https://docs.getunleash.io/reference/environments#environment-types).
	Type string `json:"type"`
	// `true` if the environment is enabled for the project, otherwise `false`
	Enabled bool `json:"enabled"`
	// `true` if the environment is protected, otherwise `false`. A *protected* environment can not be deleted.
	Protected bool `json:"protected"`
	// Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear
	SortOrder int32 `json:"sortOrder"`
	// The number of client and front-end API tokens that have access to this project
	ProjectApiTokenCount *int32 `json:"projectApiTokenCount,omitempty"`
	// The number of features enabled in this environment for this project
	ProjectEnabledToggleCount *int32                       `json:"projectEnabledToggleCount,omitempty"`
	DefaultStrategy           *CreateFeatureStrategySchema `json:"defaultStrategy,omitempty"`
}

// NewEnvironmentProjectSchema instantiates a new EnvironmentProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentProjectSchema(name string, type_ string, enabled bool, protected bool, sortOrder int32) *EnvironmentProjectSchema {
	this := EnvironmentProjectSchema{}
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	this.Protected = protected
	this.SortOrder = sortOrder
	return &this
}

// NewEnvironmentProjectSchemaWithDefaults instantiates a new EnvironmentProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentProjectSchemaWithDefaults() *EnvironmentProjectSchema {
	this := EnvironmentProjectSchema{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentProjectSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentProjectSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *EnvironmentProjectSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnvironmentProjectSchema) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *EnvironmentProjectSchema) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EnvironmentProjectSchema) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProtected returns the Protected field value
func (o *EnvironmentProjectSchema) GetProtected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value
func (o *EnvironmentProjectSchema) SetProtected(v bool) {
	o.Protected = v
}

// GetSortOrder returns the SortOrder field value
func (o *EnvironmentProjectSchema) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *EnvironmentProjectSchema) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetProjectApiTokenCount returns the ProjectApiTokenCount field value if set, zero value otherwise.
func (o *EnvironmentProjectSchema) GetProjectApiTokenCount() int32 {
	if o == nil || IsNil(o.ProjectApiTokenCount) {
		var ret int32
		return ret
	}
	return *o.ProjectApiTokenCount
}

// GetProjectApiTokenCountOk returns a tuple with the ProjectApiTokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetProjectApiTokenCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectApiTokenCount) {
		return nil, false
	}
	return o.ProjectApiTokenCount, true
}

// HasProjectApiTokenCount returns a boolean if a field has been set.
func (o *EnvironmentProjectSchema) HasProjectApiTokenCount() bool {
	if o != nil && !IsNil(o.ProjectApiTokenCount) {
		return true
	}

	return false
}

// SetProjectApiTokenCount gets a reference to the given int32 and assigns it to the ProjectApiTokenCount field.
func (o *EnvironmentProjectSchema) SetProjectApiTokenCount(v int32) {
	o.ProjectApiTokenCount = &v
}

// GetProjectEnabledToggleCount returns the ProjectEnabledToggleCount field value if set, zero value otherwise.
func (o *EnvironmentProjectSchema) GetProjectEnabledToggleCount() int32 {
	if o == nil || IsNil(o.ProjectEnabledToggleCount) {
		var ret int32
		return ret
	}
	return *o.ProjectEnabledToggleCount
}

// GetProjectEnabledToggleCountOk returns a tuple with the ProjectEnabledToggleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetProjectEnabledToggleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectEnabledToggleCount) {
		return nil, false
	}
	return o.ProjectEnabledToggleCount, true
}

// HasProjectEnabledToggleCount returns a boolean if a field has been set.
func (o *EnvironmentProjectSchema) HasProjectEnabledToggleCount() bool {
	if o != nil && !IsNil(o.ProjectEnabledToggleCount) {
		return true
	}

	return false
}

// SetProjectEnabledToggleCount gets a reference to the given int32 and assigns it to the ProjectEnabledToggleCount field.
func (o *EnvironmentProjectSchema) SetProjectEnabledToggleCount(v int32) {
	o.ProjectEnabledToggleCount = &v
}

// GetDefaultStrategy returns the DefaultStrategy field value if set, zero value otherwise.
func (o *EnvironmentProjectSchema) GetDefaultStrategy() CreateFeatureStrategySchema {
	if o == nil || IsNil(o.DefaultStrategy) {
		var ret CreateFeatureStrategySchema
		return ret
	}
	return *o.DefaultStrategy
}

// GetDefaultStrategyOk returns a tuple with the DefaultStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentProjectSchema) GetDefaultStrategyOk() (*CreateFeatureStrategySchema, bool) {
	if o == nil || IsNil(o.DefaultStrategy) {
		return nil, false
	}
	return o.DefaultStrategy, true
}

// HasDefaultStrategy returns a boolean if a field has been set.
func (o *EnvironmentProjectSchema) HasDefaultStrategy() bool {
	if o != nil && !IsNil(o.DefaultStrategy) {
		return true
	}

	return false
}

// SetDefaultStrategy gets a reference to the given CreateFeatureStrategySchema and assigns it to the DefaultStrategy field.
func (o *EnvironmentProjectSchema) SetDefaultStrategy(v CreateFeatureStrategySchema) {
	o.DefaultStrategy = &v
}

func (o EnvironmentProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["enabled"] = o.Enabled
	toSerialize["protected"] = o.Protected
	toSerialize["sortOrder"] = o.SortOrder
	if !IsNil(o.ProjectApiTokenCount) {
		toSerialize["projectApiTokenCount"] = o.ProjectApiTokenCount
	}
	if !IsNil(o.ProjectEnabledToggleCount) {
		toSerialize["projectEnabledToggleCount"] = o.ProjectEnabledToggleCount
	}
	if !IsNil(o.DefaultStrategy) {
		toSerialize["defaultStrategy"] = o.DefaultStrategy
	}
	return toSerialize, nil
}

type NullableEnvironmentProjectSchema struct {
	value *EnvironmentProjectSchema
	isSet bool
}

func (v NullableEnvironmentProjectSchema) Get() *EnvironmentProjectSchema {
	return v.value
}

func (v *NullableEnvironmentProjectSchema) Set(val *EnvironmentProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentProjectSchema(val *EnvironmentProjectSchema) *NullableEnvironmentProjectSchema {
	return &NullableEnvironmentProjectSchema{value: val, isSet: true}
}

func (v NullableEnvironmentProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}