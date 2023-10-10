# GroupWithProjectRoleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the group | [optional] 
**Id** | **int32** | The group&#39;s ID in the Unleash system | 
**AddedAt** | Pointer to **time.Time** | When this group was added to the project | [optional] 
**RoleId** | Pointer to **int32** | The ID of the role this group has in the given project | [optional] 
**Roles** | Pointer to **[]int32** | A list of roles this user has in the given project | [optional] 
**Description** | Pointer to **NullableString** | A custom description of the group | [optional] 
**MappingsSSO** | Pointer to **[]string** | A list of SSO groups that should map to this Unleash group | [optional] 
**RootRole** | Pointer to **NullableFloat32** | A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role. | [optional] 
**CreatedBy** | Pointer to **NullableString** | A user who created this group | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When was this group created | [optional] 
**Users** | Pointer to [**[]GroupUserModelSchema**](GroupUserModelSchema.md) | A list of users belonging to this group | [optional] 

## Methods

### NewGroupWithProjectRoleSchema

`func NewGroupWithProjectRoleSchema(id int32, ) *GroupWithProjectRoleSchema`

NewGroupWithProjectRoleSchema instantiates a new GroupWithProjectRoleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithProjectRoleSchemaWithDefaults

`func NewGroupWithProjectRoleSchemaWithDefaults() *GroupWithProjectRoleSchema`

NewGroupWithProjectRoleSchemaWithDefaults instantiates a new GroupWithProjectRoleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupWithProjectRoleSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupWithProjectRoleSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupWithProjectRoleSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupWithProjectRoleSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *GroupWithProjectRoleSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupWithProjectRoleSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupWithProjectRoleSchema) SetId(v int32)`

SetId sets Id field to given value.


### GetAddedAt

`func (o *GroupWithProjectRoleSchema) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *GroupWithProjectRoleSchema) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *GroupWithProjectRoleSchema) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *GroupWithProjectRoleSchema) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetRoleId

`func (o *GroupWithProjectRoleSchema) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GroupWithProjectRoleSchema) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GroupWithProjectRoleSchema) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *GroupWithProjectRoleSchema) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoles

`func (o *GroupWithProjectRoleSchema) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupWithProjectRoleSchema) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupWithProjectRoleSchema) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GroupWithProjectRoleSchema) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetDescription

`func (o *GroupWithProjectRoleSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupWithProjectRoleSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupWithProjectRoleSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupWithProjectRoleSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GroupWithProjectRoleSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GroupWithProjectRoleSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMappingsSSO

`func (o *GroupWithProjectRoleSchema) GetMappingsSSO() []string`

GetMappingsSSO returns the MappingsSSO field if non-nil, zero value otherwise.

### GetMappingsSSOOk

`func (o *GroupWithProjectRoleSchema) GetMappingsSSOOk() (*[]string, bool)`

GetMappingsSSOOk returns a tuple with the MappingsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingsSSO

`func (o *GroupWithProjectRoleSchema) SetMappingsSSO(v []string)`

SetMappingsSSO sets MappingsSSO field to given value.

### HasMappingsSSO

`func (o *GroupWithProjectRoleSchema) HasMappingsSSO() bool`

HasMappingsSSO returns a boolean if a field has been set.

### GetRootRole

`func (o *GroupWithProjectRoleSchema) GetRootRole() float32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *GroupWithProjectRoleSchema) GetRootRoleOk() (*float32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *GroupWithProjectRoleSchema) SetRootRole(v float32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *GroupWithProjectRoleSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### SetRootRoleNil

`func (o *GroupWithProjectRoleSchema) SetRootRoleNil(b bool)`

 SetRootRoleNil sets the value for RootRole to be an explicit nil

### UnsetRootRole
`func (o *GroupWithProjectRoleSchema) UnsetRootRole()`

UnsetRootRole ensures that no value is present for RootRole, not even an explicit nil
### GetCreatedBy

`func (o *GroupWithProjectRoleSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GroupWithProjectRoleSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GroupWithProjectRoleSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GroupWithProjectRoleSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GroupWithProjectRoleSchema) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GroupWithProjectRoleSchema) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *GroupWithProjectRoleSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupWithProjectRoleSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupWithProjectRoleSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupWithProjectRoleSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupWithProjectRoleSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupWithProjectRoleSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUsers

`func (o *GroupWithProjectRoleSchema) GetUsers() []GroupUserModelSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupWithProjectRoleSchema) GetUsersOk() (*[]GroupUserModelSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupWithProjectRoleSchema) SetUsers(v []GroupUserModelSchema)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupWithProjectRoleSchema) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


