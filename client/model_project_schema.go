/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the ProjectSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectSchema{}

// ProjectSchema A definition of the project used for projects listing purposes
type ProjectSchema struct {
	// The id of this project
	Id string `json:"id"`
	// The name of this project
	Name string `json:"name"`
	// Additional information about the project
	Description NullableString `json:"description,omitempty"`
	// An indicator of the [project's health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100
	Health *float32 `json:"health,omitempty"`
	// The number of features this project has
	FeatureCount *float32 `json:"featureCount,omitempty"`
	// The number of members this project has
	MemberCount *float32     `json:"memberCount,omitempty"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   NullableTime `json:"updatedAt,omitempty"`
	// `true` if the project was favorited, otherwise `false`.
	Favorite *bool `json:"favorite,omitempty"`
	// The project's [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not.
	Mode *string `json:"mode,omitempty"`
	// A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy
	DefaultStickiness *string `json:"defaultStickiness,omitempty"`
}

// NewProjectSchema instantiates a new ProjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSchema(id string, name string) *ProjectSchema {
	this := ProjectSchema{}
	this.Id = id
	this.Name = name
	return &this
}

// NewProjectSchemaWithDefaults instantiates a new ProjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSchemaWithDefaults() *ProjectSchema {
	this := ProjectSchema{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectSchema) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ProjectSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectSchema) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSchema) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSchema) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectSchema) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ProjectSchema) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ProjectSchema) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ProjectSchema) UnsetDescription() {
	o.Description.Unset()
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ProjectSchema) GetHealth() float32 {
	if o == nil || IsNil(o.Health) {
		var ret float32
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetHealthOk() (*float32, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ProjectSchema) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given float32 and assigns it to the Health field.
func (o *ProjectSchema) SetHealth(v float32) {
	o.Health = &v
}

// GetFeatureCount returns the FeatureCount field value if set, zero value otherwise.
func (o *ProjectSchema) GetFeatureCount() float32 {
	if o == nil || IsNil(o.FeatureCount) {
		var ret float32
		return ret
	}
	return *o.FeatureCount
}

// GetFeatureCountOk returns a tuple with the FeatureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetFeatureCountOk() (*float32, bool) {
	if o == nil || IsNil(o.FeatureCount) {
		return nil, false
	}
	return o.FeatureCount, true
}

// HasFeatureCount returns a boolean if a field has been set.
func (o *ProjectSchema) HasFeatureCount() bool {
	if o != nil && !IsNil(o.FeatureCount) {
		return true
	}

	return false
}

// SetFeatureCount gets a reference to the given float32 and assigns it to the FeatureCount field.
func (o *ProjectSchema) SetFeatureCount(v float32) {
	o.FeatureCount = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *ProjectSchema) GetMemberCount() float32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret float32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetMemberCountOk() (*float32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *ProjectSchema) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given float32 and assigns it to the MemberCount field.
func (o *ProjectSchema) SetMemberCount(v float32) {
	o.MemberCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectSchema) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectSchema) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectSchema) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectSchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectSchema) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *ProjectSchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ProjectSchema) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ProjectSchema) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *ProjectSchema) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *ProjectSchema) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *ProjectSchema) SetFavorite(v bool) {
	o.Favorite = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ProjectSchema) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ProjectSchema) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ProjectSchema) SetMode(v string) {
	o.Mode = &v
}

// GetDefaultStickiness returns the DefaultStickiness field value if set, zero value otherwise.
func (o *ProjectSchema) GetDefaultStickiness() string {
	if o == nil || IsNil(o.DefaultStickiness) {
		var ret string
		return ret
	}
	return *o.DefaultStickiness
}

// GetDefaultStickinessOk returns a tuple with the DefaultStickiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSchema) GetDefaultStickinessOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultStickiness) {
		return nil, false
	}
	return o.DefaultStickiness, true
}

// HasDefaultStickiness returns a boolean if a field has been set.
func (o *ProjectSchema) HasDefaultStickiness() bool {
	if o != nil && !IsNil(o.DefaultStickiness) {
		return true
	}

	return false
}

// SetDefaultStickiness gets a reference to the given string and assigns it to the DefaultStickiness field.
func (o *ProjectSchema) SetDefaultStickiness(v string) {
	o.DefaultStickiness = &v
}

func (o ProjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.FeatureCount) {
		toSerialize["featureCount"] = o.FeatureCount
	}
	if !IsNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.DefaultStickiness) {
		toSerialize["defaultStickiness"] = o.DefaultStickiness
	}
	return toSerialize, nil
}

type NullableProjectSchema struct {
	value *ProjectSchema
	isSet bool
}

func (v NullableProjectSchema) Get() *ProjectSchema {
	return v.value
}

func (v *NullableProjectSchema) Set(val *ProjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSchema(val *ProjectSchema) *NullableProjectSchema {
	return &NullableProjectSchema{value: val, isSet: true}
}

func (v NullableProjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}