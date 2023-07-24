/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SegmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentSchema{}

// SegmentSchema Represents a segment of users defined by a set of constraints.
type SegmentSchema struct {
	// The segment's id.
	Id float32 `json:"id"`
	// The name of the segment.
	Name *string `json:"name,omitempty"`
	// The description of the segment.
	Description NullableString `json:"description,omitempty"`
	// List of constraints that determine which users are part of the segment
	Constraints []ConstraintSchema `json:"constraints"`
}

// NewSegmentSchema instantiates a new SegmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentSchema(id float32, constraints []ConstraintSchema) *SegmentSchema {
	this := SegmentSchema{}
	this.Id = id
	this.Constraints = constraints
	return &this
}

// NewSegmentSchemaWithDefaults instantiates a new SegmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentSchemaWithDefaults() *SegmentSchema {
	this := SegmentSchema{}
	return &this
}

// GetId returns the Id field value
func (o *SegmentSchema) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SegmentSchema) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SegmentSchema) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentSchema) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentSchema) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentSchema) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentSchema) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SegmentSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SegmentSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SegmentSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SegmentSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SegmentSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SegmentSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetConstraints returns the Constraints field value
func (o *SegmentSchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		var ret []ConstraintSchema
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *SegmentSchema) GetConstraintsOk() ([]ConstraintSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *SegmentSchema) SetConstraints(v []ConstraintSchema) {
	o.Constraints = v
}

func (o SegmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["constraints"] = o.Constraints
	return toSerialize, nil
}

type NullableSegmentSchema struct {
	value *SegmentSchema
	isSet bool
}

func (v NullableSegmentSchema) Get() *SegmentSchema {
	return v.value
}

func (v *NullableSegmentSchema) Set(val *SegmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentSchema(val *SegmentSchema) *NullableSegmentSchema {
	return &NullableSegmentSchema{value: val, isSet: true}
}

func (v NullableSegmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
