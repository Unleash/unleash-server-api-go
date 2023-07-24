/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.2.0-main
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the OidcSettingsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcSettingsSchema{}

// OidcSettingsSchema Settings for configuring OpenID Connect as a login provider for Unleash
type OidcSettingsSchema struct {
	// `true` if OpenID connect is turned on for this instance, otherwise `false`
	Enabled *bool `json:"enabled,omitempty"`
	// The [.well-known OpenID discover URL](https://swagger.io/docs/specification/authentication/openid-connect-discovery/)
	DiscoverUrl *string `json:"discoverUrl,omitempty"`
	// The OIDC client ID of this application.
	ClientId string `json:"clientId"`
	// Shared secret from OpenID server. Used to authenticate login requests
	Secret string `json:"secret"`
	// Auto create users based on email addresses from login tokens
	AutoCreate *bool `json:"autoCreate,omitempty"`
	// Support Single sign out when user clicks logout in Unleash. If `true` user is signed out of all OpenID Connect sessions against the clientId they may have active
	EnableSingleSignOut *bool `json:"enableSingleSignOut,omitempty"`
	// [Default role](https://docs.getunleash.io/reference/rbac#standard-roles) granted to users auto-created from email. Only relevant if autoCreate is `true`
	DefaultRootRole *string `json:"defaultRootRole,omitempty"`
	// Comma separated list of email domains that are automatically approved for an account in the server. Only relevant if autoCreate is `true`
	EmailDomains *string `json:"emailDomains,omitempty"`
	// Authentication Context Class Reference, used to request extra values in the acr claim returned from the server. If multiple values are required, they should be space separated.   Consult [the OIDC reference](https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationEndpoint) for more information
	AcrValues *string `json:"acrValues,omitempty"`
	// The signing algorithm used to sign our token. Refer to the [JWT signatures](https://jwt.io/introduction) documentation for more information.
	IdTokenSigningAlgorithm *string `json:"idTokenSigningAlgorithm,omitempty"`
}

// NewOidcSettingsSchema instantiates a new OidcSettingsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcSettingsSchema(clientId string, secret string) *OidcSettingsSchema {
	this := OidcSettingsSchema{}
	this.ClientId = clientId
	this.Secret = secret
	return &this
}

// NewOidcSettingsSchemaWithDefaults instantiates a new OidcSettingsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcSettingsSchemaWithDefaults() *OidcSettingsSchema {
	this := OidcSettingsSchema{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *OidcSettingsSchema) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDiscoverUrl returns the DiscoverUrl field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetDiscoverUrl() string {
	if o == nil || IsNil(o.DiscoverUrl) {
		var ret string
		return ret
	}
	return *o.DiscoverUrl
}

// GetDiscoverUrlOk returns a tuple with the DiscoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetDiscoverUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DiscoverUrl) {
		return nil, false
	}
	return o.DiscoverUrl, true
}

// HasDiscoverUrl returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasDiscoverUrl() bool {
	if o != nil && !IsNil(o.DiscoverUrl) {
		return true
	}

	return false
}

// SetDiscoverUrl gets a reference to the given string and assigns it to the DiscoverUrl field.
func (o *OidcSettingsSchema) SetDiscoverUrl(v string) {
	o.DiscoverUrl = &v
}

// GetClientId returns the ClientId field value
func (o *OidcSettingsSchema) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OidcSettingsSchema) SetClientId(v string) {
	o.ClientId = v
}

// GetSecret returns the Secret field value
func (o *OidcSettingsSchema) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *OidcSettingsSchema) SetSecret(v string) {
	o.Secret = v
}

// GetAutoCreate returns the AutoCreate field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetAutoCreate() bool {
	if o == nil || IsNil(o.AutoCreate) {
		var ret bool
		return ret
	}
	return *o.AutoCreate
}

// GetAutoCreateOk returns a tuple with the AutoCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetAutoCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreate) {
		return nil, false
	}
	return o.AutoCreate, true
}

// HasAutoCreate returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasAutoCreate() bool {
	if o != nil && !IsNil(o.AutoCreate) {
		return true
	}

	return false
}

// SetAutoCreate gets a reference to the given bool and assigns it to the AutoCreate field.
func (o *OidcSettingsSchema) SetAutoCreate(v bool) {
	o.AutoCreate = &v
}

// GetEnableSingleSignOut returns the EnableSingleSignOut field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetEnableSingleSignOut() bool {
	if o == nil || IsNil(o.EnableSingleSignOut) {
		var ret bool
		return ret
	}
	return *o.EnableSingleSignOut
}

// GetEnableSingleSignOutOk returns a tuple with the EnableSingleSignOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetEnableSingleSignOutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSingleSignOut) {
		return nil, false
	}
	return o.EnableSingleSignOut, true
}

// HasEnableSingleSignOut returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasEnableSingleSignOut() bool {
	if o != nil && !IsNil(o.EnableSingleSignOut) {
		return true
	}

	return false
}

// SetEnableSingleSignOut gets a reference to the given bool and assigns it to the EnableSingleSignOut field.
func (o *OidcSettingsSchema) SetEnableSingleSignOut(v bool) {
	o.EnableSingleSignOut = &v
}

// GetDefaultRootRole returns the DefaultRootRole field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetDefaultRootRole() string {
	if o == nil || IsNil(o.DefaultRootRole) {
		var ret string
		return ret
	}
	return *o.DefaultRootRole
}

// GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetDefaultRootRoleOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRootRole) {
		return nil, false
	}
	return o.DefaultRootRole, true
}

// HasDefaultRootRole returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasDefaultRootRole() bool {
	if o != nil && !IsNil(o.DefaultRootRole) {
		return true
	}

	return false
}

// SetDefaultRootRole gets a reference to the given string and assigns it to the DefaultRootRole field.
func (o *OidcSettingsSchema) SetDefaultRootRole(v string) {
	o.DefaultRootRole = &v
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetEmailDomains() string {
	if o == nil || IsNil(o.EmailDomains) {
		var ret string
		return ret
	}
	return *o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetEmailDomainsOk() (*string, bool) {
	if o == nil || IsNil(o.EmailDomains) {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasEmailDomains() bool {
	if o != nil && !IsNil(o.EmailDomains) {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given string and assigns it to the EmailDomains field.
func (o *OidcSettingsSchema) SetEmailDomains(v string) {
	o.EmailDomains = &v
}

// GetAcrValues returns the AcrValues field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetAcrValues() string {
	if o == nil || IsNil(o.AcrValues) {
		var ret string
		return ret
	}
	return *o.AcrValues
}

// GetAcrValuesOk returns a tuple with the AcrValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetAcrValuesOk() (*string, bool) {
	if o == nil || IsNil(o.AcrValues) {
		return nil, false
	}
	return o.AcrValues, true
}

// HasAcrValues returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasAcrValues() bool {
	if o != nil && !IsNil(o.AcrValues) {
		return true
	}

	return false
}

// SetAcrValues gets a reference to the given string and assigns it to the AcrValues field.
func (o *OidcSettingsSchema) SetAcrValues(v string) {
	o.AcrValues = &v
}

// GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field value if set, zero value otherwise.
func (o *OidcSettingsSchema) GetIdTokenSigningAlgorithm() string {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenSigningAlgorithm
}

// GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettingsSchema) GetIdTokenSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		return nil, false
	}
	return o.IdTokenSigningAlgorithm, true
}

// HasIdTokenSigningAlgorithm returns a boolean if a field has been set.
func (o *OidcSettingsSchema) HasIdTokenSigningAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenSigningAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenSigningAlgorithm gets a reference to the given string and assigns it to the IdTokenSigningAlgorithm field.
func (o *OidcSettingsSchema) SetIdTokenSigningAlgorithm(v string) {
	o.IdTokenSigningAlgorithm = &v
}

func (o OidcSettingsSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcSettingsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DiscoverUrl) {
		toSerialize["discoverUrl"] = o.DiscoverUrl
	}
	toSerialize["clientId"] = o.ClientId
	toSerialize["secret"] = o.Secret
	if !IsNil(o.AutoCreate) {
		toSerialize["autoCreate"] = o.AutoCreate
	}
	if !IsNil(o.EnableSingleSignOut) {
		toSerialize["enableSingleSignOut"] = o.EnableSingleSignOut
	}
	if !IsNil(o.DefaultRootRole) {
		toSerialize["defaultRootRole"] = o.DefaultRootRole
	}
	if !IsNil(o.EmailDomains) {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	if !IsNil(o.AcrValues) {
		toSerialize["acrValues"] = o.AcrValues
	}
	if !IsNil(o.IdTokenSigningAlgorithm) {
		toSerialize["idTokenSigningAlgorithm"] = o.IdTokenSigningAlgorithm
	}
	return toSerialize, nil
}

type NullableOidcSettingsSchema struct {
	value *OidcSettingsSchema
	isSet bool
}

func (v NullableOidcSettingsSchema) Get() *OidcSettingsSchema {
	return v.value
}

func (v *NullableOidcSettingsSchema) Set(val *OidcSettingsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcSettingsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcSettingsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcSettingsSchema(val *OidcSettingsSchema) *NullableOidcSettingsSchema {
	return &NullableOidcSettingsSchema{value: val, isSet: true}
}

func (v NullableOidcSettingsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcSettingsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
