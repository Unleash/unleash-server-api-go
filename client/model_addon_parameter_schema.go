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

// checks if the AddonParameterSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonParameterSchema{}

// AddonParameterSchema An addon parameter definition.
type AddonParameterSchema struct {
	// The name of the parameter as it is used in code. References to this parameter should use this value.
	Name string `json:"name"`
	// The name of the parameter as it is shown to the end user in the Admin UI.
	DisplayName string `json:"displayName"`
	// The type of the parameter. Corresponds roughly to [HTML `input` field types](https://developer.mozilla.org/docs/Web/HTML/Element/Input#input_types). Multi-line inut fields are indicated as `textfield` (equivalent to the HTML `textarea` tag).
	Type string `json:"type"`
	// A description of the parameter. This should explain to the end user what the parameter is used for.
	Description *string `json:"description,omitempty"`
	// The default value for this parameter. This value is used if no other value is provided.
	Placeholder *string `json:"placeholder,omitempty"`
	// Whether this parameter is required or not. If a parameter is required, you must give it a value when you create the addon. If it is not required it can be left out. It may receive a default value in those cases.
	Required bool `json:"required"`
	// Indicates whether this parameter is **sensitive** or not. Unleash will not return sensitive parameters to API requests. It will instead use a number of asterisks to indicate that a value is set, e.g. \"******\". The number of asterisks does not correlate to the parameter's value.
	Sensitive bool `json:"sensitive"`
}

// NewAddonParameterSchema instantiates a new AddonParameterSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonParameterSchema(name string, displayName string, type_ string, required bool, sensitive bool) *AddonParameterSchema {
	this := AddonParameterSchema{}
	this.Name = name
	this.DisplayName = displayName
	this.Type = type_
	this.Required = required
	this.Sensitive = sensitive
	return &this
}

// NewAddonParameterSchemaWithDefaults instantiates a new AddonParameterSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonParameterSchemaWithDefaults() *AddonParameterSchema {
	this := AddonParameterSchema{}
	return &this
}

// GetName returns the Name field value
func (o *AddonParameterSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddonParameterSchema) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *AddonParameterSchema) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AddonParameterSchema) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetType returns the Type field value
func (o *AddonParameterSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddonParameterSchema) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddonParameterSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddonParameterSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddonParameterSchema) SetDescription(v string) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *AddonParameterSchema) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *AddonParameterSchema) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *AddonParameterSchema) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetRequired returns the Required field value
func (o *AddonParameterSchema) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *AddonParameterSchema) SetRequired(v bool) {
	o.Required = v
}

// GetSensitive returns the Sensitive field value
func (o *AddonParameterSchema) GetSensitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value
// and a boolean to check if the value has been set.
func (o *AddonParameterSchema) GetSensitiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sensitive, true
}

// SetSensitive sets field value
func (o *AddonParameterSchema) SetSensitive(v bool) {
	o.Sensitive = v
}

func (o AddonParameterSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonParameterSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	toSerialize["required"] = o.Required
	toSerialize["sensitive"] = o.Sensitive
	return toSerialize, nil
}

type NullableAddonParameterSchema struct {
	value *AddonParameterSchema
	isSet bool
}

func (v NullableAddonParameterSchema) Get() *AddonParameterSchema {
	return v.value
}

func (v *NullableAddonParameterSchema) Set(val *AddonParameterSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonParameterSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonParameterSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonParameterSchema(val *AddonParameterSchema) *NullableAddonParameterSchema {
	return &NullableAddonParameterSchema{value: val, isSet: true}
}

func (v NullableAddonParameterSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonParameterSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
