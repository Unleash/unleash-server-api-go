/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PlaygroundFeatureSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundFeatureSchema{}

// PlaygroundFeatureSchema A simplified feature toggle model intended for the Unleash playground.
type PlaygroundFeatureSchema struct {
	// The feature's name.
	Name string `json:"name"`
	// The ID of the project that contains this feature.
	ProjectId string `json:"projectId"`
	Strategies PlaygroundFeatureSchemaStrategies `json:"strategies"`
	// Whether the feature is active and would be evaluated in the provided environment in a normal SDK context.
	IsEnabledInCurrentEnvironment bool `json:"isEnabledInCurrentEnvironment"`
	// Whether this feature is enabled or not in the current environment.                           If a feature can't be fully evaluated (that is, `strategies.result` is `unknown`),                           this will be `false` to align with how client SDKs treat unresolved feature states.
	IsEnabled bool `json:"isEnabled"`
	Variant NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant `json:"variant"`
	// The feature variants.
	Variants []VariantSchema `json:"variants"`
}

// NewPlaygroundFeatureSchema instantiates a new PlaygroundFeatureSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundFeatureSchema(name string, projectId string, strategies PlaygroundFeatureSchemaStrategies, isEnabledInCurrentEnvironment bool, isEnabled bool, variant NullableAdvancedPlaygroundEnvironmentFeatureSchemaVariant, variants []VariantSchema) *PlaygroundFeatureSchema {
	this := PlaygroundFeatureSchema{}
	this.Name = name
	this.ProjectId = projectId
	this.Strategies = strategies
	this.IsEnabledInCurrentEnvironment = isEnabledInCurrentEnvironment
	this.IsEnabled = isEnabled
	this.Variant = variant
	this.Variants = variants
	return &this
}

// NewPlaygroundFeatureSchemaWithDefaults instantiates a new PlaygroundFeatureSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundFeatureSchemaWithDefaults() *PlaygroundFeatureSchema {
	this := PlaygroundFeatureSchema{}
	return &this
}

// GetName returns the Name field value
func (o *PlaygroundFeatureSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlaygroundFeatureSchema) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *PlaygroundFeatureSchema) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PlaygroundFeatureSchema) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStrategies returns the Strategies field value
func (o *PlaygroundFeatureSchema) GetStrategies() PlaygroundFeatureSchemaStrategies {
	if o == nil {
		var ret PlaygroundFeatureSchemaStrategies
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetStrategiesOk() (*PlaygroundFeatureSchemaStrategies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategies, true
}

// SetStrategies sets field value
func (o *PlaygroundFeatureSchema) SetStrategies(v PlaygroundFeatureSchemaStrategies) {
	o.Strategies = v
}

// GetIsEnabledInCurrentEnvironment returns the IsEnabledInCurrentEnvironment field value
func (o *PlaygroundFeatureSchema) GetIsEnabledInCurrentEnvironment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabledInCurrentEnvironment
}

// GetIsEnabledInCurrentEnvironmentOk returns a tuple with the IsEnabledInCurrentEnvironment field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetIsEnabledInCurrentEnvironmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabledInCurrentEnvironment, true
}

// SetIsEnabledInCurrentEnvironment sets field value
func (o *PlaygroundFeatureSchema) SetIsEnabledInCurrentEnvironment(v bool) {
	o.IsEnabledInCurrentEnvironment = v
}

// GetIsEnabled returns the IsEnabled field value
func (o *PlaygroundFeatureSchema) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *PlaygroundFeatureSchema) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetVariant returns the Variant field value
// If the value is explicit nil, the zero value for AdvancedPlaygroundEnvironmentFeatureSchemaVariant will be returned
func (o *PlaygroundFeatureSchema) GetVariant() AdvancedPlaygroundEnvironmentFeatureSchemaVariant {
	if o == nil || o.Variant.Get() == nil {
		var ret AdvancedPlaygroundEnvironmentFeatureSchemaVariant
		return ret
	}

	return *o.Variant.Get()
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaygroundFeatureSchema) GetVariantOk() (*AdvancedPlaygroundEnvironmentFeatureSchemaVariant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variant.Get(), o.Variant.IsSet()
}

// SetVariant sets field value
func (o *PlaygroundFeatureSchema) SetVariant(v AdvancedPlaygroundEnvironmentFeatureSchemaVariant) {
	o.Variant.Set(&v)
}

// GetVariants returns the Variants field value
func (o *PlaygroundFeatureSchema) GetVariants() []VariantSchema {
	if o == nil {
		var ret []VariantSchema
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *PlaygroundFeatureSchema) GetVariantsOk() ([]VariantSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *PlaygroundFeatureSchema) SetVariants(v []VariantSchema) {
	o.Variants = v
}

func (o PlaygroundFeatureSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundFeatureSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	toSerialize["strategies"] = o.Strategies
	toSerialize["isEnabledInCurrentEnvironment"] = o.IsEnabledInCurrentEnvironment
	toSerialize["isEnabled"] = o.IsEnabled
	toSerialize["variant"] = o.Variant.Get()
	toSerialize["variants"] = o.Variants
	return toSerialize, nil
}

type NullablePlaygroundFeatureSchema struct {
	value *PlaygroundFeatureSchema
	isSet bool
}

func (v NullablePlaygroundFeatureSchema) Get() *PlaygroundFeatureSchema {
	return v.value
}

func (v *NullablePlaygroundFeatureSchema) Set(val *PlaygroundFeatureSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundFeatureSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundFeatureSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundFeatureSchema(val *PlaygroundFeatureSchema) *NullablePlaygroundFeatureSchema {
	return &NullablePlaygroundFeatureSchema{value: val, isSet: true}
}

func (v NullablePlaygroundFeatureSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundFeatureSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


