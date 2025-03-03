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

// checks if the ChangeRequestEnvironmentConfigSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestEnvironmentConfigSchema{}

// ChangeRequestEnvironmentConfigSchema The [change request](https://docs.getunleash.io/reference/change-requests) configuration for a specific environment.
type ChangeRequestEnvironmentConfigSchema struct {
	// The environment that this configuration applies to.
	Environment string `json:"environment"`
	// The [type of the environment](https://docs.getunleash.io/reference/environments#environment-types) listed in `environment`.
	Type string `json:"type"`
	// `true` if this environment has change requests enabled, otherwise `false`.
	ChangeRequestEnabled bool `json:"changeRequestEnabled"`
	// The number of approvals that are required for a change request to be fully approved and ready to be applied in this environment.
	RequiredApprovals    NullableFloat32 `json:"requiredApprovals"`
	AdditionalProperties map[string]interface{}
}

type _ChangeRequestEnvironmentConfigSchema ChangeRequestEnvironmentConfigSchema

// NewChangeRequestEnvironmentConfigSchema instantiates a new ChangeRequestEnvironmentConfigSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestEnvironmentConfigSchema(environment string, type_ string, changeRequestEnabled bool, requiredApprovals NullableFloat32) *ChangeRequestEnvironmentConfigSchema {
	this := ChangeRequestEnvironmentConfigSchema{}
	this.Environment = environment
	this.Type = type_
	this.ChangeRequestEnabled = changeRequestEnabled
	this.RequiredApprovals = requiredApprovals
	return &this
}

// NewChangeRequestEnvironmentConfigSchemaWithDefaults instantiates a new ChangeRequestEnvironmentConfigSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestEnvironmentConfigSchemaWithDefaults() *ChangeRequestEnvironmentConfigSchema {
	this := ChangeRequestEnvironmentConfigSchema{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *ChangeRequestEnvironmentConfigSchema) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestEnvironmentConfigSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ChangeRequestEnvironmentConfigSchema) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *ChangeRequestEnvironmentConfigSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestEnvironmentConfigSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChangeRequestEnvironmentConfigSchema) SetType(v string) {
	o.Type = v
}

// GetChangeRequestEnabled returns the ChangeRequestEnabled field value
func (o *ChangeRequestEnvironmentConfigSchema) GetChangeRequestEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ChangeRequestEnabled
}

// GetChangeRequestEnabledOk returns a tuple with the ChangeRequestEnabled field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestEnvironmentConfigSchema) GetChangeRequestEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestEnabled, true
}

// SetChangeRequestEnabled sets field value
func (o *ChangeRequestEnvironmentConfigSchema) SetChangeRequestEnabled(v bool) {
	o.ChangeRequestEnabled = v
}

// GetRequiredApprovals returns the RequiredApprovals field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ChangeRequestEnvironmentConfigSchema) GetRequiredApprovals() float32 {
	if o == nil || o.RequiredApprovals.Get() == nil {
		var ret float32
		return ret
	}

	return *o.RequiredApprovals.Get()
}

// GetRequiredApprovalsOk returns a tuple with the RequiredApprovals field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeRequestEnvironmentConfigSchema) GetRequiredApprovalsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredApprovals.Get(), o.RequiredApprovals.IsSet()
}

// SetRequiredApprovals sets field value
func (o *ChangeRequestEnvironmentConfigSchema) SetRequiredApprovals(v float32) {
	o.RequiredApprovals.Set(&v)
}

func (o ChangeRequestEnvironmentConfigSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestEnvironmentConfigSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	toSerialize["changeRequestEnabled"] = o.ChangeRequestEnabled
	toSerialize["requiredApprovals"] = o.RequiredApprovals.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChangeRequestEnvironmentConfigSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"type",
		"changeRequestEnabled",
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

	varChangeRequestEnvironmentConfigSchema := _ChangeRequestEnvironmentConfigSchema{}

	err = json.Unmarshal(data, &varChangeRequestEnvironmentConfigSchema)

	if err != nil {
		return err
	}

	*o = ChangeRequestEnvironmentConfigSchema(varChangeRequestEnvironmentConfigSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "type")
		delete(additionalProperties, "changeRequestEnabled")
		delete(additionalProperties, "requiredApprovals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChangeRequestEnvironmentConfigSchema struct {
	value *ChangeRequestEnvironmentConfigSchema
	isSet bool
}

func (v NullableChangeRequestEnvironmentConfigSchema) Get() *ChangeRequestEnvironmentConfigSchema {
	return v.value
}

func (v *NullableChangeRequestEnvironmentConfigSchema) Set(val *ChangeRequestEnvironmentConfigSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestEnvironmentConfigSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestEnvironmentConfigSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestEnvironmentConfigSchema(val *ChangeRequestEnvironmentConfigSchema) *NullableChangeRequestEnvironmentConfigSchema {
	return &NullableChangeRequestEnvironmentConfigSchema{value: val, isSet: true}
}

func (v NullableChangeRequestEnvironmentConfigSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestEnvironmentConfigSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
