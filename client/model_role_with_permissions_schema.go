/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.10+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RoleWithPermissionsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleWithPermissionsSchema{}

// RoleWithPermissionsSchema A read model for the role and permissions to allow Unleash to decide what actions a role holder is allowed to perform
type RoleWithPermissionsSchema struct {
	// The role id
	Id float32 `json:"id"`
	// A role can either be a global `root` role, or a `project` role or a `custom` project role or a custom global `root-custom` role
	Type string `json:"type"`
	// The name of the role
	Name string `json:"name"`
	// A more detailed description of the role and what use it's intended for
	Description *string `json:"description,omitempty"`
	// A list of permissions assigned to this role
	Permissions []AdminPermissionSchema `json:"permissions"`
}

type _RoleWithPermissionsSchema RoleWithPermissionsSchema

// NewRoleWithPermissionsSchema instantiates a new RoleWithPermissionsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleWithPermissionsSchema(id float32, type_ string, name string, permissions []AdminPermissionSchema) *RoleWithPermissionsSchema {
	this := RoleWithPermissionsSchema{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Permissions = permissions
	return &this
}

// NewRoleWithPermissionsSchemaWithDefaults instantiates a new RoleWithPermissionsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithPermissionsSchemaWithDefaults() *RoleWithPermissionsSchema {
	this := RoleWithPermissionsSchema{}
	return &this
}

// GetId returns the Id field value
func (o *RoleWithPermissionsSchema) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleWithPermissionsSchema) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleWithPermissionsSchema) SetId(v float32) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RoleWithPermissionsSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleWithPermissionsSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleWithPermissionsSchema) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *RoleWithPermissionsSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleWithPermissionsSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleWithPermissionsSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleWithPermissionsSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleWithPermissionsSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleWithPermissionsSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleWithPermissionsSchema) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value
func (o *RoleWithPermissionsSchema) GetPermissions() []AdminPermissionSchema {
	if o == nil {
		var ret []AdminPermissionSchema
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *RoleWithPermissionsSchema) GetPermissionsOk() ([]AdminPermissionSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *RoleWithPermissionsSchema) SetPermissions(v []AdminPermissionSchema) {
	o.Permissions = v
}

func (o RoleWithPermissionsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleWithPermissionsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

func (o *RoleWithPermissionsSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"name",
		"permissions",
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

	varRoleWithPermissionsSchema := _RoleWithPermissionsSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoleWithPermissionsSchema)

	if err != nil {
		return err
	}

	*o = RoleWithPermissionsSchema(varRoleWithPermissionsSchema)

	return err
}

type NullableRoleWithPermissionsSchema struct {
	value *RoleWithPermissionsSchema
	isSet bool
}

func (v NullableRoleWithPermissionsSchema) Get() *RoleWithPermissionsSchema {
	return v.value
}

func (v *NullableRoleWithPermissionsSchema) Set(val *RoleWithPermissionsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleWithPermissionsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleWithPermissionsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleWithPermissionsSchema(val *RoleWithPermissionsSchema) *NullableRoleWithPermissionsSchema {
	return &NullableRoleWithPermissionsSchema{value: val, isSet: true}
}

func (v NullableRoleWithPermissionsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleWithPermissionsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
