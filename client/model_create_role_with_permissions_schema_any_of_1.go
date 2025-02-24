/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.7.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateRoleWithPermissionsSchemaAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleWithPermissionsSchemaAnyOf1{}

// CreateRoleWithPermissionsSchemaAnyOf1 struct for CreateRoleWithPermissionsSchemaAnyOf1
type CreateRoleWithPermissionsSchemaAnyOf1 struct {
	// The name of the custom role
	Name string `json:"name"`
	// A more detailed description of the custom role and what use it's intended for
	Description *string `json:"description,omitempty"`
	// [Custom project roles](https://docs.getunleash.io/reference/rbac#custom-project-roles) contain a specific set of permissions for project resources.
	Type *string `json:"type,omitempty"`
	// A list of permissions assigned to this role
	Permissions          []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRoleWithPermissionsSchemaAnyOf1 CreateRoleWithPermissionsSchemaAnyOf1

// NewCreateRoleWithPermissionsSchemaAnyOf1 instantiates a new CreateRoleWithPermissionsSchemaAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleWithPermissionsSchemaAnyOf1(name string) *CreateRoleWithPermissionsSchemaAnyOf1 {
	this := CreateRoleWithPermissionsSchemaAnyOf1{}
	this.Name = name
	return &this
}

// NewCreateRoleWithPermissionsSchemaAnyOf1WithDefaults instantiates a new CreateRoleWithPermissionsSchemaAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithPermissionsSchemaAnyOf1WithDefaults() *CreateRoleWithPermissionsSchemaAnyOf1 {
	this := CreateRoleWithPermissionsSchemaAnyOf1{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetType(v string) {
	o.Type = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetPermissions() []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) GetPermissionsOk() ([]CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner and assigns it to the Permissions field.
func (o *CreateRoleWithPermissionsSchemaAnyOf1) SetPermissions(v []CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) {
	o.Permissions = v
}

func (o CreateRoleWithPermissionsSchemaAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleWithPermissionsSchemaAnyOf1) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRoleWithPermissionsSchemaAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRoleWithPermissionsSchemaAnyOf1 := _CreateRoleWithPermissionsSchemaAnyOf1{}

	err = json.Unmarshal(data, &varCreateRoleWithPermissionsSchemaAnyOf1)

	if err != nil {
		return err
	}

	*o = CreateRoleWithPermissionsSchemaAnyOf1(varCreateRoleWithPermissionsSchemaAnyOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRoleWithPermissionsSchemaAnyOf1 struct {
	value *CreateRoleWithPermissionsSchemaAnyOf1
	isSet bool
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1) Get() *CreateRoleWithPermissionsSchemaAnyOf1 {
	return v.value
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1) Set(val *CreateRoleWithPermissionsSchemaAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleWithPermissionsSchemaAnyOf1(val *CreateRoleWithPermissionsSchemaAnyOf1) *NullableCreateRoleWithPermissionsSchemaAnyOf1 {
	return &NullableCreateRoleWithPermissionsSchemaAnyOf1{value: val, isSet: true}
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
