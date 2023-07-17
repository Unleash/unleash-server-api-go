/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the CreateApiTokenSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiTokenSchema{}

// CreateApiTokenSchema The data required to create an [Unleash API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys).
type CreateApiTokenSchema struct {
	// The time when this token should expire.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewCreateApiTokenSchema instantiates a new CreateApiTokenSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiTokenSchema() *CreateApiTokenSchema {
	this := CreateApiTokenSchema{}
	return &this
}

// NewCreateApiTokenSchemaWithDefaults instantiates a new CreateApiTokenSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiTokenSchemaWithDefaults() *CreateApiTokenSchema {
	this := CreateApiTokenSchema{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateApiTokenSchema) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchema) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateApiTokenSchema) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreateApiTokenSchema) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o CreateApiTokenSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiTokenSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableCreateApiTokenSchema struct {
	value *CreateApiTokenSchema
	isSet bool
}

func (v NullableCreateApiTokenSchema) Get() *CreateApiTokenSchema {
	return v.value
}

func (v *NullableCreateApiTokenSchema) Set(val *CreateApiTokenSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiTokenSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiTokenSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiTokenSchema(val *CreateApiTokenSchema) *NullableCreateApiTokenSchema {
	return &NullableCreateApiTokenSchema{value: val, isSet: true}
}

func (v NullableCreateApiTokenSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiTokenSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
