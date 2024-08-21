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

// checks if the CreateRole409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRole409Response{}

// CreateRole409Response struct for CreateRole409Response
type CreateRole409Response struct {
	// The ID of the error instance
	Id *string `json:"id,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
	// A description of what went wrong.
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRole409Response CreateRole409Response

// NewCreateRole409Response instantiates a new CreateRole409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRole409Response() *CreateRole409Response {
	this := CreateRole409Response{}
	return &this
}

// NewCreateRole409ResponseWithDefaults instantiates a new CreateRole409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRole409ResponseWithDefaults() *CreateRole409Response {
	this := CreateRole409Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRole409Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole409Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRole409Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateRole409Response) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRole409Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole409Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRole409Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRole409Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateRole409Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole409Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateRole409Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateRole409Response) SetMessage(v string) {
	o.Message = &v
}

func (o CreateRole409Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRole409Response) ToMap() (map[string]interface{}, error) {
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

func (o *CreateRole409Response) UnmarshalJSON(data []byte) (err error) {
	varCreateRole409Response := _CreateRole409Response{}

	err = json.Unmarshal(data, &varCreateRole409Response)

	if err != nil {
		return err
	}

	*o = CreateRole409Response(varCreateRole409Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRole409Response struct {
	value *CreateRole409Response
	isSet bool
}

func (v NullableCreateRole409Response) Get() *CreateRole409Response {
	return v.value
}

func (v *NullableCreateRole409Response) Set(val *CreateRole409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRole409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRole409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRole409Response(val *CreateRole409Response) *NullableCreateRole409Response {
	return &NullableCreateRole409Response{value: val, isSet: true}
}

func (v NullableCreateRole409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRole409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
