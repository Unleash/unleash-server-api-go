# UserWithProjectRoleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAPI** | Pointer to **bool** | Whether this user is authenticated through Unleash tokens or logged in with a session | [optional] 
**Name** | Pointer to **string** | The name of the user | [optional] 
**Email** | Pointer to **NullableString** | The user&#39;s email address | [optional] 
**Id** | **int32** | The user&#39;s ID in the Unleash system | 
**ImageUrl** | Pointer to **NullableString** | A URL pointing to the user&#39;s image. | [optional] 
**AddedAt** | Pointer to **time.Time** | When this user was added to the project | [optional] 
**RoleId** | Pointer to **int32** | The ID of the role this user has in the given project | [optional] 
**Roles** | Pointer to **[]int32** | A list of roles this user has in the given project | [optional] 

## Methods

### NewUserWithProjectRoleSchema

`func NewUserWithProjectRoleSchema(id int32, ) *UserWithProjectRoleSchema`

NewUserWithProjectRoleSchema instantiates a new UserWithProjectRoleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithProjectRoleSchemaWithDefaults

`func NewUserWithProjectRoleSchemaWithDefaults() *UserWithProjectRoleSchema`

NewUserWithProjectRoleSchemaWithDefaults instantiates a new UserWithProjectRoleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAPI

`func (o *UserWithProjectRoleSchema) GetIsAPI() bool`

GetIsAPI returns the IsAPI field if non-nil, zero value otherwise.

### GetIsAPIOk

`func (o *UserWithProjectRoleSchema) GetIsAPIOk() (*bool, bool)`

GetIsAPIOk returns a tuple with the IsAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAPI

`func (o *UserWithProjectRoleSchema) SetIsAPI(v bool)`

SetIsAPI sets IsAPI field to given value.

### HasIsAPI

`func (o *UserWithProjectRoleSchema) HasIsAPI() bool`

HasIsAPI returns a boolean if a field has been set.

### GetName

`func (o *UserWithProjectRoleSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserWithProjectRoleSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserWithProjectRoleSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserWithProjectRoleSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserWithProjectRoleSchema) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserWithProjectRoleSchema) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserWithProjectRoleSchema) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserWithProjectRoleSchema) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserWithProjectRoleSchema) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserWithProjectRoleSchema) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *UserWithProjectRoleSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWithProjectRoleSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWithProjectRoleSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *UserWithProjectRoleSchema) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UserWithProjectRoleSchema) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UserWithProjectRoleSchema) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UserWithProjectRoleSchema) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *UserWithProjectRoleSchema) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *UserWithProjectRoleSchema) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetAddedAt

`func (o *UserWithProjectRoleSchema) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *UserWithProjectRoleSchema) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *UserWithProjectRoleSchema) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *UserWithProjectRoleSchema) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetRoleId

`func (o *UserWithProjectRoleSchema) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UserWithProjectRoleSchema) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UserWithProjectRoleSchema) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UserWithProjectRoleSchema) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoles

`func (o *UserWithProjectRoleSchema) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserWithProjectRoleSchema) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserWithProjectRoleSchema) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserWithProjectRoleSchema) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


