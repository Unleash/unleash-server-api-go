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

// checks if the UpdateFeatureStrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFeatureStrategySchema{}

// UpdateFeatureStrategySchema Update a strategy configuration in a feature
type UpdateFeatureStrategySchema struct {
	// The name of the strategy type
	Name *string `json:"name,omitempty"`
	// The order of the strategy in the list in feature environment configuration
	SortOrder *float32 `json:"sortOrder,omitempty"`
	// A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints
	Constraints []ConstraintSchema `json:"constraints,omitempty"`
	// A descriptive title for the strategy
	Title NullableString `json:"title,omitempty"`
	// A toggle to disable the strategy. defaults to true. Disabled strategies are not evaluated or returned to the SDKs
	Disabled NullableBool `json:"disabled,omitempty"`
	// A list of parameters for a strategy
	Parameters *map[string]string `json:"parameters,omitempty"`
}

// NewUpdateFeatureStrategySchema instantiates a new UpdateFeatureStrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFeatureStrategySchema() *UpdateFeatureStrategySchema {
	this := UpdateFeatureStrategySchema{}
	return &this
}

// NewUpdateFeatureStrategySchemaWithDefaults instantiates a new UpdateFeatureStrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFeatureStrategySchemaWithDefaults() *UpdateFeatureStrategySchema {
	this := UpdateFeatureStrategySchema{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFeatureStrategySchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureStrategySchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFeatureStrategySchema) SetName(v string) {
	o.Name = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *UpdateFeatureStrategySchema) GetSortOrder() float32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret float32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureStrategySchema) GetSortOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given float32 and assigns it to the SortOrder field.
func (o *UpdateFeatureStrategySchema) SetSortOrder(v float32) {
	o.SortOrder = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *UpdateFeatureStrategySchema) GetConstraints() []ConstraintSchema {
	if o == nil || IsNil(o.Constraints) {
		var ret []ConstraintSchema
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureStrategySchema) GetConstraintsOk() ([]ConstraintSchema, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []ConstraintSchema and assigns it to the Constraints field.
func (o *UpdateFeatureStrategySchema) SetConstraints(v []ConstraintSchema) {
	o.Constraints = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureStrategySchema) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureStrategySchema) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UpdateFeatureStrategySchema) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *UpdateFeatureStrategySchema) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UpdateFeatureStrategySchema) UnsetTitle() {
	o.Title.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFeatureStrategySchema) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFeatureStrategySchema) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *UpdateFeatureStrategySchema) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *UpdateFeatureStrategySchema) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *UpdateFeatureStrategySchema) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *UpdateFeatureStrategySchema) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFeatureStrategySchema) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *UpdateFeatureStrategySchema) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *UpdateFeatureStrategySchema) SetParameters(v map[string]string) {
	o.Parameters = &v
}

func (o UpdateFeatureStrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFeatureStrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableUpdateFeatureStrategySchema struct {
	value *UpdateFeatureStrategySchema
	isSet bool
}

func (v NullableUpdateFeatureStrategySchema) Get() *UpdateFeatureStrategySchema {
	return v.value
}

func (v *NullableUpdateFeatureStrategySchema) Set(val *UpdateFeatureStrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFeatureStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFeatureStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFeatureStrategySchema(val *UpdateFeatureStrategySchema) *NullableUpdateFeatureStrategySchema {
	return &NullableUpdateFeatureStrategySchema{value: val, isSet: true}
}

func (v NullableUpdateFeatureStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFeatureStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
