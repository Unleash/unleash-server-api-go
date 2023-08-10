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

// checks if the ConstraintSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConstraintSchema{}

// ConstraintSchema A strategy constraint. For more information, refer to [the strategy constraint reference documentation](https://docs.getunleash.io/reference/strategy-constraints)
type ConstraintSchema struct {
	// The name of the context field that this constraint should apply to.
	ContextName string `json:"contextName"`
	// The operator to use when evaluating this constraint. For more information about the various operators, refer to [the strategy constraint operator documentation](https://docs.getunleash.io/reference/strategy-constraints#strategy-constraint-operators).
	Operator string `json:"operator"`
	// Whether the operator should be case sensitive or not. Defaults to `false` (being case sensitive).
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`
	// Whether the result should be negated or not. If `true`, will turn a `true` result into a `false` result and vice versa.
	Inverted *bool `json:"inverted,omitempty"`
	// The context values that should be used for constraint evaluation. Use this property instead of `value` for properties that accept multiple values.
	Values []string `json:"values,omitempty"`
	// The context value that should be used for constraint evaluation. Use this property instead of `values` for properties that only accept single values.
	Value *string `json:"value,omitempty"`
}

// NewConstraintSchema instantiates a new ConstraintSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraintSchema(contextName string, operator string) *ConstraintSchema {
	this := ConstraintSchema{}
	this.ContextName = contextName
	this.Operator = operator
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	var inverted bool = false
	this.Inverted = &inverted
	return &this
}

// NewConstraintSchemaWithDefaults instantiates a new ConstraintSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintSchemaWithDefaults() *ConstraintSchema {
	this := ConstraintSchema{}
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	var inverted bool = false
	this.Inverted = &inverted
	return &this
}

// GetContextName returns the ContextName field value
func (o *ConstraintSchema) GetContextName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextName
}

// GetContextNameOk returns a tuple with the ContextName field value
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetContextNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextName, true
}

// SetContextName sets field value
func (o *ConstraintSchema) SetContextName(v string) {
	o.ContextName = v
}

// GetOperator returns the Operator field value
func (o *ConstraintSchema) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ConstraintSchema) SetOperator(v string) {
	o.Operator = v
}

// GetCaseInsensitive returns the CaseInsensitive field value if set, zero value otherwise.
func (o *ConstraintSchema) GetCaseInsensitive() bool {
	if o == nil || IsNil(o.CaseInsensitive) {
		var ret bool
		return ret
	}
	return *o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseInsensitive) {
		return nil, false
	}
	return o.CaseInsensitive, true
}

// HasCaseInsensitive returns a boolean if a field has been set.
func (o *ConstraintSchema) HasCaseInsensitive() bool {
	if o != nil && !IsNil(o.CaseInsensitive) {
		return true
	}

	return false
}

// SetCaseInsensitive gets a reference to the given bool and assigns it to the CaseInsensitive field.
func (o *ConstraintSchema) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = &v
}

// GetInverted returns the Inverted field value if set, zero value otherwise.
func (o *ConstraintSchema) GetInverted() bool {
	if o == nil || IsNil(o.Inverted) {
		var ret bool
		return ret
	}
	return *o.Inverted
}

// GetInvertedOk returns a tuple with the Inverted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetInvertedOk() (*bool, bool) {
	if o == nil || IsNil(o.Inverted) {
		return nil, false
	}
	return o.Inverted, true
}

// HasInverted returns a boolean if a field has been set.
func (o *ConstraintSchema) HasInverted() bool {
	if o != nil && !IsNil(o.Inverted) {
		return true
	}

	return false
}

// SetInverted gets a reference to the given bool and assigns it to the Inverted field.
func (o *ConstraintSchema) SetInverted(v bool) {
	o.Inverted = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ConstraintSchema) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ConstraintSchema) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ConstraintSchema) SetValues(v []string) {
	o.Values = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConstraintSchema) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConstraintSchema) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConstraintSchema) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConstraintSchema) SetValue(v string) {
	o.Value = &v
}

func (o ConstraintSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConstraintSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contextName"] = o.ContextName
	toSerialize["operator"] = o.Operator
	if !IsNil(o.CaseInsensitive) {
		toSerialize["caseInsensitive"] = o.CaseInsensitive
	}
	if !IsNil(o.Inverted) {
		toSerialize["inverted"] = o.Inverted
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableConstraintSchema struct {
	value *ConstraintSchema
	isSet bool
}

func (v NullableConstraintSchema) Get() *ConstraintSchema {
	return v.value
}

func (v *NullableConstraintSchema) Set(val *ConstraintSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintSchema(val *ConstraintSchema) *NullableConstraintSchema {
	return &NullableConstraintSchema{value: val, isSet: true}
}

func (v NullableConstraintSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraintSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
