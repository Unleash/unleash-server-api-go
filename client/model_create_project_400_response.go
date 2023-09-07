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

// checks if the CreateProject400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProject400Response{}

// CreateProject400Response struct for CreateProject400Response
type CreateProject400Response struct {
	// The ID of the error instance
	Id *string `json:"id,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
}

// NewCreateProject400Response instantiates a new CreateProject400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProject400Response() *CreateProject400Response {
	this := CreateProject400Response{}
	return &this
}

// NewCreateProject400ResponseWithDefaults instantiates a new CreateProject400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProject400ResponseWithDefaults() *CreateProject400Response {
	this := CreateProject400Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateProject400Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProject400Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateProject400Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateProject400Response) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateProject400Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProject400Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateProject400Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateProject400Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateProject400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProject400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateProject400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateProject400Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateProject400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProject400Response) ToMap() (map[string]interface{}, error) {
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

type NullableCreateProject400Response struct {
	value *CreateProject400Response
	isSet bool
}

func (v NullableCreateProject400Response) Get() *CreateProject400Response {
	return v.value
}

func (v *NullableCreateProject400Response) Set(val *CreateProject400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProject400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProject400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProject400Response(val *CreateProject400Response) *NullableCreateProject400Response {
	return &NullableCreateProject400Response{value: val, isSet: true}
}

func (v NullableCreateProject400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProject400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}