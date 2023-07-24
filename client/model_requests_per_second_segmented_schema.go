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

// checks if the RequestsPerSecondSegmentedSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestsPerSecondSegmentedSchema{}

// RequestsPerSecondSegmentedSchema struct for RequestsPerSecondSegmentedSchema
type RequestsPerSecondSegmentedSchema struct {
	ClientMetrics *RequestsPerSecondSchema `json:"clientMetrics,omitempty"`
	AdminMetrics  *RequestsPerSecondSchema `json:"adminMetrics,omitempty"`
}

// NewRequestsPerSecondSegmentedSchema instantiates a new RequestsPerSecondSegmentedSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPerSecondSegmentedSchema() *RequestsPerSecondSegmentedSchema {
	this := RequestsPerSecondSegmentedSchema{}
	return &this
}

// NewRequestsPerSecondSegmentedSchemaWithDefaults instantiates a new RequestsPerSecondSegmentedSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPerSecondSegmentedSchemaWithDefaults() *RequestsPerSecondSegmentedSchema {
	this := RequestsPerSecondSegmentedSchema{}
	return &this
}

// GetClientMetrics returns the ClientMetrics field value if set, zero value otherwise.
func (o *RequestsPerSecondSegmentedSchema) GetClientMetrics() RequestsPerSecondSchema {
	if o == nil || IsNil(o.ClientMetrics) {
		var ret RequestsPerSecondSchema
		return ret
	}
	return *o.ClientMetrics
}

// GetClientMetricsOk returns a tuple with the ClientMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPerSecondSegmentedSchema) GetClientMetricsOk() (*RequestsPerSecondSchema, bool) {
	if o == nil || IsNil(o.ClientMetrics) {
		return nil, false
	}
	return o.ClientMetrics, true
}

// HasClientMetrics returns a boolean if a field has been set.
func (o *RequestsPerSecondSegmentedSchema) HasClientMetrics() bool {
	if o != nil && !IsNil(o.ClientMetrics) {
		return true
	}

	return false
}

// SetClientMetrics gets a reference to the given RequestsPerSecondSchema and assigns it to the ClientMetrics field.
func (o *RequestsPerSecondSegmentedSchema) SetClientMetrics(v RequestsPerSecondSchema) {
	o.ClientMetrics = &v
}

// GetAdminMetrics returns the AdminMetrics field value if set, zero value otherwise.
func (o *RequestsPerSecondSegmentedSchema) GetAdminMetrics() RequestsPerSecondSchema {
	if o == nil || IsNil(o.AdminMetrics) {
		var ret RequestsPerSecondSchema
		return ret
	}
	return *o.AdminMetrics
}

// GetAdminMetricsOk returns a tuple with the AdminMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPerSecondSegmentedSchema) GetAdminMetricsOk() (*RequestsPerSecondSchema, bool) {
	if o == nil || IsNil(o.AdminMetrics) {
		return nil, false
	}
	return o.AdminMetrics, true
}

// HasAdminMetrics returns a boolean if a field has been set.
func (o *RequestsPerSecondSegmentedSchema) HasAdminMetrics() bool {
	if o != nil && !IsNil(o.AdminMetrics) {
		return true
	}

	return false
}

// SetAdminMetrics gets a reference to the given RequestsPerSecondSchema and assigns it to the AdminMetrics field.
func (o *RequestsPerSecondSegmentedSchema) SetAdminMetrics(v RequestsPerSecondSchema) {
	o.AdminMetrics = &v
}

func (o RequestsPerSecondSegmentedSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestsPerSecondSegmentedSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientMetrics) {
		toSerialize["clientMetrics"] = o.ClientMetrics
	}
	if !IsNil(o.AdminMetrics) {
		toSerialize["adminMetrics"] = o.AdminMetrics
	}
	return toSerialize, nil
}

type NullableRequestsPerSecondSegmentedSchema struct {
	value *RequestsPerSecondSegmentedSchema
	isSet bool
}

func (v NullableRequestsPerSecondSegmentedSchema) Get() *RequestsPerSecondSegmentedSchema {
	return v.value
}

func (v *NullableRequestsPerSecondSegmentedSchema) Set(val *RequestsPerSecondSegmentedSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPerSecondSegmentedSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPerSecondSegmentedSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPerSecondSegmentedSchema(val *RequestsPerSecondSegmentedSchema) *NullableRequestsPerSecondSegmentedSchema {
	return &NullableRequestsPerSecondSegmentedSchema{value: val, isSet: true}
}

func (v NullableRequestsPerSecondSegmentedSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPerSecondSegmentedSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}