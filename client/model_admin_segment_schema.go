/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.4.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the AdminSegmentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminSegmentSchema{}

// AdminSegmentSchema A description of a [segment](https://docs.getunleash.io/reference/segments)
type AdminSegmentSchema struct {
	// The ID of this segment
	Id int32 `json:"id"`
	// The name of this segment
	Name string `json:"name"`
	// The description for this segment
	Description NullableString `json:"description,omitempty"`
	// The list of constraints that are used in this segment
	Constraints []ConstraintSchema `json:"constraints"`
	// The number of projects that use this segment
	UsedInFeatures NullableInt32 `json:"usedInFeatures,omitempty"`
	// The number of projects that use this segment
	UsedInProjects NullableInt32 `json:"usedInProjects,omitempty"`
	// The project the segment belongs to. Only present if the segment is a project-specific segment.
	Project NullableString `json:"project,omitempty"`
	// The creator's email or username
	CreatedBy NullableString `json:"createdBy,omitempty"`
	// When the segment was created
	CreatedAt time.Time `json:"createdAt"`
}

// NewAdminSegmentSchema instantiates a new AdminSegmentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminSegmentSchema(id int32, name string, constraints []ConstraintSchema, createdAt time.Time) *AdminSegmentSchema {
	this := AdminSegmentSchema{}
	this.Id = id
	this.Name = name
	this.Constraints = constraints
	this.CreatedAt = createdAt
	return &this
}

// NewAdminSegmentSchemaWithDefaults instantiates a new AdminSegmentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminSegmentSchemaWithDefaults() *AdminSegmentSchema {
	this := AdminSegmentSchema{}
	return &this
}

// GetId returns the Id field value
func (o *AdminSegmentSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AdminSegmentSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AdminSegmentSchema) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AdminSegmentSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminSegmentSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdminSegmentSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSegmentSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSegmentSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AdminSegmentSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AdminSegmentSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AdminSegmentSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AdminSegmentSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetConstraints returns the Constraints field value
func (o *AdminSegmentSchema) GetConstraints() []ConstraintSchema {
	if o == nil {
		var ret []ConstraintSchema
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *AdminSegmentSchema) GetConstraintsOk() ([]ConstraintSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *AdminSegmentSchema) SetConstraints(v []ConstraintSchema) {
	o.Constraints = v
}

// GetUsedInFeatures returns the UsedInFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSegmentSchema) GetUsedInFeatures() int32 {
	if o == nil || IsNil(o.UsedInFeatures.Get()) {
		var ret int32
		return ret
	}
	return *o.UsedInFeatures.Get()
}

// GetUsedInFeaturesOk returns a tuple with the UsedInFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSegmentSchema) GetUsedInFeaturesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedInFeatures.Get(), o.UsedInFeatures.IsSet()
}

// HasUsedInFeatures returns a boolean if a field has been set.
func (o *AdminSegmentSchema) HasUsedInFeatures() bool {
	if o != nil && o.UsedInFeatures.IsSet() {
		return true
	}

	return false
}

// SetUsedInFeatures gets a reference to the given NullableInt32 and assigns it to the UsedInFeatures field.
func (o *AdminSegmentSchema) SetUsedInFeatures(v int32) {
	o.UsedInFeatures.Set(&v)
}

// SetUsedInFeaturesNil sets the value for UsedInFeatures to be an explicit nil
func (o *AdminSegmentSchema) SetUsedInFeaturesNil() {
	o.UsedInFeatures.Set(nil)
}

// UnsetUsedInFeatures ensures that no value is present for UsedInFeatures, not even an explicit nil
func (o *AdminSegmentSchema) UnsetUsedInFeatures() {
	o.UsedInFeatures.Unset()
}

// GetUsedInProjects returns the UsedInProjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSegmentSchema) GetUsedInProjects() int32 {
	if o == nil || IsNil(o.UsedInProjects.Get()) {
		var ret int32
		return ret
	}
	return *o.UsedInProjects.Get()
}

// GetUsedInProjectsOk returns a tuple with the UsedInProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSegmentSchema) GetUsedInProjectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedInProjects.Get(), o.UsedInProjects.IsSet()
}

// HasUsedInProjects returns a boolean if a field has been set.
func (o *AdminSegmentSchema) HasUsedInProjects() bool {
	if o != nil && o.UsedInProjects.IsSet() {
		return true
	}

	return false
}

// SetUsedInProjects gets a reference to the given NullableInt32 and assigns it to the UsedInProjects field.
func (o *AdminSegmentSchema) SetUsedInProjects(v int32) {
	o.UsedInProjects.Set(&v)
}

// SetUsedInProjectsNil sets the value for UsedInProjects to be an explicit nil
func (o *AdminSegmentSchema) SetUsedInProjectsNil() {
	o.UsedInProjects.Set(nil)
}

// UnsetUsedInProjects ensures that no value is present for UsedInProjects, not even an explicit nil
func (o *AdminSegmentSchema) UnsetUsedInProjects() {
	o.UsedInProjects.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSegmentSchema) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSegmentSchema) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *AdminSegmentSchema) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *AdminSegmentSchema) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *AdminSegmentSchema) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *AdminSegmentSchema) UnsetProject() {
	o.Project.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdminSegmentSchema) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdminSegmentSchema) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AdminSegmentSchema) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *AdminSegmentSchema) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *AdminSegmentSchema) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *AdminSegmentSchema) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *AdminSegmentSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AdminSegmentSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AdminSegmentSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o AdminSegmentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminSegmentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["constraints"] = o.Constraints
	if o.UsedInFeatures.IsSet() {
		toSerialize["usedInFeatures"] = o.UsedInFeatures.Get()
	}
	if o.UsedInProjects.IsSet() {
		toSerialize["usedInProjects"] = o.UsedInProjects.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableAdminSegmentSchema struct {
	value *AdminSegmentSchema
	isSet bool
}

func (v NullableAdminSegmentSchema) Get() *AdminSegmentSchema {
	return v.value
}

func (v *NullableAdminSegmentSchema) Set(val *AdminSegmentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminSegmentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminSegmentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminSegmentSchema(val *AdminSegmentSchema) *NullableAdminSegmentSchema {
	return &NullableAdminSegmentSchema{value: val, isSet: true}
}

func (v NullableAdminSegmentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminSegmentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
