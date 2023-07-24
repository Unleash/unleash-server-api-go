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

// checks if the ChangeRequestStateSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestStateSchema{}

// ChangeRequestStateSchema struct for ChangeRequestStateSchema
type ChangeRequestStateSchema struct {
	State   string  `json:"state"`
	Comment *string `json:"comment,omitempty"`
}

// NewChangeRequestStateSchema instantiates a new ChangeRequestStateSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestStateSchema(state string) *ChangeRequestStateSchema {
	this := ChangeRequestStateSchema{}
	this.State = state
	return &this
}

// NewChangeRequestStateSchemaWithDefaults instantiates a new ChangeRequestStateSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestStateSchemaWithDefaults() *ChangeRequestStateSchema {
	this := ChangeRequestStateSchema{}
	return &this
}

// GetState returns the State field value
func (o *ChangeRequestStateSchema) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestStateSchema) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ChangeRequestStateSchema) SetState(v string) {
	o.State = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ChangeRequestStateSchema) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestStateSchema) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ChangeRequestStateSchema) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ChangeRequestStateSchema) SetComment(v string) {
	o.Comment = &v
}

func (o ChangeRequestStateSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestStateSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableChangeRequestStateSchema struct {
	value *ChangeRequestStateSchema
	isSet bool
}

func (v NullableChangeRequestStateSchema) Get() *ChangeRequestStateSchema {
	return v.value
}

func (v *NullableChangeRequestStateSchema) Set(val *ChangeRequestStateSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestStateSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestStateSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestStateSchema(val *ChangeRequestStateSchema) *NullableChangeRequestStateSchema {
	return &NullableChangeRequestStateSchema{value: val, isSet: true}
}

func (v NullableChangeRequestStateSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestStateSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
