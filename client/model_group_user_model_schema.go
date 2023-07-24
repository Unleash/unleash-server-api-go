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

// checks if the GroupUserModelSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUserModelSchema{}

// GroupUserModelSchema struct for GroupUserModelSchema
type GroupUserModelSchema struct {
	JoinedAt  *time.Time     `json:"joinedAt,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	User      UserSchema     `json:"user"`
}

// NewGroupUserModelSchema instantiates a new GroupUserModelSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUserModelSchema(user UserSchema) *GroupUserModelSchema {
	this := GroupUserModelSchema{}
	this.User = user
	return &this
}

// NewGroupUserModelSchemaWithDefaults instantiates a new GroupUserModelSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserModelSchemaWithDefaults() *GroupUserModelSchema {
	this := GroupUserModelSchema{}
	return &this
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise.
func (o *GroupUserModelSchema) GetJoinedAt() time.Time {
	if o == nil || IsNil(o.JoinedAt) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUserModelSchema) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.JoinedAt) {
		return nil, false
	}
	return o.JoinedAt, true
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *GroupUserModelSchema) HasJoinedAt() bool {
	if o != nil && !IsNil(o.JoinedAt) {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given time.Time and assigns it to the JoinedAt field.
func (o *GroupUserModelSchema) SetJoinedAt(v time.Time) {
	o.JoinedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUserModelSchema) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUserModelSchema) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GroupUserModelSchema) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *GroupUserModelSchema) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *GroupUserModelSchema) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *GroupUserModelSchema) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUser returns the User field value
func (o *GroupUserModelSchema) GetUser() UserSchema {
	if o == nil {
		var ret UserSchema
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GroupUserModelSchema) GetUserOk() (*UserSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GroupUserModelSchema) SetUser(v UserSchema) {
	o.User = v
}

func (o GroupUserModelSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUserModelSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JoinedAt) {
		toSerialize["joinedAt"] = o.JoinedAt
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableGroupUserModelSchema struct {
	value *GroupUserModelSchema
	isSet bool
}

func (v NullableGroupUserModelSchema) Get() *GroupUserModelSchema {
	return v.value
}

func (v *NullableGroupUserModelSchema) Set(val *GroupUserModelSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUserModelSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUserModelSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUserModelSchema(val *GroupUserModelSchema) *NullableGroupUserModelSchema {
	return &NullableGroupUserModelSchema{value: val, isSet: true}
}

func (v NullableGroupUserModelSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUserModelSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
