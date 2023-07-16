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

// checks if the AddonTypeSchemaAlertsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonTypeSchemaAlertsInner{}

// AddonTypeSchemaAlertsInner struct for AddonTypeSchemaAlertsInner
type AddonTypeSchemaAlertsInner struct {
	// The type of alert. This determines the color of the alert.
	Type string `json:"type"`
	// The text of the alert. This is what will be displayed to the user.
	Text string `json:"text"`
}

// NewAddonTypeSchemaAlertsInner instantiates a new AddonTypeSchemaAlertsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonTypeSchemaAlertsInner(type_ string, text string) *AddonTypeSchemaAlertsInner {
	this := AddonTypeSchemaAlertsInner{}
	this.Type = type_
	this.Text = text
	return &this
}

// NewAddonTypeSchemaAlertsInnerWithDefaults instantiates a new AddonTypeSchemaAlertsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonTypeSchemaAlertsInnerWithDefaults() *AddonTypeSchemaAlertsInner {
	this := AddonTypeSchemaAlertsInner{}
	return &this
}

// GetType returns the Type field value
func (o *AddonTypeSchemaAlertsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchemaAlertsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddonTypeSchemaAlertsInner) SetType(v string) {
	o.Type = v
}

// GetText returns the Text field value
func (o *AddonTypeSchemaAlertsInner) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AddonTypeSchemaAlertsInner) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AddonTypeSchemaAlertsInner) SetText(v string) {
	o.Text = v
}

func (o AddonTypeSchemaAlertsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonTypeSchemaAlertsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableAddonTypeSchemaAlertsInner struct {
	value *AddonTypeSchemaAlertsInner
	isSet bool
}

func (v NullableAddonTypeSchemaAlertsInner) Get() *AddonTypeSchemaAlertsInner {
	return v.value
}

func (v *NullableAddonTypeSchemaAlertsInner) Set(val *AddonTypeSchemaAlertsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonTypeSchemaAlertsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonTypeSchemaAlertsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonTypeSchemaAlertsInner(val *AddonTypeSchemaAlertsInner) *NullableAddonTypeSchemaAlertsInner {
	return &NullableAddonTypeSchemaAlertsInner{value: val, isSet: true}
}

func (v NullableAddonTypeSchemaAlertsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonTypeSchemaAlertsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
