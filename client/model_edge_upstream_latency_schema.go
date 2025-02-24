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

// checks if the EdgeUpstreamLatencySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeUpstreamLatencySchema{}

// EdgeUpstreamLatencySchema Latencies for upstream actions from Edge (downloading/syncing new features, uploading metrics, uploading instance data)
type EdgeUpstreamLatencySchema struct {
	Features             EdgeLatencyMetricsSchema `json:"features"`
	Metrics              EdgeLatencyMetricsSchema `json:"metrics"`
	Edge                 EdgeLatencyMetricsSchema `json:"edge"`
	AdditionalProperties map[string]interface{}
}

type _EdgeUpstreamLatencySchema EdgeUpstreamLatencySchema

// NewEdgeUpstreamLatencySchema instantiates a new EdgeUpstreamLatencySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeUpstreamLatencySchema(features EdgeLatencyMetricsSchema, metrics EdgeLatencyMetricsSchema, edge EdgeLatencyMetricsSchema) *EdgeUpstreamLatencySchema {
	this := EdgeUpstreamLatencySchema{}
	this.Features = features
	this.Metrics = metrics
	this.Edge = edge
	return &this
}

// NewEdgeUpstreamLatencySchemaWithDefaults instantiates a new EdgeUpstreamLatencySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeUpstreamLatencySchemaWithDefaults() *EdgeUpstreamLatencySchema {
	this := EdgeUpstreamLatencySchema{}
	return &this
}

// GetFeatures returns the Features field value
func (o *EdgeUpstreamLatencySchema) GetFeatures() EdgeLatencyMetricsSchema {
	if o == nil {
		var ret EdgeLatencyMetricsSchema
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *EdgeUpstreamLatencySchema) GetFeaturesOk() (*EdgeLatencyMetricsSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Features, true
}

// SetFeatures sets field value
func (o *EdgeUpstreamLatencySchema) SetFeatures(v EdgeLatencyMetricsSchema) {
	o.Features = v
}

// GetMetrics returns the Metrics field value
func (o *EdgeUpstreamLatencySchema) GetMetrics() EdgeLatencyMetricsSchema {
	if o == nil {
		var ret EdgeLatencyMetricsSchema
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *EdgeUpstreamLatencySchema) GetMetricsOk() (*EdgeLatencyMetricsSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *EdgeUpstreamLatencySchema) SetMetrics(v EdgeLatencyMetricsSchema) {
	o.Metrics = v
}

// GetEdge returns the Edge field value
func (o *EdgeUpstreamLatencySchema) GetEdge() EdgeLatencyMetricsSchema {
	if o == nil {
		var ret EdgeLatencyMetricsSchema
		return ret
	}

	return o.Edge
}

// GetEdgeOk returns a tuple with the Edge field value
// and a boolean to check if the value has been set.
func (o *EdgeUpstreamLatencySchema) GetEdgeOk() (*EdgeLatencyMetricsSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edge, true
}

// SetEdge sets field value
func (o *EdgeUpstreamLatencySchema) SetEdge(v EdgeLatencyMetricsSchema) {
	o.Edge = v
}

func (o EdgeUpstreamLatencySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeUpstreamLatencySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["features"] = o.Features
	toSerialize["metrics"] = o.Metrics
	toSerialize["edge"] = o.Edge

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeUpstreamLatencySchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"features",
		"metrics",
		"edge",
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

	varEdgeUpstreamLatencySchema := _EdgeUpstreamLatencySchema{}

	err = json.Unmarshal(data, &varEdgeUpstreamLatencySchema)

	if err != nil {
		return err
	}

	*o = EdgeUpstreamLatencySchema(varEdgeUpstreamLatencySchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "features")
		delete(additionalProperties, "metrics")
		delete(additionalProperties, "edge")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeUpstreamLatencySchema struct {
	value *EdgeUpstreamLatencySchema
	isSet bool
}

func (v NullableEdgeUpstreamLatencySchema) Get() *EdgeUpstreamLatencySchema {
	return v.value
}

func (v *NullableEdgeUpstreamLatencySchema) Set(val *EdgeUpstreamLatencySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeUpstreamLatencySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeUpstreamLatencySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeUpstreamLatencySchema(val *EdgeUpstreamLatencySchema) *NullableEdgeUpstreamLatencySchema {
	return &NullableEdgeUpstreamLatencySchema{value: val, isSet: true}
}

func (v NullableEdgeUpstreamLatencySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeUpstreamLatencySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
