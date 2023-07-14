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

// checks if the HealthCheckSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckSchema{}

// HealthCheckSchema Used by service orchestrators to decide whether this Unleash instance should be marked as healthy or unhealthy
type HealthCheckSchema struct {
	// The state this Unleash instance is in. GOOD if everything is ok, BAD if the instance should be restarted
	Health string `json:"health"`
}

// NewHealthCheckSchema instantiates a new HealthCheckSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckSchema(health string) *HealthCheckSchema {
	this := HealthCheckSchema{}
	this.Health = health
	return &this
}

// NewHealthCheckSchemaWithDefaults instantiates a new HealthCheckSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckSchemaWithDefaults() *HealthCheckSchema {
	this := HealthCheckSchema{}
	return &this
}

// GetHealth returns the Health field value
func (o *HealthCheckSchema) GetHealth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *HealthCheckSchema) GetHealthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *HealthCheckSchema) SetHealth(v string) {
	o.Health = v
}

func (o HealthCheckSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["health"] = o.Health
	return toSerialize, nil
}

type NullableHealthCheckSchema struct {
	value *HealthCheckSchema
	isSet bool
}

func (v NullableHealthCheckSchema) Get() *HealthCheckSchema {
	return v.value
}

func (v *NullableHealthCheckSchema) Set(val *HealthCheckSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckSchema(val *HealthCheckSchema) *NullableHealthCheckSchema {
	return &NullableHealthCheckSchema{value: val, isSet: true}
}

func (v NullableHealthCheckSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
