# OidcSettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether to enable or disable OpenID Connect for this instance | 
**DiscoverUrl** | Pointer to **string** | The [.well-known OpenID discover URL](https://swagger.io/docs/specification/authentication/openid-connect-discovery/) | [optional] 
**ClientId** | **string** | The OIDC client ID of this application. | 
**Secret** | **string** | Shared secret from OpenID server. Used to authenticate login requests | 
**AutoCreate** | Pointer to **bool** | Auto create users based on email addresses from login tokens | [optional] 
**EnableSingleSignOut** | Pointer to **bool** | Support Single sign out when user clicks logout in Unleash. If &#x60;true&#x60; user is signed out of all OpenID Connect sessions against the clientId they may have active | [optional] 
**DefaultRootRole** | Pointer to **string** | [Default role](https://docs.getunleash.io/reference/rbac#standard-roles) granted to users auto-created from email. Only relevant if autoCreate is &#x60;true&#x60; | [optional] 
**DefaultRootRoleId** | Pointer to **float32** | Assign this root role to auto created users. Should be a role ID and takes precedence over &#x60;defaultRootRole&#x60;. | [optional] 
**EmailDomains** | Pointer to **string** | Comma separated list of email domains that are automatically approved for an account in the server. Only relevant if autoCreate is &#x60;true&#x60; | [optional] 
**AcrValues** | Pointer to **string** | Authentication Context Class Reference, used to request extra values in the acr claim returned from the server. If multiple values are required, they should be space separated.   Consult [the OIDC reference](https://openid.net/specs/openid-connect-core-1_0.html#AuthorizationEndpoint) for more information   | [optional] 
**IdTokenSigningAlgorithm** | Pointer to **string** | The signing algorithm used to sign our token. Refer to the [JWT signatures](https://jwt.io/introduction) documentation for more information. | [optional] 
**EnableGroupSyncing** | Pointer to **bool** | Should we enable group syncing. Refer to the documentation [Group syncing](https://docs.getunleash.io/how-to/how-to-set-up-group-sso-sync) | [optional] 
**GroupJsonPath** | Pointer to **string** | Specifies the path in the OIDC token response to read which groups the user belongs to from. | [optional] 
**AddGroupsScope** | Pointer to **bool** | When enabled Unleash will also request the &#39;groups&#39; scope as part of the login request. | [optional] 

## Methods

### NewOidcSettingsSchema

`func NewOidcSettingsSchema(enabled bool, clientId string, secret string, ) *OidcSettingsSchema`

NewOidcSettingsSchema instantiates a new OidcSettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcSettingsSchemaWithDefaults

`func NewOidcSettingsSchemaWithDefaults() *OidcSettingsSchema`

NewOidcSettingsSchemaWithDefaults instantiates a new OidcSettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OidcSettingsSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OidcSettingsSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OidcSettingsSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDiscoverUrl

`func (o *OidcSettingsSchema) GetDiscoverUrl() string`

GetDiscoverUrl returns the DiscoverUrl field if non-nil, zero value otherwise.

### GetDiscoverUrlOk

`func (o *OidcSettingsSchema) GetDiscoverUrlOk() (*string, bool)`

GetDiscoverUrlOk returns a tuple with the DiscoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverUrl

`func (o *OidcSettingsSchema) SetDiscoverUrl(v string)`

SetDiscoverUrl sets DiscoverUrl field to given value.

### HasDiscoverUrl

`func (o *OidcSettingsSchema) HasDiscoverUrl() bool`

HasDiscoverUrl returns a boolean if a field has been set.

### GetClientId

`func (o *OidcSettingsSchema) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcSettingsSchema) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcSettingsSchema) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSecret

`func (o *OidcSettingsSchema) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OidcSettingsSchema) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OidcSettingsSchema) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetAutoCreate

`func (o *OidcSettingsSchema) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *OidcSettingsSchema) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *OidcSettingsSchema) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *OidcSettingsSchema) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEnableSingleSignOut

`func (o *OidcSettingsSchema) GetEnableSingleSignOut() bool`

GetEnableSingleSignOut returns the EnableSingleSignOut field if non-nil, zero value otherwise.

### GetEnableSingleSignOutOk

`func (o *OidcSettingsSchema) GetEnableSingleSignOutOk() (*bool, bool)`

GetEnableSingleSignOutOk returns a tuple with the EnableSingleSignOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSignOut

`func (o *OidcSettingsSchema) SetEnableSingleSignOut(v bool)`

SetEnableSingleSignOut sets EnableSingleSignOut field to given value.

### HasEnableSingleSignOut

`func (o *OidcSettingsSchema) HasEnableSingleSignOut() bool`

HasEnableSingleSignOut returns a boolean if a field has been set.

### GetDefaultRootRole

`func (o *OidcSettingsSchema) GetDefaultRootRole() string`

GetDefaultRootRole returns the DefaultRootRole field if non-nil, zero value otherwise.

### GetDefaultRootRoleOk

`func (o *OidcSettingsSchema) GetDefaultRootRoleOk() (*string, bool)`

GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRole

`func (o *OidcSettingsSchema) SetDefaultRootRole(v string)`

SetDefaultRootRole sets DefaultRootRole field to given value.

### HasDefaultRootRole

`func (o *OidcSettingsSchema) HasDefaultRootRole() bool`

HasDefaultRootRole returns a boolean if a field has been set.

### GetDefaultRootRoleId

`func (o *OidcSettingsSchema) GetDefaultRootRoleId() float32`

GetDefaultRootRoleId returns the DefaultRootRoleId field if non-nil, zero value otherwise.

### GetDefaultRootRoleIdOk

`func (o *OidcSettingsSchema) GetDefaultRootRoleIdOk() (*float32, bool)`

GetDefaultRootRoleIdOk returns a tuple with the DefaultRootRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRoleId

`func (o *OidcSettingsSchema) SetDefaultRootRoleId(v float32)`

SetDefaultRootRoleId sets DefaultRootRoleId field to given value.

### HasDefaultRootRoleId

`func (o *OidcSettingsSchema) HasDefaultRootRoleId() bool`

HasDefaultRootRoleId returns a boolean if a field has been set.

### GetEmailDomains

`func (o *OidcSettingsSchema) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *OidcSettingsSchema) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *OidcSettingsSchema) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *OidcSettingsSchema) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetAcrValues

`func (o *OidcSettingsSchema) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *OidcSettingsSchema) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *OidcSettingsSchema) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *OidcSettingsSchema) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetIdTokenSigningAlgorithm

`func (o *OidcSettingsSchema) GetIdTokenSigningAlgorithm() string`

GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgorithmOk

`func (o *OidcSettingsSchema) GetIdTokenSigningAlgorithmOk() (*string, bool)`

GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgorithm

`func (o *OidcSettingsSchema) SetIdTokenSigningAlgorithm(v string)`

SetIdTokenSigningAlgorithm sets IdTokenSigningAlgorithm field to given value.

### HasIdTokenSigningAlgorithm

`func (o *OidcSettingsSchema) HasIdTokenSigningAlgorithm() bool`

HasIdTokenSigningAlgorithm returns a boolean if a field has been set.

### GetEnableGroupSyncing

`func (o *OidcSettingsSchema) GetEnableGroupSyncing() bool`

GetEnableGroupSyncing returns the EnableGroupSyncing field if non-nil, zero value otherwise.

### GetEnableGroupSyncingOk

`func (o *OidcSettingsSchema) GetEnableGroupSyncingOk() (*bool, bool)`

GetEnableGroupSyncingOk returns a tuple with the EnableGroupSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupSyncing

`func (o *OidcSettingsSchema) SetEnableGroupSyncing(v bool)`

SetEnableGroupSyncing sets EnableGroupSyncing field to given value.

### HasEnableGroupSyncing

`func (o *OidcSettingsSchema) HasEnableGroupSyncing() bool`

HasEnableGroupSyncing returns a boolean if a field has been set.

### GetGroupJsonPath

`func (o *OidcSettingsSchema) GetGroupJsonPath() string`

GetGroupJsonPath returns the GroupJsonPath field if non-nil, zero value otherwise.

### GetGroupJsonPathOk

`func (o *OidcSettingsSchema) GetGroupJsonPathOk() (*string, bool)`

GetGroupJsonPathOk returns a tuple with the GroupJsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupJsonPath

`func (o *OidcSettingsSchema) SetGroupJsonPath(v string)`

SetGroupJsonPath sets GroupJsonPath field to given value.

### HasGroupJsonPath

`func (o *OidcSettingsSchema) HasGroupJsonPath() bool`

HasGroupJsonPath returns a boolean if a field has been set.

### GetAddGroupsScope

`func (o *OidcSettingsSchema) GetAddGroupsScope() bool`

GetAddGroupsScope returns the AddGroupsScope field if non-nil, zero value otherwise.

### GetAddGroupsScopeOk

`func (o *OidcSettingsSchema) GetAddGroupsScopeOk() (*bool, bool)`

GetAddGroupsScopeOk returns a tuple with the AddGroupsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGroupsScope

`func (o *OidcSettingsSchema) SetAddGroupsScope(v bool)`

SetAddGroupsScope sets AddGroupsScope field to given value.

### HasAddGroupsScope

`func (o *OidcSettingsSchema) HasAddGroupsScope() bool`

HasAddGroupsScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


