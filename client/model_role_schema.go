/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the RoleSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleSchema{}

// RoleSchema struct for RoleSchema
type RoleSchema struct {
	Id          float32 `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewRoleSchema instantiates a new RoleSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleSchema(id float32, type_ string, name string) *RoleSchema {
	this := RoleSchema{}
	this.Id = id
	this.Type = type_
	this.Name = name
	return &this
}

// NewRoleSchemaWithDefaults instantiates a new RoleSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleSchemaWithDefaults() *RoleSchema {
	this := RoleSchema{}
	return &this
}

// GetId returns the Id field value
func (o *RoleSchema) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleSchema) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleSchema) SetId(v float32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RoleSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleSchema) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *RoleSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleSchema) SetDescription(v string) {
	o.Description = &v
}

func (o RoleSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableRoleSchema struct {
	value *RoleSchema
	isSet bool
}

func (v NullableRoleSchema) Get() *RoleSchema {
	return v.value
}

func (v *NullableRoleSchema) Set(val *RoleSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleSchema(val *RoleSchema) *NullableRoleSchema {
	return &NullableRoleSchema{value: val, isSet: true}
}

func (v NullableRoleSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
