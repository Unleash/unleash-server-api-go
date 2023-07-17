# EdgeTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | **[]string** | The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as [&#x60;*&#x60;] | 
**Type** | **string** | The [API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys#api-tokens)&#39;s **type**. Unleash supports three different types of API tokens ([ADMIN](https://docs.getunleash.io/reference/api-tokens-and-client-keys#admin-tokens), [CLIENT](https://docs.getunleash.io/reference/api-tokens-and-client-keys#client-tokens), [FRONTEND](https://docs.getunleash.io/reference/api-tokens-and-client-keys#front-end-tokens)). They all have varying access, so when validating a token it&#39;s important to know what kind you&#39;re dealing with | 
**Token** | **string** | The actual token value. [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys) are comprised of three parts. &lt;project(s)&gt;:&lt;environment&gt;.randomcharacters | 

## Methods

### NewEdgeTokenSchema

`func NewEdgeTokenSchema(projects []string, type_ string, token string, ) *EdgeTokenSchema`

NewEdgeTokenSchema instantiates a new EdgeTokenSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeTokenSchemaWithDefaults

`func NewEdgeTokenSchemaWithDefaults() *EdgeTokenSchema`

NewEdgeTokenSchemaWithDefaults instantiates a new EdgeTokenSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *EdgeTokenSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *EdgeTokenSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *EdgeTokenSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.


### GetType

`func (o *EdgeTokenSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EdgeTokenSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EdgeTokenSchema) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *EdgeTokenSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EdgeTokenSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EdgeTokenSchema) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


