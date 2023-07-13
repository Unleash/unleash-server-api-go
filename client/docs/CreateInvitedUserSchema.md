# CreateInvitedUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewCreateInvitedUserSchema

`func NewCreateInvitedUserSchema(email string, name string, password string, ) *CreateInvitedUserSchema`

NewCreateInvitedUserSchema instantiates a new CreateInvitedUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInvitedUserSchemaWithDefaults

`func NewCreateInvitedUserSchemaWithDefaults() *CreateInvitedUserSchema`

NewCreateInvitedUserSchemaWithDefaults instantiates a new CreateInvitedUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateInvitedUserSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateInvitedUserSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateInvitedUserSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateInvitedUserSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *CreateInvitedUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateInvitedUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateInvitedUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CreateInvitedUserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInvitedUserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInvitedUserSchema) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateInvitedUserSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateInvitedUserSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateInvitedUserSchema) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


