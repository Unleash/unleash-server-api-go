/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the FeedbackResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackResponseSchema{}

// FeedbackResponseSchema User feedback information about a particular feedback item.
type FeedbackResponseSchema struct {
	// The ID of the user that gave the feedback.
	UserId *int32 `json:"userId,omitempty"`
	// `true` if the user has asked never to see this feedback questionnaire again.
	NeverShow *bool `json:"neverShow,omitempty"`
	// When this feedback was given
	Given NullableTime `json:"given,omitempty"`
	// The name of the feedback session
	FeedbackId *string `json:"feedbackId,omitempty"`
}

// NewFeedbackResponseSchema instantiates a new FeedbackResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackResponseSchema() *FeedbackResponseSchema {
	this := FeedbackResponseSchema{}
	return &this
}

// NewFeedbackResponseSchemaWithDefaults instantiates a new FeedbackResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackResponseSchemaWithDefaults() *FeedbackResponseSchema {
	this := FeedbackResponseSchema{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FeedbackResponseSchema) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackResponseSchema) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FeedbackResponseSchema) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *FeedbackResponseSchema) SetUserId(v int32) {
	o.UserId = &v
}

// GetNeverShow returns the NeverShow field value if set, zero value otherwise.
func (o *FeedbackResponseSchema) GetNeverShow() bool {
	if o == nil || IsNil(o.NeverShow) {
		var ret bool
		return ret
	}
	return *o.NeverShow
}

// GetNeverShowOk returns a tuple with the NeverShow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackResponseSchema) GetNeverShowOk() (*bool, bool) {
	if o == nil || IsNil(o.NeverShow) {
		return nil, false
	}
	return o.NeverShow, true
}

// HasNeverShow returns a boolean if a field has been set.
func (o *FeedbackResponseSchema) HasNeverShow() bool {
	if o != nil && !IsNil(o.NeverShow) {
		return true
	}

	return false
}

// SetNeverShow gets a reference to the given bool and assigns it to the NeverShow field.
func (o *FeedbackResponseSchema) SetNeverShow(v bool) {
	o.NeverShow = &v
}

// GetGiven returns the Given field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackResponseSchema) GetGiven() time.Time {
	if o == nil || IsNil(o.Given.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Given.Get()
}

// GetGivenOk returns a tuple with the Given field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackResponseSchema) GetGivenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Given.Get(), o.Given.IsSet()
}

// HasGiven returns a boolean if a field has been set.
func (o *FeedbackResponseSchema) HasGiven() bool {
	if o != nil && o.Given.IsSet() {
		return true
	}

	return false
}

// SetGiven gets a reference to the given NullableTime and assigns it to the Given field.
func (o *FeedbackResponseSchema) SetGiven(v time.Time) {
	o.Given.Set(&v)
}

// SetGivenNil sets the value for Given to be an explicit nil
func (o *FeedbackResponseSchema) SetGivenNil() {
	o.Given.Set(nil)
}

// UnsetGiven ensures that no value is present for Given, not even an explicit nil
func (o *FeedbackResponseSchema) UnsetGiven() {
	o.Given.Unset()
}

// GetFeedbackId returns the FeedbackId field value if set, zero value otherwise.
func (o *FeedbackResponseSchema) GetFeedbackId() string {
	if o == nil || IsNil(o.FeedbackId) {
		var ret string
		return ret
	}
	return *o.FeedbackId
}

// GetFeedbackIdOk returns a tuple with the FeedbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackResponseSchema) GetFeedbackIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeedbackId) {
		return nil, false
	}
	return o.FeedbackId, true
}

// HasFeedbackId returns a boolean if a field has been set.
func (o *FeedbackResponseSchema) HasFeedbackId() bool {
	if o != nil && !IsNil(o.FeedbackId) {
		return true
	}

	return false
}

// SetFeedbackId gets a reference to the given string and assigns it to the FeedbackId field.
func (o *FeedbackResponseSchema) SetFeedbackId(v string) {
	o.FeedbackId = &v
}

func (o FeedbackResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackResponseSchema) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FeedbackId) {
		toSerialize["feedbackId"] = o.FeedbackId
	}
	return toSerialize, nil
}

type NullableFeedbackResponseSchema struct {
	value *FeedbackResponseSchema
	isSet bool
}

func (v NullableFeedbackResponseSchema) Get() *FeedbackResponseSchema {
	return v.value
}

func (v *NullableFeedbackResponseSchema) Set(val *FeedbackResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackResponseSchema(val *FeedbackResponseSchema) *NullableFeedbackResponseSchema {
	return &NullableFeedbackResponseSchema{value: val, isSet: true}
}

func (v NullableFeedbackResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
