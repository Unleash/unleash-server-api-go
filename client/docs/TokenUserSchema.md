# TokenUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The user id | 
**Name** | **string** | The name of the user | 
**Email** | **string** | The email of the user | 
**Token** | **string** | A token uniquely identifying a user | 
**CreatedBy** | **NullableString** | A username or email identifying which user created this token | 
**Role** | [**RoleSchema**](RoleSchema.md) |  | 

## Methods

### NewTokenUserSchema

`func NewTokenUserSchema(id int32, name string, email string, token string, createdBy NullableString, role RoleSchema, ) *TokenUserSchema`

NewTokenUserSchema instantiates a new TokenUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUserSchemaWithDefaults

`func NewTokenUserSchemaWithDefaults() *TokenUserSchema`

NewTokenUserSchemaWithDefaults instantiates a new TokenUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenUserSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenUserSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenUserSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TokenUserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenUserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenUserSchema) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *TokenUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TokenUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TokenUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetToken

`func (o *TokenUserSchema) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenUserSchema) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenUserSchema) SetToken(v string)`

SetToken sets Token field to given value.


### GetCreatedBy

`func (o *TokenUserSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TokenUserSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TokenUserSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *TokenUserSchema) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *TokenUserSchema) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetRole

`func (o *TokenUserSchema) GetRole() RoleSchema`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TokenUserSchema) GetRoleOk() (*RoleSchema, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TokenUserSchema) SetRole(v RoleSchema)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


