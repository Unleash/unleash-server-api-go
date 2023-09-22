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

// checks if the CreateRoleWithPermissionsSchemaOneOfPermissionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoleWithPermissionsSchemaOneOfPermissionsInner{}

// CreateRoleWithPermissionsSchemaOneOfPermissionsInner struct for CreateRoleWithPermissionsSchemaOneOfPermissionsInner
type CreateRoleWithPermissionsSchemaOneOfPermissionsInner struct {
	// The name of the permission
	Name string `json:"name"`
	// The environments of the permission if the permission is environment specific
	Environment *string `json:"environment,omitempty"`
}

// NewCreateRoleWithPermissionsSchemaOneOfPermissionsInner instantiates a new CreateRoleWithPermissionsSchemaOneOfPermissionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoleWithPermissionsSchemaOneOfPermissionsInner(name string) *CreateRoleWithPermissionsSchemaOneOfPermissionsInner {
	this := CreateRoleWithPermissionsSchemaOneOfPermissionsInner{}
	this.Name = name
	return &this
}

// NewCreateRoleWithPermissionsSchemaOneOfPermissionsInnerWithDefaults instantiates a new CreateRoleWithPermissionsSchemaOneOfPermissionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithPermissionsSchemaOneOfPermissionsInnerWithDefaults() *CreateRoleWithPermissionsSchemaOneOfPermissionsInner {
	this := CreateRoleWithPermissionsSchemaOneOfPermissionsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) SetName(v string) {
	o.Name = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) SetEnvironment(v string) {
	o.Environment = &v
}

func (o CreateRoleWithPermissionsSchemaOneOfPermissionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoleWithPermissionsSchemaOneOfPermissionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner struct {
	value *CreateRoleWithPermissionsSchemaOneOfPermissionsInner
	isSet bool
}

func (v NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) Get() *CreateRoleWithPermissionsSchemaOneOfPermissionsInner {
	return v.value
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) Set(val *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner(val *CreateRoleWithPermissionsSchemaOneOfPermissionsInner) *NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner {
	return &NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner{value: val, isSet: true}
}

func (v NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoleWithPermissionsSchemaOneOfPermissionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
