/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectCreatedSchemaChangeRequestEnvironmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectCreatedSchemaChangeRequestEnvironmentsInner{}

// ProjectCreatedSchemaChangeRequestEnvironmentsInner struct for ProjectCreatedSchemaChangeRequestEnvironmentsInner
type ProjectCreatedSchemaChangeRequestEnvironmentsInner struct {
	// The name of the environment this change request configuration applies to.
	Name string `json:"name"`
	// The number of approvals required for a change request to be fully approved and ready to applied in this environment. If no value is provided, it will be set to the default number, which is 1. The value must be greater than or equal to 1.
	RequiredApprovals    int32 `json:"requiredApprovals"`
	AdditionalProperties map[string]interface{}
}

type _ProjectCreatedSchemaChangeRequestEnvironmentsInner ProjectCreatedSchemaChangeRequestEnvironmentsInner

// NewProjectCreatedSchemaChangeRequestEnvironmentsInner instantiates a new ProjectCreatedSchemaChangeRequestEnvironmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCreatedSchemaChangeRequestEnvironmentsInner(name string, requiredApprovals int32) *ProjectCreatedSchemaChangeRequestEnvironmentsInner {
	this := ProjectCreatedSchemaChangeRequestEnvironmentsInner{}
	this.Name = name
	this.RequiredApprovals = requiredApprovals
	return &this
}

// NewProjectCreatedSchemaChangeRequestEnvironmentsInnerWithDefaults instantiates a new ProjectCreatedSchemaChangeRequestEnvironmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCreatedSchemaChangeRequestEnvironmentsInnerWithDefaults() *ProjectCreatedSchemaChangeRequestEnvironmentsInner {
	this := ProjectCreatedSchemaChangeRequestEnvironmentsInner{}
	var requiredApprovals int32 = 1
	this.RequiredApprovals = requiredApprovals
	return &this
}

// GetName returns the Name field value
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) SetName(v string) {
	o.Name = v
}

// GetRequiredApprovals returns the RequiredApprovals field value
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) GetRequiredApprovals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value
// and a boolean to check if the value has been set.
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) GetRequiredApprovalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredApprovals, true
}

// SetRequiredApprovals sets field value
func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) SetRequiredApprovals(v int32) {
	o.RequiredApprovals = v
}

func (o ProjectCreatedSchemaChangeRequestEnvironmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectCreatedSchemaChangeRequestEnvironmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["requiredApprovals"] = o.RequiredApprovals

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectCreatedSchemaChangeRequestEnvironmentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"requiredApprovals",
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

	varProjectCreatedSchemaChangeRequestEnvironmentsInner := _ProjectCreatedSchemaChangeRequestEnvironmentsInner{}

	err = json.Unmarshal(data, &varProjectCreatedSchemaChangeRequestEnvironmentsInner)

	if err != nil {
		return err
	}

	*o = ProjectCreatedSchemaChangeRequestEnvironmentsInner(varProjectCreatedSchemaChangeRequestEnvironmentsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "requiredApprovals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectCreatedSchemaChangeRequestEnvironmentsInner struct {
	value *ProjectCreatedSchemaChangeRequestEnvironmentsInner
	isSet bool
}

func (v NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) Get() *ProjectCreatedSchemaChangeRequestEnvironmentsInner {
	return v.value
}

func (v *NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) Set(val *ProjectCreatedSchemaChangeRequestEnvironmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCreatedSchemaChangeRequestEnvironmentsInner(val *ProjectCreatedSchemaChangeRequestEnvironmentsInner) *NullableProjectCreatedSchemaChangeRequestEnvironmentsInner {
	return &NullableProjectCreatedSchemaChangeRequestEnvironmentsInner{value: val, isSet: true}
}

func (v NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCreatedSchemaChangeRequestEnvironmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
