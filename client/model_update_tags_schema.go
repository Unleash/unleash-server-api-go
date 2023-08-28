/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateTagsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTagsSchema{}

// UpdateTagsSchema Represents a set of changes to a feature's tags, such as adding or removing tags.
type UpdateTagsSchema struct {
	// Tags to add to the feature.
	AddedTags []TagSchema `json:"addedTags"`
	// Tags to remove from the feature.
	RemovedTags []TagSchema `json:"removedTags"`
}

// NewUpdateTagsSchema instantiates a new UpdateTagsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTagsSchema(addedTags []TagSchema, removedTags []TagSchema) *UpdateTagsSchema {
	this := UpdateTagsSchema{}
	this.AddedTags = addedTags
	this.RemovedTags = removedTags
	return &this
}

// NewUpdateTagsSchemaWithDefaults instantiates a new UpdateTagsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTagsSchemaWithDefaults() *UpdateTagsSchema {
	this := UpdateTagsSchema{}
	return &this
}

// GetAddedTags returns the AddedTags field value
func (o *UpdateTagsSchema) GetAddedTags() []TagSchema {
	if o == nil {
		var ret []TagSchema
		return ret
	}

	return o.AddedTags
}

// GetAddedTagsOk returns a tuple with the AddedTags field value
// and a boolean to check if the value has been set.
func (o *UpdateTagsSchema) GetAddedTagsOk() ([]TagSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedTags, true
}

// SetAddedTags sets field value
func (o *UpdateTagsSchema) SetAddedTags(v []TagSchema) {
	o.AddedTags = v
}

// GetRemovedTags returns the RemovedTags field value
func (o *UpdateTagsSchema) GetRemovedTags() []TagSchema {
	if o == nil {
		var ret []TagSchema
		return ret
	}

	return o.RemovedTags
}

// GetRemovedTagsOk returns a tuple with the RemovedTags field value
// and a boolean to check if the value has been set.
func (o *UpdateTagsSchema) GetRemovedTagsOk() ([]TagSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedTags, true
}

// SetRemovedTags sets field value
func (o *UpdateTagsSchema) SetRemovedTags(v []TagSchema) {
	o.RemovedTags = v
}

func (o UpdateTagsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTagsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addedTags"] = o.AddedTags
	toSerialize["removedTags"] = o.RemovedTags
	return toSerialize, nil
}

type NullableUpdateTagsSchema struct {
	value *UpdateTagsSchema
	isSet bool
}

func (v NullableUpdateTagsSchema) Get() *UpdateTagsSchema {
	return v.value
}

func (v *NullableUpdateTagsSchema) Set(val *UpdateTagsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTagsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTagsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTagsSchema(val *UpdateTagsSchema) *NullableUpdateTagsSchema {
	return &NullableUpdateTagsSchema{value: val, isSet: true}
}

func (v NullableUpdateTagsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTagsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
