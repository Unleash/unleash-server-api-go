# CreateApiTokenSchemaOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | The time when this token should expire. | [optional] 
**Type** | **string** | An admin token. Must be the string \&quot;admin\&quot; (not case sensitive). | 
**Username** | **string** | The name of the token. This property was deprecated in v5. Use &#x60;tokenName&#x60; instead. | 

## Methods

### NewCreateApiTokenSchemaOneOf1

`func NewCreateApiTokenSchemaOneOf1(type_ string, username string, ) *CreateApiTokenSchemaOneOf1`

NewCreateApiTokenSchemaOneOf1 instantiates a new CreateApiTokenSchemaOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiTokenSchemaOneOf1WithDefaults

`func NewCreateApiTokenSchemaOneOf1WithDefaults() *CreateApiTokenSchemaOneOf1`

NewCreateApiTokenSchemaOneOf1WithDefaults instantiates a new CreateApiTokenSchemaOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CreateApiTokenSchemaOneOf1) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateApiTokenSchemaOneOf1) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateApiTokenSchemaOneOf1) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateApiTokenSchemaOneOf1) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetType

`func (o *CreateApiTokenSchemaOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateApiTokenSchemaOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateApiTokenSchemaOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetUsername

`func (o *CreateApiTokenSchemaOneOf1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateApiTokenSchemaOneOf1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateApiTokenSchemaOneOf1) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


