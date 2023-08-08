# ServiceAccountSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | The service account id | 
**IsAPI** | Pointer to **bool** | Deprecated: for internal use only, should not be exposed through the API | [optional] 
**Name** | Pointer to **string** | The name of the service account | [optional] 
**Email** | Pointer to **string** | Deprecated: service accounts don&#39;t have emails associated with them | [optional] 
**Username** | Pointer to **string** | The service account username | [optional] 
**ImageUrl** | Pointer to **string** | The service account image url | [optional] 
**InviteLink** | Pointer to **string** | Deprecated: service accounts cannot be invited via an invitation link | [optional] 
**LoginAttempts** | Pointer to **float32** | Deprecated: service accounts cannot log in to Unleash | [optional] 
**EmailSent** | Pointer to **bool** | Deprecated: internal use only | [optional] 
**RootRole** | Pointer to **int32** | The root role id associated with the service account | [optional] 
**SeenAt** | Pointer to **NullableTime** | Deprecated. This property is always &#x60;null&#x60;. To find out when a service account was last seen, check its &#x60;tokens&#x60; list and refer to each token&#39;s &#x60;lastSeen&#x60; property instead. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The service account creation date | [optional] 
**Tokens** | Pointer to [**[]PatSchema**](PatSchema.md) | The list of tokens associated with the service account | [optional] 

## Methods

### NewServiceAccountSchema

`func NewServiceAccountSchema(id float32, ) *ServiceAccountSchema`

NewServiceAccountSchema instantiates a new ServiceAccountSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSchemaWithDefaults

`func NewServiceAccountSchemaWithDefaults() *ServiceAccountSchema`

NewServiceAccountSchemaWithDefaults instantiates a new ServiceAccountSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccountSchema) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountSchema) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountSchema) SetId(v float32)`

SetId sets Id field to given value.


### GetIsAPI

`func (o *ServiceAccountSchema) GetIsAPI() bool`

GetIsAPI returns the IsAPI field if non-nil, zero value otherwise.

### GetIsAPIOk

`func (o *ServiceAccountSchema) GetIsAPIOk() (*bool, bool)`

GetIsAPIOk returns a tuple with the IsAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPI

`func (o *ServiceAccountSchema) SetIsAPI(v bool)`

SetIsAPI sets IsAPI field to given value.

### HasIsAPI

`func (o *ServiceAccountSchema) HasIsAPI() bool`

HasIsAPI returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceAccountSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccountSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccountSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceAccountSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *ServiceAccountSchema) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccountSchema) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccountSchema) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceAccountSchema) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetImageUrl

`func (o *ServiceAccountSchema) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ServiceAccountSchema) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ServiceAccountSchema) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ServiceAccountSchema) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetInviteLink

`func (o *ServiceAccountSchema) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *ServiceAccountSchema) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *ServiceAccountSchema) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *ServiceAccountSchema) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetLoginAttempts

`func (o *ServiceAccountSchema) GetLoginAttempts() float32`

GetLoginAttempts returns the LoginAttempts field if non-nil, zero value otherwise.

### GetLoginAttemptsOk

`func (o *ServiceAccountSchema) GetLoginAttemptsOk() (*float32, bool)`

GetLoginAttemptsOk returns a tuple with the LoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttempts

`func (o *ServiceAccountSchema) SetLoginAttempts(v float32)`

SetLoginAttempts sets LoginAttempts field to given value.

### HasLoginAttempts

`func (o *ServiceAccountSchema) HasLoginAttempts() bool`

HasLoginAttempts returns a boolean if a field has been set.

### GetEmailSent

`func (o *ServiceAccountSchema) GetEmailSent() bool`

GetEmailSent returns the EmailSent field if non-nil, zero value otherwise.

### GetEmailSentOk

`func (o *ServiceAccountSchema) GetEmailSentOk() (*bool, bool)`

GetEmailSentOk returns a tuple with the EmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSent

`func (o *ServiceAccountSchema) SetEmailSent(v bool)`

SetEmailSent sets EmailSent field to given value.

### HasEmailSent

`func (o *ServiceAccountSchema) HasEmailSent() bool`

HasEmailSent returns a boolean if a field has been set.

### GetRootRole

`func (o *ServiceAccountSchema) GetRootRole() int32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *ServiceAccountSchema) GetRootRoleOk() (*int32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *ServiceAccountSchema) SetRootRole(v int32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *ServiceAccountSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### GetSeenAt

`func (o *ServiceAccountSchema) GetSeenAt() time.Time`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *ServiceAccountSchema) GetSeenAtOk() (*time.Time, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *ServiceAccountSchema) SetSeenAt(v time.Time)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *ServiceAccountSchema) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### SetSeenAtNil

`func (o *ServiceAccountSchema) SetSeenAtNil(b bool)`

 SetSeenAtNil sets the value for SeenAt to be an explicit nil

### UnsetSeenAt
`func (o *ServiceAccountSchema) UnsetSeenAt()`

UnsetSeenAt ensures that no value is present for SeenAt, not even an explicit nil
### GetCreatedAt

`func (o *ServiceAccountSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceAccountSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceAccountSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceAccountSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTokens

`func (o *ServiceAccountSchema) GetTokens() []PatSchema`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ServiceAccountSchema) GetTokensOk() (*[]PatSchema, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ServiceAccountSchema) SetTokens(v []PatSchema)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ServiceAccountSchema) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


