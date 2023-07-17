# PatSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identification number for this Personal Access Token. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] 
**Secret** | Pointer to **string** | The token used for authentication. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The token&#39;s expiration date. | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the token was created. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] 
**SeenAt** | Pointer to **NullableTime** | When the token was last seen/used to authenticate with. &#x60;null&#x60; if it has not been used yet. (This property is set by Unleash when the token is created and cannot be set manually: if you provide a value when creating a PAT, Unleash will ignore it.) | [optional] 

## Methods

### NewPatSchema

`func NewPatSchema() *PatSchema`

NewPatSchema instantiates a new PatSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatSchemaWithDefaults

`func NewPatSchemaWithDefaults() *PatSchema`

NewPatSchemaWithDefaults instantiates a new PatSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatSchema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecret

`func (o *PatSchema) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatSchema) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatSchema) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatSchema) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PatSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PatSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PatSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PatSchema) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSeenAt

`func (o *PatSchema) GetSeenAt() time.Time`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *PatSchema) GetSeenAtOk() (*time.Time, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *PatSchema) SetSeenAt(v time.Time)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *PatSchema) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *PatSchema) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *PatSchema) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


