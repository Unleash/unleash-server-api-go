# CreateUserResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The user id | 
**IsAPI** | Pointer to **bool** | (Deprecated): Used internally to know which operations the user should be allowed to perform | [optional] 
**Name** | Pointer to **string** | Name of the user | [optional] 
**Email** | Pointer to **string** | Email of the user | [optional] 
**Username** | Pointer to **string** | A unique username for the user | [optional] 
**ImageUrl** | Pointer to **string** | URL used for the userprofile image | [optional] 
**InviteLink** | Pointer to **string** | If the user is actively inviting other users, this is the link that can be shared with other users | [optional] 
**LoginAttempts** | Pointer to **int32** | How many unsuccessful attempts at logging in has the user made | [optional] 
**EmailSent** | Pointer to **bool** | Is the welcome email sent to the user or not | [optional] 
**RootRole** | Pointer to [**CreateUserResponseSchemaRootRole**](CreateUserResponseSchemaRootRole.md) |  | [optional] 
**SeenAt** | Pointer to **NullableTime** | The last time this user logged in | [optional] 
**CreatedAt** | Pointer to **time.Time** | The user was created at this time | [optional] 
**AccountType** | Pointer to **string** | A user is either an actual User or a Service Account | [optional] 
**Permissions** | Pointer to **[]string** | Deprecated | [optional] 

## Methods

### NewCreateUserResponseSchema

`func NewCreateUserResponseSchema(id int32, ) *CreateUserResponseSchema`

NewCreateUserResponseSchema instantiates a new CreateUserResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserResponseSchemaWithDefaults

`func NewCreateUserResponseSchemaWithDefaults() *CreateUserResponseSchema`

NewCreateUserResponseSchemaWithDefaults instantiates a new CreateUserResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateUserResponseSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateUserResponseSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateUserResponseSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetIsAPI

`func (o *CreateUserResponseSchema) GetIsAPI() bool`

GetIsAPI returns the IsAPI field if non-nil, zero value otherwise.

### GetIsAPIOk

`func (o *CreateUserResponseSchema) GetIsAPIOk() (*bool, bool)`

GetIsAPIOk returns a tuple with the IsAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPI

`func (o *CreateUserResponseSchema) SetIsAPI(v bool)`

SetIsAPI sets IsAPI field to given value.

### HasIsAPI

`func (o *CreateUserResponseSchema) HasIsAPI() bool`

HasIsAPI returns a boolean if a field has been set.

### GetName

`func (o *CreateUserResponseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserResponseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserResponseSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserResponseSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserResponseSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserResponseSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserResponseSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserResponseSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *CreateUserResponseSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserResponseSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserResponseSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateUserResponseSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateUserResponseSchema) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateUserResponseSchema) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateUserResponseSchema) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateUserResponseSchema) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetInviteLink

`func (o *CreateUserResponseSchema) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *CreateUserResponseSchema) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *CreateUserResponseSchema) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *CreateUserResponseSchema) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *CreateUserResponseSchema) GetLoginAttempts() int32`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *CreateUserResponseSchema) GetLoginAttemptsOk() (*int32, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *CreateUserResponseSchema) SetLoginAttempts(v int32)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *CreateUserResponseSchema) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetEmailSent

`func (o *CreateUserResponseSchema) GetEmailSent() bool`

GetEmailSent returns the EmailSent field if non-nil, zero value otherwise.

### GetEmailSentOk

`func (o *CreateUserResponseSchema) GetEmailSentOk() (*bool, bool)`

GetEmailSentOk returns a tuple with the EmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSent

`func (o *CreateUserResponseSchema) SetEmailSent(v bool)`

SetEmailSent sets EmailSent field to given value.

### HasEmailSent

`func (o *CreateUserResponseSchema) HasEmailSent() bool`

HasEmailSent returns a boolean if a field has been set.

### GetRootRole

`func (o *CreateUserResponseSchema) GetRootRole() CreateUserResponseSchemaRootRole`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *CreateUserResponseSchema) GetRootRoleOk() (*CreateUserResponseSchemaRootRole, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *CreateUserResponseSchema) SetRootRole(v CreateUserResponseSchemaRootRole)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *CreateUserResponseSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### GetSeenAt

`func (o *CreateUserResponseSchema) GetSeenAt() time.Time`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *CreateUserResponseSchema) GetSeenAtOk() (*time.Time, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *CreateUserResponseSchema) SetSeenAt(v time.Time)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *CreateUserResponseSchema) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *CreateUserResponseSchema) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *CreateUserResponseSchema) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
### GetCreatedAt

`func (o *CreateUserResponseSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateUserResponseSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateUserResponseSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateUserResponseSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccountType

`func (o *CreateUserResponseSchema) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *CreateUserResponseSchema) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *CreateUserResponseSchema) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *CreateUserResponseSchema) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateUserResponseSchema) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateUserResponseSchema) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateUserResponseSchema) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateUserResponseSchema) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


