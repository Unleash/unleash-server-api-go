# PublicSignupTokenCreateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The token&#39;s name. | 
**ExpiresAt** | **time.Time** | The token&#39;s expiration date. | 

## Methods

### NewPublicSignupTokenCreateSchema

`func NewPublicSignupTokenCreateSchema(name string, expiresAt time.Time, ) *PublicSignupTokenCreateSchema`

NewPublicSignupTokenCreateSchema instantiates a new PublicSignupTokenCreateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSignupTokenCreateSchemaWithDefaults

`func NewPublicSignupTokenCreateSchemaWithDefaults() *PublicSignupTokenCreateSchema`

NewPublicSignupTokenCreateSchemaWithDefaults instantiates a new PublicSignupTokenCreateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicSignupTokenCreateSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSignupTokenCreateSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSignupTokenCreateSchema) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresAt

`func (o *PublicSignupTokenCreateSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicSignupTokenCreateSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicSignupTokenCreateSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

