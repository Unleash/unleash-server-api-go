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

// checks if the SegmentsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentsSchema{}

// SegmentsSchema struct for SegmentsSchema
type SegmentsSchema struct {
	Segments []AdminSegmentSchema `json:"segments,omitempty"`
}

// NewSegmentsSchema instantiates a new SegmentsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentsSchema() *SegmentsSchema {
	this := SegmentsSchema{}
	return &this
}

// NewSegmentsSchemaWithDefaults instantiates a new SegmentsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentsSchemaWithDefaults() *SegmentsSchema {
	this := SegmentsSchema{}
	return &this
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *SegmentsSchema) GetSegments() []AdminSegmentSchema {
	if o == nil || IsNil(o.Segments) {
		var ret []AdminSegmentSchema
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentsSchema) GetSegmentsOk() ([]AdminSegmentSchema, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *SegmentsSchema) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []AdminSegmentSchema and assigns it to the Segments field.
func (o *SegmentsSchema) SetSegments(v []AdminSegmentSchema) {
	o.Segments = v
}

func (o SegmentsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	return toSerialize, nil
}

type NullableSegmentsSchema struct {
	value *SegmentsSchema
	isSet bool
}

func (v NullableSegmentsSchema) Get() *SegmentsSchema {
	return v.value
}

func (v *NullableSegmentsSchema) Set(val *SegmentsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentsSchema(val *SegmentsSchema) *NullableSegmentsSchema {
	return &NullableSegmentsSchema{value: val, isSet: true}
}

func (v NullableSegmentsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}