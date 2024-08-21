/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.1.10+main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSchema{}

// ProjectSchema A definition of the project used for projects listing purposes
type ProjectSchema struct {
	// The id of this project
	Id string `json:"id"`
	// The name of this project
	Name string `json:"name"`
	// Additional information about the project
	Description NullableString `json:"description,omitempty"`
	// An indicator of the [project's health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100
	Health *float32 `json:"health,omitempty"`
	// The number of features this project has
	FeatureCount         *float32 `json:"featureCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProjectSchema ProjectSchema

// NewProjectSchema instantiates a new ProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSchema(id string, name string) *ProjectSchema {
	this := ProjectSchema{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProjectSchemaWithDefaults instantiates a new ProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSchemaWithDefaults() *ProjectSchema {
	this := ProjectSchema{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectSchema) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ProjectSchema) GetHealth() float32 {
	if o == nil || IsNil(o.Health) {
		var ret float32
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ProjectSchema) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given float32 and assigns it to the Health field.
func (o *ProjectSchema) SetHealth(v float32) {
	o.Health = &v
}

// GetFeatureCount returns the FeatureCount field value if set, zero value otherwise.
func (o *ProjectSchema) GetFeatureCount() float32 {
	if o == nil || IsNil(o.FeatureCount) {
		var ret float32
		return ret
	}
	return *o.FeatureCount
}

// GetFeatureCountOk returns a tuple with the FeatureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetFeatureCountOk() (*float32, bool) {
	if o == nil || IsNil(o.FeatureCount) {
		return nil, false
	}
	return o.FeatureCount, true
}

// HasFeatureCount returns a boolean if a field has been set.
func (o *ProjectSchema) HasFeatureCount() bool {
	if o != nil && !IsNil(o.FeatureCount) {
		return true
	}

	return false
}

// SetFeatureCount gets a reference to the given float32 and assigns it to the FeatureCount field.
func (o *ProjectSchema) SetFeatureCount(v float32) {
	o.FeatureCount = &v
}

func (o ProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.FeatureCount) {
		toSerialize["featureCount"] = o.FeatureCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varProjectSchema := _ProjectSchema{}

	err = json.Unmarshal(data, &varProjectSchema)

	if err != nil {
		return err
	}

	*o = ProjectSchema(varProjectSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "health")
		delete(additionalProperties, "featureCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectSchema struct {
	value *ProjectSchema
	isSet bool
}

func (v NullableProjectSchema) Get() *ProjectSchema {
	return v.value
}

func (v *NullableProjectSchema) Set(val *ProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSchema(val *ProjectSchema) *NullableProjectSchema {
	return &NullableProjectSchema{value: val, isSet: true}
}

func (v NullableProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
