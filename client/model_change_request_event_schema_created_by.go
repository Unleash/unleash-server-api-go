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

// checks if the ChangeRequestEventSchemaCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestEventSchemaCreatedBy{}

// ChangeRequestEventSchemaCreatedBy struct for ChangeRequestEventSchemaCreatedBy
type ChangeRequestEventSchemaCreatedBy struct {
	Username NullableString `json:"username,omitempty"`
	ImageUrl NullableString `json:"imageUrl,omitempty"`
}

// NewChangeRequestEventSchemaCreatedBy instantiates a new ChangeRequestEventSchemaCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestEventSchemaCreatedBy() *ChangeRequestEventSchemaCreatedBy {
	this := ChangeRequestEventSchemaCreatedBy{}
	return &this
}

// NewChangeRequestEventSchemaCreatedByWithDefaults instantiates a new ChangeRequestEventSchemaCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestEventSchemaCreatedByWithDefaults() *ChangeRequestEventSchemaCreatedBy {
	this := ChangeRequestEventSchemaCreatedBy{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeRequestEventSchemaCreatedBy) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeRequestEventSchemaCreatedBy) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ChangeRequestEventSchemaCreatedBy) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *ChangeRequestEventSchemaCreatedBy) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ChangeRequestEventSchemaCreatedBy) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ChangeRequestEventSchemaCreatedBy) UnsetUsername() {
	o.Username.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeRequestEventSchemaCreatedBy) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeRequestEventSchemaCreatedBy) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ChangeRequestEventSchemaCreatedBy) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *ChangeRequestEventSchemaCreatedBy) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *ChangeRequestEventSchemaCreatedBy) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *ChangeRequestEventSchemaCreatedBy) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

func (o ChangeRequestEventSchemaCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestEventSchemaCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	return toSerialize, nil
}

type NullableChangeRequestEventSchemaCreatedBy struct {
	value *ChangeRequestEventSchemaCreatedBy
	isSet bool
}

func (v NullableChangeRequestEventSchemaCreatedBy) Get() *ChangeRequestEventSchemaCreatedBy {
	return v.value
}

func (v *NullableChangeRequestEventSchemaCreatedBy) Set(val *ChangeRequestEventSchemaCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestEventSchemaCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestEventSchemaCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestEventSchemaCreatedBy(val *ChangeRequestEventSchemaCreatedBy) *NullableChangeRequestEventSchemaCreatedBy {
	return &NullableChangeRequestEventSchemaCreatedBy{value: val, isSet: true}
}

func (v NullableChangeRequestEventSchemaCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestEventSchemaCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
