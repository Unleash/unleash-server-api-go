/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ExportResultSchemaSegmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportResultSchemaSegmentsInner{}

// ExportResultSchemaSegmentsInner struct for ExportResultSchemaSegmentsInner
type ExportResultSchemaSegmentsInner struct {
	Id float32 `json:"id"`
	Name *string `json:"name,omitempty"`
}

// NewExportResultSchemaSegmentsInner instantiates a new ExportResultSchemaSegmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportResultSchemaSegmentsInner(id float32) *ExportResultSchemaSegmentsInner {
	this := ExportResultSchemaSegmentsInner{}
	this.Id = id
	return &this
}

// NewExportResultSchemaSegmentsInnerWithDefaults instantiates a new ExportResultSchemaSegmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportResultSchemaSegmentsInnerWithDefaults() *ExportResultSchemaSegmentsInner {
	this := ExportResultSchemaSegmentsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ExportResultSchemaSegmentsInner) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportResultSchemaSegmentsInner) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportResultSchemaSegmentsInner) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExportResultSchemaSegmentsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportResultSchemaSegmentsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExportResultSchemaSegmentsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExportResultSchemaSegmentsInner) SetName(v string) {
	o.Name = &v
}

func (o ExportResultSchemaSegmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportResultSchemaSegmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableExportResultSchemaSegmentsInner struct {
	value *ExportResultSchemaSegmentsInner
	isSet bool
}

func (v NullableExportResultSchemaSegmentsInner) Get() *ExportResultSchemaSegmentsInner {
	return v.value
}

func (v *NullableExportResultSchemaSegmentsInner) Set(val *ExportResultSchemaSegmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExportResultSchemaSegmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExportResultSchemaSegmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportResultSchemaSegmentsInner(val *ExportResultSchemaSegmentsInner) *NullableExportResultSchemaSegmentsInner {
	return &NullableExportResultSchemaSegmentsInner{value: val, isSet: true}
}

func (v NullableExportResultSchemaSegmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportResultSchemaSegmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


