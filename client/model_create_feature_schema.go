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

// checks if the CreateFeatureSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeatureSchema{}

// CreateFeatureSchema struct for CreateFeatureSchema
type CreateFeatureSchema struct {
	Name           string  `json:"name"`
	Type           *string `json:"type,omitempty"`
	Description    *string `json:"description,omitempty"`
	ImpressionData *bool   `json:"impressionData,omitempty"`
}

// NewCreateFeatureSchema instantiates a new CreateFeatureSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeatureSchema(name string) *CreateFeatureSchema {
	this := CreateFeatureSchema{}
	this.Name = name
	return &this
}

// NewCreateFeatureSchemaWithDefaults instantiates a new CreateFeatureSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeatureSchemaWithDefaults() *CreateFeatureSchema {
	this := CreateFeatureSchema{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFeatureSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFeatureSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateFeatureSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateFeatureSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateFeatureSchema) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFeatureSchema) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFeatureSchema) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFeatureSchema) SetDescription(v string) {
	o.Description = &v
}

// GetImpressionData returns the ImpressionData field value if set, zero value otherwise.
func (o *CreateFeatureSchema) GetImpressionData() bool {
	if o == nil || IsNil(o.ImpressionData) {
		var ret bool
		return ret
	}
	return *o.ImpressionData
}

// GetImpressionDataOk returns a tuple with the ImpressionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureSchema) GetImpressionDataOk() (*bool, bool) {
	if o == nil || IsNil(o.ImpressionData) {
		return nil, false
	}
	return o.ImpressionData, true
}

// HasImpressionData returns a boolean if a field has been set.
func (o *CreateFeatureSchema) HasImpressionData() bool {
	if o != nil && !IsNil(o.ImpressionData) {
		return true
	}

	return false
}

// SetImpressionData gets a reference to the given bool and assigns it to the ImpressionData field.
func (o *CreateFeatureSchema) SetImpressionData(v bool) {
	o.ImpressionData = &v
}

func (o CreateFeatureSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFeatureSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImpressionData) {
		toSerialize["impressionData"] = o.ImpressionData
	}
	return toSerialize, nil
}

type NullableCreateFeatureSchema struct {
	value *CreateFeatureSchema
	isSet bool
}

func (v NullableCreateFeatureSchema) Get() *CreateFeatureSchema {
	return v.value
}

func (v *NullableCreateFeatureSchema) Set(val *CreateFeatureSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeatureSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeatureSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeatureSchema(val *CreateFeatureSchema) *NullableCreateFeatureSchema {
	return &NullableCreateFeatureSchema{value: val, isSet: true}
}

func (v NullableCreateFeatureSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFeatureSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
