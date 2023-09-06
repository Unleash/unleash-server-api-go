# CreateApiTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when this token should expire. | [optional] 
**Type** | **string** | A client or frontend token. Must be one of the strings \&quot;client\&quot; or \&quot;frontend\&quot; (not case sensitive). | 
**TokenName** | **string** | The name of the token. | 
**Username** | **string** | The name of the token. This property is deprecated. Use &#x60;tokenName&#x60; instead. | 
**Environment** | Pointer to **string** | The environment that the token should be valid for. Defaults to \&quot;default\&quot; | [optional] 
**Project** | Pointer to **string** | The project that the token should be valid for. Defaults to \&quot;*\&quot; meaning every project. This property is mutually incompatible with the &#x60;projects&#x60; property. If you specify one, you cannot specify the other. | [optional] 
**Projects** | Pointer to **[]string** | A list of projects that the token should be valid for. This property is mutually incompatible with the &#x60;project&#x60; property. If you specify one, you cannot specify the other. | [optional] 

## Methods

### NewCreateApiTokenSchema

`func NewCreateApiTokenSchema(type_ string, tokenName string, username string, ) *CreateApiTokenSchema`

NewCreateApiTokenSchema instantiates a new CreateApiTokenSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiTokenSchemaWithDefaults

`func NewCreateApiTokenSchemaWithDefaults() *CreateApiTokenSchema`

NewCreateApiTokenSchemaWithDefaults instantiates a new CreateApiTokenSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CreateApiTokenSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateApiTokenSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateApiTokenSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateApiTokenSchema) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetType

`func (o *CreateApiTokenSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApiTokenSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApiTokenSchema) SetType(v string)`

SetType sets Type field to given value.


### GetTokenName

`func (o *CreateApiTokenSchema) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CreateApiTokenSchema) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CreateApiTokenSchema) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.


### GetUsername

`func (o *CreateApiTokenSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateApiTokenSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateApiTokenSchema) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEnvironment

`func (o *CreateApiTokenSchema) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateApiTokenSchema) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateApiTokenSchema) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateApiTokenSchema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProject

`func (o *CreateApiTokenSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateApiTokenSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateApiTokenSchema) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateApiTokenSchema) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *CreateApiTokenSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateApiTokenSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateApiTokenSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateApiTokenSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


