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

// checks if the CreateEnvironmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEnvironmentSchema{}

// CreateEnvironmentSchema Data required to create a new [environment](https://docs.getunleash.io/reference/environments)
type CreateEnvironmentSchema struct {
	// The name of the environment. Must be a URL-friendly string according to [RFC 3968, section 2.3](https://www.rfc-editor.org/rfc/rfc3986#section-2.3)
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9~_.-]+$"`
	// The [type of environment](https://docs.getunleash.io/reference/environments#environment-types) you would like to create. Unleash officially recognizes the following values: - `development` - `test` - `preproduction` - `production`  If you pass a string that is not one of the recognized values, Unleash will accept it, but it will carry no special semantics.
	Type string `json:"type"`
	// Newly created environments are enabled by default. Set this property to `false` to create the environment in a disabled state.
	Enabled *bool `json:"enabled,omitempty"`
	// Defines where in the list of environments to place this environment. The list uses an ascending sort, so lower numbers are shown first. You can change this value later.
	SortOrder            *int32 `json:"sortOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateEnvironmentSchema CreateEnvironmentSchema

// NewCreateEnvironmentSchema instantiates a new CreateEnvironmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentSchema(name string, type_ string) *CreateEnvironmentSchema {
	this := CreateEnvironmentSchema{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateEnvironmentSchemaWithDefaults instantiates a new CreateEnvironmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentSchemaWithDefaults() *CreateEnvironmentSchema {
	this := CreateEnvironmentSchema{}
	return &this
}

// GetName returns the Name field value
func (o *CreateEnvironmentSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateEnvironmentSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CreateEnvironmentSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateEnvironmentSchema) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateEnvironmentSchema) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentSchema) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateEnvironmentSchema) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateEnvironmentSchema) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *CreateEnvironmentSchema) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentSchema) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *CreateEnvironmentSchema) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *CreateEnvironmentSchema) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o CreateEnvironmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEnvironmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateEnvironmentSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varCreateEnvironmentSchema := _CreateEnvironmentSchema{}

	err = json.Unmarshal(data, &varCreateEnvironmentSchema)

	if err != nil {
		return err
	}

	*o = CreateEnvironmentSchema(varCreateEnvironmentSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "sortOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateEnvironmentSchema struct {
	value *CreateEnvironmentSchema
	isSet bool
}

func (v NullableCreateEnvironmentSchema) Get() *CreateEnvironmentSchema {
	return v.value
}

func (v *NullableCreateEnvironmentSchema) Set(val *CreateEnvironmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentSchema(val *CreateEnvironmentSchema) *NullableCreateEnvironmentSchema {
	return &NullableCreateEnvironmentSchema{value: val, isSet: true}
}

func (v NullableCreateEnvironmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
