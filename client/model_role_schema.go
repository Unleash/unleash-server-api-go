/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the RoleSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleSchema{}

// RoleSchema A role holds permissions to allow Unleash to decide what actions a role holder is allowed to perform
type RoleSchema struct {
	// The role id
	Id int32 `json:"id"`
	// A role can either be a global root role (applies to all projects) or a project role
	Type string `json:"type"`
	// The name of the role
	Name string `json:"name"`
	// A more detailed description of the role and what use it's intended for
	Description *string `json:"description,omitempty"`
	// What project the role belongs to
	Project              NullableString `json:"project,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleSchema RoleSchema

// NewRoleSchema instantiates a new RoleSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleSchema(id int32, type_ string, name string) *RoleSchema {
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
func (o *RoleSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoleSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoleSchema) SetId(v int32) {
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

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleSchema) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleSchema) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *RoleSchema) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *RoleSchema) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *RoleSchema) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *RoleSchema) UnsetProject() {
	o.Project.Unset()
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
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoleSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varRoleSchema := _RoleSchema{}

	err = json.Unmarshal(data, &varRoleSchema)

	if err != nil {
		return err
	}

	*o = RoleSchema(varRoleSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "project")
		o.AdditionalProperties = additionalProperties
	}

	return err
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
