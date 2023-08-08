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

// checks if the GoogleSettingsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoogleSettingsSchema{}

// GoogleSettingsSchema Configuration for using Google Authentication
type GoogleSettingsSchema struct {
	// Is google OIDC enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The google client id, used to authenticate against google
	ClientId string `json:"clientId"`
	// The client secret used to authenticate the OAuth session used to log the user in
	ClientSecret string `json:"clientSecret"`
	// Name of the host allowed to access the Google authentication flow
	UnleashHostname string `json:"unleashHostname"`
	// Should Unleash create users based on the emails coming back in the authentication reply from Google
	AutoCreate *bool `json:"autoCreate,omitempty"`
	// A comma separated list of email domains that Unleash will auto create user accounts for.
	EmailDomains *string `json:"emailDomains,omitempty"`
}

// NewGoogleSettingsSchema instantiates a new GoogleSettingsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleSettingsSchema(clientId string, clientSecret string, unleashHostname string) *GoogleSettingsSchema {
	this := GoogleSettingsSchema{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.UnleashHostname = unleashHostname
	return &this
}

// NewGoogleSettingsSchemaWithDefaults instantiates a new GoogleSettingsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleSettingsSchemaWithDefaults() *GoogleSettingsSchema {
	this := GoogleSettingsSchema{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GoogleSettingsSchema) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GoogleSettingsSchema) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GoogleSettingsSchema) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetClientId returns the ClientId field value
func (o *GoogleSettingsSchema) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *GoogleSettingsSchema) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *GoogleSettingsSchema) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *GoogleSettingsSchema) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetUnleashHostname returns the UnleashHostname field value
func (o *GoogleSettingsSchema) GetUnleashHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnleashHostname
}

// GetUnleashHostnameOk returns a tuple with the UnleashHostname field value
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetUnleashHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnleashHostname, true
}

// SetUnleashHostname sets field value
func (o *GoogleSettingsSchema) SetUnleashHostname(v string) {
	o.UnleashHostname = v
}

// GetAutoCreate returns the AutoCreate field value if set, zero value otherwise.
func (o *GoogleSettingsSchema) GetAutoCreate() bool {
	if o == nil || IsNil(o.AutoCreate) {
		var ret bool
		return ret
	}
	return *o.AutoCreate
}

// GetAutoCreateOk returns a tuple with the AutoCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetAutoCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreate) {
		return nil, false
	}
	return o.AutoCreate, true
}

// HasAutoCreate returns a boolean if a field has been set.
func (o *GoogleSettingsSchema) HasAutoCreate() bool {
	if o != nil && !IsNil(o.AutoCreate) {
		return true
	}

	return false
}

// SetAutoCreate gets a reference to the given bool and assigns it to the AutoCreate field.
func (o *GoogleSettingsSchema) SetAutoCreate(v bool) {
	o.AutoCreate = &v
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *GoogleSettingsSchema) GetEmailDomains() string {
	if o == nil || IsNil(o.EmailDomains) {
		var ret string
		return ret
	}
	return *o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleSettingsSchema) GetEmailDomainsOk() (*string, bool) {
	if o == nil || IsNil(o.EmailDomains) {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *GoogleSettingsSchema) HasEmailDomains() bool {
	if o != nil && !IsNil(o.EmailDomains) {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given string and assigns it to the EmailDomains field.
func (o *GoogleSettingsSchema) SetEmailDomains(v string) {
	o.EmailDomains = &v
}

func (o GoogleSettingsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoogleSettingsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["unleashHostname"] = o.UnleashHostname
	if !IsNil(o.AutoCreate) {
		toSerialize["autoCreate"] = o.AutoCreate
	}
	if !IsNil(o.EmailDomains) {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	return toSerialize, nil
}

type NullableGoogleSettingsSchema struct {
	value *GoogleSettingsSchema
	isSet bool
}

func (v NullableGoogleSettingsSchema) Get() *GoogleSettingsSchema {
	return v.value
}

func (v *NullableGoogleSettingsSchema) Set(val *GoogleSettingsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleSettingsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleSettingsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleSettingsSchema(val *GoogleSettingsSchema) *NullableGoogleSettingsSchema {
	return &NullableGoogleSettingsSchema{value: val, isSet: true}
}

func (v NullableGoogleSettingsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleSettingsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
