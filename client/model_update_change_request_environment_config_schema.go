/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateChangeRequestEnvironmentConfigSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateChangeRequestEnvironmentConfigSchema{}

// UpdateChangeRequestEnvironmentConfigSchema Data used to update change request in an environment
type UpdateChangeRequestEnvironmentConfigSchema struct {
	// `true` if change requests should be enabled, otherwise `false`.
	ChangeRequestsEnabled bool `json:"changeRequestsEnabled"`
	// The number of approvals required before a change request can be applied in this environment.
	RequiredApprovals *int32 `json:"requiredApprovals,omitempty"`
}

// NewUpdateChangeRequestEnvironmentConfigSchema instantiates a new UpdateChangeRequestEnvironmentConfigSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateChangeRequestEnvironmentConfigSchema(changeRequestsEnabled bool) *UpdateChangeRequestEnvironmentConfigSchema {
	this := UpdateChangeRequestEnvironmentConfigSchema{}
	this.ChangeRequestsEnabled = changeRequestsEnabled
	return &this
}

// NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults instantiates a new UpdateChangeRequestEnvironmentConfigSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateChangeRequestEnvironmentConfigSchemaWithDefaults() *UpdateChangeRequestEnvironmentConfigSchema {
	this := UpdateChangeRequestEnvironmentConfigSchema{}
	return &this
}

// GetChangeRequestsEnabled returns the ChangeRequestsEnabled field value
func (o *UpdateChangeRequestEnvironmentConfigSchema) GetChangeRequestsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ChangeRequestsEnabled
}

// GetChangeRequestsEnabledOk returns a tuple with the ChangeRequestsEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateChangeRequestEnvironmentConfigSchema) GetChangeRequestsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestsEnabled, true
}

// SetChangeRequestsEnabled sets field value
func (o *UpdateChangeRequestEnvironmentConfigSchema) SetChangeRequestsEnabled(v bool) {
	o.ChangeRequestsEnabled = v
}

// GetRequiredApprovals returns the RequiredApprovals field value if set, zero value otherwise.
func (o *UpdateChangeRequestEnvironmentConfigSchema) GetRequiredApprovals() int32 {
	if o == nil || IsNil(o.RequiredApprovals) {
		var ret int32
		return ret
	}
	return *o.RequiredApprovals
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateChangeRequestEnvironmentConfigSchema) GetRequiredApprovalsOk() (*int32, bool) {
	if o == nil || IsNil(o.RequiredApprovals) {
		return nil, false
	}
	return o.RequiredApprovals, true
}

// HasRequiredApprovals returns a boolean if a field has been set.
func (o *UpdateChangeRequestEnvironmentConfigSchema) HasRequiredApprovals() bool {
	if o != nil && !IsNil(o.RequiredApprovals) {
		return true
	}

	return false
}

// SetRequiredApprovals gets a reference to the given int32 and assigns it to the RequiredApprovals field.
func (o *UpdateChangeRequestEnvironmentConfigSchema) SetRequiredApprovals(v int32) {
	o.RequiredApprovals = &v
}

func (o UpdateChangeRequestEnvironmentConfigSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateChangeRequestEnvironmentConfigSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["changeRequestsEnabled"] = o.ChangeRequestsEnabled
	if !IsNil(o.RequiredApprovals) {
		toSerialize["requiredApprovals"] = o.RequiredApprovals
	}
	return toSerialize, nil
}

type NullableUpdateChangeRequestEnvironmentConfigSchema struct {
	value *UpdateChangeRequestEnvironmentConfigSchema
	isSet bool
}

func (v NullableUpdateChangeRequestEnvironmentConfigSchema) Get() *UpdateChangeRequestEnvironmentConfigSchema {
	return v.value
}

func (v *NullableUpdateChangeRequestEnvironmentConfigSchema) Set(val *UpdateChangeRequestEnvironmentConfigSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateChangeRequestEnvironmentConfigSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateChangeRequestEnvironmentConfigSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateChangeRequestEnvironmentConfigSchema(val *UpdateChangeRequestEnvironmentConfigSchema) *NullableUpdateChangeRequestEnvironmentConfigSchema {
	return &NullableUpdateChangeRequestEnvironmentConfigSchema{value: val, isSet: true}
}

func (v NullableUpdateChangeRequestEnvironmentConfigSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateChangeRequestEnvironmentConfigSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
