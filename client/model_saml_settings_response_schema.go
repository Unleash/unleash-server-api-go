/*
Unleash API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 6.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SamlSettingsResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlSettingsResponseSchema{}

// SamlSettingsResponseSchema Response for SAML settings
type SamlSettingsResponseSchema struct {
	// Whether to enable or disable SAML 2.0 for this instance
	Enabled *bool `json:"enabled,omitempty"`
	// The SAML 2.0 entity ID
	EntityId *string `json:"entityId,omitempty"`
	// Which URL to use for Single Sign On
	SignOnUrl *string `json:"signOnUrl,omitempty"`
	// The X509 certificate used to validate requests
	Certificate *string `json:"certificate,omitempty"`
	// Which URL to use for Single Sign Out
	SignOutUrl *string `json:"signOutUrl,omitempty"`
	// Signing certificate for sign out requests
	SpCertificate *string `json:"spCertificate,omitempty"`
	// Should Unleash create users based on the emails coming back in the authentication reply from the SAML server
	AutoCreate *bool `json:"autoCreate,omitempty"`
	// A comma separated list of email domains that Unleash will auto create user accounts for.
	EmailDomains *string `json:"emailDomains,omitempty"`
	// Assign this root role to auto created users
	DefaultRootRole *string `json:"defaultRootRole,omitempty"`
	// Assign this root role to auto created users. Should be a role ID and takes precedence over `defaultRootRole`.
	DefaultRootRoleId *float32 `json:"defaultRootRoleId,omitempty"`
	// Should we enable group syncing. Refer to the documentation [Group syncing](https://docs.getunleash.io/how-to/how-to-set-up-group-sso-sync)
	EnableGroupSyncing *bool `json:"enableGroupSyncing,omitempty"`
	// Specifies the path in the SAML token response from which to read the groups the user belongs to.
	GroupJsonPath        *string `json:"groupJsonPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlSettingsResponseSchema SamlSettingsResponseSchema

// NewSamlSettingsResponseSchema instantiates a new SamlSettingsResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlSettingsResponseSchema() *SamlSettingsResponseSchema {
	this := SamlSettingsResponseSchema{}
	return &this
}

// NewSamlSettingsResponseSchemaWithDefaults instantiates a new SamlSettingsResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlSettingsResponseSchemaWithDefaults() *SamlSettingsResponseSchema {
	this := SamlSettingsResponseSchema{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SamlSettingsResponseSchema) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SamlSettingsResponseSchema) SetEntityId(v string) {
	o.EntityId = &v
}

// GetSignOnUrl returns the SignOnUrl field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetSignOnUrl() string {
	if o == nil || IsNil(o.SignOnUrl) {
		var ret string
		return ret
	}
	return *o.SignOnUrl
}

// GetSignOnUrlOk returns a tuple with the SignOnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetSignOnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignOnUrl) {
		return nil, false
	}
	return o.SignOnUrl, true
}

// HasSignOnUrl returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasSignOnUrl() bool {
	if o != nil && !IsNil(o.SignOnUrl) {
		return true
	}

	return false
}

// SetSignOnUrl gets a reference to the given string and assigns it to the SignOnUrl field.
func (o *SamlSettingsResponseSchema) SetSignOnUrl(v string) {
	o.SignOnUrl = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *SamlSettingsResponseSchema) SetCertificate(v string) {
	o.Certificate = &v
}

// GetSignOutUrl returns the SignOutUrl field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetSignOutUrl() string {
	if o == nil || IsNil(o.SignOutUrl) {
		var ret string
		return ret
	}
	return *o.SignOutUrl
}

// GetSignOutUrlOk returns a tuple with the SignOutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetSignOutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignOutUrl) {
		return nil, false
	}
	return o.SignOutUrl, true
}

// HasSignOutUrl returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasSignOutUrl() bool {
	if o != nil && !IsNil(o.SignOutUrl) {
		return true
	}

	return false
}

// SetSignOutUrl gets a reference to the given string and assigns it to the SignOutUrl field.
func (o *SamlSettingsResponseSchema) SetSignOutUrl(v string) {
	o.SignOutUrl = &v
}

// GetSpCertificate returns the SpCertificate field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetSpCertificate() string {
	if o == nil || IsNil(o.SpCertificate) {
		var ret string
		return ret
	}
	return *o.SpCertificate
}

// GetSpCertificateOk returns a tuple with the SpCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetSpCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SpCertificate) {
		return nil, false
	}
	return o.SpCertificate, true
}

// HasSpCertificate returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasSpCertificate() bool {
	if o != nil && !IsNil(o.SpCertificate) {
		return true
	}

	return false
}

// SetSpCertificate gets a reference to the given string and assigns it to the SpCertificate field.
func (o *SamlSettingsResponseSchema) SetSpCertificate(v string) {
	o.SpCertificate = &v
}

// GetAutoCreate returns the AutoCreate field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetAutoCreate() bool {
	if o == nil || IsNil(o.AutoCreate) {
		var ret bool
		return ret
	}
	return *o.AutoCreate
}

// GetAutoCreateOk returns a tuple with the AutoCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetAutoCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreate) {
		return nil, false
	}
	return o.AutoCreate, true
}

// HasAutoCreate returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasAutoCreate() bool {
	if o != nil && !IsNil(o.AutoCreate) {
		return true
	}

	return false
}

// SetAutoCreate gets a reference to the given bool and assigns it to the AutoCreate field.
func (o *SamlSettingsResponseSchema) SetAutoCreate(v bool) {
	o.AutoCreate = &v
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetEmailDomains() string {
	if o == nil || IsNil(o.EmailDomains) {
		var ret string
		return ret
	}
	return *o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetEmailDomainsOk() (*string, bool) {
	if o == nil || IsNil(o.EmailDomains) {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasEmailDomains() bool {
	if o != nil && !IsNil(o.EmailDomains) {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given string and assigns it to the EmailDomains field.
func (o *SamlSettingsResponseSchema) SetEmailDomains(v string) {
	o.EmailDomains = &v
}

// GetDefaultRootRole returns the DefaultRootRole field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetDefaultRootRole() string {
	if o == nil || IsNil(o.DefaultRootRole) {
		var ret string
		return ret
	}
	return *o.DefaultRootRole
}

// GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetDefaultRootRoleOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRootRole) {
		return nil, false
	}
	return o.DefaultRootRole, true
}

// HasDefaultRootRole returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasDefaultRootRole() bool {
	if o != nil && !IsNil(o.DefaultRootRole) {
		return true
	}

	return false
}

// SetDefaultRootRole gets a reference to the given string and assigns it to the DefaultRootRole field.
func (o *SamlSettingsResponseSchema) SetDefaultRootRole(v string) {
	o.DefaultRootRole = &v
}

// GetDefaultRootRoleId returns the DefaultRootRoleId field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetDefaultRootRoleId() float32 {
	if o == nil || IsNil(o.DefaultRootRoleId) {
		var ret float32
		return ret
	}
	return *o.DefaultRootRoleId
}

// GetDefaultRootRoleIdOk returns a tuple with the DefaultRootRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetDefaultRootRoleIdOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultRootRoleId) {
		return nil, false
	}
	return o.DefaultRootRoleId, true
}

// HasDefaultRootRoleId returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasDefaultRootRoleId() bool {
	if o != nil && !IsNil(o.DefaultRootRoleId) {
		return true
	}

	return false
}

// SetDefaultRootRoleId gets a reference to the given float32 and assigns it to the DefaultRootRoleId field.
func (o *SamlSettingsResponseSchema) SetDefaultRootRoleId(v float32) {
	o.DefaultRootRoleId = &v
}

// GetEnableGroupSyncing returns the EnableGroupSyncing field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetEnableGroupSyncing() bool {
	if o == nil || IsNil(o.EnableGroupSyncing) {
		var ret bool
		return ret
	}
	return *o.EnableGroupSyncing
}

// GetEnableGroupSyncingOk returns a tuple with the EnableGroupSyncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetEnableGroupSyncingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableGroupSyncing) {
		return nil, false
	}
	return o.EnableGroupSyncing, true
}

// HasEnableGroupSyncing returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasEnableGroupSyncing() bool {
	if o != nil && !IsNil(o.EnableGroupSyncing) {
		return true
	}

	return false
}

// SetEnableGroupSyncing gets a reference to the given bool and assigns it to the EnableGroupSyncing field.
func (o *SamlSettingsResponseSchema) SetEnableGroupSyncing(v bool) {
	o.EnableGroupSyncing = &v
}

// GetGroupJsonPath returns the GroupJsonPath field value if set, zero value otherwise.
func (o *SamlSettingsResponseSchema) GetGroupJsonPath() string {
	if o == nil || IsNil(o.GroupJsonPath) {
		var ret string
		return ret
	}
	return *o.GroupJsonPath
}

// GetGroupJsonPathOk returns a tuple with the GroupJsonPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettingsResponseSchema) GetGroupJsonPathOk() (*string, bool) {
	if o == nil || IsNil(o.GroupJsonPath) {
		return nil, false
	}
	return o.GroupJsonPath, true
}

// HasGroupJsonPath returns a boolean if a field has been set.
func (o *SamlSettingsResponseSchema) HasGroupJsonPath() bool {
	if o != nil && !IsNil(o.GroupJsonPath) {
		return true
	}

	return false
}

// SetGroupJsonPath gets a reference to the given string and assigns it to the GroupJsonPath field.
func (o *SamlSettingsResponseSchema) SetGroupJsonPath(v string) {
	o.GroupJsonPath = &v
}

func (o SamlSettingsResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlSettingsResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.SignOnUrl) {
		toSerialize["signOnUrl"] = o.SignOnUrl
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.SignOutUrl) {
		toSerialize["signOutUrl"] = o.SignOutUrl
	}
	if !IsNil(o.SpCertificate) {
		toSerialize["spCertificate"] = o.SpCertificate
	}
	if !IsNil(o.AutoCreate) {
		toSerialize["autoCreate"] = o.AutoCreate
	}
	if !IsNil(o.EmailDomains) {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	if !IsNil(o.DefaultRootRole) {
		toSerialize["defaultRootRole"] = o.DefaultRootRole
	}
	if !IsNil(o.DefaultRootRoleId) {
		toSerialize["defaultRootRoleId"] = o.DefaultRootRoleId
	}
	if !IsNil(o.EnableGroupSyncing) {
		toSerialize["enableGroupSyncing"] = o.EnableGroupSyncing
	}
	if !IsNil(o.GroupJsonPath) {
		toSerialize["groupJsonPath"] = o.GroupJsonPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlSettingsResponseSchema) UnmarshalJSON(data []byte) (err error) {
	varSamlSettingsResponseSchema := _SamlSettingsResponseSchema{}

	err = json.Unmarshal(data, &varSamlSettingsResponseSchema)

	if err != nil {
		return err
	}

	*o = SamlSettingsResponseSchema(varSamlSettingsResponseSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "entityId")
		delete(additionalProperties, "signOnUrl")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "signOutUrl")
		delete(additionalProperties, "spCertificate")
		delete(additionalProperties, "autoCreate")
		delete(additionalProperties, "emailDomains")
		delete(additionalProperties, "defaultRootRole")
		delete(additionalProperties, "defaultRootRoleId")
		delete(additionalProperties, "enableGroupSyncing")
		delete(additionalProperties, "groupJsonPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlSettingsResponseSchema struct {
	value *SamlSettingsResponseSchema
	isSet bool
}

func (v NullableSamlSettingsResponseSchema) Get() *SamlSettingsResponseSchema {
	return v.value
}

func (v *NullableSamlSettingsResponseSchema) Set(val *SamlSettingsResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSettingsResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSettingsResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSettingsResponseSchema(val *SamlSettingsResponseSchema) *NullableSamlSettingsResponseSchema {
	return &NullableSamlSettingsResponseSchema{value: val, isSet: true}
}

func (v NullableSamlSettingsResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSettingsResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
