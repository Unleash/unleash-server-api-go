/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateEnvironmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateEnvironmentSchema{}

// UpdateEnvironmentSchema struct for UpdateEnvironmentSchema
type UpdateEnvironmentSchema struct {
	// Updates the type of environment (i.e. development or production).
	Type *string `json:"type,omitempty"`
	// Changes the sort order of this environment.
	SortOrder *int32 `json:"sortOrder,omitempty"`
}

// NewUpdateEnvironmentSchema instantiates a new UpdateEnvironmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentSchema() *UpdateEnvironmentSchema {
	this := UpdateEnvironmentSchema{}
	return &this
}

// NewUpdateEnvironmentSchemaWithDefaults instantiates a new UpdateEnvironmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentSchemaWithDefaults() *UpdateEnvironmentSchema {
	this := UpdateEnvironmentSchema{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateEnvironmentSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateEnvironmentSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateEnvironmentSchema) SetType(v string) {
	o.Type = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *UpdateEnvironmentSchema) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentSchema) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *UpdateEnvironmentSchema) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *UpdateEnvironmentSchema) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o UpdateEnvironmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateEnvironmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableUpdateEnvironmentSchema struct {
	value *UpdateEnvironmentSchema
	isSet bool
}

func (v NullableUpdateEnvironmentSchema) Get() *UpdateEnvironmentSchema {
	return v.value
}

func (v *NullableUpdateEnvironmentSchema) Set(val *UpdateEnvironmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentSchema(val *UpdateEnvironmentSchema) *NullableUpdateEnvironmentSchema {
	return &NullableUpdateEnvironmentSchema{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}