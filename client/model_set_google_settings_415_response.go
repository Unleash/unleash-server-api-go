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

// checks if the SetGoogleSettings415Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetGoogleSettings415Response{}

// SetGoogleSettings415Response struct for SetGoogleSettings415Response
type SetGoogleSettings415Response struct {
	// The ID of the error instance
	Id *string `json:"id,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
}

// NewSetGoogleSettings415Response instantiates a new SetGoogleSettings415Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetGoogleSettings415Response() *SetGoogleSettings415Response {
	this := SetGoogleSettings415Response{}
	return &this
}

// NewSetGoogleSettings415ResponseWithDefaults instantiates a new SetGoogleSettings415Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetGoogleSettings415ResponseWithDefaults() *SetGoogleSettings415Response {
	this := SetGoogleSettings415Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SetGoogleSettings415Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetGoogleSettings415Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SetGoogleSettings415Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SetGoogleSettings415Response) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetGoogleSettings415Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetGoogleSettings415Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetGoogleSettings415Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetGoogleSettings415Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SetGoogleSettings415Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetGoogleSettings415Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SetGoogleSettings415Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SetGoogleSettings415Response) SetMessage(v string) {
	o.Message = &v
}

func (o SetGoogleSettings415Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetGoogleSettings415Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSetGoogleSettings415Response struct {
	value *SetGoogleSettings415Response
	isSet bool
}

func (v NullableSetGoogleSettings415Response) Get() *SetGoogleSettings415Response {
	return v.value
}

func (v *NullableSetGoogleSettings415Response) Set(val *SetGoogleSettings415Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSetGoogleSettings415Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSetGoogleSettings415Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetGoogleSettings415Response(val *SetGoogleSettings415Response) *NullableSetGoogleSettings415Response {
	return &NullableSetGoogleSettings415Response{value: val, isSet: true}
}

func (v NullableSetGoogleSettings415Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetGoogleSettings415Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
