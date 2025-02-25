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

// checks if the CreateFeatureNamingPatternSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeatureNamingPatternSchema{}

// CreateFeatureNamingPatternSchema Create a feature naming pattern
type CreateFeatureNamingPatternSchema struct {
	// A JavaScript regular expression pattern, without the start and end delimiters. Optional flags are not allowed.
	Pattern NullableString `json:"pattern"`
	// An example of a feature name that matches the pattern. Must itself match the pattern supplied.
	Example NullableString `json:"example,omitempty"`
	// A description of the pattern in a human-readable format. Will be shown to users when they create a new feature flag.
	Description          NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFeatureNamingPatternSchema CreateFeatureNamingPatternSchema

// NewCreateFeatureNamingPatternSchema instantiates a new CreateFeatureNamingPatternSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeatureNamingPatternSchema(pattern NullableString) *CreateFeatureNamingPatternSchema {
	this := CreateFeatureNamingPatternSchema{}
	this.Pattern = pattern
	return &this
}

// NewCreateFeatureNamingPatternSchemaWithDefaults instantiates a new CreateFeatureNamingPatternSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeatureNamingPatternSchemaWithDefaults() *CreateFeatureNamingPatternSchema {
	this := CreateFeatureNamingPatternSchema{}
	return &this
}

// GetPattern returns the Pattern field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateFeatureNamingPatternSchema) GetPattern() string {
	if o == nil || o.Pattern.Get() == nil {
		var ret string
		return ret
	}

	return *o.Pattern.Get()
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureNamingPatternSchema) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pattern.Get(), o.Pattern.IsSet()
}

// SetPattern sets field value
func (o *CreateFeatureNamingPatternSchema) SetPattern(v string) {
	o.Pattern.Set(&v)
}

// GetExample returns the Example field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureNamingPatternSchema) GetExample() string {
	if o == nil || IsNil(o.Example.Get()) {
		var ret string
		return ret
	}
	return *o.Example.Get()
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureNamingPatternSchema) GetExampleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Example.Get(), o.Example.IsSet()
}

// HasExample returns a boolean if a field has been set.
func (o *CreateFeatureNamingPatternSchema) HasExample() bool {
	if o != nil && o.Example.IsSet() {
		return true
	}

	return false
}

// SetExample gets a reference to the given NullableString and assigns it to the Example field.
func (o *CreateFeatureNamingPatternSchema) SetExample(v string) {
	o.Example.Set(&v)
}

// SetExampleNil sets the value for Example to be an explicit nil
func (o *CreateFeatureNamingPatternSchema) SetExampleNil() {
	o.Example.Set(nil)
}

// UnsetExample ensures that no value is present for Example, not even an explicit nil
func (o *CreateFeatureNamingPatternSchema) UnsetExample() {
	o.Example.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureNamingPatternSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureNamingPatternSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFeatureNamingPatternSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateFeatureNamingPatternSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateFeatureNamingPatternSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateFeatureNamingPatternSchema) UnsetDescription() {
	o.Description.Unset()
}

func (o CreateFeatureNamingPatternSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFeatureNamingPatternSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pattern"] = o.Pattern.Get()
	if o.Example.IsSet() {
		toSerialize["example"] = o.Example.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFeatureNamingPatternSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pattern",
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

	varCreateFeatureNamingPatternSchema := _CreateFeatureNamingPatternSchema{}

	err = json.Unmarshal(data, &varCreateFeatureNamingPatternSchema)

	if err != nil {
		return err
	}

	*o = CreateFeatureNamingPatternSchema(varCreateFeatureNamingPatternSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pattern")
		delete(additionalProperties, "example")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFeatureNamingPatternSchema struct {
	value *CreateFeatureNamingPatternSchema
	isSet bool
}

func (v NullableCreateFeatureNamingPatternSchema) Get() *CreateFeatureNamingPatternSchema {
	return v.value
}

func (v *NullableCreateFeatureNamingPatternSchema) Set(val *CreateFeatureNamingPatternSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeatureNamingPatternSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeatureNamingPatternSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeatureNamingPatternSchema(val *CreateFeatureNamingPatternSchema) *NullableCreateFeatureNamingPatternSchema {
	return &NullableCreateFeatureNamingPatternSchema{value: val, isSet: true}
}

func (v NullableCreateFeatureNamingPatternSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFeatureNamingPatternSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
