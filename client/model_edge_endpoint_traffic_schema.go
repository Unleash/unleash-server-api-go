/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the EdgeEndpointTrafficSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeEndpointTrafficSchema{}

// EdgeEndpointTrafficSchema Represents traffic data for a single endpoint in Edge
type EdgeEndpointTrafficSchema struct {
	// Number of 20x requests
	Requests200 *float32 `json:"requests200,omitempty"`
	// Number of 30x requests
	Requests304          *float32 `json:"requests304,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EdgeEndpointTrafficSchema EdgeEndpointTrafficSchema

// NewEdgeEndpointTrafficSchema instantiates a new EdgeEndpointTrafficSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeEndpointTrafficSchema() *EdgeEndpointTrafficSchema {
	this := EdgeEndpointTrafficSchema{}
	return &this
}

// NewEdgeEndpointTrafficSchemaWithDefaults instantiates a new EdgeEndpointTrafficSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeEndpointTrafficSchemaWithDefaults() *EdgeEndpointTrafficSchema {
	this := EdgeEndpointTrafficSchema{}
	return &this
}

// GetRequests200 returns the Requests200 field value if set, zero value otherwise.
func (o *EdgeEndpointTrafficSchema) GetRequests200() float32 {
	if o == nil || IsNil(o.Requests200) {
		var ret float32
		return ret
	}
	return *o.Requests200
}

// GetRequests200Ok returns a tuple with the Requests200 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeEndpointTrafficSchema) GetRequests200Ok() (*float32, bool) {
	if o == nil || IsNil(o.Requests200) {
		return nil, false
	}
	return o.Requests200, true
}

// HasRequests200 returns a boolean if a field has been set.
func (o *EdgeEndpointTrafficSchema) HasRequests200() bool {
	if o != nil && !IsNil(o.Requests200) {
		return true
	}

	return false
}

// SetRequests200 gets a reference to the given float32 and assigns it to the Requests200 field.
func (o *EdgeEndpointTrafficSchema) SetRequests200(v float32) {
	o.Requests200 = &v
}

// GetRequests304 returns the Requests304 field value if set, zero value otherwise.
func (o *EdgeEndpointTrafficSchema) GetRequests304() float32 {
	if o == nil || IsNil(o.Requests304) {
		var ret float32
		return ret
	}
	return *o.Requests304
}

// GetRequests304Ok returns a tuple with the Requests304 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeEndpointTrafficSchema) GetRequests304Ok() (*float32, bool) {
	if o == nil || IsNil(o.Requests304) {
		return nil, false
	}
	return o.Requests304, true
}

// HasRequests304 returns a boolean if a field has been set.
func (o *EdgeEndpointTrafficSchema) HasRequests304() bool {
	if o != nil && !IsNil(o.Requests304) {
		return true
	}

	return false
}

// SetRequests304 gets a reference to the given float32 and assigns it to the Requests304 field.
func (o *EdgeEndpointTrafficSchema) SetRequests304(v float32) {
	o.Requests304 = &v
}

func (o EdgeEndpointTrafficSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeEndpointTrafficSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests200) {
		toSerialize["requests200"] = o.Requests200
	}
	if !IsNil(o.Requests304) {
		toSerialize["requests304"] = o.Requests304
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EdgeEndpointTrafficSchema) UnmarshalJSON(data []byte) (err error) {
	varEdgeEndpointTrafficSchema := _EdgeEndpointTrafficSchema{}

	err = json.Unmarshal(data, &varEdgeEndpointTrafficSchema)

	if err != nil {
		return err
	}

	*o = EdgeEndpointTrafficSchema(varEdgeEndpointTrafficSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requests200")
		delete(additionalProperties, "requests304")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEdgeEndpointTrafficSchema struct {
	value *EdgeEndpointTrafficSchema
	isSet bool
}

func (v NullableEdgeEndpointTrafficSchema) Get() *EdgeEndpointTrafficSchema {
	return v.value
}

func (v *NullableEdgeEndpointTrafficSchema) Set(val *EdgeEndpointTrafficSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeEndpointTrafficSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeEndpointTrafficSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeEndpointTrafficSchema(val *EdgeEndpointTrafficSchema) *NullableEdgeEndpointTrafficSchema {
	return &NullableEdgeEndpointTrafficSchema{value: val, isSet: true}
}

func (v NullableEdgeEndpointTrafficSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeEndpointTrafficSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
