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

// checks if the ClientMetricsEnvSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientMetricsEnvSchema{}

// ClientMetricsEnvSchema Used for reporting feature evaluation results from SDKs
type ClientMetricsEnvSchema struct {
	// Name of the feature checked by the SDK
	FeatureName string `json:"featureName"`
	// The name of the application the SDK is being used in
	AppName string `json:"appName"`
	// Which environment the SDK is being used in
	Environment string      `json:"environment"`
	Timestamp   *DateSchema `json:"timestamp,omitempty"`
	// How many times the toggle evaluated to true
	Yes *int32 `json:"yes,omitempty"`
	// How many times the toggle evaluated to false
	No *int32 `json:"no,omitempty"`
	// How many times each variant was returned
	Variants             *map[string]int32 `json:"variants,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientMetricsEnvSchema ClientMetricsEnvSchema

// NewClientMetricsEnvSchema instantiates a new ClientMetricsEnvSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientMetricsEnvSchema(featureName string, appName string, environment string) *ClientMetricsEnvSchema {
	this := ClientMetricsEnvSchema{}
	this.FeatureName = featureName
	this.AppName = appName
	this.Environment = environment
	return &this
}

// NewClientMetricsEnvSchemaWithDefaults instantiates a new ClientMetricsEnvSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientMetricsEnvSchemaWithDefaults() *ClientMetricsEnvSchema {
	this := ClientMetricsEnvSchema{}
	return &this
}

// GetFeatureName returns the FeatureName field value
func (o *ClientMetricsEnvSchema) GetFeatureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeatureName
}

// GetFeatureNameOk returns a tuple with the FeatureName field value
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetFeatureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeatureName, true
}

// SetFeatureName sets field value
func (o *ClientMetricsEnvSchema) SetFeatureName(v string) {
	o.FeatureName = v
}

// GetAppName returns the AppName field value
func (o *ClientMetricsEnvSchema) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *ClientMetricsEnvSchema) SetAppName(v string) {
	o.AppName = v
}

// GetEnvironment returns the Environment field value
func (o *ClientMetricsEnvSchema) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ClientMetricsEnvSchema) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ClientMetricsEnvSchema) GetTimestamp() DateSchema {
	if o == nil || IsNil(o.Timestamp) {
		var ret DateSchema
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetTimestampOk() (*DateSchema, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ClientMetricsEnvSchema) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given DateSchema and assigns it to the Timestamp field.
func (o *ClientMetricsEnvSchema) SetTimestamp(v DateSchema) {
	o.Timestamp = &v
}

// GetYes returns the Yes field value if set, zero value otherwise.
func (o *ClientMetricsEnvSchema) GetYes() int32 {
	if o == nil || IsNil(o.Yes) {
		var ret int32
		return ret
	}
	return *o.Yes
}

// GetYesOk returns a tuple with the Yes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetYesOk() (*int32, bool) {
	if o == nil || IsNil(o.Yes) {
		return nil, false
	}
	return o.Yes, true
}

// HasYes returns a boolean if a field has been set.
func (o *ClientMetricsEnvSchema) HasYes() bool {
	if o != nil && !IsNil(o.Yes) {
		return true
	}

	return false
}

// SetYes gets a reference to the given int32 and assigns it to the Yes field.
func (o *ClientMetricsEnvSchema) SetYes(v int32) {
	o.Yes = &v
}

// GetNo returns the No field value if set, zero value otherwise.
func (o *ClientMetricsEnvSchema) GetNo() int32 {
	if o == nil || IsNil(o.No) {
		var ret int32
		return ret
	}
	return *o.No
}

// GetNoOk returns a tuple with the No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetNoOk() (*int32, bool) {
	if o == nil || IsNil(o.No) {
		return nil, false
	}
	return o.No, true
}

// HasNo returns a boolean if a field has been set.
func (o *ClientMetricsEnvSchema) HasNo() bool {
	if o != nil && !IsNil(o.No) {
		return true
	}

	return false
}

// SetNo gets a reference to the given int32 and assigns it to the No field.
func (o *ClientMetricsEnvSchema) SetNo(v int32) {
	o.No = &v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ClientMetricsEnvSchema) GetVariants() map[string]int32 {
	if o == nil || IsNil(o.Variants) {
		var ret map[string]int32
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMetricsEnvSchema) GetVariantsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ClientMetricsEnvSchema) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given map[string]int32 and assigns it to the Variants field.
func (o *ClientMetricsEnvSchema) SetVariants(v map[string]int32) {
	o.Variants = &v
}

func (o ClientMetricsEnvSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientMetricsEnvSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["featureName"] = o.FeatureName
	toSerialize["appName"] = o.AppName
	toSerialize["environment"] = o.Environment
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Yes) {
		toSerialize["yes"] = o.Yes
	}
	if !IsNil(o.No) {
		toSerialize["no"] = o.No
	}
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClientMetricsEnvSchema) UnmarshalJSON(bytes []byte) (err error) {
	varClientMetricsEnvSchema := _ClientMetricsEnvSchema{}

	if err = json.Unmarshal(bytes, &varClientMetricsEnvSchema); err == nil {
		*o = ClientMetricsEnvSchema(varClientMetricsEnvSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "featureName")
		delete(additionalProperties, "appName")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "yes")
		delete(additionalProperties, "no")
		delete(additionalProperties, "variants")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientMetricsEnvSchema struct {
	value *ClientMetricsEnvSchema
	isSet bool
}

func (v NullableClientMetricsEnvSchema) Get() *ClientMetricsEnvSchema {
	return v.value
}

func (v *NullableClientMetricsEnvSchema) Set(val *ClientMetricsEnvSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableClientMetricsEnvSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableClientMetricsEnvSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientMetricsEnvSchema(val *ClientMetricsEnvSchema) *NullableClientMetricsEnvSchema {
	return &NullableClientMetricsEnvSchema{value: val, isSet: true}
}

func (v NullableClientMetricsEnvSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientMetricsEnvSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
