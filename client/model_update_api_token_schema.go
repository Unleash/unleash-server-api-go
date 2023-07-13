/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the UpdateApiTokenSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApiTokenSchema{}

// UpdateApiTokenSchema An object with fields to updated for a given API token.
type UpdateApiTokenSchema struct {
	// The new time when this token should expire.
	ExpiresAt time.Time `json:"expiresAt"`
}

// NewUpdateApiTokenSchema instantiates a new UpdateApiTokenSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiTokenSchema(expiresAt time.Time) *UpdateApiTokenSchema {
	this := UpdateApiTokenSchema{}
	this.ExpiresAt = expiresAt
	return &this
}

// NewUpdateApiTokenSchemaWithDefaults instantiates a new UpdateApiTokenSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiTokenSchemaWithDefaults() *UpdateApiTokenSchema {
	this := UpdateApiTokenSchema{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value
func (o *UpdateApiTokenSchema) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *UpdateApiTokenSchema) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *UpdateApiTokenSchema) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o UpdateApiTokenSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApiTokenSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiresAt"] = o.ExpiresAt
	return toSerialize, nil
}

type NullableUpdateApiTokenSchema struct {
	value *UpdateApiTokenSchema
	isSet bool
}

func (v NullableUpdateApiTokenSchema) Get() *UpdateApiTokenSchema {
	return v.value
}

func (v *NullableUpdateApiTokenSchema) Set(val *UpdateApiTokenSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiTokenSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiTokenSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiTokenSchema(val *UpdateApiTokenSchema) *NullableUpdateApiTokenSchema {
	return &NullableUpdateApiTokenSchema{value: val, isSet: true}
}

func (v NullableUpdateApiTokenSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiTokenSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


