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

// checks if the TagsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagsSchema{}

// TagsSchema A list of tags with a version number
type TagsSchema struct {
	// The version of the schema used to model the tags.
	Version int32 `json:"version"`
	// A list of tags.
	Tags []TagSchema `json:"tags"`
}

// NewTagsSchema instantiates a new TagsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsSchema(version int32, tags []TagSchema) *TagsSchema {
	this := TagsSchema{}
	this.Version = version
	this.Tags = tags
	return &this
}

// NewTagsSchemaWithDefaults instantiates a new TagsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsSchemaWithDefaults() *TagsSchema {
	this := TagsSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *TagsSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagsSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TagsSchema) SetVersion(v int32) {
	o.Version = v
}

// GetTags returns the Tags field value
func (o *TagsSchema) GetTags() []TagSchema {
	if o == nil {
		var ret []TagSchema
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagsSchema) GetTagsOk() ([]TagSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagsSchema) SetTags(v []TagSchema) {
	o.Tags = v
}

func (o TagsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableTagsSchema struct {
	value *TagsSchema
	isSet bool
}

func (v NullableTagsSchema) Get() *TagsSchema {
	return v.value
}

func (v *NullableTagsSchema) Set(val *TagsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsSchema(val *TagsSchema) *NullableTagsSchema {
	return &NullableTagsSchema{value: val, isSet: true}
}

func (v NullableTagsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
