/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.10+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateRole400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRole400Response{}

// CreateRole400Response struct for CreateRole400Response
type CreateRole400Response struct {
	// The ID of the error instance
	Id *string `json:"id,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
	// A description of what went wrong.
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRole400Response CreateRole400Response

// NewCreateRole400Response instantiates a new CreateRole400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRole400Response() *CreateRole400Response {
	this := CreateRole400Response{}
	return &this
}

// NewCreateRole400ResponseWithDefaults instantiates a new CreateRole400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRole400ResponseWithDefaults() *CreateRole400Response {
	this := CreateRole400Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRole400Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole400Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRole400Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateRole400Response) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRole400Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole400Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRole400Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRole400Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRole400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRole400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRole400Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRole400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRole400Response) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRole400Response) UnmarshalJSON(data []byte) (err error) {
	varCreateRole400Response := _CreateRole400Response{}

	err = json.Unmarshal(data, &varCreateRole400Response)

	if err != nil {
		return err
	}

	*o = CreateRole400Response(varCreateRole400Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRole400Response struct {
	value *CreateRole400Response
	isSet bool
}

func (v NullableCreateRole400Response) Get() *CreateRole400Response {
	return v.value
}

func (v *NullableCreateRole400Response) Set(val *CreateRole400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRole400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRole400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRole400Response(val *CreateRole400Response) *NullableCreateRole400Response {
	return &NullableCreateRole400Response{value: val, isSet: true}
}

func (v NullableCreateRole400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRole400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
