# ApiTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | **string** | The token used for authentication. | 
**Username** | Pointer to **string** | This property was deprecated in Unleash v5. Prefer the &#x60;tokenName&#x60; property instead. | [optional] 
**TokenName** | **string** | A unique name for this particular token | 
**Type** | **string** | The type of API token | 
**Environment** | Pointer to **string** | The environment the token has access to. &#x60;*&#x60; if it has access to all environments. | [optional] 
**Project** | Pointer to **string** | The project this token belongs to. | [optional] 
**Projects** | Pointer to **[]string** | The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as &#x60;[*]&#x60; | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | The token&#39;s expiration date. NULL if the token doesn&#39;t have an expiration set. | [optional] 
**CreatedAt** | **time.Time** | When the token was created. | 
**SeenAt** | Pointer to **NullableTime** | When the token was last seen/used to authenticate with. NULL if the token has not yet been used for authentication. | [optional] 
**Alias** | Pointer to **NullableString** | Alias is no longer in active use and will often be NULL. It&#39;s kept around as a way of allowing old proxy tokens created with the old metadata format to keep working. | [optional] 

## Methods

### NewApiTokenSchema

`func NewApiTokenSchema(secret string, tokenName string, type_ string, createdAt time.Time, ) *ApiTokenSchema`

NewApiTokenSchema instantiates a new ApiTokenSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokenSchemaWithDefaults

`func NewApiTokenSchemaWithDefaults() *ApiTokenSchema`

NewApiTokenSchemaWithDefaults instantiates a new ApiTokenSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ApiTokenSchema) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApiTokenSchema) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApiTokenSchema) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetUsername

`func (o *ApiTokenSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiTokenSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiTokenSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiTokenSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetTokenName

`func (o *ApiTokenSchema) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *ApiTokenSchema) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *ApiTokenSchema) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetType

`func (o *ApiTokenSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiTokenSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiTokenSchema) SetType(v string)`

SetType sets Type field to given value.


### GetEnvironment

`func (o *ApiTokenSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApiTokenSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApiTokenSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApiTokenSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProject

`func (o *ApiTokenSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ApiTokenSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ApiTokenSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ApiTokenSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *ApiTokenSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ApiTokenSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ApiTokenSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ApiTokenSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiTokenSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiTokenSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiTokenSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiTokenSchema) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ApiTokenSchema) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ApiTokenSchema) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *ApiTokenSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiTokenSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiTokenSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSeenAt

`func (o *ApiTokenSchema) GetSeenAt() time.Time`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *ApiTokenSchema) GetSeenAtOk() (*time.Time, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *ApiTokenSchema) SetSeenAt(v time.Time)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *ApiTokenSchema) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *ApiTokenSchema) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *ApiTokenSchema) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
### GetAlias

`func (o *ApiTokenSchema) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ApiTokenSchema) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ApiTokenSchema) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ApiTokenSchema) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *ApiTokenSchema) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *ApiTokenSchema) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


