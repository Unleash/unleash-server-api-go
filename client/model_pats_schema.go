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

// checks if the PatsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatsSchema{}

// PatsSchema Contains a collection of [Personal Access Tokens](https://docs.getunleash.io/how-to/how-to-create-personal-access-tokens).
type PatsSchema struct {
	// A collection of Personal Access Tokens
	Pats []PatSchema `json:"pats,omitempty"`
}

// NewPatsSchema instantiates a new PatsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatsSchema() *PatsSchema {
	this := PatsSchema{}
	return &this
}

// NewPatsSchemaWithDefaults instantiates a new PatsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatsSchemaWithDefaults() *PatsSchema {
	this := PatsSchema{}
	return &this
}

// GetPats returns the Pats field value if set, zero value otherwise.
func (o *PatsSchema) GetPats() []PatSchema {
	if o == nil || IsNil(o.Pats) {
		var ret []PatSchema
		return ret
	}
	return o.Pats
}

// GetPatsOk returns a tuple with the Pats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatsSchema) GetPatsOk() ([]PatSchema, bool) {
	if o == nil || IsNil(o.Pats) {
		return nil, false
	}
	return o.Pats, true
}

// HasPats returns a boolean if a field has been set.
func (o *PatsSchema) HasPats() bool {
	if o != nil && !IsNil(o.Pats) {
		return true
	}

	return false
}

// SetPats gets a reference to the given []PatSchema and assigns it to the Pats field.
func (o *PatsSchema) SetPats(v []PatSchema) {
	o.Pats = v
}

func (o PatsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pats) {
		toSerialize["pats"] = o.Pats
	}
	return toSerialize, nil
}

type NullablePatsSchema struct {
	value *PatsSchema
	isSet bool
}

func (v NullablePatsSchema) Get() *PatsSchema {
	return v.value
}

func (v *NullablePatsSchema) Set(val *PatsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePatsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePatsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatsSchema(val *PatsSchema) *NullablePatsSchema {
	return &NullablePatsSchema{value: val, isSet: true}
}

func (v NullablePatsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
