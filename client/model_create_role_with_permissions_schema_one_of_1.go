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

// checks if the CreateRoleWithPermissionsSchemaOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleWithPermissionsSchemaOneOf1{}

// CreateRoleWithPermissionsSchemaOneOf1 struct for CreateRoleWithPermissionsSchemaOneOf1
type CreateRoleWithPermissionsSchemaOneOf1 struct {
	// The name of the custom role
	Name string `json:"name"`
	// A more detailed description of the custom role and what use it's intended for
	Description *string `json:"description,omitempty"`
	// [Custom project roles](https://docs.getunleash.io/reference/rbac#custom-project-roles) contain a specific set of permissions for project resources.
	Type *string `json:"type,omitempty"`
	// A list of permissions assigned to this role
	Permissions []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner `json:"permissions,omitempty"`
}

// NewCreateRoleWithPermissionsSchemaOneOf1 instantiates a new CreateRoleWithPermissionsSchemaOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleWithPermissionsSchemaOneOf1(name string) *CreateRoleWithPermissionsSchemaOneOf1 {
	this := CreateRoleWithPermissionsSchemaOneOf1{}
	this.Name = name
	return &this
}

// NewCreateRoleWithPermissionsSchemaOneOf1WithDefaults instantiates a new CreateRoleWithPermissionsSchemaOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithPermissionsSchemaOneOf1WithDefaults() *CreateRoleWithPermissionsSchemaOneOf1 {
	this := CreateRoleWithPermissionsSchemaOneOf1{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRoleWithPermissionsSchemaOneOf1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRoleWithPermissionsSchemaOneOf1) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRoleWithPermissionsSchemaOneOf1) SetType(v string) {
	o.Type = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetPermissions() []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) GetPermissionsOk() ([]CreateRoleWithPermissionsSchemaOneOf1PermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaOneOf1) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner and assigns it to the Permissions field.
func (o *CreateRoleWithPermissionsSchemaOneOf1) SetPermissions(v []CreateRoleWithPermissionsSchemaOneOf1PermissionsInner) {
	o.Permissions = v
}

func (o CreateRoleWithPermissionsSchemaOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleWithPermissionsSchemaOneOf1) ToMap() (map[string]interface{}, error) {
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

type NullableCreateRoleWithPermissionsSchemaOneOf1 struct {
	value *CreateRoleWithPermissionsSchemaOneOf1
	isSet bool
}

func (v NullableCreateRoleWithPermissionsSchemaOneOf1) Get() *CreateRoleWithPermissionsSchemaOneOf1 {
	return v.value
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOf1) Set(val *CreateRoleWithPermissionsSchemaOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleWithPermissionsSchemaOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleWithPermissionsSchemaOneOf1(val *CreateRoleWithPermissionsSchemaOneOf1) *NullableCreateRoleWithPermissionsSchemaOneOf1 {
	return &NullableCreateRoleWithPermissionsSchemaOneOf1{value: val, isSet: true}
}

func (v NullableCreateRoleWithPermissionsSchemaOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}