/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.12+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreateApiTokenSchemaOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiTokenSchemaOneOf{}

// CreateApiTokenSchemaOneOf struct for CreateApiTokenSchemaOneOf
type CreateApiTokenSchemaOneOf struct {
	// The time when this token should expire.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// An admin token. Must be the string \"admin\" (not case sensitive).
	Type string `json:"type" validate:"regexp=^[Aa][Dd][Mm][Ii][Nn]$"`
	// The name of the token.
	TokenName            string `json:"tokenName"`
	AdditionalProperties map[string]interface{}
}

type _CreateApiTokenSchemaOneOf CreateApiTokenSchemaOneOf

// NewCreateApiTokenSchemaOneOf instantiates a new CreateApiTokenSchemaOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiTokenSchemaOneOf(type_ string, tokenName string) *CreateApiTokenSchemaOneOf {
	this := CreateApiTokenSchemaOneOf{}
	this.Type = type_
	this.TokenName = tokenName
	return &this
}

// NewCreateApiTokenSchemaOneOfWithDefaults instantiates a new CreateApiTokenSchemaOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiTokenSchemaOneOfWithDefaults() *CreateApiTokenSchemaOneOf {
	this := CreateApiTokenSchemaOneOf{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateApiTokenSchemaOneOf) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateApiTokenSchemaOneOf) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreateApiTokenSchemaOneOf) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetType returns the Type field value
func (o *CreateApiTokenSchemaOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateApiTokenSchemaOneOf) SetType(v string) {
	o.Type = v
}

// GetTokenName returns the TokenName field value
func (o *CreateApiTokenSchemaOneOf) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *CreateApiTokenSchemaOneOf) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *CreateApiTokenSchemaOneOf) SetTokenName(v string) {
	o.TokenName = v
}

func (o CreateApiTokenSchemaOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiTokenSchemaOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	toSerialize["type"] = o.Type
	toSerialize["tokenName"] = o.TokenName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateApiTokenSchemaOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"tokenName",
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

	varCreateApiTokenSchemaOneOf := _CreateApiTokenSchemaOneOf{}

	err = json.Unmarshal(data, &varCreateApiTokenSchemaOneOf)

	if err != nil {
		return err
	}

	*o = CreateApiTokenSchemaOneOf(varCreateApiTokenSchemaOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "type")
		delete(additionalProperties, "tokenName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateApiTokenSchemaOneOf struct {
	value *CreateApiTokenSchemaOneOf
	isSet bool
}

func (v NullableCreateApiTokenSchemaOneOf) Get() *CreateApiTokenSchemaOneOf {
	return v.value
}

func (v *NullableCreateApiTokenSchemaOneOf) Set(val *CreateApiTokenSchemaOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiTokenSchemaOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiTokenSchemaOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiTokenSchemaOneOf(val *CreateApiTokenSchemaOneOf) *NullableCreateApiTokenSchemaOneOf {
	return &NullableCreateApiTokenSchemaOneOf{value: val, isSet: true}
}

func (v NullableCreateApiTokenSchemaOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiTokenSchemaOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
