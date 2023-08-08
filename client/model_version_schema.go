/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the VersionSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionSchema{}

// VersionSchema Detailed information about an Unleash version
type VersionSchema struct {
	Current VersionSchemaCurrent `json:"current"`
	Latest  VersionSchemaLatest  `json:"latest"`
	// Whether the Unleash server is running the latest release (`true`) or if there are updates available (`false`)
	IsLatest bool `json:"isLatest"`
	// The instance identifier of the Unleash instance
	InstanceId *string `json:"instanceId,omitempty"`
}

// NewVersionSchema instantiates a new VersionSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionSchema(current VersionSchemaCurrent, latest VersionSchemaLatest, isLatest bool) *VersionSchema {
	this := VersionSchema{}
	this.Current = current
	this.Latest = latest
	this.IsLatest = isLatest
	return &this
}

// NewVersionSchemaWithDefaults instantiates a new VersionSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionSchemaWithDefaults() *VersionSchema {
	this := VersionSchema{}
	return &this
}

// GetCurrent returns the Current field value
func (o *VersionSchema) GetCurrent() VersionSchemaCurrent {
	if o == nil {
		var ret VersionSchemaCurrent
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *VersionSchema) GetCurrentOk() (*VersionSchemaCurrent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *VersionSchema) SetCurrent(v VersionSchemaCurrent) {
	o.Current = v
}

// GetLatest returns the Latest field value
func (o *VersionSchema) GetLatest() VersionSchemaLatest {
	if o == nil {
		var ret VersionSchemaLatest
		return ret
	}

	return o.Latest
}

// GetLatestOk returns a tuple with the Latest field value
// and a boolean to check if the value has been set.
func (o *VersionSchema) GetLatestOk() (*VersionSchemaLatest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latest, true
}

// SetLatest sets field value
func (o *VersionSchema) SetLatest(v VersionSchemaLatest) {
	o.Latest = v
}

// GetIsLatest returns the IsLatest field value
func (o *VersionSchema) GetIsLatest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLatest
}

// GetIsLatestOk returns a tuple with the IsLatest field value
// and a boolean to check if the value has been set.
func (o *VersionSchema) GetIsLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLatest, true
}

// SetIsLatest sets field value
func (o *VersionSchema) SetIsLatest(v bool) {
	o.IsLatest = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *VersionSchema) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionSchema) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *VersionSchema) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *VersionSchema) SetInstanceId(v string) {
	o.InstanceId = &v
}

func (o VersionSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	toSerialize["latest"] = o.Latest
	toSerialize["isLatest"] = o.IsLatest
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	return toSerialize, nil
}

type NullableVersionSchema struct {
	value *VersionSchema
	isSet bool
}

func (v NullableVersionSchema) Get() *VersionSchema {
	return v.value
}

func (v *NullableVersionSchema) Set(val *VersionSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionSchema(val *VersionSchema) *NullableVersionSchema {
	return &NullableVersionSchema{value: val, isSet: true}
}

func (v NullableVersionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
