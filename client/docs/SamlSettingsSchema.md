# SamlSettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is SAML authentication enabled | [optional] 
**EntityId** | **string** | The SAML 2.0 entity ID | 
**SignOnUrl** | **string** | Which URL to use for Single Sign On | 
**Certificate** | **string** | The X509 certificate used to validate requests | 
**SignOutUrl** | Pointer to **string** | Which URL to use for Single Sign Out | [optional] 
**SpCertificate** | Pointer to **string** | Signing certificate for sign out requests | [optional] 
**AutoCreate** | Pointer to **bool** | Should Unleash create users based on the emails coming back in the authentication reply from the SAML server | [optional] 
**EmailDomains** | Pointer to **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] 
**DefaultRootRole** | Pointer to **string** | Assign this root role to auto created users | [optional] 

## Methods

### NewSamlSettingsSchema

`func NewSamlSettingsSchema(entityId string, signOnUrl string, certificate string, ) *SamlSettingsSchema`

NewSamlSettingsSchema instantiates a new SamlSettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlSettingsSchemaWithDefaults

`func NewSamlSettingsSchemaWithDefaults() *SamlSettingsSchema`

NewSamlSettingsSchemaWithDefaults instantiates a new SamlSettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SamlSettingsSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SamlSettingsSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SamlSettingsSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SamlSettingsSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlSettingsSchema) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlSettingsSchema) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlSettingsSchema) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetSignOnUrl

`func (o *SamlSettingsSchema) GetSignOnUrl() string`

GetSignOnUrl returns the SignOnUrl field if non-nil, zero value otherwise.

### GetSignOnUrlOk

`func (o *SamlSettingsSchema) GetSignOnUrlOk() (*string, bool)`

GetSignOnUrlOk returns a tuple with the SignOnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnUrl

`func (o *SamlSettingsSchema) SetSignOnUrl(v string)`

SetSignOnUrl sets SignOnUrl field to given value.


### GetCertificate

`func (o *SamlSettingsSchema) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SamlSettingsSchema) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SamlSettingsSchema) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetSignOutUrl

`func (o *SamlSettingsSchema) GetSignOutUrl() string`

GetSignOutUrl returns the SignOutUrl field if non-nil, zero value otherwise.

### GetSignOutUrlOk

`func (o *SamlSettingsSchema) GetSignOutUrlOk() (*string, bool)`

GetSignOutUrlOk returns a tuple with the SignOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOutUrl

`func (o *SamlSettingsSchema) SetSignOutUrl(v string)`

SetSignOutUrl sets SignOutUrl field to given value.

### HasSignOutUrl

`func (o *SamlSettingsSchema) HasSignOutUrl() bool`

HasSignOutUrl returns a boolean if a field has been set.

### GetSpCertificate

`func (o *SamlSettingsSchema) GetSpCertificate() string`

GetSpCertificate returns the SpCertificate field if non-nil, zero value otherwise.

### GetSpCertificateOk

`func (o *SamlSettingsSchema) GetSpCertificateOk() (*string, bool)`

GetSpCertificateOk returns a tuple with the SpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpCertificate

`func (o *SamlSettingsSchema) SetSpCertificate(v string)`

SetSpCertificate sets SpCertificate field to given value.

### HasSpCertificate

`func (o *SamlSettingsSchema) HasSpCertificate() bool`

HasSpCertificate returns a boolean if a field has been set.

### GetAutoCreate

`func (o *SamlSettingsSchema) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *SamlSettingsSchema) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *SamlSettingsSchema) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *SamlSettingsSchema) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEmailDomains

`func (o *SamlSettingsSchema) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *SamlSettingsSchema) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *SamlSettingsSchema) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *SamlSettingsSchema) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetDefaultRootRole

`func (o *SamlSettingsSchema) GetDefaultRootRole() string`

GetDefaultRootRole returns the DefaultRootRole field if non-nil, zero value otherwise.

### GetDefaultRootRoleOk

`func (o *SamlSettingsSchema) GetDefaultRootRoleOk() (*string, bool)`

GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRole

`func (o *SamlSettingsSchema) SetDefaultRootRole(v string)`

SetDefaultRootRole sets DefaultRootRole field to given value.

### HasDefaultRootRole

`func (o *SamlSettingsSchema) HasDefaultRootRole() bool`

HasDefaultRootRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


