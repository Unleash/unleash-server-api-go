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

// checks if the CreateRoleWithPermissionsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleWithPermissionsSchema{}

// CreateRoleWithPermissionsSchema struct for CreateRoleWithPermissionsSchema
type CreateRoleWithPermissionsSchema struct {
	Name        string                                            `json:"name"`
	Description *string                                           `json:"description,omitempty"`
	Type        *string                                           `json:"type,omitempty"`
	Permissions []CreateRoleWithPermissionsSchemaPermissionsInner `json:"permissions,omitempty"`
}

// NewCreateRoleWithPermissionsSchema instantiates a new CreateRoleWithPermissionsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleWithPermissionsSchema(name string) *CreateRoleWithPermissionsSchema {
	this := CreateRoleWithPermissionsSchema{}
	this.Name = name
	return &this
}

// NewCreateRoleWithPermissionsSchemaWithDefaults instantiates a new CreateRoleWithPermissionsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithPermissionsSchemaWithDefaults() *CreateRoleWithPermissionsSchema {
	this := CreateRoleWithPermissionsSchema{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRoleWithPermissionsSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRoleWithPermissionsSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRoleWithPermissionsSchema) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRoleWithPermissionsSchema) SetType(v string) {
	o.Type = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchema) GetPermissions() []CreateRoleWithPermissionsSchemaPermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []CreateRoleWithPermissionsSchemaPermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchema) GetPermissionsOk() ([]CreateRoleWithPermissionsSchemaPermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchema) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CreateRoleWithPermissionsSchemaPermissionsInner and assigns it to the Permissions field.
func (o *CreateRoleWithPermissionsSchema) SetPermissions(v []CreateRoleWithPermissionsSchemaPermissionsInner) {
	o.Permissions = v
}

func (o CreateRoleWithPermissionsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleWithPermissionsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableCreateRoleWithPermissionsSchema struct {
	value *CreateRoleWithPermissionsSchema
	isSet bool
}

func (v NullableCreateRoleWithPermissionsSchema) Get() *CreateRoleWithPermissionsSchema {
	return v.value
}

func (v *NullableCreateRoleWithPermissionsSchema) Set(val *CreateRoleWithPermissionsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleWithPermissionsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleWithPermissionsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleWithPermissionsSchema(val *CreateRoleWithPermissionsSchema) *NullableCreateRoleWithPermissionsSchema {
	return &NullableCreateRoleWithPermissionsSchema{value: val, isSet: true}
}

func (v NullableCreateRoleWithPermissionsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleWithPermissionsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
