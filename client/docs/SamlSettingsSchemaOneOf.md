# SamlSettingsSchemaOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable or disable SAML 2.0 for this instance | [optional] 
**EntityId** | **string** | The SAML 2.0 entity ID | 
**SignOnUrl** | **string** | Which URL to use for Single Sign On | 
**Certificate** | **string** | The X509 certificate used to validate requests | 
**SignOutUrl** | Pointer to **string** | Which URL to use for Single Sign Out | [optional] 
**SpCertificate** | Pointer to **string** | Signing certificate for sign out requests | [optional] 
**AutoCreate** | Pointer to **bool** | Should Unleash create users based on the emails coming back in the authentication reply from the SAML server | [optional] 
**EmailDomains** | Pointer to **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] 
**DefaultRootRole** | Pointer to **string** | Assign this root role to auto created users | [optional] 
**DefaultRootRoleId** | Pointer to **float32** | Assign this root role to auto created users. Should be a role ID and takes precedence over &#x60;defaultRootRole&#x60;. | [optional] 
**EnableGroupSyncing** | Pointer to **bool** | Should we enable group syncing. Refer to the documentation [Group syncing](https://docs.getunleash.io/how-to/how-to-set-up-group-sso-sync) | [optional] 
**GroupJsonPath** | Pointer to **string** | Specifies the path in the SAML token response from which to read the groups the user belongs to. | [optional] 

## Methods

### NewSamlSettingsSchemaOneOf

`func NewSamlSettingsSchemaOneOf(entityId string, signOnUrl string, certificate string, ) *SamlSettingsSchemaOneOf`

NewSamlSettingsSchemaOneOf instantiates a new SamlSettingsSchemaOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSettingsSchemaOneOfWithDefaults

`func NewSamlSettingsSchemaOneOfWithDefaults() *SamlSettingsSchemaOneOf`

NewSamlSettingsSchemaOneOfWithDefaults instantiates a new SamlSettingsSchemaOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SamlSettingsSchemaOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlSettingsSchemaOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlSettingsSchemaOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SamlSettingsSchemaOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlSettingsSchemaOneOf) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlSettingsSchemaOneOf) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlSettingsSchemaOneOf) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetSignOnUrl

`func (o *SamlSettingsSchemaOneOf) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *SamlSettingsSchemaOneOf) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *SamlSettingsSchemaOneOf) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.


### GetCertificate

`func (o *SamlSettingsSchemaOneOf) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SamlSettingsSchemaOneOf) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SamlSettingsSchemaOneOf) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetSignOutUrl

`func (o *SamlSettingsSchemaOneOf) GetSignOutUrl() string`

GetSignOutUrl returns the SignOutUrl field if non-nil, zero value otherwise.

### GetSignOutUrlOk

`func (o *SamlSettingsSchemaOneOf) GetSignOutUrlOk() (*string, bool)`

GetSignOutUrlOk returns a tuple with the SignOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutUrl

`func (o *SamlSettingsSchemaOneOf) SetSignOutUrl(v string)`

SetSignOutUrl sets SignOutUrl field to given value.

### HasSignOutUrl

`func (o *SamlSettingsSchemaOneOf) HasSignOutUrl() bool`

HasSignOutUrl returns a boolean if a field has been set.

### GetSpCertificate

`func (o *SamlSettingsSchemaOneOf) GetSpCertificate() string`

GetSpCertificate returns the SpCertificate field if non-nil, zero value otherwise.

### GetSpCertificateOk

`func (o *SamlSettingsSchemaOneOf) GetSpCertificateOk() (*string, bool)`

GetSpCertificateOk returns a tuple with the SpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificate

`func (o *SamlSettingsSchemaOneOf) SetSpCertificate(v string)`

SetSpCertificate sets SpCertificate field to given value.

### HasSpCertificate

`func (o *SamlSettingsSchemaOneOf) HasSpCertificate() bool`

HasSpCertificate returns a boolean if a field has been set.

### GetAutoCreate

`func (o *SamlSettingsSchemaOneOf) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *SamlSettingsSchemaOneOf) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *SamlSettingsSchemaOneOf) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *SamlSettingsSchemaOneOf) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEmailDomains

`func (o *SamlSettingsSchemaOneOf) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *SamlSettingsSchemaOneOf) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *SamlSettingsSchemaOneOf) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *SamlSettingsSchemaOneOf) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetDefaultRootRole

`func (o *SamlSettingsSchemaOneOf) GetDefaultRootRole() string`

GetDefaultRootRole returns the DefaultRootRole field if non-nil, zero value otherwise.

### GetDefaultRootRoleOk

`func (o *SamlSettingsSchemaOneOf) GetDefaultRootRoleOk() (*string, bool)`

GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRole

`func (o *SamlSettingsSchemaOneOf) SetDefaultRootRole(v string)`

SetDefaultRootRole sets DefaultRootRole field to given value.

### HasDefaultRootRole

`func (o *SamlSettingsSchemaOneOf) HasDefaultRootRole() bool`

HasDefaultRootRole returns a boolean if a field has been set.

### GetDefaultRootRoleId

`func (o *SamlSettingsSchemaOneOf) GetDefaultRootRoleId() float32`

GetDefaultRootRoleId returns the DefaultRootRoleId field if non-nil, zero value otherwise.

### GetDefaultRootRoleIdOk

`func (o *SamlSettingsSchemaOneOf) GetDefaultRootRoleIdOk() (*float32, bool)`

GetDefaultRootRoleIdOk returns a tuple with the DefaultRootRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRoleId

`func (o *SamlSettingsSchemaOneOf) SetDefaultRootRoleId(v float32)`

SetDefaultRootRoleId sets DefaultRootRoleId field to given value.

### HasDefaultRootRoleId

`func (o *SamlSettingsSchemaOneOf) HasDefaultRootRoleId() bool`

HasDefaultRootRoleId returns a boolean if a field has been set.

### GetEnableGroupSyncing

`func (o *SamlSettingsSchemaOneOf) GetEnableGroupSyncing() bool`

GetEnableGroupSyncing returns the EnableGroupSyncing field if non-nil, zero value otherwise.

### GetEnableGroupSyncingOk

`func (o *SamlSettingsSchemaOneOf) GetEnableGroupSyncingOk() (*bool, bool)`

GetEnableGroupSyncingOk returns a tuple with the EnableGroupSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupSyncing

`func (o *SamlSettingsSchemaOneOf) SetEnableGroupSyncing(v bool)`

SetEnableGroupSyncing sets EnableGroupSyncing field to given value.

### HasEnableGroupSyncing

`func (o *SamlSettingsSchemaOneOf) HasEnableGroupSyncing() bool`

HasEnableGroupSyncing returns a boolean if a field has been set.

### GetGroupJsonPath

`func (o *SamlSettingsSchemaOneOf) GetGroupJsonPath() string`

GetGroupJsonPath returns the GroupJsonPath field if non-nil, zero value otherwise.

### GetGroupJsonPathOk

`func (o *SamlSettingsSchemaOneOf) GetGroupJsonPathOk() (*string, bool)`

GetGroupJsonPathOk returns a tuple with the GroupJsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupJsonPath

`func (o *SamlSettingsSchemaOneOf) SetGroupJsonPath(v string)`

SetGroupJsonPath sets GroupJsonPath field to given value.

### HasGroupJsonPath

`func (o *SamlSettingsSchemaOneOf) HasGroupJsonPath() bool`

HasGroupJsonPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


