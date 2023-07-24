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

// checks if the BulkMetricsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkMetricsSchema{}

// BulkMetricsSchema A batch of metrics accumulated by Edge (or other compatible applications). Includes both application registrations as well usage metrics from clients
type BulkMetricsSchema struct {
	// A list of applications registered by an Unleash SDK
	Applications []BulkRegistrationSchema `json:"applications"`
	// a list of client usage metrics registered by downstream providers. (Typically Unleash Edge)
	Metrics []ClientMetricsEnvSchema `json:"metrics"`
}

// NewBulkMetricsSchema instantiates a new BulkMetricsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkMetricsSchema(applications []BulkRegistrationSchema, metrics []ClientMetricsEnvSchema) *BulkMetricsSchema {
	this := BulkMetricsSchema{}
	this.Applications = applications
	this.Metrics = metrics
	return &this
}

// NewBulkMetricsSchemaWithDefaults instantiates a new BulkMetricsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkMetricsSchemaWithDefaults() *BulkMetricsSchema {
	this := BulkMetricsSchema{}
	return &this
}

// GetApplications returns the Applications field value
func (o *BulkMetricsSchema) GetApplications() []BulkRegistrationSchema {
	if o == nil {
		var ret []BulkRegistrationSchema
		return ret
	}

	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value
// and a boolean to check if the value has been set.
func (o *BulkMetricsSchema) GetApplicationsOk() ([]BulkRegistrationSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Applications, true
}

// SetApplications sets field value
func (o *BulkMetricsSchema) SetApplications(v []BulkRegistrationSchema) {
	o.Applications = v
}

// GetMetrics returns the Metrics field value
func (o *BulkMetricsSchema) GetMetrics() []ClientMetricsEnvSchema {
	if o == nil {
		var ret []ClientMetricsEnvSchema
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *BulkMetricsSchema) GetMetricsOk() ([]ClientMetricsEnvSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *BulkMetricsSchema) SetMetrics(v []ClientMetricsEnvSchema) {
	o.Metrics = v
}

func (o BulkMetricsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkMetricsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applications"] = o.Applications
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

type NullableBulkMetricsSchema struct {
	value *BulkMetricsSchema
	isSet bool
}

func (v NullableBulkMetricsSchema) Get() *BulkMetricsSchema {
	return v.value
}

func (v *NullableBulkMetricsSchema) Set(val *BulkMetricsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkMetricsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkMetricsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkMetricsSchema(val *BulkMetricsSchema) *NullableBulkMetricsSchema {
	return &NullableBulkMetricsSchema{value: val, isSet: true}
}

func (v NullableBulkMetricsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkMetricsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}