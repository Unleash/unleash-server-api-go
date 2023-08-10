/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ChangeRequestApprovalSchemaCreatedBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestApprovalSchemaCreatedBy{}

// ChangeRequestApprovalSchemaCreatedBy struct for ChangeRequestApprovalSchemaCreatedBy
type ChangeRequestApprovalSchemaCreatedBy struct {
	Id       *float32 `json:"id,omitempty"`
	Username *string  `json:"username,omitempty"`
	ImageUrl *string  `json:"imageUrl,omitempty"`
}

// NewChangeRequestApprovalSchemaCreatedBy instantiates a new ChangeRequestApprovalSchemaCreatedBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestApprovalSchemaCreatedBy() *ChangeRequestApprovalSchemaCreatedBy {
	this := ChangeRequestApprovalSchemaCreatedBy{}
	return &this
}

// NewChangeRequestApprovalSchemaCreatedByWithDefaults instantiates a new ChangeRequestApprovalSchemaCreatedBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestApprovalSchemaCreatedByWithDefaults() *ChangeRequestApprovalSchemaCreatedBy {
	this := ChangeRequestApprovalSchemaCreatedBy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ChangeRequestApprovalSchemaCreatedBy) SetId(v float32) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ChangeRequestApprovalSchemaCreatedBy) SetUsername(v string) {
	o.Username = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ChangeRequestApprovalSchemaCreatedBy) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ChangeRequestApprovalSchemaCreatedBy) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o ChangeRequestApprovalSchemaCreatedBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestApprovalSchemaCreatedBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	return toSerialize, nil
}

type NullableChangeRequestApprovalSchemaCreatedBy struct {
	value *ChangeRequestApprovalSchemaCreatedBy
	isSet bool
}

func (v NullableChangeRequestApprovalSchemaCreatedBy) Get() *ChangeRequestApprovalSchemaCreatedBy {
	return v.value
}

func (v *NullableChangeRequestApprovalSchemaCreatedBy) Set(val *ChangeRequestApprovalSchemaCreatedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestApprovalSchemaCreatedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestApprovalSchemaCreatedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestApprovalSchemaCreatedBy(val *ChangeRequestApprovalSchemaCreatedBy) *NullableChangeRequestApprovalSchemaCreatedBy {
	return &NullableChangeRequestApprovalSchemaCreatedBy{value: val, isSet: true}
}

func (v NullableChangeRequestApprovalSchemaCreatedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestApprovalSchemaCreatedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}