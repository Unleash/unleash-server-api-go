# CreateApiTokenSchemaOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A client or frontend token. Must be one of the strings \&quot;client\&quot; or \&quot;frontend\&quot; (not case sensitive). | 
**Environment** | Pointer to **string** | The environment that the token should be valid for. Defaults to \&quot;default\&quot; | [optional] 
**Project** | Pointer to **string** | The project that the token should be valid for. Defaults to \&quot;*\&quot; meaning every project. This property is mutually incompatible with the &#x60;projects&#x60; property. If you specify one, you cannot specify the other. | [optional] 
**Projects** | Pointer to **[]string** | A list of projects that the token should be valid for. This property is mutually incompatible with the &#x60;project&#x60; property. If you specify one, you cannot specify the other. | [optional] 
**TokenName** | **string** | The name of the token. | 

## Methods

### NewCreateApiTokenSchemaOneOf2

`func NewCreateApiTokenSchemaOneOf2(type_ string, tokenName string, ) *CreateApiTokenSchemaOneOf2`

NewCreateApiTokenSchemaOneOf2 instantiates a new CreateApiTokenSchemaOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiTokenSchemaOneOf2WithDefaults

`func NewCreateApiTokenSchemaOneOf2WithDefaults() *CreateApiTokenSchemaOneOf2`

NewCreateApiTokenSchemaOneOf2WithDefaults instantiates a new CreateApiTokenSchemaOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateApiTokenSchemaOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApiTokenSchemaOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApiTokenSchemaOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetEnvironment

`func (o *CreateApiTokenSchemaOneOf2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateApiTokenSchemaOneOf2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateApiTokenSchemaOneOf2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateApiTokenSchemaOneOf2) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProject

`func (o *CreateApiTokenSchemaOneOf2) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CreateApiTokenSchemaOneOf2) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CreateApiTokenSchemaOneOf2) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CreateApiTokenSchemaOneOf2) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *CreateApiTokenSchemaOneOf2) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateApiTokenSchemaOneOf2) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateApiTokenSchemaOneOf2) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateApiTokenSchemaOneOf2) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTokenName

`func (o *CreateApiTokenSchemaOneOf2) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CreateApiTokenSchemaOneOf2) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CreateApiTokenSchemaOneOf2) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


