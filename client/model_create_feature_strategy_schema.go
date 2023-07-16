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

// checks if the CreateFeatureStrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFeatureStrategySchema{}

// CreateFeatureStrategySchema Create a strategy configuration in a feature
type CreateFeatureStrategySchema struct {
	// The name of the strategy type
	Name string `json:"name"`
	// A descriptive title for the strategy
	Title NullableString `json:"title,omitempty"`
	// A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs
	Disabled NullableBool `json:"disabled,omitempty"`
	// The order of the strategy in the list
	SortOrder *float32 `json:"sortOrder,omitempty"`
	// A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints
	Constraints []ConstraintSchema `json:"constraints,omitempty"`
	// Strategy level variants
	Variants []CreateStrategyVariantSchema `json:"variants,omitempty"`
	// A list of parameters for a strategy
	Parameters *map[string]string `json:"parameters,omitempty"`
	// Ids of segments to use for this strategy
	Segments []float32 `json:"segments,omitempty"`
}

// NewCreateFeatureStrategySchema instantiates a new CreateFeatureStrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFeatureStrategySchema(name string) *CreateFeatureStrategySchema {
	this := CreateFeatureStrategySchema{}
	this.Name = name
	return &this
}

// NewCreateFeatureStrategySchemaWithDefaults instantiates a new CreateFeatureStrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFeatureStrategySchemaWithDefaults() *CreateFeatureStrategySchema {
	this := CreateFeatureStrategySchema{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFeatureStrategySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFeatureStrategySchema) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureStrategySchema) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureStrategySchema) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *CreateFeatureStrategySchema) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *CreateFeatureStrategySchema) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *CreateFeatureStrategySchema) UnsetTitle() {
	o.Title.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFeatureStrategySchema) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFeatureStrategySchema) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// HasDisabled returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasDisabled() bool {
	if o != nil && o.Disabled.IsSet() {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given NullableBool and assigns it to the Disabled field.
func (o *CreateFeatureStrategySchema) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// SetDisabledNil sets the value for Disabled to be an explicit nil
func (o *CreateFeatureStrategySchema) SetDisabledNil() {
	o.Disabled.Set(nil)
}

// UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
func (o *CreateFeatureStrategySchema) UnsetDisabled() {
	o.Disabled.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *CreateFeatureStrategySchema) GetSortOrder() float32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret float32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetSortOrderOk() (*float32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given float32 and assigns it to the SortOrder field.
func (o *CreateFeatureStrategySchema) SetSortOrder(v float32) {
	o.SortOrder = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *CreateFeatureStrategySchema) GetConstraints() []ConstraintSchema {
	if o == nil || IsNil(o.Constraints) {
		var ret []ConstraintSchema
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetConstraintsOk() ([]ConstraintSchema, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []ConstraintSchema and assigns it to the Constraints field.
func (o *CreateFeatureStrategySchema) SetConstraints(v []ConstraintSchema) {
	o.Constraints = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *CreateFeatureStrategySchema) GetVariants() []CreateStrategyVariantSchema {
	if o == nil || IsNil(o.Variants) {
		var ret []CreateStrategyVariantSchema
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetVariantsOk() ([]CreateStrategyVariantSchema, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []CreateStrategyVariantSchema and assigns it to the Variants field.
func (o *CreateFeatureStrategySchema) SetVariants(v []CreateStrategyVariantSchema) {
	o.Variants = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateFeatureStrategySchema) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CreateFeatureStrategySchema) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetSegments returns the Segments field value if set, zero value otherwise.
func (o *CreateFeatureStrategySchema) GetSegments() []float32 {
	if o == nil || IsNil(o.Segments) {
		var ret []float32
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFeatureStrategySchema) GetSegmentsOk() ([]float32, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *CreateFeatureStrategySchema) HasSegments() bool {
	if o != nil && !IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []float32 and assigns it to the Segments field.
func (o *CreateFeatureStrategySchema) SetSegments(v []float32) {
	o.Segments = v
}

func (o CreateFeatureStrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFeatureStrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Disabled.IsSet() {
		toSerialize["disabled"] = o.Disabled.Get()
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Segments) {
		toSerialize["segments"] = o.Segments
	}
	return toSerialize, nil
}

type NullableCreateFeatureStrategySchema struct {
	value *CreateFeatureStrategySchema
	isSet bool
}

func (v NullableCreateFeatureStrategySchema) Get() *CreateFeatureStrategySchema {
	return v.value
}

func (v *NullableCreateFeatureStrategySchema) Set(val *CreateFeatureStrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFeatureStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFeatureStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFeatureStrategySchema(val *CreateFeatureStrategySchema) *NullableCreateFeatureStrategySchema {
	return &NullableCreateFeatureStrategySchema{value: val, isSet: true}
}

func (v NullableCreateFeatureStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFeatureStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
