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

// checks if the PlaygroundConstraintSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundConstraintSchema{}

// PlaygroundConstraintSchema A strategy constraint. For more information, refer to [the strategy constraint reference documentation](https://docs.getunleash.io/reference/strategy-constraints)
type PlaygroundConstraintSchema struct {
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
	// Whether this was evaluated as true or false.
	Result bool `json:"result"`
}

// NewPlaygroundConstraintSchema instantiates a new PlaygroundConstraintSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundConstraintSchema(contextName string, operator string, result bool) *PlaygroundConstraintSchema {
	this := PlaygroundConstraintSchema{}
	this.ContextName = contextName
	this.Operator = operator
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	var inverted bool = false
	this.Inverted = &inverted
	this.Result = result
	return &this
}

// NewPlaygroundConstraintSchemaWithDefaults instantiates a new PlaygroundConstraintSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundConstraintSchemaWithDefaults() *PlaygroundConstraintSchema {
	this := PlaygroundConstraintSchema{}
	var caseInsensitive bool = false
	this.CaseInsensitive = &caseInsensitive
	var inverted bool = false
	this.Inverted = &inverted
	return &this
}

// GetContextName returns the ContextName field value
func (o *PlaygroundConstraintSchema) GetContextName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextName
}

// GetContextNameOk returns a tuple with the ContextName field value
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetContextNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextName, true
}

// SetContextName sets field value
func (o *PlaygroundConstraintSchema) SetContextName(v string) {
	o.ContextName = v
}

// GetOperator returns the Operator field value
func (o *PlaygroundConstraintSchema) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PlaygroundConstraintSchema) SetOperator(v string) {
	o.Operator = v
}

// GetCaseInsensitive returns the CaseInsensitive field value if set, zero value otherwise.
func (o *PlaygroundConstraintSchema) GetCaseInsensitive() bool {
	if o == nil || IsNil(o.CaseInsensitive) {
		var ret bool
		return ret
	}
	return *o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseInsensitive) {
		return nil, false
	}
	return o.CaseInsensitive, true
}

// HasCaseInsensitive returns a boolean if a field has been set.
func (o *PlaygroundConstraintSchema) HasCaseInsensitive() bool {
	if o != nil && !IsNil(o.CaseInsensitive) {
		return true
	}

	return false
}

// SetCaseInsensitive gets a reference to the given bool and assigns it to the CaseInsensitive field.
func (o *PlaygroundConstraintSchema) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = &v
}

// GetInverted returns the Inverted field value if set, zero value otherwise.
func (o *PlaygroundConstraintSchema) GetInverted() bool {
	if o == nil || IsNil(o.Inverted) {
		var ret bool
		return ret
	}
	return *o.Inverted
}

// GetInvertedOk returns a tuple with the Inverted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetInvertedOk() (*bool, bool) {
	if o == nil || IsNil(o.Inverted) {
		return nil, false
	}
	return o.Inverted, true
}

// HasInverted returns a boolean if a field has been set.
func (o *PlaygroundConstraintSchema) HasInverted() bool {
	if o != nil && !IsNil(o.Inverted) {
		return true
	}

	return false
}

// SetInverted gets a reference to the given bool and assigns it to the Inverted field.
func (o *PlaygroundConstraintSchema) SetInverted(v bool) {
	o.Inverted = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PlaygroundConstraintSchema) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PlaygroundConstraintSchema) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *PlaygroundConstraintSchema) SetValues(v []string) {
	o.Values = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PlaygroundConstraintSchema) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlaygroundConstraintSchema) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PlaygroundConstraintSchema) SetValue(v string) {
	o.Value = &v
}

// GetResult returns the Result field value
func (o *PlaygroundConstraintSchema) GetResult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PlaygroundConstraintSchema) GetResultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PlaygroundConstraintSchema) SetResult(v bool) {
	o.Result = v
}

func (o PlaygroundConstraintSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundConstraintSchema) ToMap() (map[string]interface{}, error) {
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
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

type NullablePlaygroundConstraintSchema struct {
	value *PlaygroundConstraintSchema
	isSet bool
}

func (v NullablePlaygroundConstraintSchema) Get() *PlaygroundConstraintSchema {
	return v.value
}

func (v *NullablePlaygroundConstraintSchema) Set(val *PlaygroundConstraintSchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundConstraintSchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundConstraintSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundConstraintSchema(val *PlaygroundConstraintSchema) *NullablePlaygroundConstraintSchema {
	return &NullablePlaygroundConstraintSchema{value: val, isSet: true}
}

func (v NullablePlaygroundConstraintSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundConstraintSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
