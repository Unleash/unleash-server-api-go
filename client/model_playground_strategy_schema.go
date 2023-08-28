/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the PlaygroundStrategySchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaygroundStrategySchema{}

// PlaygroundStrategySchema An evaluated feature toggle strategy as used by the Playground
type PlaygroundStrategySchema struct {
	// The strategy's name.
	Name string `json:"name"`
	// Description of the feature's purpose.
	Title *string `json:"title,omitempty"`
	// The strategy's id.
	Id     string                         `json:"id"`
	Result PlaygroundStrategySchemaResult `json:"result"`
	// The strategy's status. Disabled strategies are not evaluated
	Disabled NullableBool `json:"disabled"`
	// The strategy's segments and their evaluation results.
	Segments []PlaygroundSegmentSchema `json:"segments"`
	// The strategy's constraints and their evaluation results.
	Constraints []PlaygroundConstraintSchema `json:"constraints"`
	// A list of parameters for a strategy
	Parameters map[string]string             `json:"parameters"`
	Links      PlaygroundStrategySchemaLinks `json:"links"`
}

// NewPlaygroundStrategySchema instantiates a new PlaygroundStrategySchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaygroundStrategySchema(name string, id string, result PlaygroundStrategySchemaResult, disabled NullableBool, segments []PlaygroundSegmentSchema, constraints []PlaygroundConstraintSchema, parameters map[string]string, links PlaygroundStrategySchemaLinks) *PlaygroundStrategySchema {
	this := PlaygroundStrategySchema{}
	this.Name = name
	this.Id = id
	this.Result = result
	this.Disabled = disabled
	this.Segments = segments
	this.Constraints = constraints
	this.Parameters = parameters
	this.Links = links
	return &this
}

// NewPlaygroundStrategySchemaWithDefaults instantiates a new PlaygroundStrategySchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaygroundStrategySchemaWithDefaults() *PlaygroundStrategySchema {
	this := PlaygroundStrategySchema{}
	return &this
}

// GetName returns the Name field value
func (o *PlaygroundStrategySchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlaygroundStrategySchema) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlaygroundStrategySchema) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlaygroundStrategySchema) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlaygroundStrategySchema) SetTitle(v string) {
	o.Title = &v
}

// GetId returns the Id field value
func (o *PlaygroundStrategySchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PlaygroundStrategySchema) SetId(v string) {
	o.Id = v
}

// GetResult returns the Result field value
func (o *PlaygroundStrategySchema) GetResult() PlaygroundStrategySchemaResult {
	if o == nil {
		var ret PlaygroundStrategySchemaResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetResultOk() (*PlaygroundStrategySchemaResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PlaygroundStrategySchema) SetResult(v PlaygroundStrategySchemaResult) {
	o.Result = v
}

// GetDisabled returns the Disabled field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PlaygroundStrategySchema) GetDisabled() bool {
	if o == nil || o.Disabled.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Disabled.Get()
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaygroundStrategySchema) GetDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disabled.Get(), o.Disabled.IsSet()
}

// SetDisabled sets field value
func (o *PlaygroundStrategySchema) SetDisabled(v bool) {
	o.Disabled.Set(&v)
}

// GetSegments returns the Segments field value
func (o *PlaygroundStrategySchema) GetSegments() []PlaygroundSegmentSchema {
	if o == nil {
		var ret []PlaygroundSegmentSchema
		return ret
	}

	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetSegmentsOk() ([]PlaygroundSegmentSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Segments, true
}

// SetSegments sets field value
func (o *PlaygroundStrategySchema) SetSegments(v []PlaygroundSegmentSchema) {
	o.Segments = v
}

// GetConstraints returns the Constraints field value
func (o *PlaygroundStrategySchema) GetConstraints() []PlaygroundConstraintSchema {
	if o == nil {
		var ret []PlaygroundConstraintSchema
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetConstraintsOk() ([]PlaygroundConstraintSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *PlaygroundStrategySchema) SetConstraints(v []PlaygroundConstraintSchema) {
	o.Constraints = v
}

// GetParameters returns the Parameters field value
func (o *PlaygroundStrategySchema) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetParametersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *PlaygroundStrategySchema) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetLinks returns the Links field value
func (o *PlaygroundStrategySchema) GetLinks() PlaygroundStrategySchemaLinks {
	if o == nil {
		var ret PlaygroundStrategySchemaLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PlaygroundStrategySchema) GetLinksOk() (*PlaygroundStrategySchemaLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PlaygroundStrategySchema) SetLinks(v PlaygroundStrategySchemaLinks) {
	o.Links = v
}

func (o PlaygroundStrategySchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaygroundStrategySchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["id"] = o.Id
	toSerialize["result"] = o.Result
	toSerialize["disabled"] = o.Disabled.Get()
	toSerialize["segments"] = o.Segments
	toSerialize["constraints"] = o.Constraints
	toSerialize["parameters"] = o.Parameters
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullablePlaygroundStrategySchema struct {
	value *PlaygroundStrategySchema
	isSet bool
}

func (v NullablePlaygroundStrategySchema) Get() *PlaygroundStrategySchema {
	return v.value
}

func (v *NullablePlaygroundStrategySchema) Set(val *PlaygroundStrategySchema) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaygroundStrategySchema) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaygroundStrategySchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaygroundStrategySchema(val *PlaygroundStrategySchema) *NullablePlaygroundStrategySchema {
	return &NullablePlaygroundStrategySchema{value: val, isSet: true}
}

func (v NullablePlaygroundStrategySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaygroundStrategySchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
