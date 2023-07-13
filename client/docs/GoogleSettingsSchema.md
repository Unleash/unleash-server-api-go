# GoogleSettingsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is google OIDC enabled | [optional] 
**ClientId** | **string** | The google client id, used to authenticate against google | 
**ClientSecret** | **string** | The client secret used to authenticate the OAuth session used to log the user in | 
**UnleashHostname** | **string** | Name of the host allowed to access the Google authentication flow | 
**AutoCreate** | Pointer to **bool** | Should Unleash create users based on the emails coming back in the authentication reply from Google | [optional] 
**EmailDomains** | Pointer to **string** | A comma separated list of email domains that Unleash will auto create user accounts for. | [optional] 

## Methods

### NewGoogleSettingsSchema

`func NewGoogleSettingsSchema(clientId string, clientSecret string, unleashHostname string, ) *GoogleSettingsSchema`

NewGoogleSettingsSchema instantiates a new GoogleSettingsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleSettingsSchemaWithDefaults

`func NewGoogleSettingsSchemaWithDefaults() *GoogleSettingsSchema`

NewGoogleSettingsSchemaWithDefaults instantiates a new GoogleSettingsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GoogleSettingsSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GoogleSettingsSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GoogleSettingsSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GoogleSettingsSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetClientId

`func (o *GoogleSettingsSchema) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GoogleSettingsSchema) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GoogleSettingsSchema) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *GoogleSettingsSchema) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GoogleSettingsSchema) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GoogleSettingsSchema) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetUnleashHostname

`func (o *GoogleSettingsSchema) GetUnleashHostname() string`

GetUnleashHostname returns the UnleashHostname field if non-nil, zero value otherwise.

### GetUnleashHostnameOk

`func (o *GoogleSettingsSchema) GetUnleashHostnameOk() (*string, bool)`

GetUnleashHostnameOk returns a tuple with the UnleashHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnleashHostname

`func (o *GoogleSettingsSchema) SetUnleashHostname(v string)`

SetUnleashHostname sets UnleashHostname field to given value.


### GetAutoCreate

`func (o *GoogleSettingsSchema) GetAutoCreate() bool`

GetAutoCreate returns the AutoCreate field if non-nil, zero value otherwise.

### GetAutoCreateOk

`func (o *GoogleSettingsSchema) GetAutoCreateOk() (*bool, bool)`

GetAutoCreateOk returns a tuple with the AutoCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreate

`func (o *GoogleSettingsSchema) SetAutoCreate(v bool)`

SetAutoCreate sets AutoCreate field to given value.

### HasAutoCreate

`func (o *GoogleSettingsSchema) HasAutoCreate() bool`

HasAutoCreate returns a boolean if a field has been set.

### GetEmailDomains

`func (o *GoogleSettingsSchema) GetEmailDomains() string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *GoogleSettingsSchema) GetEmailDomainsOk() (*string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *GoogleSettingsSchema) SetEmailDomains(v string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *GoogleSettingsSchema) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


