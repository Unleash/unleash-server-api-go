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

// checks if the UpsertStrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertStrategySchema{}

// UpsertStrategySchema struct for UpsertStrategySchema
type UpsertStrategySchema struct {
	Name        string                                `json:"name"`
	Description *string                               `json:"description,omitempty"`
	Editable    *bool                                 `json:"editable,omitempty"`
	Parameters  []UpsertStrategySchemaParametersInner `json:"parameters,omitempty"`
}

// NewUpsertStrategySchema instantiates a new UpsertStrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertStrategySchema(name string) *UpsertStrategySchema {
	this := UpsertStrategySchema{}
	this.Name = name
	return &this
}

// NewUpsertStrategySchemaWithDefaults instantiates a new UpsertStrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertStrategySchemaWithDefaults() *UpsertStrategySchema {
	this := UpsertStrategySchema{}
	return &this
}

// GetName returns the Name field value
func (o *UpsertStrategySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpsertStrategySchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpsertStrategySchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpsertStrategySchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertStrategySchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpsertStrategySchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpsertStrategySchema) SetDescription(v string) {
	o.Description = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *UpsertStrategySchema) GetEditable() bool {
	if o == nil || IsNil(o.Editable) {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertStrategySchema) GetEditableOk() (*bool, bool) {
	if o == nil || IsNil(o.Editable) {
		return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *UpsertStrategySchema) HasEditable() bool {
	if o != nil && !IsNil(o.Editable) {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *UpsertStrategySchema) SetEditable(v bool) {
	o.Editable = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UpsertStrategySchema) GetParameters() []UpsertStrategySchemaParametersInner {
	if o == nil || IsNil(o.Parameters) {
		var ret []UpsertStrategySchemaParametersInner
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertStrategySchema) GetParametersOk() ([]UpsertStrategySchemaParametersInner, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UpsertStrategySchema) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []UpsertStrategySchemaParametersInner and assigns it to the Parameters field.
func (o *UpsertStrategySchema) SetParameters(v []UpsertStrategySchemaParametersInner) {
	o.Parameters = v
}

func (o UpsertStrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertStrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Editable) {
		toSerialize["editable"] = o.Editable
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableUpsertStrategySchema struct {
	value *UpsertStrategySchema
	isSet bool
}

func (v NullableUpsertStrategySchema) Get() *UpsertStrategySchema {
	return v.value
}

func (v *NullableUpsertStrategySchema) Set(val *UpsertStrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertStrategySchema(val *UpsertStrategySchema) *NullableUpsertStrategySchema {
	return &NullableUpsertStrategySchema{value: val, isSet: true}
}

func (v NullableUpsertStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
