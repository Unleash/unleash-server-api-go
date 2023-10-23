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

// checks if the VersionSchemaCurrent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionSchemaCurrent{}

// VersionSchemaCurrent The current version of Unleash.
type VersionSchemaCurrent struct {
	// The OSS version used when building this Unleash instance, represented as a git revision belonging to the [main Unleash git repo](https://github.com/Unleash/unleash/)
	Oss *string `json:"oss,omitempty"`
	// The Enterpris version of Unleash used to build this instance, represented as a git revision belonging to the [Unleash Enterprise](https://github.com/ivarconr/unleash-enterprise) repository. Will be an empty string if no enterprise version was used,
	Enterprise *string `json:"enterprise,omitempty"`
}

// NewVersionSchemaCurrent instantiates a new VersionSchemaCurrent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionSchemaCurrent() *VersionSchemaCurrent {
	this := VersionSchemaCurrent{}
	return &this
}

// NewVersionSchemaCurrentWithDefaults instantiates a new VersionSchemaCurrent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionSchemaCurrentWithDefaults() *VersionSchemaCurrent {
	this := VersionSchemaCurrent{}
	return &this
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *VersionSchemaCurrent) GetOss() string {
	if o == nil || IsNil(o.Oss) {
		var ret string
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionSchemaCurrent) GetOssOk() (*string, bool) {
	if o == nil || IsNil(o.Oss) {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *VersionSchemaCurrent) HasOss() bool {
	if o != nil && !IsNil(o.Oss) {
		return true
	}

	return false
}

// SetOss gets a reference to the given string and assigns it to the Oss field.
func (o *VersionSchemaCurrent) SetOss(v string) {
	o.Oss = &v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise.
func (o *VersionSchemaCurrent) GetEnterprise() string {
	if o == nil || IsNil(o.Enterprise) {
		var ret string
		return ret
	}
	return *o.Enterprise
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionSchemaCurrent) GetEnterpriseOk() (*string, bool) {
	if o == nil || IsNil(o.Enterprise) {
		return nil, false
	}
	return o.Enterprise, true
}

// HasEnterprise returns a boolean if a field has been set.
func (o *VersionSchemaCurrent) HasEnterprise() bool {
	if o != nil && !IsNil(o.Enterprise) {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given string and assigns it to the Enterprise field.
func (o *VersionSchemaCurrent) SetEnterprise(v string) {
	o.Enterprise = &v
}

func (o VersionSchemaCurrent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionSchemaCurrent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oss) {
		toSerialize["oss"] = o.Oss
	}
	if !IsNil(o.Enterprise) {
		toSerialize["enterprise"] = o.Enterprise
	}
	return toSerialize, nil
}

type NullableVersionSchemaCurrent struct {
	value *VersionSchemaCurrent
	isSet bool
}

func (v NullableVersionSchemaCurrent) Get() *VersionSchemaCurrent {
	return v.value
}

func (v *NullableVersionSchemaCurrent) Set(val *VersionSchemaCurrent) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionSchemaCurrent) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionSchemaCurrent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionSchemaCurrent(val *VersionSchemaCurrent) *NullableVersionSchemaCurrent {
	return &NullableVersionSchemaCurrent{value: val, isSet: true}
}

func (v NullableVersionSchemaCurrent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionSchemaCurrent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
