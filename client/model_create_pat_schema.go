/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreatePatSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePatSchema{}

// CreatePatSchema Describes the properties required to create a [personal access token](https://docs.getunleash.io/reference/api-tokens-and-client-keys#personal-access-tokens), or PAT. PATs are automatically scoped to the authenticated user.
type CreatePatSchema struct {
	// The PAT's description.
	Description string `json:"description"`
	// The PAT's expiration date.
	ExpiresAt            time.Time `json:"expiresAt"`
	AdditionalProperties map[string]interface{}
}

type _CreatePatSchema CreatePatSchema

// NewCreatePatSchema instantiates a new CreatePatSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePatSchema(description string, expiresAt time.Time) *CreatePatSchema {
	this := CreatePatSchema{}
	this.Description = description
	this.ExpiresAt = expiresAt
	return &this
}

// NewCreatePatSchemaWithDefaults instantiates a new CreatePatSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePatSchemaWithDefaults() *CreatePatSchema {
	this := CreatePatSchema{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreatePatSchema) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreatePatSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreatePatSchema) SetDescription(v string) {
	o.Description = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *CreatePatSchema) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *CreatePatSchema) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *CreatePatSchema) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

func (o CreatePatSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePatSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["expiresAt"] = o.ExpiresAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePatSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"expiresAt",
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

	varCreatePatSchema := _CreatePatSchema{}

	err = json.Unmarshal(data, &varCreatePatSchema)

	if err != nil {
		return err
	}

	*o = CreatePatSchema(varCreatePatSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "expiresAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePatSchema struct {
	value *CreatePatSchema
	isSet bool
}

func (v NullableCreatePatSchema) Get() *CreatePatSchema {
	return v.value
}

func (v *NullableCreatePatSchema) Set(val *CreatePatSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePatSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePatSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePatSchema(val *CreatePatSchema) *NullableCreatePatSchema {
	return &NullableCreatePatSchema{value: val, isSet: true}
}

func (v NullableCreatePatSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePatSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
