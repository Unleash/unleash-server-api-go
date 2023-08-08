/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ImportTogglesValidateItemSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportTogglesValidateItemSchema{}

// ImportTogglesValidateItemSchema A description of an error or warning pertaining to a feature toggle import job.
type ImportTogglesValidateItemSchema struct {
	// The validation error message
	Message string `json:"message"`
	// The items affected by this error message
	AffectedItems []string `json:"affectedItems"`
}

// NewImportTogglesValidateItemSchema instantiates a new ImportTogglesValidateItemSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportTogglesValidateItemSchema(message string, affectedItems []string) *ImportTogglesValidateItemSchema {
	this := ImportTogglesValidateItemSchema{}
	this.Message = message
	this.AffectedItems = affectedItems
	return &this
}

// NewImportTogglesValidateItemSchemaWithDefaults instantiates a new ImportTogglesValidateItemSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportTogglesValidateItemSchemaWithDefaults() *ImportTogglesValidateItemSchema {
	this := ImportTogglesValidateItemSchema{}
	return &this
}

// GetMessage returns the Message field value
func (o *ImportTogglesValidateItemSchema) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ImportTogglesValidateItemSchema) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ImportTogglesValidateItemSchema) SetMessage(v string) {
	o.Message = v
}

// GetAffectedItems returns the AffectedItems field value
func (o *ImportTogglesValidateItemSchema) GetAffectedItems() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AffectedItems
}

// GetAffectedItemsOk returns a tuple with the AffectedItems field value
// and a boolean to check if the value has been set.
func (o *ImportTogglesValidateItemSchema) GetAffectedItemsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffectedItems, true
}

// SetAffectedItems sets field value
func (o *ImportTogglesValidateItemSchema) SetAffectedItems(v []string) {
	o.AffectedItems = v
}

func (o ImportTogglesValidateItemSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportTogglesValidateItemSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["affectedItems"] = o.AffectedItems
	return toSerialize, nil
}

type NullableImportTogglesValidateItemSchema struct {
	value *ImportTogglesValidateItemSchema
	isSet bool
}

func (v NullableImportTogglesValidateItemSchema) Get() *ImportTogglesValidateItemSchema {
	return v.value
}

func (v *NullableImportTogglesValidateItemSchema) Set(val *ImportTogglesValidateItemSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableImportTogglesValidateItemSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableImportTogglesValidateItemSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportTogglesValidateItemSchema(val *ImportTogglesValidateItemSchema) *NullableImportTogglesValidateItemSchema {
	return &NullableImportTogglesValidateItemSchema{value: val, isSet: true}
}

func (v NullableImportTogglesValidateItemSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportTogglesValidateItemSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
