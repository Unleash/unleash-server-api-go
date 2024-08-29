# OidcSettingsSchemaOneOf

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

### NewOidcSettingsSchemaOneOf

`func NewOidcSettingsSchemaOneOf(enabled bool, clientId string, secret string, ) *OidcSettingsSchemaOneOf`

NewOidcSettingsSchemaOneOf instantiates a new OidcSettingsSchemaOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcSettingsSchemaOneOfWithDefaults

`func NewOidcSettingsSchemaOneOfWithDefaults() *OidcSettingsSchemaOneOf`

NewOidcSettingsSchemaOneOfWithDefaults instantiates a new OidcSettingsSchemaOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *OidcSettingsSchemaOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OidcSettingsSchemaOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OidcSettingsSchemaOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDiscoverUrl

`func (o *OidcSettingsSchemaOneOf) GetDiscoverUrl() string`

GetDiscoverUrl returns the DiscoverUrl field if non-nil, zero value otherwise.

### GetDiscoverUrlOk

`func (o *OidcSettingsSchemaOneOf) GetDiscoverUrlOk() (*string, bool)`

GetDiscoverUrlOk returns a tuple with the DiscoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverUrl

`func (o *OidcSettingsSchemaOneOf) SetDiscoverUrl(v string)`

SetDiscoverUrl sets DiscoverUrl field to given value.

### HasDiscoverUrl

`func (o *OidcSettingsSchemaOneOf) HasDiscoverUrl() bool`

HasDiscoverUrl returns a boolean if a field has been set.

### GetClientId

`func (o *OidcSettingsSchemaOneOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OidcSettingsSchemaOneOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OidcSettingsSchemaOneOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSecret

`func (o *OidcSettingsSchemaOneOf) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *OidcSettingsSchemaOneOf) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *OidcSettingsSchemaOneOf) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetAutoCreate

`func (o *OidcSettingsSchemaOneOf) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *OidcSettingsSchemaOneOf) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *OidcSettingsSchemaOneOf) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *OidcSettingsSchemaOneOf) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEnableSingleSignOut

`func (o *OidcSettingsSchemaOneOf) GetEnableSingleSignOut() bool`

GetEnableSingleSignOut returns the EnableSingleSignOut field if non-nil, zero value otherwise.

### GetEnableSingleSignOutOk

`func (o *OidcSettingsSchemaOneOf) GetEnableSingleSignOutOk() (*bool, bool)`

GetEnableSingleSignOutOk returns a tuple with the EnableSingleSignOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSignOut

`func (o *OidcSettingsSchemaOneOf) SetEnableSingleSignOut(v bool)`

SetEnableSingleSignOut sets EnableSingleSignOut field to given value.

### HasEnableSingleSignOut

`func (o *OidcSettingsSchemaOneOf) HasEnableSingleSignOut() bool`

HasEnableSingleSignOut returns a boolean if a field has been set.

### GetDefaultRootRole

`func (o *OidcSettingsSchemaOneOf) GetDefaultRootRole() string`

GetDefaultRootRole returns the DefaultRootRole field if non-nil, zero value otherwise.

### GetDefaultRootRoleOk

`func (o *OidcSettingsSchemaOneOf) GetDefaultRootRoleOk() (*string, bool)`

GetDefaultRootRoleOk returns a tuple with the DefaultRootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRole

`func (o *OidcSettingsSchemaOneOf) SetDefaultRootRole(v string)`

SetDefaultRootRole sets DefaultRootRole field to given value.

### HasDefaultRootRole

`func (o *OidcSettingsSchemaOneOf) HasDefaultRootRole() bool`

HasDefaultRootRole returns a boolean if a field has been set.

### GetDefaultRootRoleId

`func (o *OidcSettingsSchemaOneOf) GetDefaultRootRoleId() float32`

GetDefaultRootRoleId returns the DefaultRootRoleId field if non-nil, zero value otherwise.

### GetDefaultRootRoleIdOk

`func (o *OidcSettingsSchemaOneOf) GetDefaultRootRoleIdOk() (*float32, bool)`

GetDefaultRootRoleIdOk returns a tuple with the DefaultRootRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRootRoleId

`func (o *OidcSettingsSchemaOneOf) SetDefaultRootRoleId(v float32)`

SetDefaultRootRoleId sets DefaultRootRoleId field to given value.

### HasDefaultRootRoleId

`func (o *OidcSettingsSchemaOneOf) HasDefaultRootRoleId() bool`

HasDefaultRootRoleId returns a boolean if a field has been set.

### GetEmailDomains

`func (o *OidcSettingsSchemaOneOf) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *OidcSettingsSchemaOneOf) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *OidcSettingsSchemaOneOf) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *OidcSettingsSchemaOneOf) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetAcrValues

`func (o *OidcSettingsSchemaOneOf) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *OidcSettingsSchemaOneOf) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *OidcSettingsSchemaOneOf) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *OidcSettingsSchemaOneOf) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetIdTokenSigningAlgorithm

`func (o *OidcSettingsSchemaOneOf) GetIdTokenSigningAlgorithm() string`

GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgorithmOk

`func (o *OidcSettingsSchemaOneOf) GetIdTokenSigningAlgorithmOk() (*string, bool)`

GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgorithm

`func (o *OidcSettingsSchemaOneOf) SetIdTokenSigningAlgorithm(v string)`

SetIdTokenSigningAlgorithm sets IdTokenSigningAlgorithm field to given value.

### HasIdTokenSigningAlgorithm

`func (o *OidcSettingsSchemaOneOf) HasIdTokenSigningAlgorithm() bool`

HasIdTokenSigningAlgorithm returns a boolean if a field has been set.

### GetEnableGroupSyncing

`func (o *OidcSettingsSchemaOneOf) GetEnableGroupSyncing() bool`

GetEnableGroupSyncing returns the EnableGroupSyncing field if non-nil, zero value otherwise.

### GetEnableGroupSyncingOk

`func (o *OidcSettingsSchemaOneOf) GetEnableGroupSyncingOk() (*bool, bool)`

GetEnableGroupSyncingOk returns a tuple with the EnableGroupSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupSyncing

`func (o *OidcSettingsSchemaOneOf) SetEnableGroupSyncing(v bool)`

SetEnableGroupSyncing sets EnableGroupSyncing field to given value.

### HasEnableGroupSyncing

`func (o *OidcSettingsSchemaOneOf) HasEnableGroupSyncing() bool`

HasEnableGroupSyncing returns a boolean if a field has been set.

### GetGroupJsonPath

`func (o *OidcSettingsSchemaOneOf) GetGroupJsonPath() string`

GetGroupJsonPath returns the GroupJsonPath field if non-nil, zero value otherwise.

### GetGroupJsonPathOk

`func (o *OidcSettingsSchemaOneOf) GetGroupJsonPathOk() (*string, bool)`

GetGroupJsonPathOk returns a tuple with the GroupJsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupJsonPath

`func (o *OidcSettingsSchemaOneOf) SetGroupJsonPath(v string)`

SetGroupJsonPath sets GroupJsonPath field to given value.

### HasGroupJsonPath

`func (o *OidcSettingsSchemaOneOf) HasGroupJsonPath() bool`

HasGroupJsonPath returns a boolean if a field has been set.

### GetAddGroupsScope

`func (o *OidcSettingsSchemaOneOf) GetAddGroupsScope() bool`

GetAddGroupsScope returns the AddGroupsScope field if non-nil, zero value otherwise.

### GetAddGroupsScopeOk

`func (o *OidcSettingsSchemaOneOf) GetAddGroupsScopeOk() (*bool, bool)`

GetAddGroupsScopeOk returns a tuple with the AddGroupsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddGroupsScope

`func (o *OidcSettingsSchemaOneOf) SetAddGroupsScope(v bool)`

SetAddGroupsScope sets AddGroupsScope field to given value.

### HasAddGroupsScope

`func (o *OidcSettingsSchemaOneOf) HasAddGroupsScope() bool`

HasAddGroupsScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


