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

// checks if the TagWithVersionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagWithVersionSchema{}

// TagWithVersionSchema A tag with a version number representing the schema used to model the tag.
type TagWithVersionSchema struct {
	// The version of the schema used to model the tag.
	Version int32     `json:"version"`
	Tag     TagSchema `json:"tag"`
}

// NewTagWithVersionSchema instantiates a new TagWithVersionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagWithVersionSchema(version int32, tag TagSchema) *TagWithVersionSchema {
	this := TagWithVersionSchema{}
	this.Version = version
	this.Tag = tag
	return &this
}

// NewTagWithVersionSchemaWithDefaults instantiates a new TagWithVersionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithVersionSchemaWithDefaults() *TagWithVersionSchema {
	this := TagWithVersionSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *TagWithVersionSchema) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagWithVersionSchema) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *TagWithVersionSchema) SetVersion(v int32) {
	o.Version = v
}

// GetTag returns the Tag field value
func (o *TagWithVersionSchema) GetTag() TagSchema {
	if o == nil {
		var ret TagSchema
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *TagWithVersionSchema) GetTagOk() (*TagSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *TagWithVersionSchema) SetTag(v TagSchema) {
	o.Tag = v
}

func (o TagWithVersionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagWithVersionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["tag"] = o.Tag
	return toSerialize, nil
}

type NullableTagWithVersionSchema struct {
	value *TagWithVersionSchema
	isSet bool
}

func (v NullableTagWithVersionSchema) Get() *TagWithVersionSchema {
	return v.value
}

func (v *NullableTagWithVersionSchema) Set(val *TagWithVersionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableTagWithVersionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableTagWithVersionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagWithVersionSchema(val *TagWithVersionSchema) *NullableTagWithVersionSchema {
	return &NullableTagWithVersionSchema{value: val, isSet: true}
}

func (v NullableTagWithVersionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagWithVersionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
