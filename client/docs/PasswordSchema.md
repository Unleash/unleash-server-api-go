# PasswordSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**OldPassword** | Pointer to **string** |  | [optional] 
**ConfirmPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewPasswordSchema

`func NewPasswordSchema(password string, ) *PasswordSchema`

NewPasswordSchema instantiates a new PasswordSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordSchemaWithDefaults

`func NewPasswordSchemaWithDefaults() *PasswordSchema`

NewPasswordSchemaWithDefaults instantiates a new PasswordSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordSchema) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOldPassword

`func (o *PasswordSchema) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PasswordSchema) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PasswordSchema) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *PasswordSchema) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetConfirmPassword

`func (o *PasswordSchema) GetConfirmPassword() string`

GetConfirmPassword returns the ConfirmPassword field if non-nil, zero value otherwise.

### GetConfirmPasswordOk

`func (o *PasswordSchema) GetConfirmPasswordOk() (*string, bool)`

GetConfirmPasswordOk returns a tuple with the ConfirmPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmPassword

`func (o *PasswordSchema) SetConfirmPassword(v string)`

SetConfirmPassword sets ConfirmPassword field to given value.

### HasConfirmPassword

`func (o *PasswordSchema) HasConfirmPassword() bool`

HasConfirmPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


