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

// checks if the ClientApplicationSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientApplicationSchema{}

// ClientApplicationSchema A client application is an instance of one of our SDKs
type ClientApplicationSchema struct {
	// An identifier for the app that uses the sdk, should be static across SDK restarts
	AppName string `json:"appName"`
	// A unique identifier identifying the instance of the application running the SDK. Often changes based on execution environment. For instance: two pods in Kubernetes will have two different instanceIds
	InstanceId *string `json:"instanceId,omitempty"`
	// An SDK version identifier. Usually formatted as \"unleash-client-<language>:<version>\"
	SdkVersion *string `json:"sdkVersion,omitempty"`
	// The SDK's configured 'environment' property. Deprecated. This property  does **not** control which Unleash environment the SDK gets toggles for. To control Unleash environments, use the SDKs API key.
	// Deprecated
	Environment *string `json:"environment,omitempty"`
	// How often (in seconds) does the client refresh its toggles
	Interval float32 `json:"interval"`
	Started ClientApplicationSchemaStarted `json:"started"`
	// Which strategies the SDKs runtime knows about
	Strategies []string `json:"strategies"`
}

// NewClientApplicationSchema instantiates a new ClientApplicationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientApplicationSchema(appName string, interval float32, started ClientApplicationSchemaStarted, strategies []string) *ClientApplicationSchema {
	this := ClientApplicationSchema{}
	this.AppName = appName
	this.Interval = interval
	this.Started = started
	this.Strategies = strategies
	return &this
}

// NewClientApplicationSchemaWithDefaults instantiates a new ClientApplicationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientApplicationSchemaWithDefaults() *ClientApplicationSchema {
	this := ClientApplicationSchema{}
	return &this
}

// GetAppName returns the AppName field value
func (o *ClientApplicationSchema) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
func (o *ClientApplicationSchema) SetAppName(v string) {
	o.AppName = v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *ClientApplicationSchema) GetInstanceId() string {
	if o == nil || IsNil(o.InstanceId) {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceId) {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *ClientApplicationSchema) HasInstanceId() bool {
	if o != nil && !IsNil(o.InstanceId) {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *ClientApplicationSchema) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetSdkVersion returns the SdkVersion field value if set, zero value otherwise.
func (o *ClientApplicationSchema) GetSdkVersion() string {
	if o == nil || IsNil(o.SdkVersion) {
		var ret string
		return ret
	}
	return *o.SdkVersion
}

// GetSdkVersionOk returns a tuple with the SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetSdkVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SdkVersion) {
		return nil, false
	}
	return o.SdkVersion, true
}

// HasSdkVersion returns a boolean if a field has been set.
func (o *ClientApplicationSchema) HasSdkVersion() bool {
	if o != nil && !IsNil(o.SdkVersion) {
		return true
	}

	return false
}

// SetSdkVersion gets a reference to the given string and assigns it to the SdkVersion field.
func (o *ClientApplicationSchema) SetSdkVersion(v string) {
	o.SdkVersion = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
// Deprecated
func (o *ClientApplicationSchema) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ClientApplicationSchema) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ClientApplicationSchema) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
// Deprecated
func (o *ClientApplicationSchema) SetEnvironment(v string) {
	o.Environment = &v
}

// GetInterval returns the Interval field value
func (o *ClientApplicationSchema) GetInterval() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetIntervalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ClientApplicationSchema) SetInterval(v float32) {
	o.Interval = v
}

// GetStarted returns the Started field value
func (o *ClientApplicationSchema) GetStarted() ClientApplicationSchemaStarted {
	if o == nil {
		var ret ClientApplicationSchemaStarted
		return ret
	}

	return o.Started
}

// GetStartedOk returns a tuple with the Started field value
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetStartedOk() (*ClientApplicationSchemaStarted, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Started, true
}

// SetStarted sets field value
func (o *ClientApplicationSchema) SetStarted(v ClientApplicationSchemaStarted) {
	o.Started = v
}

// GetStrategies returns the Strategies field value
func (o *ClientApplicationSchema) GetStrategies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *ClientApplicationSchema) GetStrategiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategies, true
}

// SetStrategies sets field value
func (o *ClientApplicationSchema) SetStrategies(v []string) {
	o.Strategies = v
}

func (o ClientApplicationSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientApplicationSchema) ToMap() (map[string]interface{}, error) {
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

type NullableClientApplicationSchema struct {
	value *ClientApplicationSchema
	isSet bool
}

func (v NullableClientApplicationSchema) Get() *ClientApplicationSchema {
	return v.value
}

func (v *NullableClientApplicationSchema) Set(val *ClientApplicationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableClientApplicationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableClientApplicationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientApplicationSchema(val *ClientApplicationSchema) *NullableClientApplicationSchema {
	return &NullableClientApplicationSchema{value: val, isSet: true}
}

func (v NullableClientApplicationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientApplicationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


