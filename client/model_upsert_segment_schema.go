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

// checks if the UpsertSegmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertSegmentSchema{}

// UpsertSegmentSchema struct for UpsertSegmentSchema
type UpsertSegmentSchema struct {
	Name        string             `json:"name"`
	Description NullableString     `json:"description,omitempty"`
	Project     NullableString     `json:"project,omitempty"`
	Constraints []ConstraintSchema `json:"constraints"`
}

// NewUpsertSegmentSchema instantiates a new UpsertSegmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertSegmentSchema(name string, constraints []ConstraintSchema) *UpsertSegmentSchema {
	this := UpsertSegmentSchema{}
	this.Name = name
	this.Constraints = constraints
	return &this
}

// NewUpsertSegmentSchemaWithDefaults instantiates a new UpsertSegmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertSegmentSchemaWithDefaults() *UpsertSegmentSchema {
	this := UpsertSegmentSchema{}
	return &this
}

// GetName returns the Name field value
func (o *UpsertSegmentSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpsertSegmentSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpsertSegmentSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertSegmentSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertSegmentSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpsertSegmentSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpsertSegmentSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpsertSegmentSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpsertSegmentSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertSegmentSchema) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertSegmentSchema) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *UpsertSegmentSchema) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *UpsertSegmentSchema) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *UpsertSegmentSchema) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *UpsertSegmentSchema) UnsetProject() {
	o.Project.Unset()
}

// GetConstraints returns the Constraints field value
func (o *UpsertSegmentSchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		var ret []ConstraintSchema
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *UpsertSegmentSchema) GetConstraintsOk() ([]ConstraintSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *UpsertSegmentSchema) SetConstraints(v []ConstraintSchema) {
	o.Constraints = v
}

func (o UpsertSegmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertSegmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	toSerialize["constraints"] = o.Constraints
	return toSerialize, nil
}

type NullableUpsertSegmentSchema struct {
	value *UpsertSegmentSchema
	isSet bool
}

func (v NullableUpsertSegmentSchema) Get() *UpsertSegmentSchema {
	return v.value
}

func (v *NullableUpsertSegmentSchema) Set(val *UpsertSegmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertSegmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertSegmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertSegmentSchema(val *UpsertSegmentSchema) *NullableUpsertSegmentSchema {
	return &NullableUpsertSegmentSchema{value: val, isSet: true}
}

func (v NullableUpsertSegmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertSegmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
