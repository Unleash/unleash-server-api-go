# CreateApiTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when this token should expire. | [optional] 
**Type** | **string** | An admin token. Must be the string \&quot;admin\&quot; (not case sensitive). | 
**TokenName** | **string** | The name of the token. | 

## Methods

### NewCreateApiTokenSchema

`func NewCreateApiTokenSchema(type_ string, tokenName string, ) *CreateApiTokenSchema`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


