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

// checks if the FeatureEventsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEventsSchema{}

// FeatureEventsSchema struct for FeatureEventsSchema
type FeatureEventsSchema struct {
	Version     *float32      `json:"version,omitempty"`
	ToggleName  *string       `json:"toggleName,omitempty"`
	Events      []EventSchema `json:"events"`
	TotalEvents *int32        `json:"totalEvents,omitempty"`
}

// NewFeatureEventsSchema instantiates a new FeatureEventsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEventsSchema(events []EventSchema) *FeatureEventsSchema {
	this := FeatureEventsSchema{}
	this.Events = events
	return &this
}

// NewFeatureEventsSchemaWithDefaults instantiates a new FeatureEventsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEventsSchemaWithDefaults() *FeatureEventsSchema {
	this := FeatureEventsSchema{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FeatureEventsSchema) GetVersion() float32 {
	if o == nil || IsNil(o.Version) {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEventsSchema) GetVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FeatureEventsSchema) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *FeatureEventsSchema) SetVersion(v float32) {
	o.Version = &v
}

// GetToggleName returns the ToggleName field value if set, zero value otherwise.
func (o *FeatureEventsSchema) GetToggleName() string {
	if o == nil || IsNil(o.ToggleName) {
		var ret string
		return ret
	}
	return *o.ToggleName
}

// GetToggleNameOk returns a tuple with the ToggleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEventsSchema) GetToggleNameOk() (*string, bool) {
	if o == nil || IsNil(o.ToggleName) {
		return nil, false
	}
	return o.ToggleName, true
}

// HasToggleName returns a boolean if a field has been set.
func (o *FeatureEventsSchema) HasToggleName() bool {
	if o != nil && !IsNil(o.ToggleName) {
		return true
	}

	return false
}

// SetToggleName gets a reference to the given string and assigns it to the ToggleName field.
func (o *FeatureEventsSchema) SetToggleName(v string) {
	o.ToggleName = &v
}

// GetEvents returns the Events field value
func (o *FeatureEventsSchema) GetEvents() []EventSchema {
	if o == nil {
		var ret []EventSchema
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *FeatureEventsSchema) GetEventsOk() ([]EventSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *FeatureEventsSchema) SetEvents(v []EventSchema) {
	o.Events = v
}

// GetTotalEvents returns the TotalEvents field value if set, zero value otherwise.
func (o *FeatureEventsSchema) GetTotalEvents() int32 {
	if o == nil || IsNil(o.TotalEvents) {
		var ret int32
		return ret
	}
	return *o.TotalEvents
}

// GetTotalEventsOk returns a tuple with the TotalEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEventsSchema) GetTotalEventsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalEvents) {
		return nil, false
	}
	return o.TotalEvents, true
}

// HasTotalEvents returns a boolean if a field has been set.
func (o *FeatureEventsSchema) HasTotalEvents() bool {
	if o != nil && !IsNil(o.TotalEvents) {
		return true
	}

	return false
}

// SetTotalEvents gets a reference to the given int32 and assigns it to the TotalEvents field.
func (o *FeatureEventsSchema) SetTotalEvents(v int32) {
	o.TotalEvents = &v
}

func (o FeatureEventsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureEventsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ToggleName) {
		toSerialize["toggleName"] = o.ToggleName
	}
	toSerialize["events"] = o.Events
	if !IsNil(o.TotalEvents) {
		toSerialize["totalEvents"] = o.TotalEvents
	}
	return toSerialize, nil
}

type NullableFeatureEventsSchema struct {
	value *FeatureEventsSchema
	isSet bool
}

func (v NullableFeatureEventsSchema) Get() *FeatureEventsSchema {
	return v.value
}

func (v *NullableFeatureEventsSchema) Set(val *FeatureEventsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEventsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEventsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEventsSchema(val *FeatureEventsSchema) *NullableFeatureEventsSchema {
	return &NullableFeatureEventsSchema{value: val, isSet: true}
}

func (v NullableFeatureEventsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureEventsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
