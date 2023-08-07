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

// checks if the SearchEventsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEventsSchema{}

// SearchEventsSchema          Search for events by type, project, feature, free-text query,         or a combination thereof. Pass an empty object to fetch all events.
type SearchEventsSchema struct {
	// Find events by event type (case-sensitive).
	Type *string `json:"type,omitempty"`
	// Find events by project ID (case-sensitive).
	Project *string `json:"project,omitempty"`
	// Find events by feature toggle name (case-sensitive).
	Feature *string `json:"feature,omitempty"`
	//                  Find events by a free-text search query.                 The query will be matched against the event type,                 the username or email that created the event (if any),                 and the event data payload (if any).
	Query *string `json:"query,omitempty"`
	// The maximum amount of events to return in the search result
	Limit *int32 `json:"limit,omitempty"`
	// Which event id to start listing from
	Offset *int32 `json:"offset,omitempty"`
}

// NewSearchEventsSchema instantiates a new SearchEventsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEventsSchema() *SearchEventsSchema {
	this := SearchEventsSchema{}
	var limit int32 = 100
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewSearchEventsSchemaWithDefaults instantiates a new SearchEventsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEventsSchemaWithDefaults() *SearchEventsSchema {
	this := SearchEventsSchema{}
	var limit int32 = 100
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchEventsSchema) SetType(v string) {
	o.Type = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *SearchEventsSchema) SetProject(v string) {
	o.Project = &v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *SearchEventsSchema) SetFeature(v string) {
	o.Feature = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SearchEventsSchema) SetQuery(v string) {
	o.Query = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *SearchEventsSchema) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SearchEventsSchema) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEventsSchema) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SearchEventsSchema) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *SearchEventsSchema) SetOffset(v int32) {
	o.Offset = &v
}

func (o SearchEventsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEventsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableSearchEventsSchema struct {
	value *SearchEventsSchema
	isSet bool
}

func (v NullableSearchEventsSchema) Get() *SearchEventsSchema {
	return v.value
}

func (v *NullableSearchEventsSchema) Set(val *SearchEventsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEventsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEventsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEventsSchema(val *SearchEventsSchema) *NullableSearchEventsSchema {
	return &NullableSearchEventsSchema{value: val, isSet: true}
}

func (v NullableSearchEventsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEventsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
