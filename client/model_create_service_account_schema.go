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

// checks if the CreateServiceAccountSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateServiceAccountSchema{}

// CreateServiceAccountSchema Describes the properties required to create a new service account
type CreateServiceAccountSchema struct {
	// The username of the service account
	Username *string `json:"username,omitempty"`
	// The name of the service account
	Name *string `json:"name,omitempty"`
	// The id of the root role for the service account
	RootRole int32 `json:"rootRole"`
}

type _CreateServiceAccountSchema CreateServiceAccountSchema

// NewCreateServiceAccountSchema instantiates a new CreateServiceAccountSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceAccountSchema(rootRole int32) *CreateServiceAccountSchema {
	this := CreateServiceAccountSchema{}
	this.RootRole = rootRole
	return &this
}

// NewCreateServiceAccountSchemaWithDefaults instantiates a new CreateServiceAccountSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceAccountSchemaWithDefaults() *CreateServiceAccountSchema {
	this := CreateServiceAccountSchema{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateServiceAccountSchema) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountSchema) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateServiceAccountSchema) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateServiceAccountSchema) SetUsername(v string) {
	o.Username = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateServiceAccountSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateServiceAccountSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateServiceAccountSchema) SetName(v string) {
	o.Name = &v
}

// GetRootRole returns the RootRole field value
func (o *CreateServiceAccountSchema) GetRootRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RootRole
}

// GetRootRoleOk returns a tuple with the RootRole field value
// and a boolean to check if the value has been set.
func (o *CreateServiceAccountSchema) GetRootRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootRole, true
}

// SetRootRole sets field value
func (o *CreateServiceAccountSchema) SetRootRole(v int32) {
	o.RootRole = v
}

func (o CreateServiceAccountSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceAccountSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["rootRole"] = o.RootRole
	return toSerialize, nil
}

func (o *CreateServiceAccountSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rootRole",
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

	varCreateServiceAccountSchema := _CreateServiceAccountSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateServiceAccountSchema)

	if err != nil {
		return err
	}

	*o = CreateServiceAccountSchema(varCreateServiceAccountSchema)

	return err
}

type NullableCreateServiceAccountSchema struct {
	value *CreateServiceAccountSchema
	isSet bool
}

func (v NullableCreateServiceAccountSchema) Get() *CreateServiceAccountSchema {
	return v.value
}

func (v *NullableCreateServiceAccountSchema) Set(val *CreateServiceAccountSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceAccountSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceAccountSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceAccountSchema(val *CreateServiceAccountSchema) *NullableCreateServiceAccountSchema {
	return &NullableCreateServiceAccountSchema{value: val, isSet: true}
}

func (v NullableCreateServiceAccountSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceAccountSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
