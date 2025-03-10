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

// checks if the CreateProjectSchemaChangeRequestEnvironmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProjectSchemaChangeRequestEnvironmentsInner{}

// CreateProjectSchemaChangeRequestEnvironmentsInner struct for CreateProjectSchemaChangeRequestEnvironmentsInner
type CreateProjectSchemaChangeRequestEnvironmentsInner struct {
	// The name of the environment to configure change requests for.
	Name string `json:"name"`
	// The number of approvals required for a change request to be fully approved and ready to applied in this environment. If no value is provided, it will be set to the default number, which is 1. Values will be clamped to between 1 and 10 inclusive.
	RequiredApprovals    *int32 `json:"requiredApprovals,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateProjectSchemaChangeRequestEnvironmentsInner CreateProjectSchemaChangeRequestEnvironmentsInner

// NewCreateProjectSchemaChangeRequestEnvironmentsInner instantiates a new CreateProjectSchemaChangeRequestEnvironmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectSchemaChangeRequestEnvironmentsInner(name string) *CreateProjectSchemaChangeRequestEnvironmentsInner {
	this := CreateProjectSchemaChangeRequestEnvironmentsInner{}
	this.Name = name
	var requiredApprovals int32 = 1
	this.RequiredApprovals = &requiredApprovals
	return &this
}

// NewCreateProjectSchemaChangeRequestEnvironmentsInnerWithDefaults instantiates a new CreateProjectSchemaChangeRequestEnvironmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectSchemaChangeRequestEnvironmentsInnerWithDefaults() *CreateProjectSchemaChangeRequestEnvironmentsInner {
	this := CreateProjectSchemaChangeRequestEnvironmentsInner{}
	var requiredApprovals int32 = 1
	this.RequiredApprovals = &requiredApprovals
	return &this
}

// GetName returns the Name field value
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) SetName(v string) {
	o.Name = v
}

// GetRequiredApprovals returns the RequiredApprovals field value if set, zero value otherwise.
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetRequiredApprovals() int32 {
	if o == nil || IsNil(o.RequiredApprovals) {
		var ret int32
		return ret
	}
	return *o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) GetRequiredApprovalsOk() (*int32, bool) {
	if o == nil || IsNil(o.RequiredApprovals) {
		return nil, false
	}
	return o.RequiredApprovals, true
}

// HasRequiredApprovals returns a boolean if a field has been set.
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) HasRequiredApprovals() bool {
	if o != nil && !IsNil(o.RequiredApprovals) {
		return true
	}

	return false
}

// SetRequiredApprovals gets a reference to the given int32 and assigns it to the RequiredApprovals field.
func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) SetRequiredApprovals(v int32) {
	o.RequiredApprovals = &v
}

func (o CreateProjectSchemaChangeRequestEnvironmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectSchemaChangeRequestEnvironmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.RequiredApprovals) {
		toSerialize["requiredApprovals"] = o.RequiredApprovals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateProjectSchemaChangeRequestEnvironmentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateProjectSchemaChangeRequestEnvironmentsInner := _CreateProjectSchemaChangeRequestEnvironmentsInner{}

	err = json.Unmarshal(data, &varCreateProjectSchemaChangeRequestEnvironmentsInner)

	if err != nil {
		return err
	}

	*o = CreateProjectSchemaChangeRequestEnvironmentsInner(varCreateProjectSchemaChangeRequestEnvironmentsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "requiredApprovals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateProjectSchemaChangeRequestEnvironmentsInner struct {
	value *CreateProjectSchemaChangeRequestEnvironmentsInner
	isSet bool
}

func (v NullableCreateProjectSchemaChangeRequestEnvironmentsInner) Get() *CreateProjectSchemaChangeRequestEnvironmentsInner {
	return v.value
}

func (v *NullableCreateProjectSchemaChangeRequestEnvironmentsInner) Set(val *CreateProjectSchemaChangeRequestEnvironmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectSchemaChangeRequestEnvironmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectSchemaChangeRequestEnvironmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectSchemaChangeRequestEnvironmentsInner(val *CreateProjectSchemaChangeRequestEnvironmentsInner) *NullableCreateProjectSchemaChangeRequestEnvironmentsInner {
	return &NullableCreateProjectSchemaChangeRequestEnvironmentsInner{value: val, isSet: true}
}

func (v NullableCreateProjectSchemaChangeRequestEnvironmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectSchemaChangeRequestEnvironmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
