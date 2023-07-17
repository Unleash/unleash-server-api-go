# CreateApiTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when this token should expire. | [optional] 

## Methods

### NewCreateApiTokenSchema

`func NewCreateApiTokenSchema() *CreateApiTokenSchema`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


