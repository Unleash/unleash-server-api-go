# UserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The user id | 
**IsAPI** | Pointer to **bool** | Deprecated in v5. Used internally to know which operations the user should be allowed to perform | [optional] 
**Name** | Pointer to **NullableString** | Name of the user | [optional] 
**Email** | Pointer to **string** | Email of the user | [optional] 
**Username** | Pointer to **NullableString** | A unique username for the user | [optional] 
**ImageUrl** | Pointer to **string** | URL used for the user profile image | [optional] 
**InviteLink** | Pointer to **string** | If the user is actively inviting other users, this is the link that can be shared with other users | [optional] 
**LoginAttempts** | Pointer to **int32** | How many unsuccessful attempts at logging in has the user made | [optional] 
**EmailSent** | Pointer to **bool** | Is the welcome email sent to the user or not | [optional] 
**RootRole** | Pointer to **int32** | Which [root role](https://docs.getunleash.io/reference/rbac#predefined-roles) this user is assigned | [optional] 
**SeenAt** | Pointer to **NullableTime** | The last time this user logged in | [optional] 
**CreatedAt** | Pointer to **time.Time** | The user was created at this time | [optional] 
**AccountType** | Pointer to **string** | A user is either an actual User or a Service Account | [optional] 
**Permissions** | Pointer to **[]string** | Deprecated | [optional] 
**ScimId** | Pointer to **NullableString** | The SCIM ID of the user, only present if managed by SCIM | [optional] 
**ActiveSessions** | Pointer to **NullableInt32** | Count of active browser sessions for this user | [optional] 
**DeletedSessions** | Pointer to **float32** | Experimental. The number of deleted browser sessions after last login | [optional] 

## Methods

### NewUserSchema

`func NewUserSchema(id int32, ) *UserSchema`

NewUserSchema instantiates a new UserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaWithDefaults

`func NewUserSchemaWithDefaults() *UserSchema`

NewUserSchemaWithDefaults instantiates a new UserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetIsAPI

`func (o *UserSchema) GetIsAPI() bool`

GetIsAPI returns the IsAPI field if non-nil, zero value otherwise.

### GetIsAPIOk

`func (o *UserSchema) GetIsAPIOk() (*bool, bool)`

GetIsAPIOk returns a tuple with the IsAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPI

`func (o *UserSchema) SetIsAPI(v bool)`

SetIsAPI sets IsAPI field to given value.

### HasIsAPI

`func (o *UserSchema) HasIsAPI() bool`

HasIsAPI returns a boolean if a field has been set.

### GetName

`func (o *UserSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserSchema) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserSchema) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *UserSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *UserSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserSchema) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserSchema) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetImageUrl

`func (o *UserSchema) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UserSchema) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UserSchema) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UserSchema) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetInviteLink

`func (o *UserSchema) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *UserSchema) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *UserSchema) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *UserSchema) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *UserSchema) GetLoginAttempts() int32`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *UserSchema) GetLoginAttemptsOk() (*int32, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *UserSchema) SetLoginAttempts(v int32)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *UserSchema) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetEmailSent

`func (o *UserSchema) GetEmailSent() bool`

GetEmailSent returns the EmailSent field if non-nil, zero value otherwise.

### GetEmailSentOk

`func (o *UserSchema) GetEmailSentOk() (*bool, bool)`

GetEmailSentOk returns a tuple with the EmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSent

`func (o *UserSchema) SetEmailSent(v bool)`

SetEmailSent sets EmailSent field to given value.

### HasEmailSent

`func (o *UserSchema) HasEmailSent() bool`

HasEmailSent returns a boolean if a field has been set.

### GetRootRole

`func (o *UserSchema) GetRootRole() int32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *UserSchema) GetRootRoleOk() (*int32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *UserSchema) SetRootRole(v int32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *UserSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### GetSeenAt

`func (o *UserSchema) GetSeenAt() time.Time`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *UserSchema) GetSeenAtOk() (*time.Time, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *UserSchema) SetSeenAt(v time.Time)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *UserSchema) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *UserSchema) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *UserSchema) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
### GetCreatedAt

`func (o *UserSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccountType

`func (o *UserSchema) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UserSchema) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UserSchema) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UserSchema) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetPermissions

`func (o *UserSchema) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserSchema) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserSchema) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserSchema) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetScimId

`func (o *UserSchema) GetScimId() string`

GetScimId returns the ScimId field if non-nil, zero value otherwise.

### GetScimIdOk

`func (o *UserSchema) GetScimIdOk() (*string, bool)`

GetScimIdOk returns a tuple with the ScimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimId

`func (o *UserSchema) SetScimId(v string)`

SetScimId sets ScimId field to given value.

### HasScimId

`func (o *UserSchema) HasScimId() bool`

HasScimId returns a boolean if a field has been set.

### SetScimIdNil

`func (o *UserSchema) SetScimIdNil(b bool)`

 SetScimIdNil sets the value for ScimId to be an explicit nil

### UnsetScimId
`func (o *UserSchema) UnsetScimId()`

UnsetScimId ensures that no value is present for ScimId, not even an explicit nil
### GetActiveSessions

`func (o *UserSchema) GetActiveSessions() int32`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *UserSchema) GetActiveSessionsOk() (*int32, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *UserSchema) SetActiveSessions(v int32)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *UserSchema) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### SetActiveSessionsNil

`func (o *UserSchema) SetActiveSessionsNil(b bool)`

 SetActiveSessionsNil sets the value for ActiveSessions to be an explicit nil

### UnsetActiveSessions
`func (o *UserSchema) UnsetActiveSessions()`

UnsetActiveSessions ensures that no value is present for ActiveSessions, not even an explicit nil
### GetDeletedSessions

`func (o *UserSchema) GetDeletedSessions() float32`

GetDeletedSessions returns the DeletedSessions field if non-nil, zero value otherwise.

### GetDeletedSessionsOk

`func (o *UserSchema) GetDeletedSessionsOk() (*float32, bool)`

GetDeletedSessionsOk returns a tuple with the DeletedSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedSessions

`func (o *UserSchema) SetDeletedSessions(v float32)`

SetDeletedSessions sets DeletedSessions field to given value.

### HasDeletedSessions

`func (o *UserSchema) HasDeletedSessions() bool`

HasDeletedSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


