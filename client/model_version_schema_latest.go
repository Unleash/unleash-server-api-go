/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.6.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VersionSchemaLatest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionSchemaLatest{}

// VersionSchemaLatest Information about the latest available Unleash releases. Will be an empty object if no data is available.
type VersionSchemaLatest struct {
	// The latest available OSS version of Unleash
	Oss *string `json:"oss,omitempty"`
	// The latest available Enterprise version of Unleash
	Enterprise *string `json:"enterprise,omitempty"`
}

// NewVersionSchemaLatest instantiates a new VersionSchemaLatest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionSchemaLatest() *VersionSchemaLatest {
	this := VersionSchemaLatest{}
	return &this
}

// NewVersionSchemaLatestWithDefaults instantiates a new VersionSchemaLatest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionSchemaLatestWithDefaults() *VersionSchemaLatest {
	this := VersionSchemaLatest{}
	return &this
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *VersionSchemaLatest) GetOss() string {
	if o == nil || IsNil(o.Oss) {
		var ret string
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionSchemaLatest) GetOssOk() (*string, bool) {
	if o == nil || IsNil(o.Oss) {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *VersionSchemaLatest) HasOss() bool {
	if o != nil && !IsNil(o.Oss) {
		return true
	}

	return false
}

// SetOss gets a reference to the given string and assigns it to the Oss field.
func (o *VersionSchemaLatest) SetOss(v string) {
	o.Oss = &v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise.
func (o *VersionSchemaLatest) GetEnterprise() string {
	if o == nil || IsNil(o.Enterprise) {
		var ret string
		return ret
	}
	return *o.Enterprise
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionSchemaLatest) GetEnterpriseOk() (*string, bool) {
	if o == nil || IsNil(o.Enterprise) {
		return nil, false
	}
	return o.Enterprise, true
}

// HasEnterprise returns a boolean if a field has been set.
func (o *VersionSchemaLatest) HasEnterprise() bool {
	if o != nil && !IsNil(o.Enterprise) {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given string and assigns it to the Enterprise field.
func (o *VersionSchemaLatest) SetEnterprise(v string) {
	o.Enterprise = &v
}

func (o VersionSchemaLatest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionSchemaLatest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oss) {
		toSerialize["oss"] = o.Oss
	}
	if !IsNil(o.Enterprise) {
		toSerialize["enterprise"] = o.Enterprise
	}
	return toSerialize, nil
}

type NullableVersionSchemaLatest struct {
	value *VersionSchemaLatest
	isSet bool
}

func (v NullableVersionSchemaLatest) Get() *VersionSchemaLatest {
	return v.value
}

func (v *NullableVersionSchemaLatest) Set(val *VersionSchemaLatest) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionSchemaLatest) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionSchemaLatest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionSchemaLatest(val *VersionSchemaLatest) *NullableVersionSchemaLatest {
	return &NullableVersionSchemaLatest{value: val, isSet: true}
}

func (v NullableVersionSchemaLatest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionSchemaLatest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
