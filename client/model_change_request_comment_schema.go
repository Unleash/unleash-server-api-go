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

// checks if the ChangeRequestCommentSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeRequestCommentSchema{}

// ChangeRequestCommentSchema A comment belonging to a [change request](https://docs.getunleash.io/reference/change-requests).
type ChangeRequestCommentSchema struct {
	// The comment's ID. Unique per change request.
	Id *float32 `json:"id,omitempty"`
	// The content of the comment.
	Text      string                              `json:"text"`
	CreatedBy ChangeRequestCommentSchemaCreatedBy `json:"createdBy"`
	// When the comment was made.
	CreatedAt time.Time `json:"createdAt"`
}

// NewChangeRequestCommentSchema instantiates a new ChangeRequestCommentSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeRequestCommentSchema(text string, createdBy ChangeRequestCommentSchemaCreatedBy, createdAt time.Time) *ChangeRequestCommentSchema {
	this := ChangeRequestCommentSchema{}
	this.Text = text
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	return &this
}

// NewChangeRequestCommentSchemaWithDefaults instantiates a new ChangeRequestCommentSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeRequestCommentSchemaWithDefaults() *ChangeRequestCommentSchema {
	this := ChangeRequestCommentSchema{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChangeRequestCommentSchema) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestCommentSchema) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChangeRequestCommentSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ChangeRequestCommentSchema) SetId(v float32) {
	o.Id = &v
}

// GetText returns the Text field value
func (o *ChangeRequestCommentSchema) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestCommentSchema) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *ChangeRequestCommentSchema) SetText(v string) {
	o.Text = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ChangeRequestCommentSchema) GetCreatedBy() ChangeRequestCommentSchemaCreatedBy {
	if o == nil {
		var ret ChangeRequestCommentSchemaCreatedBy
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestCommentSchema) GetCreatedByOk() (*ChangeRequestCommentSchemaCreatedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ChangeRequestCommentSchema) SetCreatedBy(v ChangeRequestCommentSchemaCreatedBy) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChangeRequestCommentSchema) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestCommentSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChangeRequestCommentSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o ChangeRequestCommentSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeRequestCommentSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["text"] = o.Text
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableChangeRequestCommentSchema struct {
	value *ChangeRequestCommentSchema
	isSet bool
}

func (v NullableChangeRequestCommentSchema) Get() *ChangeRequestCommentSchema {
	return v.value
}

func (v *NullableChangeRequestCommentSchema) Set(val *ChangeRequestCommentSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeRequestCommentSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeRequestCommentSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeRequestCommentSchema(val *ChangeRequestCommentSchema) *NullableChangeRequestCommentSchema {
	return &NullableChangeRequestCommentSchema{value: val, isSet: true}
}

func (v NullableChangeRequestCommentSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeRequestCommentSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
