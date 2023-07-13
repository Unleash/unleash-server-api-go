/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MaintenanceSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceSchema{}

// MaintenanceSchema struct for MaintenanceSchema
type MaintenanceSchema struct {
	Enabled bool `json:"enabled"`
}

// NewMaintenanceSchema instantiates a new MaintenanceSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceSchema(enabled bool) *MaintenanceSchema {
	this := MaintenanceSchema{}
	this.Enabled = enabled
	return &this
}

// NewMaintenanceSchemaWithDefaults instantiates a new MaintenanceSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceSchemaWithDefaults() *MaintenanceSchema {
	this := MaintenanceSchema{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *MaintenanceSchema) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MaintenanceSchema) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MaintenanceSchema) SetEnabled(v bool) {
	o.Enabled = v
}

func (o MaintenanceSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableMaintenanceSchema struct {
	value *MaintenanceSchema
	isSet bool
}

func (v NullableMaintenanceSchema) Get() *MaintenanceSchema {
	return v.value
}

func (v *NullableMaintenanceSchema) Set(val *MaintenanceSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceSchema(val *MaintenanceSchema) *NullableMaintenanceSchema {
	return &NullableMaintenanceSchema{value: val, isSet: true}
}

func (v NullableMaintenanceSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


