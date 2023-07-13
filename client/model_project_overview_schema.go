/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ProjectOverviewSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectOverviewSchema{}

// ProjectOverviewSchema A high-level overview of a project. It contains information such as project statistics, the name of the project, what members and what features it contains, etc.
type ProjectOverviewSchema struct {
	Stats *ProjectStatsSchema `json:"stats,omitempty"`
	Version float32 `json:"version"`
	// The name of this project
	Name string `json:"name"`
	// Additional information about the project
	Description NullableString `json:"description,omitempty"`
	// A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy
	DefaultStickiness *string `json:"defaultStickiness,omitempty"`
	// The project's [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not.
	Mode *string `json:"mode,omitempty"`
	// The number of members this project has
	Members *float32 `json:"members,omitempty"`
	// An indicator of the [project's health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100
	Health *float32 `json:"health,omitempty"`
	// The environments that are enabled for this project
	Environments []ProjectEnvironmentSchema `json:"environments,omitempty"`
	// The full list of features in this project (excluding archived features)
	Features []FeatureSchema `json:"features,omitempty"`
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
	CreatedAt NullableTime `json:"createdAt,omitempty"`
	// `true` if the project was favorited, otherwise `false`.
	Favorite *bool `json:"favorite,omitempty"`
}

// NewProjectOverviewSchema instantiates a new ProjectOverviewSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectOverviewSchema(version float32, name string) *ProjectOverviewSchema {
	this := ProjectOverviewSchema{}
	this.Version = version
	this.Name = name
	return &this
}

// NewProjectOverviewSchemaWithDefaults instantiates a new ProjectOverviewSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectOverviewSchemaWithDefaults() *ProjectOverviewSchema {
	this := ProjectOverviewSchema{}
	return &this
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetStats() ProjectStatsSchema {
	if o == nil || IsNil(o.Stats) {
		var ret ProjectStatsSchema
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetStatsOk() (*ProjectStatsSchema, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ProjectStatsSchema and assigns it to the Stats field.
func (o *ProjectOverviewSchema) SetStats(v ProjectStatsSchema) {
	o.Stats = &v
}

// GetVersion returns the Version field value
func (o *ProjectOverviewSchema) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ProjectOverviewSchema) SetVersion(v float32) {
	o.Version = v
}

// GetName returns the Name field value
func (o *ProjectOverviewSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectOverviewSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectOverviewSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectOverviewSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectOverviewSchema) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectOverviewSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectOverviewSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetDefaultStickiness returns the DefaultStickiness field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetDefaultStickiness() string {
	if o == nil || IsNil(o.DefaultStickiness) {
		var ret string
		return ret
	}
	return *o.DefaultStickiness
}

// GetDefaultStickinessOk returns a tuple with the DefaultStickiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetDefaultStickinessOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultStickiness) {
		return nil, false
	}
	return o.DefaultStickiness, true
}

// HasDefaultStickiness returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasDefaultStickiness() bool {
	if o != nil && !IsNil(o.DefaultStickiness) {
		return true
	}

	return false
}

// SetDefaultStickiness gets a reference to the given string and assigns it to the DefaultStickiness field.
func (o *ProjectOverviewSchema) SetDefaultStickiness(v string) {
	o.DefaultStickiness = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ProjectOverviewSchema) SetMode(v string) {
	o.Mode = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetMembers() float32 {
	if o == nil || IsNil(o.Members) {
		var ret float32
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetMembersOk() (*float32, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given float32 and assigns it to the Members field.
func (o *ProjectOverviewSchema) SetMembers(v float32) {
	o.Members = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetHealth() float32 {
	if o == nil || IsNil(o.Health) {
		var ret float32
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given float32 and assigns it to the Health field.
func (o *ProjectOverviewSchema) SetHealth(v float32) {
	o.Health = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetEnvironments() []ProjectEnvironmentSchema {
	if o == nil || IsNil(o.Environments) {
		var ret []ProjectEnvironmentSchema
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetEnvironmentsOk() ([]ProjectEnvironmentSchema, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []ProjectEnvironmentSchema and assigns it to the Environments field.
func (o *ProjectOverviewSchema) SetEnvironments(v []ProjectEnvironmentSchema) {
	o.Environments = v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetFeatures() []FeatureSchema {
	if o == nil || IsNil(o.Features) {
		var ret []FeatureSchema
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetFeaturesOk() ([]FeatureSchema, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FeatureSchema and assigns it to the Features field.
func (o *ProjectOverviewSchema) SetFeatures(v []FeatureSchema) {
	o.Features = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectOverviewSchema) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectOverviewSchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ProjectOverviewSchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ProjectOverviewSchema) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ProjectOverviewSchema) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectOverviewSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectOverviewSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *ProjectOverviewSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ProjectOverviewSchema) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ProjectOverviewSchema) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ProjectOverviewSchema) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectOverviewSchema) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ProjectOverviewSchema) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ProjectOverviewSchema) SetFavorite(v bool) {
	o.Favorite = &v
}

func (o ProjectOverviewSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectOverviewSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	toSerialize["version"] = o.Version
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.DefaultStickiness) {
		toSerialize["defaultStickiness"] = o.DefaultStickiness
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	return toSerialize, nil
}

type NullableProjectOverviewSchema struct {
	value *ProjectOverviewSchema
	isSet bool
}

func (v NullableProjectOverviewSchema) Get() *ProjectOverviewSchema {
	return v.value
}

func (v *NullableProjectOverviewSchema) Set(val *ProjectOverviewSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectOverviewSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectOverviewSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectOverviewSchema(val *ProjectOverviewSchema) *NullableProjectOverviewSchema {
	return &NullableProjectOverviewSchema{value: val, isSet: true}
}

func (v NullableProjectOverviewSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectOverviewSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


