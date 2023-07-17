/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the EventSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSchema{}

// EventSchema An event describing something happening in the system
type EventSchema struct {
	// The ID of the event. An increasing natural number.
	Id int32 `json:"id"`
	// The time the event happened as a RFC 3339-conformant timestamp.
	CreatedAt time.Time `json:"createdAt"`
	// What [type](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#event-type-description) of event this is
	Type string `json:"type"`
	// Which user created this event
	CreatedBy string `json:"createdBy"`
	// The feature toggle environment the event relates to, if applicable.
	Environment NullableString `json:"environment,omitempty"`
	// The project the event relates to, if applicable.
	Project NullableString `json:"project,omitempty"`
	// The name of the feature toggle the event relates to, if applicable.
	FeatureName NullableString             `json:"featureName,omitempty"`
	Data        *EventSchemaData           `json:"data,omitempty"`
	PreData     NullableEventSchemaPreData `json:"preData,omitempty"`
	// Any tags related to the event, if applicable.
	Tags []TagSchema `json:"tags,omitempty"`
}

// NewEventSchema instantiates a new EventSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSchema(id int32, createdAt time.Time, type_ string, createdBy string) *EventSchema {
	this := EventSchema{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Type = type_
	this.CreatedBy = createdBy
	return &this
}

// NewEventSchemaWithDefaults instantiates a new EventSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSchemaWithDefaults() *EventSchema {
	this := EventSchema{}
	return &this
}

// GetId returns the Id field value
func (o *EventSchema) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventSchema) SetId(v int32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EventSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EventSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetType returns the Type field value
func (o *EventSchema) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventSchema) SetType(v string) {
	o.Type = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *EventSchema) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *EventSchema) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *EventSchema) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EventSchema) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *EventSchema) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *EventSchema) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *EventSchema) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSchema) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSchema) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *EventSchema) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *EventSchema) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *EventSchema) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *EventSchema) UnsetProject() {
	o.Project.Unset()
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSchema) GetFeatureName() string {
	if o == nil || IsNil(o.FeatureName.Get()) {
		var ret string
		return ret
	}
	return *o.FeatureName.Get()
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSchema) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureName.Get(), o.FeatureName.IsSet()
}

// HasFeatureName returns a boolean if a field has been set.
func (o *EventSchema) HasFeatureName() bool {
	if o != nil && o.FeatureName.IsSet() {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given NullableString and assigns it to the FeatureName field.
func (o *EventSchema) SetFeatureName(v string) {
	o.FeatureName.Set(&v)
}

// SetFeatureNameNil sets the value for FeatureName to be an explicit nil
func (o *EventSchema) SetFeatureNameNil() {
	o.FeatureName.Set(nil)
}

// UnsetFeatureName ensures that no value is present for FeatureName, not even an explicit nil
func (o *EventSchema) UnsetFeatureName() {
	o.FeatureName.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EventSchema) GetData() EventSchemaData {
	if o == nil || IsNil(o.Data) {
		var ret EventSchemaData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSchema) GetDataOk() (*EventSchemaData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventSchema) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EventSchemaData and assigns it to the Data field.
func (o *EventSchema) SetData(v EventSchemaData) {
	o.Data = &v
}

// GetPreData returns the PreData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSchema) GetPreData() EventSchemaPreData {
	if o == nil || IsNil(o.PreData.Get()) {
		var ret EventSchemaPreData
		return ret
	}
	return *o.PreData.Get()
}

// GetPreDataOk returns a tuple with the PreData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSchema) GetPreDataOk() (*EventSchemaPreData, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreData.Get(), o.PreData.IsSet()
}

// HasPreData returns a boolean if a field has been set.
func (o *EventSchema) HasPreData() bool {
	if o != nil && o.PreData.IsSet() {
		return true
	}

	return false
}

// SetPreData gets a reference to the given NullableEventSchemaPreData and assigns it to the PreData field.
func (o *EventSchema) SetPreData(v EventSchemaPreData) {
	o.PreData.Set(&v)
}

// SetPreDataNil sets the value for PreData to be an explicit nil
func (o *EventSchema) SetPreDataNil() {
	o.PreData.Set(nil)
}

// UnsetPreData ensures that no value is present for PreData, not even an explicit nil
func (o *EventSchema) UnsetPreData() {
	o.PreData.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSchema) GetTags() []TagSchema {
	if o == nil {
		var ret []TagSchema
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSchema) GetTagsOk() ([]TagSchema, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventSchema) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagSchema and assigns it to the Tags field.
func (o *EventSchema) SetTags(v []TagSchema) {
	o.Tags = v
}

func (o EventSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["type"] = o.Type
	toSerialize["createdBy"] = o.CreatedBy
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.FeatureName.IsSet() {
		toSerialize["featureName"] = o.FeatureName.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.PreData.IsSet() {
		toSerialize["preData"] = o.PreData.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableEventSchema struct {
	value *EventSchema
	isSet bool
}

func (v NullableEventSchema) Get() *EventSchema {
	return v.value
}

func (v *NullableEventSchema) Set(val *EventSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSchema(val *EventSchema) *NullableEventSchema {
	return &NullableEventSchema{value: val, isSet: true}
}

func (v NullableEventSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
