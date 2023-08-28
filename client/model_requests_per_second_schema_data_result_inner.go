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

// checks if the RequestsPerSecondSchemaDataResultInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestsPerSecondSchemaDataResultInner{}

// RequestsPerSecondSchemaDataResultInner A representation of a single metric to build a line in a graph
type RequestsPerSecondSchemaDataResultInner struct {
	Metric *RequestsPerSecondSchemaDataResultInnerMetric `json:"metric,omitempty"`
	// An array of arrays. Each element of the array is an array of size 2 consisting of the 2 axis for the graph: in position zero the x axis represented as a number and position one the y axis represented as string
	Values [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner `json:"values,omitempty"`
}

// NewRequestsPerSecondSchemaDataResultInner instantiates a new RequestsPerSecondSchemaDataResultInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPerSecondSchemaDataResultInner() *RequestsPerSecondSchemaDataResultInner {
	this := RequestsPerSecondSchemaDataResultInner{}
	return &this
}

// NewRequestsPerSecondSchemaDataResultInnerWithDefaults instantiates a new RequestsPerSecondSchemaDataResultInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPerSecondSchemaDataResultInnerWithDefaults() *RequestsPerSecondSchemaDataResultInner {
	this := RequestsPerSecondSchemaDataResultInner{}
	return &this
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *RequestsPerSecondSchemaDataResultInner) GetMetric() RequestsPerSecondSchemaDataResultInnerMetric {
	if o == nil || IsNil(o.Metric) {
		var ret RequestsPerSecondSchemaDataResultInnerMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPerSecondSchemaDataResultInner) GetMetricOk() (*RequestsPerSecondSchemaDataResultInnerMetric, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *RequestsPerSecondSchemaDataResultInner) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given RequestsPerSecondSchemaDataResultInnerMetric and assigns it to the Metric field.
func (o *RequestsPerSecondSchemaDataResultInner) SetMetric(v RequestsPerSecondSchemaDataResultInnerMetric) {
	o.Metric = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RequestsPerSecondSchemaDataResultInner) GetValues() [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner {
	if o == nil || IsNil(o.Values) {
		var ret [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPerSecondSchemaDataResultInner) GetValuesOk() ([][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RequestsPerSecondSchemaDataResultInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner and assigns it to the Values field.
func (o *RequestsPerSecondSchemaDataResultInner) SetValues(v [][]RequestsPerSecondSchemaDataResultInnerValuesInnerInner) {
	o.Values = v
}

func (o RequestsPerSecondSchemaDataResultInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestsPerSecondSchemaDataResultInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableRequestsPerSecondSchemaDataResultInner struct {
	value *RequestsPerSecondSchemaDataResultInner
	isSet bool
}

func (v NullableRequestsPerSecondSchemaDataResultInner) Get() *RequestsPerSecondSchemaDataResultInner {
	return v.value
}

func (v *NullableRequestsPerSecondSchemaDataResultInner) Set(val *RequestsPerSecondSchemaDataResultInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPerSecondSchemaDataResultInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPerSecondSchemaDataResultInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPerSecondSchemaDataResultInner(val *RequestsPerSecondSchemaDataResultInner) *NullableRequestsPerSecondSchemaDataResultInner {
	return &NullableRequestsPerSecondSchemaDataResultInner{value: val, isSet: true}
}

func (v NullableRequestsPerSecondSchemaDataResultInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPerSecondSchemaDataResultInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
