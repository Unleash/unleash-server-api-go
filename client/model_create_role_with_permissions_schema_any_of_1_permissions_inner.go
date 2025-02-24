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

// checks if the CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{}

// CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner struct for CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner
type CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner struct {
	// The id of the permission
	Id float32 `json:"id"`
	// The name of the permission
	Name *string `json:"name,omitempty"`
	// The environments of the permission if the permission is environment specific
	Environment          NullableString `json:"environment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner

// NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner instantiates a new CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner(id float32) *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner {
	this := CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{}
	this.Id = id
	return &this
}

// NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInnerWithDefaults instantiates a new CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithPermissionsSchemaAnyOf1PermissionsInnerWithDefaults() *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner {
	this := CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{}
	return &this
}

// GetId returns the Id field value
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) SetName(v string) {
	o.Name = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) UnsetEnvironment() {
	o.Environment.Unset()
}

func (o CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner := _CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{}

	err = json.Unmarshal(data, &varCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner)

	if err != nil {
		return err
	}

	*o = CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner(varCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner struct {
	value *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner
	isSet bool
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) Get() *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner {
	return v.value
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) Set(val *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner(val *CreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) *NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner {
	return &NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner{value: val, isSet: true}
}

func (v NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleWithPermissionsSchemaAnyOf1PermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
