# ValidatedEdgeTokensSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | [**[]EdgeTokenSchema**](EdgeTokenSchema.md) | The list of Unleash token objects. Each object contains the token itself and some additional metadata. | 

## Methods

### NewValidatedEdgeTokensSchema

`func NewValidatedEdgeTokensSchema(tokens []EdgeTokenSchema, ) *ValidatedEdgeTokensSchema`

NewValidatedEdgeTokensSchema instantiates a new ValidatedEdgeTokensSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedEdgeTokensSchemaWithDefaults

`func NewValidatedEdgeTokensSchemaWithDefaults() *ValidatedEdgeTokensSchema`

NewValidatedEdgeTokensSchemaWithDefaults instantiates a new ValidatedEdgeTokensSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *ValidatedEdgeTokensSchema) GetTokens() []EdgeTokenSchema`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ValidatedEdgeTokensSchema) GetTokensOk() (*[]EdgeTokenSchema, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ValidatedEdgeTokensSchema) SetTokens(v []EdgeTokenSchema)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


