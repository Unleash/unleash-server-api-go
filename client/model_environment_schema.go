/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the EnvironmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentSchema{}

// EnvironmentSchema A definition of the project environment
type EnvironmentSchema struct {
	// The name of the environment
	Name string `json:"name"`
	// The [type of environment](https://docs.getunleash.io/reference/environments#environment-types).
	Type string `json:"type"`
	// `true` if the environment is enabled for the project, otherwise `false`.
	Enabled bool `json:"enabled"`
	// `true` if the environment is protected, otherwise `false`. A *protected* environment can not be deleted.
	Protected bool `json:"protected"`
	// Priority of the environment in a list of environments, the lower the value, the higher up in the list the environment will appear. Needs to be an integer
	SortOrder int32 `json:"sortOrder"`
	// The number of projects with this environment
	ProjectCount NullableInt32 `json:"projectCount,omitempty"`
	// The number of API tokens for the project environment
	ApiTokenCount NullableInt32 `json:"apiTokenCount,omitempty"`
	// The number of enabled toggles for the project environment
	EnabledToggleCount   NullableInt32 `json:"enabledToggleCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentSchema EnvironmentSchema

// NewEnvironmentSchema instantiates a new EnvironmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSchema(name string, type_ string, enabled bool, protected bool, sortOrder int32) *EnvironmentSchema {
	this := EnvironmentSchema{}
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	this.Protected = protected
	this.SortOrder = sortOrder
	return &this
}

// NewEnvironmentSchemaWithDefaults instantiates a new EnvironmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSchemaWithDefaults() *EnvironmentSchema {
	this := EnvironmentSchema{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *EnvironmentSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnvironmentSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnvironmentSchema) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *EnvironmentSchema) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EnvironmentSchema) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EnvironmentSchema) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProtected returns the Protected field value
func (o *EnvironmentSchema) GetProtected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *EnvironmentSchema) GetProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value
func (o *EnvironmentSchema) SetProtected(v bool) {
	o.Protected = v
}

// GetSortOrder returns the SortOrder field value
func (o *EnvironmentSchema) GetSortOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *EnvironmentSchema) GetSortOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *EnvironmentSchema) SetSortOrder(v int32) {
	o.SortOrder = v
}

// GetProjectCount returns the ProjectCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSchema) GetProjectCount() int32 {
	if o == nil || IsNil(o.ProjectCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectCount.Get()
}

// GetProjectCountOk returns a tuple with the ProjectCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSchema) GetProjectCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectCount.Get(), o.ProjectCount.IsSet()
}

// HasProjectCount returns a boolean if a field has been set.
func (o *EnvironmentSchema) HasProjectCount() bool {
	if o != nil && o.ProjectCount.IsSet() {
		return true
	}

	return false
}

// SetProjectCount gets a reference to the given NullableInt32 and assigns it to the ProjectCount field.
func (o *EnvironmentSchema) SetProjectCount(v int32) {
	o.ProjectCount.Set(&v)
}

// SetProjectCountNil sets the value for ProjectCount to be an explicit nil
func (o *EnvironmentSchema) SetProjectCountNil() {
	o.ProjectCount.Set(nil)
}

// UnsetProjectCount ensures that no value is present for ProjectCount, not even an explicit nil
func (o *EnvironmentSchema) UnsetProjectCount() {
	o.ProjectCount.Unset()
}

// GetApiTokenCount returns the ApiTokenCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSchema) GetApiTokenCount() int32 {
	if o == nil || IsNil(o.ApiTokenCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiTokenCount.Get()
}

// GetApiTokenCountOk returns a tuple with the ApiTokenCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSchema) GetApiTokenCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiTokenCount.Get(), o.ApiTokenCount.IsSet()
}

// HasApiTokenCount returns a boolean if a field has been set.
func (o *EnvironmentSchema) HasApiTokenCount() bool {
	if o != nil && o.ApiTokenCount.IsSet() {
		return true
	}

	return false
}

// SetApiTokenCount gets a reference to the given NullableInt32 and assigns it to the ApiTokenCount field.
func (o *EnvironmentSchema) SetApiTokenCount(v int32) {
	o.ApiTokenCount.Set(&v)
}

// SetApiTokenCountNil sets the value for ApiTokenCount to be an explicit nil
func (o *EnvironmentSchema) SetApiTokenCountNil() {
	o.ApiTokenCount.Set(nil)
}

// UnsetApiTokenCount ensures that no value is present for ApiTokenCount, not even an explicit nil
func (o *EnvironmentSchema) UnsetApiTokenCount() {
	o.ApiTokenCount.Unset()
}

// GetEnabledToggleCount returns the EnabledToggleCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentSchema) GetEnabledToggleCount() int32 {
	if o == nil || IsNil(o.EnabledToggleCount.Get()) {
		var ret int32
		return ret
	}
	return *o.EnabledToggleCount.Get()
}

// GetEnabledToggleCountOk returns a tuple with the EnabledToggleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentSchema) GetEnabledToggleCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledToggleCount.Get(), o.EnabledToggleCount.IsSet()
}

// HasEnabledToggleCount returns a boolean if a field has been set.
func (o *EnvironmentSchema) HasEnabledToggleCount() bool {
	if o != nil && o.EnabledToggleCount.IsSet() {
		return true
	}

	return false
}

// SetEnabledToggleCount gets a reference to the given NullableInt32 and assigns it to the EnabledToggleCount field.
func (o *EnvironmentSchema) SetEnabledToggleCount(v int32) {
	o.EnabledToggleCount.Set(&v)
}

// SetEnabledToggleCountNil sets the value for EnabledToggleCount to be an explicit nil
func (o *EnvironmentSchema) SetEnabledToggleCountNil() {
	o.EnabledToggleCount.Set(nil)
}

// UnsetEnabledToggleCount ensures that no value is present for EnabledToggleCount, not even an explicit nil
func (o *EnvironmentSchema) UnsetEnabledToggleCount() {
	o.EnabledToggleCount.Unset()
}

func (o EnvironmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["enabled"] = o.Enabled
	toSerialize["protected"] = o.Protected
	toSerialize["sortOrder"] = o.SortOrder
	if o.ProjectCount.IsSet() {
		toSerialize["projectCount"] = o.ProjectCount.Get()
	}
	if o.ApiTokenCount.IsSet() {
		toSerialize["apiTokenCount"] = o.ApiTokenCount.Get()
	}
	if o.EnabledToggleCount.IsSet() {
		toSerialize["enabledToggleCount"] = o.EnabledToggleCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"enabled",
		"protected",
		"sortOrder",
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

	varEnvironmentSchema := _EnvironmentSchema{}

	err = json.Unmarshal(data, &varEnvironmentSchema)

	if err != nil {
		return err
	}

	*o = EnvironmentSchema(varEnvironmentSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "protected")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "projectCount")
		delete(additionalProperties, "apiTokenCount")
		delete(additionalProperties, "enabledToggleCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentSchema struct {
	value *EnvironmentSchema
	isSet bool
}

func (v NullableEnvironmentSchema) Get() *EnvironmentSchema {
	return v.value
}

func (v *NullableEnvironmentSchema) Set(val *EnvironmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSchema(val *EnvironmentSchema) *NullableEnvironmentSchema {
	return &NullableEnvironmentSchema{value: val, isSet: true}
}

func (v NullableEnvironmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
