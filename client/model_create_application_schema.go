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

// checks if the CreateApplicationSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApplicationSchema{}

// CreateApplicationSchema Reported application information from Unleash SDKs
type CreateApplicationSchema struct {
	// Name of the application
	AppName *string `json:"appName,omitempty"`
	// Which SDK and version the application reporting uses. Typically represented as `<identifier>:<version>`
	SdkVersion *string `json:"sdkVersion,omitempty"`
	// Which [strategies](https://docs.getunleash.io/topics/the-anatomy-of-unleash#activation-strategies) the application has loaded. Useful when trying to figure out if your [custom strategy](https://docs.getunleash.io/reference/custom-activation-strategies) has been loaded in the SDK
	Strategies []string `json:"strategies,omitempty"`
	// A link to reference the application reporting the metrics. Could for instance be a GitHub link to the repository of the application
	Url *string `json:"url,omitempty"`
	// Css color to be used to color the application's entry in the application list
	Color *string `json:"color,omitempty"`
	// An URL to an icon file to be used for the applications's entry in the application list
	Icon                 *string `json:"icon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateApplicationSchema CreateApplicationSchema

// NewCreateApplicationSchema instantiates a new CreateApplicationSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApplicationSchema() *CreateApplicationSchema {
	this := CreateApplicationSchema{}
	return &this
}

// NewCreateApplicationSchemaWithDefaults instantiates a new CreateApplicationSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApplicationSchemaWithDefaults() *CreateApplicationSchema {
	this := CreateApplicationSchema{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *CreateApplicationSchema) SetAppName(v string) {
	o.AppName = &v
}

// GetSdkVersion returns the SdkVersion field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetSdkVersion() string {
	if o == nil || IsNil(o.SdkVersion) {
		var ret string
		return ret
	}
	return *o.SdkVersion
}

// GetSdkVersionOk returns a tuple with the SdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetSdkVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SdkVersion) {
		return nil, false
	}
	return o.SdkVersion, true
}

// HasSdkVersion returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasSdkVersion() bool {
	if o != nil && !IsNil(o.SdkVersion) {
		return true
	}

	return false
}

// SetSdkVersion gets a reference to the given string and assigns it to the SdkVersion field.
func (o *CreateApplicationSchema) SetSdkVersion(v string) {
	o.SdkVersion = &v
}

// GetStrategies returns the Strategies field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetStrategies() []string {
	if o == nil || IsNil(o.Strategies) {
		var ret []string
		return ret
	}
	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetStrategiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Strategies) {
		return nil, false
	}
	return o.Strategies, true
}

// HasStrategies returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasStrategies() bool {
	if o != nil && !IsNil(o.Strategies) {
		return true
	}

	return false
}

// SetStrategies gets a reference to the given []string and assigns it to the Strategies field.
func (o *CreateApplicationSchema) SetStrategies(v []string) {
	o.Strategies = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateApplicationSchema) SetUrl(v string) {
	o.Url = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CreateApplicationSchema) SetColor(v string) {
	o.Color = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CreateApplicationSchema) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApplicationSchema) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CreateApplicationSchema) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CreateApplicationSchema) SetIcon(v string) {
	o.Icon = &v
}

func (o CreateApplicationSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApplicationSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.SdkVersion) {
		toSerialize["sdkVersion"] = o.SdkVersion
	}
	if !IsNil(o.Strategies) {
		toSerialize["strategies"] = o.Strategies
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateApplicationSchema) UnmarshalJSON(bytes []byte) (err error) {
	varCreateApplicationSchema := _CreateApplicationSchema{}

	if err = json.Unmarshal(bytes, &varCreateApplicationSchema); err == nil {
		*o = CreateApplicationSchema(varCreateApplicationSchema)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "appName")
		delete(additionalProperties, "sdkVersion")
		delete(additionalProperties, "strategies")
		delete(additionalProperties, "url")
		delete(additionalProperties, "color")
		delete(additionalProperties, "icon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateApplicationSchema struct {
	value *CreateApplicationSchema
	isSet bool
}

func (v NullableCreateApplicationSchema) Get() *CreateApplicationSchema {
	return v.value
}

func (v *NullableCreateApplicationSchema) Set(val *CreateApplicationSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApplicationSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApplicationSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApplicationSchema(val *CreateApplicationSchema) *NullableCreateApplicationSchema {
	return &NullableCreateApplicationSchema{value: val, isSet: true}
}

func (v NullableCreateApplicationSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApplicationSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
