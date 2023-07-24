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

// checks if the ClientFeaturesSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientFeaturesSchema{}

// ClientFeaturesSchema Configuration data for server-side SDKs for evaluating feature flags.
type ClientFeaturesSchema struct {
	// A version number for the format used in the response. Most Unleash instances now return version 2, which includes segments as a separate array
	Version float32 `json:"version"`
	// A list of feature toggles with their configuration
	Features []ClientFeatureSchema `json:"features"`
	// A list of [Segments](https://docs.getunleash.io/reference/segments) configured for this Unleash instance
	Segments []SegmentSchema            `json:"segments,omitempty"`
	Query    *ClientFeaturesQuerySchema `json:"query,omitempty"`
}

// NewClientFeaturesSchema instantiates a new ClientFeaturesSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientFeaturesSchema(version float32, features []ClientFeatureSchema) *ClientFeaturesSchema {
	this := ClientFeaturesSchema{}
	this.Version = version
	this.Features = features
	return &this
}

// NewClientFeaturesSchemaWithDefaults instantiates a new ClientFeaturesSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientFeaturesSchemaWithDefaults() *ClientFeaturesSchema {
	this := ClientFeaturesSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *ClientFeaturesSchema) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ClientFeaturesSchema) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ClientFeaturesSchema) SetVersion(v float32) {
	o.Version = v
}

// GetFeatures returns the Features field value
func (o *ClientFeaturesSchema) GetFeatures() []ClientFeatureSchema {
	if o == nil {
		var ret []ClientFeatureSchema
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *ClientFeaturesSchema) GetFeaturesOk() ([]ClientFeatureSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *ClientFeaturesSchema) SetFeatures(v []ClientFeatureSchema) {
	o.Features = v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *ClientFeaturesSchema) GetSegments() []SegmentSchema {
	if o == nil || IsNil(o.Segments) {
		var ret []SegmentSchema
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientFeaturesSchema) GetSegmentsOk() ([]SegmentSchema, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *ClientFeaturesSchema) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []SegmentSchema and assigns it to the Segments field.
func (o *ClientFeaturesSchema) SetSegments(v []SegmentSchema) {
	o.Segments = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ClientFeaturesSchema) GetQuery() ClientFeaturesQuerySchema {
	if o == nil || IsNil(o.Query) {
		var ret ClientFeaturesQuerySchema
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientFeaturesSchema) GetQueryOk() (*ClientFeaturesQuerySchema, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ClientFeaturesSchema) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given ClientFeaturesQuerySchema and assigns it to the Query field.
func (o *ClientFeaturesSchema) SetQuery(v ClientFeaturesQuerySchema) {
	o.Query = &v
}

func (o ClientFeaturesSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientFeaturesSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["features"] = o.Features
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return toSerialize, nil
}

type NullableClientFeaturesSchema struct {
	value *ClientFeaturesSchema
	isSet bool
}

func (v NullableClientFeaturesSchema) Get() *ClientFeaturesSchema {
	return v.value
}

func (v *NullableClientFeaturesSchema) Set(val *ClientFeaturesSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableClientFeaturesSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableClientFeaturesSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientFeaturesSchema(val *ClientFeaturesSchema) *NullableClientFeaturesSchema {
	return &NullableClientFeaturesSchema{value: val, isSet: true}
}

func (v NullableClientFeaturesSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientFeaturesSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
