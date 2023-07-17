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

// checks if the ChangeRequestFeatureSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestFeatureSchema{}

// ChangeRequestFeatureSchema struct for ChangeRequestFeatureSchema
type ChangeRequestFeatureSchema struct {
	Name     string  `json:"name"`
	Conflict *string `json:"conflict,omitempty"`
	// List of changes inside change request. This list may be empty when listing all change requests for a project.
	Changes       []ChangeRequestEventSchema       `json:"changes"`
	DefaultChange *ChangeRequestDefaultEventSchema `json:"defaultChange,omitempty"`
}

// NewChangeRequestFeatureSchema instantiates a new ChangeRequestFeatureSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestFeatureSchema(name string, changes []ChangeRequestEventSchema) *ChangeRequestFeatureSchema {
	this := ChangeRequestFeatureSchema{}
	this.Name = name
	this.Changes = changes
	return &this
}

// NewChangeRequestFeatureSchemaWithDefaults instantiates a new ChangeRequestFeatureSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestFeatureSchemaWithDefaults() *ChangeRequestFeatureSchema {
	this := ChangeRequestFeatureSchema{}
	return &this
}

// GetName returns the Name field value
func (o *ChangeRequestFeatureSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestFeatureSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChangeRequestFeatureSchema) SetName(v string) {
	o.Name = v
}

// GetConflict returns the Conflict field value if set, zero value otherwise.
func (o *ChangeRequestFeatureSchema) GetConflict() string {
	if o == nil || IsNil(o.Conflict) {
		var ret string
		return ret
	}
	return *o.Conflict
}

// GetConflictOk returns a tuple with the Conflict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestFeatureSchema) GetConflictOk() (*string, bool) {
	if o == nil || IsNil(o.Conflict) {
		return nil, false
	}
	return o.Conflict, true
}

// HasConflict returns a boolean if a field has been set.
func (o *ChangeRequestFeatureSchema) HasConflict() bool {
	if o != nil && !IsNil(o.Conflict) {
		return true
	}

	return false
}

// SetConflict gets a reference to the given string and assigns it to the Conflict field.
func (o *ChangeRequestFeatureSchema) SetConflict(v string) {
	o.Conflict = &v
}

// GetChanges returns the Changes field value
func (o *ChangeRequestFeatureSchema) GetChanges() []ChangeRequestEventSchema {
	if o == nil {
		var ret []ChangeRequestEventSchema
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestFeatureSchema) GetChangesOk() ([]ChangeRequestEventSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *ChangeRequestFeatureSchema) SetChanges(v []ChangeRequestEventSchema) {
	o.Changes = v
}

// GetDefaultChange returns the DefaultChange field value if set, zero value otherwise.
func (o *ChangeRequestFeatureSchema) GetDefaultChange() ChangeRequestDefaultEventSchema {
	if o == nil || IsNil(o.DefaultChange) {
		var ret ChangeRequestDefaultEventSchema
		return ret
	}
	return *o.DefaultChange
}

// GetDefaultChangeOk returns a tuple with the DefaultChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestFeatureSchema) GetDefaultChangeOk() (*ChangeRequestDefaultEventSchema, bool) {
	if o == nil || IsNil(o.DefaultChange) {
		return nil, false
	}
	return o.DefaultChange, true
}

// HasDefaultChange returns a boolean if a field has been set.
func (o *ChangeRequestFeatureSchema) HasDefaultChange() bool {
	if o != nil && !IsNil(o.DefaultChange) {
		return true
	}

	return false
}

// SetDefaultChange gets a reference to the given ChangeRequestDefaultEventSchema and assigns it to the DefaultChange field.
func (o *ChangeRequestFeatureSchema) SetDefaultChange(v ChangeRequestDefaultEventSchema) {
	o.DefaultChange = &v
}

func (o ChangeRequestFeatureSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestFeatureSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Conflict) {
		toSerialize["conflict"] = o.Conflict
	}
	toSerialize["changes"] = o.Changes
	if !IsNil(o.DefaultChange) {
		toSerialize["defaultChange"] = o.DefaultChange
	}
	return toSerialize, nil
}

type NullableChangeRequestFeatureSchema struct {
	value *ChangeRequestFeatureSchema
	isSet bool
}

func (v NullableChangeRequestFeatureSchema) Get() *ChangeRequestFeatureSchema {
	return v.value
}

func (v *NullableChangeRequestFeatureSchema) Set(val *ChangeRequestFeatureSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestFeatureSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestFeatureSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestFeatureSchema(val *ChangeRequestFeatureSchema) *NullableChangeRequestFeatureSchema {
	return &NullableChangeRequestFeatureSchema{value: val, isSet: true}
}

func (v NullableChangeRequestFeatureSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestFeatureSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
