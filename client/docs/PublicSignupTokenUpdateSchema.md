# PublicSignupTokenUpdateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The token&#39;s expiration date. | [optional] 
**Enabled** | Pointer to **bool** | Whether the token is active or not. | [optional] 

## Methods

### NewPublicSignupTokenUpdateSchema

`func NewPublicSignupTokenUpdateSchema() *PublicSignupTokenUpdateSchema`

NewPublicSignupTokenUpdateSchema instantiates a new PublicSignupTokenUpdateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSignupTokenUpdateSchemaWithDefaults

`func NewPublicSignupTokenUpdateSchemaWithDefaults() *PublicSignupTokenUpdateSchema`

NewPublicSignupTokenUpdateSchemaWithDefaults instantiates a new PublicSignupTokenUpdateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *PublicSignupTokenUpdateSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicSignupTokenUpdateSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicSignupTokenUpdateSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PublicSignupTokenUpdateSchema) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetEnabled

`func (o *PublicSignupTokenUpdateSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PublicSignupTokenUpdateSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PublicSignupTokenUpdateSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PublicSignupTokenUpdateSchema) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


