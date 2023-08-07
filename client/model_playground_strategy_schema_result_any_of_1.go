/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PlaygroundStrategySchemaResultAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchemaResultAnyOf1{}

// PlaygroundStrategySchemaResultAnyOf1 struct for PlaygroundStrategySchemaResultAnyOf1
type PlaygroundStrategySchemaResultAnyOf1 struct {
	// Signals that this strategy was evaluated successfully.
	EvaluationStatus string `json:"evaluationStatus"`
	// Whether this strategy evaluates to true or not.
	Enabled bool                                                `json:"enabled"`
	Variant NullablePlaygroundStrategySchemaResultAnyOf1Variant `json:"variant,omitempty"`
	// The feature variants.
	Variants []VariantSchema `json:"variants,omitempty"`
}

// NewPlaygroundStrategySchemaResultAnyOf1 instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchemaResultAnyOf1(evaluationStatus string, enabled bool) *PlaygroundStrategySchemaResultAnyOf1 {
	this := PlaygroundStrategySchemaResultAnyOf1{}
	this.EvaluationStatus = evaluationStatus
	this.Enabled = enabled
	return &this
}

// NewPlaygroundStrategySchemaResultAnyOf1WithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaResultAnyOf1WithDefaults() *PlaygroundStrategySchemaResultAnyOf1 {
	this := PlaygroundStrategySchemaResultAnyOf1{}
	return &this
}

// GetEvaluationStatus returns the EvaluationStatus field value
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationStatus
}

// GetEvaluationStatusOk returns a tuple with the EvaluationStatus field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationStatus, true
}

// SetEvaluationStatus sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1) SetEvaluationStatus(v string) {
	o.EvaluationStatus = v
}

// GetEnabled returns the Enabled field value
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PlaygroundStrategySchemaResultAnyOf1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetVariant returns the Variant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariant() PlaygroundStrategySchemaResultAnyOf1Variant {
	if o == nil || IsNil(o.Variant.Get()) {
		var ret PlaygroundStrategySchemaResultAnyOf1Variant
		return ret
	}
	return *o.Variant.Get()
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariantOk() (*PlaygroundStrategySchemaResultAnyOf1Variant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variant.Get(), o.Variant.IsSet()
}

// HasVariant returns a boolean if a field has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) HasVariant() bool {
	if o != nil && o.Variant.IsSet() {
		return true
	}

	return false
}

// SetVariant gets a reference to the given NullablePlaygroundStrategySchemaResultAnyOf1Variant and assigns it to the Variant field.
func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariant(v PlaygroundStrategySchemaResultAnyOf1Variant) {
	o.Variant.Set(&v)
}

// SetVariantNil sets the value for Variant to be an explicit nil
func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariantNil() {
	o.Variant.Set(nil)
}

// UnsetVariant ensures that no value is present for Variant, not even an explicit nil
func (o *PlaygroundStrategySchemaResultAnyOf1) UnsetVariant() {
	o.Variant.Unset()
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariants() []VariantSchema {
	if o == nil || IsNil(o.Variants) {
		var ret []VariantSchema
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariantsOk() ([]VariantSchema, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *PlaygroundStrategySchemaResultAnyOf1) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []VariantSchema and assigns it to the Variants field.
func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariants(v []VariantSchema) {
	o.Variants = v
}

func (o PlaygroundStrategySchemaResultAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchemaResultAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["evaluationStatus"] = o.EvaluationStatus
	toSerialize["enabled"] = o.Enabled
	if o.Variant.IsSet() {
		toSerialize["variant"] = o.Variant.Get()
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return toSerialize, nil
}

type NullablePlaygroundStrategySchemaResultAnyOf1 struct {
	value *PlaygroundStrategySchemaResultAnyOf1
	isSet bool
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) Get() *PlaygroundStrategySchemaResultAnyOf1 {
	return v.value
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) Set(val *PlaygroundStrategySchemaResultAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchemaResultAnyOf1(val *PlaygroundStrategySchemaResultAnyOf1) *NullablePlaygroundStrategySchemaResultAnyOf1 {
	return &NullablePlaygroundStrategySchemaResultAnyOf1{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchemaResultAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchemaResultAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
