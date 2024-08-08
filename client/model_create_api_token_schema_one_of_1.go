/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreateApiTokenSchemaOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiTokenSchemaOneOf1{}

// CreateApiTokenSchemaOneOf1 struct for CreateApiTokenSchemaOneOf1
type CreateApiTokenSchemaOneOf1 struct {
	// The time when this token should expire.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// An admin token. Must be the string \"admin\" (not case sensitive).
	Type string `json:"type" validate:"regexp=^[Aa][Dd][Mm][Ii][Nn]$"`
	// The name of the token. This property was deprecated in v5. Use `tokenName` instead.
	// Deprecated
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _CreateApiTokenSchemaOneOf1 CreateApiTokenSchemaOneOf1

// NewCreateApiTokenSchemaOneOf1 instantiates a new CreateApiTokenSchemaOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiTokenSchemaOneOf1(type_ string, username string) *CreateApiTokenSchemaOneOf1 {
	this := CreateApiTokenSchemaOneOf1{}
	this.Type = type_
	this.Username = username
	return &this
}

// NewCreateApiTokenSchemaOneOf1WithDefaults instantiates a new CreateApiTokenSchemaOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiTokenSchemaOneOf1WithDefaults() *CreateApiTokenSchemaOneOf1 {
	this := CreateApiTokenSchemaOneOf1{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateApiTokenSchemaOneOf1) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf1) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateApiTokenSchemaOneOf1) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreateApiTokenSchemaOneOf1) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetType returns the Type field value
func (o *CreateApiTokenSchemaOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateApiTokenSchemaOneOf1) SetType(v string) {
	o.Type = v
}

// GetUsername returns the Username field value
// Deprecated
func (o *CreateApiTokenSchemaOneOf1) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateApiTokenSchemaOneOf1) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
// Deprecated
func (o *CreateApiTokenSchemaOneOf1) SetUsername(v string) {
	o.Username = v
}

func (o CreateApiTokenSchemaOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiTokenSchemaOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["type"] = o.Type
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateApiTokenSchemaOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"username",
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

	varCreateApiTokenSchemaOneOf1 := _CreateApiTokenSchemaOneOf1{}

	err = json.Unmarshal(data, &varCreateApiTokenSchemaOneOf1)

	if err != nil {
		return err
	}

	*o = CreateApiTokenSchemaOneOf1(varCreateApiTokenSchemaOneOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "type")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateApiTokenSchemaOneOf1 struct {
	value *CreateApiTokenSchemaOneOf1
	isSet bool
}

func (v NullableCreateApiTokenSchemaOneOf1) Get() *CreateApiTokenSchemaOneOf1 {
	return v.value
}

func (v *NullableCreateApiTokenSchemaOneOf1) Set(val *CreateApiTokenSchemaOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiTokenSchemaOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiTokenSchemaOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiTokenSchemaOneOf1(val *CreateApiTokenSchemaOneOf1) *NullableCreateApiTokenSchemaOneOf1 {
	return &NullableCreateApiTokenSchemaOneOf1{value: val, isSet: true}
}

func (v NullableCreateApiTokenSchemaOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiTokenSchemaOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
