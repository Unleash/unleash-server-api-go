/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FeatureEnvironmentMetricsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEnvironmentMetricsSchema{}

// FeatureEnvironmentMetricsSchema How many times `feautreName` was evaluated to `true` (yes) and `false` (no) for `appName` in `environmnet`
type FeatureEnvironmentMetricsSchema struct {
	// The name of the feature
	FeatureName *string `json:"featureName,omitempty"`
	// The name of the application the SDK is being used in
	AppName *string `json:"appName,omitempty"`
	// Which environment the SDK is being used in
	Environment string     `json:"environment"`
	Timestamp   DateSchema `json:"timestamp"`
	// How many times the toggle evaluated to true
	Yes int32 `json:"yes"`
	// How many times the toggle evaluated to false
	No int32 `json:"no"`
	// How many times each variant was returned
	Variants *map[string]int32 `json:"variants,omitempty"`
}

// NewFeatureEnvironmentMetricsSchema instantiates a new FeatureEnvironmentMetricsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEnvironmentMetricsSchema(environment string, timestamp DateSchema, yes int32, no int32) *FeatureEnvironmentMetricsSchema {
	this := FeatureEnvironmentMetricsSchema{}
	this.Environment = environment
	this.Timestamp = timestamp
	this.Yes = yes
	this.No = no
	return &this
}

// NewFeatureEnvironmentMetricsSchemaWithDefaults instantiates a new FeatureEnvironmentMetricsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEnvironmentMetricsSchemaWithDefaults() *FeatureEnvironmentMetricsSchema {
	this := FeatureEnvironmentMetricsSchema{}
	return &this
}

// GetFeatureName returns the FeatureName field value if set, zero value otherwise.
func (o *FeatureEnvironmentMetricsSchema) GetFeatureName() string {
	if o == nil || IsNil(o.FeatureName) {
		var ret string
		return ret
	}
	return *o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetFeatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureName) {
		return nil, false
	}
	return o.FeatureName, true
}

// HasFeatureName returns a boolean if a field has been set.
func (o *FeatureEnvironmentMetricsSchema) HasFeatureName() bool {
	if o != nil && !IsNil(o.FeatureName) {
		return true
	}

	return false
}

// SetFeatureName gets a reference to the given string and assigns it to the FeatureName field.
func (o *FeatureEnvironmentMetricsSchema) SetFeatureName(v string) {
	o.FeatureName = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *FeatureEnvironmentMetricsSchema) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *FeatureEnvironmentMetricsSchema) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *FeatureEnvironmentMetricsSchema) SetAppName(v string) {
	o.AppName = &v
}

// GetEnvironment returns the Environment field value
func (o *FeatureEnvironmentMetricsSchema) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *FeatureEnvironmentMetricsSchema) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value
func (o *FeatureEnvironmentMetricsSchema) GetTimestamp() DateSchema {
	if o == nil {
		var ret DateSchema
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetTimestampOk() (*DateSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *FeatureEnvironmentMetricsSchema) SetTimestamp(v DateSchema) {
	o.Timestamp = v
}

// GetYes returns the Yes field value
func (o *FeatureEnvironmentMetricsSchema) GetYes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Yes
}

// GetYesOk returns a tuple with the Yes field value
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetYesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yes, true
}

// SetYes sets field value
func (o *FeatureEnvironmentMetricsSchema) SetYes(v int32) {
	o.Yes = v
}

// GetNo returns the No field value
func (o *FeatureEnvironmentMetricsSchema) GetNo() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.No
}

// GetNoOk returns a tuple with the No field value
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetNoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.No, true
}

// SetNo sets field value
func (o *FeatureEnvironmentMetricsSchema) SetNo(v int32) {
	o.No = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *FeatureEnvironmentMetricsSchema) GetVariants() map[string]int32 {
	if o == nil || IsNil(o.Variants) {
		var ret map[string]int32
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEnvironmentMetricsSchema) GetVariantsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *FeatureEnvironmentMetricsSchema) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]int32 and assigns it to the Variants field.
func (o *FeatureEnvironmentMetricsSchema) SetVariants(v map[string]int32) {
	o.Variants = &v
}

func (o FeatureEnvironmentMetricsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureEnvironmentMetricsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureName) {
		toSerialize["featureName"] = o.FeatureName
	}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	toSerialize["environment"] = o.Environment
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["yes"] = o.Yes
	toSerialize["no"] = o.No
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	return toSerialize, nil
}

type NullableFeatureEnvironmentMetricsSchema struct {
	value *FeatureEnvironmentMetricsSchema
	isSet bool
}

func (v NullableFeatureEnvironmentMetricsSchema) Get() *FeatureEnvironmentMetricsSchema {
	return v.value
}

func (v *NullableFeatureEnvironmentMetricsSchema) Set(val *FeatureEnvironmentMetricsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEnvironmentMetricsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEnvironmentMetricsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEnvironmentMetricsSchema(val *FeatureEnvironmentMetricsSchema) *NullableFeatureEnvironmentMetricsSchema {
	return &NullableFeatureEnvironmentMetricsSchema{value: val, isSet: true}
}

func (v NullableFeatureEnvironmentMetricsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureEnvironmentMetricsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}