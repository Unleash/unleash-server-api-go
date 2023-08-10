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

// checks if the UpdateProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectSchema{}

// UpdateProjectSchema struct for UpdateProjectSchema
type UpdateProjectSchema struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewUpdateProjectSchema instantiates a new UpdateProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectSchema(name string) *UpdateProjectSchema {
	this := UpdateProjectSchema{}
	this.Name = name
	return &this
}

// NewUpdateProjectSchemaWithDefaults instantiates a new UpdateProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectSchemaWithDefaults() *UpdateProjectSchema {
	this := UpdateProjectSchema{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateProjectSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProjectSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateProjectSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProjectSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateProjectSchema) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableUpdateProjectSchema struct {
	value *UpdateProjectSchema
	isSet bool
}

func (v NullableUpdateProjectSchema) Get() *UpdateProjectSchema {
	return v.value
}

func (v *NullableUpdateProjectSchema) Set(val *UpdateProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectSchema(val *UpdateProjectSchema) *NullableUpdateProjectSchema {
	return &NullableUpdateProjectSchema{value: val, isSet: true}
}

func (v NullableUpdateProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
