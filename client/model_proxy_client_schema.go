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

// checks if the ProxyClientSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxyClientSchema{}

// ProxyClientSchema Frontend SDK client registration information
type ProxyClientSchema struct {
	// Name of the application using Unleash
	AppName string `json:"appName"`
	// Instance id for this application (typically hostname, podId or similar)
	InstanceId *string `json:"instanceId,omitempty"`
	// Optional field that describes the sdk version (name:version)
	SdkVersion *string `json:"sdkVersion,omitempty"`
	// deprecated
	// Deprecated
	Environment *string `json:"environment,omitempty"`
	// At which interval, in milliseconds, will this client be expected to send metrics
	Interval float32                  `json:"interval"`
	Started  ProxyClientSchemaStarted `json:"started"`
	// List of strategies implemented by this application
	Strategies []string `json:"strategies"`
}

// NewProxyClientSchema instantiates a new ProxyClientSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyClientSchema(appName string, interval float32, started ProxyClientSchemaStarted, strategies []string) *ProxyClientSchema {
	this := ProxyClientSchema{}
	this.AppName = appName
	this.Interval = interval
	this.Started = started
	this.Strategies = strategies
	return &this
}

// NewProxyClientSchemaWithDefaults instantiates a new ProxyClientSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyClientSchemaWithDefaults() *ProxyClientSchema {
	this := ProxyClientSchema{}
	return &this
}

// GetAppName returns the AppName field value
func (o *ProxyClientSchema) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *ProxyClientSchema) SetAppName(v string) {
	o.AppName = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ProxyClientSchema) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ProxyClientSchema) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ProxyClientSchema) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetSdkVersion returns the SdkVersion field value if set, zero value otherwise.
func (o *ProxyClientSchema) GetSdkVersion() string {
	if o == nil || IsNil(o.SdkVersion) {
		var ret string
		return ret
	}
	return *o.SdkVersion
}

// GetSdkVersionOk returns a tuple with the SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetSdkVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SdkVersion) {
		return nil, false
	}
	return o.SdkVersion, true
}

// HasSdkVersion returns a boolean if a field has been set.
func (o *ProxyClientSchema) HasSdkVersion() bool {
	if o != nil && !IsNil(o.SdkVersion) {
		return true
	}

	return false
}

// SetSdkVersion gets a reference to the given string and assigns it to the SdkVersion field.
func (o *ProxyClientSchema) SetSdkVersion(v string) {
	o.SdkVersion = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
// Deprecated
func (o *ProxyClientSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ProxyClientSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProxyClientSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
// Deprecated
func (o *ProxyClientSchema) SetEnvironment(v string) {
	o.Environment = &v
}

// GetInterval returns the Interval field value
func (o *ProxyClientSchema) GetInterval() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetIntervalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ProxyClientSchema) SetInterval(v float32) {
	o.Interval = v
}

// GetStarted returns the Started field value
func (o *ProxyClientSchema) GetStarted() ProxyClientSchemaStarted {
	if o == nil {
		var ret ProxyClientSchemaStarted
		return ret
	}

	return o.Started
}

// GetStartedOk returns a tuple with the Started field value
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetStartedOk() (*ProxyClientSchemaStarted, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Started, true
}

// SetStarted sets field value
func (o *ProxyClientSchema) SetStarted(v ProxyClientSchemaStarted) {
	o.Started = v
}

// GetStrategies returns the Strategies field value
func (o *ProxyClientSchema) GetStrategies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *ProxyClientSchema) GetStrategiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategies, true
}

// SetStrategies sets field value
func (o *ProxyClientSchema) SetStrategies(v []string) {
	o.Strategies = v
}

func (o ProxyClientSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxyClientSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appName"] = o.AppName
	if !IsNil(o.InstanceId) {
		toSerialize["instanceId"] = o.InstanceId
	}
	if !IsNil(o.SdkVersion) {
		toSerialize["sdkVersion"] = o.SdkVersion
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["interval"] = o.Interval
	toSerialize["started"] = o.Started
	toSerialize["strategies"] = o.Strategies
	return toSerialize, nil
}

type NullableProxyClientSchema struct {
	value *ProxyClientSchema
	isSet bool
}

func (v NullableProxyClientSchema) Get() *ProxyClientSchema {
	return v.value
}

func (v *NullableProxyClientSchema) Set(val *ProxyClientSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyClientSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyClientSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyClientSchema(val *ProxyClientSchema) *NullableProxyClientSchema {
	return &NullableProxyClientSchema{value: val, isSet: true}
}

func (v NullableProxyClientSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyClientSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
