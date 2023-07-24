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

// checks if the LoginHistorySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginHistorySchema{}

// LoginHistorySchema A response model with a list of login events.
type LoginHistorySchema struct {
	Events []LoginEventSchema `json:"events"`
}

// NewLoginHistorySchema instantiates a new LoginHistorySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginHistorySchema(events []LoginEventSchema) *LoginHistorySchema {
	this := LoginHistorySchema{}
	this.Events = events
	return &this
}

// NewLoginHistorySchemaWithDefaults instantiates a new LoginHistorySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginHistorySchemaWithDefaults() *LoginHistorySchema {
	this := LoginHistorySchema{}
	return &this
}

// GetEvents returns the Events field value
func (o *LoginHistorySchema) GetEvents() []LoginEventSchema {
	if o == nil {
		var ret []LoginEventSchema
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *LoginHistorySchema) GetEventsOk() ([]LoginEventSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *LoginHistorySchema) SetEvents(v []LoginEventSchema) {
	o.Events = v
}

func (o LoginHistorySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginHistorySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

type NullableLoginHistorySchema struct {
	value *LoginHistorySchema
	isSet bool
}

func (v NullableLoginHistorySchema) Get() *LoginHistorySchema {
	return v.value
}

func (v *NullableLoginHistorySchema) Set(val *LoginHistorySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginHistorySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginHistorySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginHistorySchema(val *LoginHistorySchema) *NullableLoginHistorySchema {
	return &NullableLoginHistorySchema{value: val, isSet: true}
}

func (v NullableLoginHistorySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginHistorySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}