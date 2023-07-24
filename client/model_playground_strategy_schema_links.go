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

// checks if the PlaygroundStrategySchemaLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchemaLinks{}

// PlaygroundStrategySchemaLinks A set of links to actions you can perform on this strategy
type PlaygroundStrategySchemaLinks struct {
	Edit string `json:"edit"`
}

// NewPlaygroundStrategySchemaLinks instantiates a new PlaygroundStrategySchemaLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchemaLinks(edit string) *PlaygroundStrategySchemaLinks {
	this := PlaygroundStrategySchemaLinks{}
	this.Edit = edit
	return &this
}

// NewPlaygroundStrategySchemaLinksWithDefaults instantiates a new PlaygroundStrategySchemaLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaLinksWithDefaults() *PlaygroundStrategySchemaLinks {
	this := PlaygroundStrategySchemaLinks{}
	return &this
}

// GetEdit returns the Edit field value
func (o *PlaygroundStrategySchemaLinks) GetEdit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edit
}

// GetEditOk returns a tuple with the Edit field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaLinks) GetEditOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edit, true
}

// SetEdit sets field value
func (o *PlaygroundStrategySchemaLinks) SetEdit(v string) {
	o.Edit = v
}

func (o PlaygroundStrategySchemaLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchemaLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["edit"] = o.Edit
	return toSerialize, nil
}

type NullablePlaygroundStrategySchemaLinks struct {
	value *PlaygroundStrategySchemaLinks
	isSet bool
}

func (v NullablePlaygroundStrategySchemaLinks) Get() *PlaygroundStrategySchemaLinks {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaLinks) Set(val *PlaygroundStrategySchemaLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaLinks(val *PlaygroundStrategySchemaLinks) *NullablePlaygroundStrategySchemaLinks {
	return &NullablePlaygroundStrategySchemaLinks{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
