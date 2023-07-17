# CreateUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The user&#39;s username. Must be provided if email is not provided. | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address. Must be provided if username is not provided. | [optional] 
**Name** | Pointer to **string** | The user&#39;s name (not the user&#39;s username). | [optional] 
**Password** | Pointer to **string** | Password for the user | [optional] 
**RootRole** | [**CreateUserSchemaRootRole**](CreateUserSchemaRootRole.md) |  | 
**SendEmail** | Pointer to **bool** | Whether to send a welcome email with a login link to the user or not. Defaults to &#x60;true&#x60;. | [optional] 

## Methods

### NewCreateUserSchema

`func NewCreateUserSchema(rootRole CreateUserSchemaRootRole, ) *CreateUserSchema`

NewCreateUserSchema instantiates a new CreateUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSchemaWithDefaults

`func NewCreateUserSchemaWithDefaults() *CreateUserSchema`

NewCreateUserSchemaWithDefaults instantiates a new CreateUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateUserSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateUserSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CreateUserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *CreateUserSchema) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserSchema) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserSchema) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserSchema) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRootRole

`func (o *CreateUserSchema) GetRootRole() CreateUserSchemaRootRole`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *CreateUserSchema) GetRootRoleOk() (*CreateUserSchemaRootRole, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *CreateUserSchema) SetRootRole(v CreateUserSchemaRootRole)`

SetRootRole sets RootRole field to given value.


### GetSendEmail

`func (o *CreateUserSchema) GetSendEmail() bool`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *CreateUserSchema) GetSendEmailOk() (*bool, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *CreateUserSchema) SetSendEmail(v bool)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *CreateUserSchema) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


