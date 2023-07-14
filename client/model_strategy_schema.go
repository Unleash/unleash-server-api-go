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

// checks if the StrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StrategySchema{}

// StrategySchema The [activation strategy](https://docs.getunleash.io/reference/activation-strategies) schema
type StrategySchema struct {
	// An optional title for the strategy
	Title NullableString `json:"title,omitempty"`
	// The name (type) of the strategy
	Name string `json:"name"`
	// A human friendly name for the strategy
	DisplayName NullableString `json:"displayName"`
	// A short description of the strategy
	Description NullableString `json:"description"`
	// Whether the strategy can be edited or not. Strategies bundled with Unleash cannot be edited.
	Editable bool `json:"editable"`
	//
	Deprecated bool `json:"deprecated"`
	// A list of relevant parameters for each strategy
	Parameters []StrategySchemaParametersInner `json:"parameters"`
}

// NewStrategySchema instantiates a new StrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategySchema(name string, displayName NullableString, description NullableString, editable bool, deprecated bool, parameters []StrategySchemaParametersInner) *StrategySchema {
	this := StrategySchema{}
	this.Name = name
	this.DisplayName = displayName
	this.Description = description
	this.Editable = editable
	this.Deprecated = deprecated
	this.Parameters = parameters
	return &this
}

// NewStrategySchemaWithDefaults instantiates a new StrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategySchemaWithDefaults() *StrategySchema {
	this := StrategySchema{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StrategySchema) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StrategySchema) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *StrategySchema) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *StrategySchema) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *StrategySchema) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *StrategySchema) UnsetTitle() {
	o.Title.Unset()
}

// GetName returns the Name field value
func (o *StrategySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StrategySchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StrategySchema) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StrategySchema) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StrategySchema) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// SetDisplayName sets field value
func (o *StrategySchema) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StrategySchema) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StrategySchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *StrategySchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEditable returns the Editable field value
func (o *StrategySchema) GetEditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Editable
}

// GetEditableOk returns a tuple with the Editable field value
// and a boolean to check if the value has been set.
func (o *StrategySchema) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Editable, true
}

// SetEditable sets field value
func (o *StrategySchema) SetEditable(v bool) {
	o.Editable = v
}

// GetDeprecated returns the Deprecated field value
func (o *StrategySchema) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *StrategySchema) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *StrategySchema) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetParameters returns the Parameters field value
func (o *StrategySchema) GetParameters() []StrategySchemaParametersInner {
	if o == nil {
		var ret []StrategySchemaParametersInner
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *StrategySchema) GetParametersOk() ([]StrategySchemaParametersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *StrategySchema) SetParameters(v []StrategySchemaParametersInner) {
	o.Parameters = v
}

func (o StrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["displayName"] = o.DisplayName.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["editable"] = o.Editable
	toSerialize["deprecated"] = o.Deprecated
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

type NullableStrategySchema struct {
	value *StrategySchema
	isSet bool
}

func (v NullableStrategySchema) Get() *StrategySchema {
	return v.value
}

func (v *NullableStrategySchema) Set(val *StrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategySchema(val *StrategySchema) *NullableStrategySchema {
	return &NullableStrategySchema{value: val, isSet: true}
}

func (v NullableStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
