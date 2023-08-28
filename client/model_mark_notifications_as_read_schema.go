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

// checks if the MarkNotificationsAsReadSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkNotificationsAsReadSchema{}

// MarkNotificationsAsReadSchema Data used to mark notifications as being read
type MarkNotificationsAsReadSchema struct {
	// A list of IDs belonging to the notifications you want to mark as read.
	Notifications []int32 `json:"notifications"`
}

// NewMarkNotificationsAsReadSchema instantiates a new MarkNotificationsAsReadSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkNotificationsAsReadSchema(notifications []int32) *MarkNotificationsAsReadSchema {
	this := MarkNotificationsAsReadSchema{}
	this.Notifications = notifications
	return &this
}

// NewMarkNotificationsAsReadSchemaWithDefaults instantiates a new MarkNotificationsAsReadSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkNotificationsAsReadSchemaWithDefaults() *MarkNotificationsAsReadSchema {
	this := MarkNotificationsAsReadSchema{}
	return &this
}

// GetNotifications returns the Notifications field value
func (o *MarkNotificationsAsReadSchema) GetNotifications() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *MarkNotificationsAsReadSchema) GetNotificationsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications, true
}

// SetNotifications sets field value
func (o *MarkNotificationsAsReadSchema) SetNotifications(v []int32) {
	o.Notifications = v
}

func (o MarkNotificationsAsReadSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkNotificationsAsReadSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifications"] = o.Notifications
	return toSerialize, nil
}

type NullableMarkNotificationsAsReadSchema struct {
	value *MarkNotificationsAsReadSchema
	isSet bool
}

func (v NullableMarkNotificationsAsReadSchema) Get() *MarkNotificationsAsReadSchema {
	return v.value
}

func (v *NullableMarkNotificationsAsReadSchema) Set(val *MarkNotificationsAsReadSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkNotificationsAsReadSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkNotificationsAsReadSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkNotificationsAsReadSchema(val *MarkNotificationsAsReadSchema) *NullableMarkNotificationsAsReadSchema {
	return &NullableMarkNotificationsAsReadSchema{value: val, isSet: true}
}

func (v NullableMarkNotificationsAsReadSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkNotificationsAsReadSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
