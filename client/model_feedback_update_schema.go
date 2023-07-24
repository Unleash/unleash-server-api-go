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

// checks if the FeedbackUpdateSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackUpdateSchema{}

// FeedbackUpdateSchema User feedback information to be updated.
type FeedbackUpdateSchema struct {
	// The ID of the user that gave the feedback.
	UserId *int32 `json:"userId,omitempty"`
	// `true` if the user has asked never to see this feedback questionnaire again.
	NeverShow *bool `json:"neverShow,omitempty"`
	// When this feedback was given
	Given NullableTime `json:"given,omitempty"`
}

// NewFeedbackUpdateSchema instantiates a new FeedbackUpdateSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackUpdateSchema() *FeedbackUpdateSchema {
	this := FeedbackUpdateSchema{}
	return &this
}

// NewFeedbackUpdateSchemaWithDefaults instantiates a new FeedbackUpdateSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackUpdateSchemaWithDefaults() *FeedbackUpdateSchema {
	this := FeedbackUpdateSchema{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FeedbackUpdateSchema) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackUpdateSchema) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FeedbackUpdateSchema) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *FeedbackUpdateSchema) SetUserId(v int32) {
	o.UserId = &v
}

// GetNeverShow returns the NeverShow field value if set, zero value otherwise.
func (o *FeedbackUpdateSchema) GetNeverShow() bool {
	if o == nil || IsNil(o.NeverShow) {
		var ret bool
		return ret
	}
	return *o.NeverShow
}

// GetNeverShowOk returns a tuple with the NeverShow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackUpdateSchema) GetNeverShowOk() (*bool, bool) {
	if o == nil || IsNil(o.NeverShow) {
		return nil, false
	}
	return o.NeverShow, true
}

// HasNeverShow returns a boolean if a field has been set.
func (o *FeedbackUpdateSchema) HasNeverShow() bool {
	if o != nil && !IsNil(o.NeverShow) {
		return true
	}

	return false
}

// SetNeverShow gets a reference to the given bool and assigns it to the NeverShow field.
func (o *FeedbackUpdateSchema) SetNeverShow(v bool) {
	o.NeverShow = &v
}

// GetGiven returns the Given field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackUpdateSchema) GetGiven() time.Time {
	if o == nil || IsNil(o.Given.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Given.Get()
}

// GetGivenOk returns a tuple with the Given field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackUpdateSchema) GetGivenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Given.Get(), o.Given.IsSet()
}

// HasGiven returns a boolean if a field has been set.
func (o *FeedbackUpdateSchema) HasGiven() bool {
	if o != nil && o.Given.IsSet() {
		return true
	}

	return false
}

// SetGiven gets a reference to the given NullableTime and assigns it to the Given field.
func (o *FeedbackUpdateSchema) SetGiven(v time.Time) {
	o.Given.Set(&v)
}

// SetGivenNil sets the value for Given to be an explicit nil
func (o *FeedbackUpdateSchema) SetGivenNil() {
	o.Given.Set(nil)
}

// UnsetGiven ensures that no value is present for Given, not even an explicit nil
func (o *FeedbackUpdateSchema) UnsetGiven() {
	o.Given.Unset()
}

func (o FeedbackUpdateSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackUpdateSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.NeverShow) {
		toSerialize["neverShow"] = o.NeverShow
	}
	if o.Given.IsSet() {
		toSerialize["given"] = o.Given.Get()
	}
	return toSerialize, nil
}

type NullableFeedbackUpdateSchema struct {
	value *FeedbackUpdateSchema
	isSet bool
}

func (v NullableFeedbackUpdateSchema) Get() *FeedbackUpdateSchema {
	return v.value
}

func (v *NullableFeedbackUpdateSchema) Set(val *FeedbackUpdateSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackUpdateSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackUpdateSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackUpdateSchema(val *FeedbackUpdateSchema) *NullableFeedbackUpdateSchema {
	return &NullableFeedbackUpdateSchema{value: val, isSet: true}
}

func (v NullableFeedbackUpdateSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackUpdateSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}