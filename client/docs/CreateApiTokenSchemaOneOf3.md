# CreateApiTokenSchemaOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when this token should expire. | [optional] 
**Type** | **string** | A client or frontend token. Must be one of the strings \&quot;client\&quot; or \&quot;frontend\&quot; (not case sensitive). | 
**Environment** | Pointer to **string** | The environment that the token should be valid for. Defaults to \&quot;default\&quot; | [optional] 
**Project** | Pointer to **string** | The project that the token should be valid for. Defaults to \&quot;*\&quot; meaning every project. This property is mutually incompatible with the &#x60;projects&#x60; property. If you specify one, you cannot specify the other. | [optional] 
**Projects** | Pointer to **[]string** | A list of projects that the token should be valid for. This property is mutually incompatible with the &#x60;project&#x60; property. If you specify one, you cannot specify the other. | [optional] 
**Username** | **string** | The name of the token. This property is deprecated. Use &#x60;tokenName&#x60; instead. | 

## Methods

### NewCreateApiTokenSchemaOneOf3

`func NewCreateApiTokenSchemaOneOf3(type_ string, username string, ) *CreateApiTokenSchemaOneOf3`

NewCreateApiTokenSchemaOneOf3 instantiates a new CreateApiTokenSchemaOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiTokenSchemaOneOf3WithDefaults

`func NewCreateApiTokenSchemaOneOf3WithDefaults() *CreateApiTokenSchemaOneOf3`

NewCreateApiTokenSchemaOneOf3WithDefaults instantiates a new CreateApiTokenSchemaOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CreateApiTokenSchemaOneOf3) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateApiTokenSchemaOneOf3) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateApiTokenSchemaOneOf3) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateApiTokenSchemaOneOf3) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetType

`func (o *CreateApiTokenSchemaOneOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApiTokenSchemaOneOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApiTokenSchemaOneOf3) SetType(v string)`

SetType sets Type field to given value.


### GetEnvironment

`func (o *CreateApiTokenSchemaOneOf3) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateApiTokenSchemaOneOf3) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateApiTokenSchemaOneOf3) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateApiTokenSchemaOneOf3) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProject

`func (o *CreateApiTokenSchemaOneOf3) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateApiTokenSchemaOneOf3) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateApiTokenSchemaOneOf3) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateApiTokenSchemaOneOf3) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *CreateApiTokenSchemaOneOf3) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateApiTokenSchemaOneOf3) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateApiTokenSchemaOneOf3) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateApiTokenSchemaOneOf3) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUsername

`func (o *CreateApiTokenSchemaOneOf3) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateApiTokenSchemaOneOf3) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateApiTokenSchemaOneOf3) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


