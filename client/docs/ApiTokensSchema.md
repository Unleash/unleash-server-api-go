# ApiTokensSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | [**[]ApiTokenSchema**](ApiTokenSchema.md) | A list of Unleash API tokens. | 

## Methods

### NewApiTokensSchema

`func NewApiTokensSchema(tokens []ApiTokenSchema, ) *ApiTokensSchema`

NewApiTokensSchema instantiates a new ApiTokensSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTokensSchemaWithDefaults

`func NewApiTokensSchemaWithDefaults() *ApiTokensSchema`

NewApiTokensSchemaWithDefaults instantiates a new ApiTokensSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *ApiTokensSchema) GetTokens() []ApiTokenSchema`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ApiTokensSchema) GetTokensOk() (*[]ApiTokenSchema, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ApiTokensSchema) SetTokens(v []ApiTokenSchema)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


