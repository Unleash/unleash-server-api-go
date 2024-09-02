# SamlSettingsResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable or disable SAML 2.0 for this instance | [optional] 
**EntityId** | Pointer to **string** | The SAML 2.0 entity ID | [optional] 
**SignOnUrl** | Pointer to **string** | Which URL to use for Single Sign On | [optional] 
**Certificate** | Pointer to **string** | The X509 certificate used to validate requests | [optional] 
**SignOutUrl** | Pointer to **string** | Which URL to use for Single Sign Out | [optional] 
**SpCertificate** | Pointer to **string** | Signing certificate for sign out requests | [optional] 
**AutoCreate** | Pointer to **bool** | Should Unleash create users based on the emails coming back in the authentication reply from the SAML server | [optional] 
**EmailDomains** | Pointer to **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] 
**DefaultRootRole** | Pointer to **string** | Assign this root role to auto created users | [optional] 
**DefaultRootRoleId** | Pointer to **float32** | Assign this root role to auto created users. Should be a role ID and takes precedence over &#x60;defaultRootRole&#x60;. | [optional] 
**EnableGroupSyncing** | Pointer to **bool** | Should we enable group syncing. Refer to the documentation [Group syncing](https://docs.getunleash.io/how-to/how-to-set-up-group-sso-sync) | [optional] 
**GroupJsonPath** | Pointer to **string** | Specifies the path in the SAML token response from which to read the groups the user belongs to. | [optional] 

## Methods

### NewSamlSettingsResponseSchema

`func NewSamlSettingsResponseSchema() *SamlSettingsResponseSchema`

NewSamlSettingsResponseSchema instantiates a new SamlSettingsResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSettingsResponseSchemaWithDefaults

`func NewSamlSettingsResponseSchemaWithDefaults() *SamlSettingsResponseSchema`

NewSamlSettingsResponseSchemaWithDefaults instantiates a new SamlSettingsResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SamlSettingsResponseSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlSettingsResponseSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlSettingsResponseSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SamlSettingsResponseSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlSettingsResponseSchema) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlSettingsResponseSchema) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlSettingsResponseSchema) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SamlSettingsResponseSchema) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetSignOnUrl

`func (o *SamlSettingsResponseSchema) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *SamlSettingsResponseSchema) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *SamlSettingsResponseSchema) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.

### HasSignOnUrl

`func (o *SamlSettingsResponseSchema) HasSignOnUrl() bool`

HasSignOnUrl returns a boolean if a field has been set.

### GetCertificate

`func (o *SamlSettingsResponseSchema) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SamlSettingsResponseSchema) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SamlSettingsResponseSchema) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SamlSettingsResponseSchema) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetSignOutUrl

`func (o *SamlSettingsResponseSchema) GetSignOutUrl() string`

GetSignOutUrl returns the SignOutUrl field if non-nil, zero value otherwise.

### GetSignOutUrlOk

`func (o *SamlSettingsResponseSchema) GetSignOutUrlOk() (*string, bool)`

GetSignOutUrlOk returns a tuple with the SignOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutUrl

`func (o *SamlSettingsResponseSchema) SetSignOutUrl(v string)`

SetSignOutUrl sets SignOutUrl field to given value.

### HasSignOutUrl

`func (o *SamlSettingsResponseSchema) HasSignOutUrl() bool`

HasSignOutUrl returns a boolean if a field has been set.

### GetSpCertificate

`func (o *SamlSettingsResponseSchema) GetSpCertificate() string`

GetSpCertificate returns the SpCertificate field if non-nil, zero value otherwise.

### GetSpCertificateOk

`func (o *SamlSettingsResponseSchema) GetSpCertificateOk() (*string, bool)`

GetSpCertificateOk returns a tuple with the SpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificate

`func (o *SamlSettingsResponseSchema) SetSpCertificate(v string)`

SetSpCertificate sets SpCertificate field to given value.

### HasSpCertificate

`func (o *SamlSettingsResponseSchema) HasSpCertificate() bool`

HasSpCertificate returns a boolean if a field has been set.

### GetAutoCreate

`func (o *SamlSettingsResponseSchema) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *SamlSettingsResponseSchema) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *SamlSettingsResponseSchema) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *SamlSettingsResponseSchema) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEmailDomains

`func (o *SamlSettingsResponseSchema) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *SamlSettingsResponseSchema) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *SamlSettingsResponseSchema) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *SamlSettingsResponseSchema) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetDefaultRootRole

`func (o *SamlSettingsResponseSchema) GetDefaultRootRole() string`

GetDefaultRootRole returns the DefaultRootRole field if non-nil, zero value otherwise.

### GetDefaultRootRoleOk

`func (o *SamlSettingsResponseSchema) GetDefaultRootRoleOk() (*string, bool)`

GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRole

`func (o *SamlSettingsResponseSchema) SetDefaultRootRole(v string)`

SetDefaultRootRole sets DefaultRootRole field to given value.

### HasDefaultRootRole

`func (o *SamlSettingsResponseSchema) HasDefaultRootRole() bool`

HasDefaultRootRole returns a boolean if a field has been set.

### GetDefaultRootRoleId

`func (o *SamlSettingsResponseSchema) GetDefaultRootRoleId() float32`

GetDefaultRootRoleId returns the DefaultRootRoleId field if non-nil, zero value otherwise.

### GetDefaultRootRoleIdOk

`func (o *SamlSettingsResponseSchema) GetDefaultRootRoleIdOk() (*float32, bool)`

GetDefaultRootRoleIdOk returns a tuple with the DefaultRootRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRoleId

`func (o *SamlSettingsResponseSchema) SetDefaultRootRoleId(v float32)`

SetDefaultRootRoleId sets DefaultRootRoleId field to given value.

### HasDefaultRootRoleId

`func (o *SamlSettingsResponseSchema) HasDefaultRootRoleId() bool`

HasDefaultRootRoleId returns a boolean if a field has been set.

### GetEnableGroupSyncing

`func (o *SamlSettingsResponseSchema) GetEnableGroupSyncing() bool`

GetEnableGroupSyncing returns the EnableGroupSyncing field if non-nil, zero value otherwise.

### GetEnableGroupSyncingOk

`func (o *SamlSettingsResponseSchema) GetEnableGroupSyncingOk() (*bool, bool)`

GetEnableGroupSyncingOk returns a tuple with the EnableGroupSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupSyncing

`func (o *SamlSettingsResponseSchema) SetEnableGroupSyncing(v bool)`

SetEnableGroupSyncing sets EnableGroupSyncing field to given value.

### HasEnableGroupSyncing

`func (o *SamlSettingsResponseSchema) HasEnableGroupSyncing() bool`

HasEnableGroupSyncing returns a boolean if a field has been set.

### GetGroupJsonPath

`func (o *SamlSettingsResponseSchema) GetGroupJsonPath() string`

GetGroupJsonPath returns the GroupJsonPath field if non-nil, zero value otherwise.

### GetGroupJsonPathOk

`func (o *SamlSettingsResponseSchema) GetGroupJsonPathOk() (*string, bool)`

GetGroupJsonPathOk returns a tuple with the GroupJsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupJsonPath

`func (o *SamlSettingsResponseSchema) SetGroupJsonPath(v string)`

SetGroupJsonPath sets GroupJsonPath field to given value.

### HasGroupJsonPath

`func (o *SamlSettingsResponseSchema) HasGroupJsonPath() bool`

HasGroupJsonPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


