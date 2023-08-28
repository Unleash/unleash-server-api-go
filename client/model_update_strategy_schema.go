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

// checks if the UpdateStrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStrategySchema{}

// UpdateStrategySchema The data required to update a strategy type.
type UpdateStrategySchema struct {
	// A description of the strategy type.
	Description *string `json:"description,omitempty"`
	// The parameter list lets you pass arguments to your custom activation strategy. These will be made available to your custom strategy implementation.
	Parameters []CreateStrategySchemaParametersInner `json:"parameters"`
}

// NewUpdateStrategySchema instantiates a new UpdateStrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStrategySchema(parameters []CreateStrategySchemaParametersInner) *UpdateStrategySchema {
	this := UpdateStrategySchema{}
	this.Parameters = parameters
	return &this
}

// NewUpdateStrategySchemaWithDefaults instantiates a new UpdateStrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStrategySchemaWithDefaults() *UpdateStrategySchema {
	this := UpdateStrategySchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateStrategySchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategySchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateStrategySchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateStrategySchema) SetDescription(v string) {
	o.Description = &v
}

// GetParameters returns the Parameters field value
func (o *UpdateStrategySchema) GetParameters() []CreateStrategySchemaParametersInner {
	if o == nil {
		var ret []CreateStrategySchemaParametersInner
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *UpdateStrategySchema) GetParametersOk() ([]CreateStrategySchemaParametersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *UpdateStrategySchema) SetParameters(v []CreateStrategySchemaParametersInner) {
	o.Parameters = v
}

func (o UpdateStrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

type NullableUpdateStrategySchema struct {
	value *UpdateStrategySchema
	isSet bool
}

func (v NullableUpdateStrategySchema) Get() *UpdateStrategySchema {
	return v.value
}

func (v *NullableUpdateStrategySchema) Set(val *UpdateStrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStrategySchema(val *UpdateStrategySchema) *NullableUpdateStrategySchema {
	return &NullableUpdateStrategySchema{value: val, isSet: true}
}

func (v NullableUpdateStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
