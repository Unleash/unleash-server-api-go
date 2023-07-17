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

// checks if the AddonsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonsSchema{}

// AddonsSchema An object containing two things: 1. A list of all [addons](https://docs.getunleash.io/reference/addons) defined on this Unleash instance 2. A list of all addon providers defined on this instance
type AddonsSchema struct {
	// All the addons that exist on this instance of Unleash.
	Addons []AddonSchema `json:"addons"`
	// A list of  all available addon providers, along with their parameters and descriptions.
	Providers []AddonTypeSchema `json:"providers"`
}

// NewAddonsSchema instantiates a new AddonsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonsSchema(addons []AddonSchema, providers []AddonTypeSchema) *AddonsSchema {
	this := AddonsSchema{}
	this.Addons = addons
	this.Providers = providers
	return &this
}

// NewAddonsSchemaWithDefaults instantiates a new AddonsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonsSchemaWithDefaults() *AddonsSchema {
	this := AddonsSchema{}
	return &this
}

// GetAddons returns the Addons field value
func (o *AddonsSchema) GetAddons() []AddonSchema {
	if o == nil {
		var ret []AddonSchema
		return ret
	}

	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value
// and a boolean to check if the value has been set.
func (o *AddonsSchema) GetAddonsOk() ([]AddonSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addons, true
}

// SetAddons sets field value
func (o *AddonsSchema) SetAddons(v []AddonSchema) {
	o.Addons = v
}

// GetProviders returns the Providers field value
func (o *AddonsSchema) GetProviders() []AddonTypeSchema {
	if o == nil {
		var ret []AddonTypeSchema
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *AddonsSchema) GetProvidersOk() ([]AddonTypeSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *AddonsSchema) SetProviders(v []AddonTypeSchema) {
	o.Providers = v
}

func (o AddonsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addons"] = o.Addons
	toSerialize["providers"] = o.Providers
	return toSerialize, nil
}

type NullableAddonsSchema struct {
	value *AddonsSchema
	isSet bool
}

func (v NullableAddonsSchema) Get() *AddonsSchema {
	return v.value
}

func (v *NullableAddonsSchema) Set(val *AddonsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonsSchema(val *AddonsSchema) *NullableAddonsSchema {
	return &NullableAddonsSchema{value: val, isSet: true}
}

func (v NullableAddonsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
