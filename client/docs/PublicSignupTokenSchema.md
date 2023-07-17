# PublicSignupTokenSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | **string** | The actual value of the token. This is the part that is used by Unleash to create an invite link | 
**Url** | **string** | The public signup link for the token. Users who follow this link will be taken to a signup page where they can create an Unleash user. | 
**Name** | **string** | The token&#39;s name. Only for displaying in the UI | 
**Enabled** | **bool** | Whether the token is active. This property will always be &#x60;false&#x60; for a token that has expired. | 
**ExpiresAt** | **time.Time** | The time when the token will expire. | 
**CreatedAt** | **time.Time** | When the token was created. | 
**CreatedBy** | **NullableString** | The creator&#39;s email or username | 
**Users** | Pointer to [**[]UserSchema**](UserSchema.md) | Array of users that have signed up using the token. | [optional] 
**Role** | [**RoleSchema**](RoleSchema.md) |  | 

## Methods

### NewPublicSignupTokenSchema

`func NewPublicSignupTokenSchema(secret string, url string, name string, enabled bool, expiresAt time.Time, createdAt time.Time, createdBy NullableString, role RoleSchema, ) *PublicSignupTokenSchema`

NewPublicSignupTokenSchema instantiates a new PublicSignupTokenSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSignupTokenSchemaWithDefaults

`func NewPublicSignupTokenSchemaWithDefaults() *PublicSignupTokenSchema`

NewPublicSignupTokenSchemaWithDefaults instantiates a new PublicSignupTokenSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *PublicSignupTokenSchema) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PublicSignupTokenSchema) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PublicSignupTokenSchema) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetUrl

`func (o *PublicSignupTokenSchema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicSignupTokenSchema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicSignupTokenSchema) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *PublicSignupTokenSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicSignupTokenSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicSignupTokenSchema) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *PublicSignupTokenSchema) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PublicSignupTokenSchema) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PublicSignupTokenSchema) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExpiresAt

`func (o *PublicSignupTokenSchema) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PublicSignupTokenSchema) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PublicSignupTokenSchema) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *PublicSignupTokenSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicSignupTokenSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicSignupTokenSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PublicSignupTokenSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PublicSignupTokenSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PublicSignupTokenSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *PublicSignupTokenSchema) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PublicSignupTokenSchema) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUsers

`func (o *PublicSignupTokenSchema) GetUsers() []UserSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PublicSignupTokenSchema) GetUsersOk() (*[]UserSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PublicSignupTokenSchema) SetUsers(v []UserSchema)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PublicSignupTokenSchema) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *PublicSignupTokenSchema) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *PublicSignupTokenSchema) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetRole

`func (o *PublicSignupTokenSchema) GetRole() RoleSchema`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PublicSignupTokenSchema) GetRoleOk() (*RoleSchema, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PublicSignupTokenSchema) SetRole(v RoleSchema)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


