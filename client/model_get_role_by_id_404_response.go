/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetRoleById404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleById404Response{}

// GetRoleById404Response struct for GetRoleById404Response
type GetRoleById404Response struct {
	// The ID of the error instance
	Id *string `json:"id,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
	// A description of what went wrong.
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetRoleById404Response GetRoleById404Response

// NewGetRoleById404Response instantiates a new GetRoleById404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleById404Response() *GetRoleById404Response {
	this := GetRoleById404Response{}
	return &this
}

// NewGetRoleById404ResponseWithDefaults instantiates a new GetRoleById404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleById404ResponseWithDefaults() *GetRoleById404Response {
	this := GetRoleById404Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRoleById404Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleById404Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRoleById404Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetRoleById404Response) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetRoleById404Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleById404Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetRoleById404Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetRoleById404Response) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetRoleById404Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleById404Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetRoleById404Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetRoleById404Response) SetMessage(v string) {
	o.Message = &v
}

func (o GetRoleById404Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleById404Response) ToMap() (map[string]interface{}, error) {
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

func (o *GetRoleById404Response) UnmarshalJSON(data []byte) (err error) {
	varGetRoleById404Response := _GetRoleById404Response{}

	err = json.Unmarshal(data, &varGetRoleById404Response)

	if err != nil {
		return err
	}

	*o = GetRoleById404Response(varGetRoleById404Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetRoleById404Response struct {
	value *GetRoleById404Response
	isSet bool
}

func (v NullableGetRoleById404Response) Get() *GetRoleById404Response {
	return v.value
}

func (v *NullableGetRoleById404Response) Set(val *GetRoleById404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleById404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleById404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleById404Response(val *GetRoleById404Response) *NullableGetRoleById404Response {
	return &NullableGetRoleById404Response{value: val, isSet: true}
}

func (v NullableGetRoleById404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleById404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
