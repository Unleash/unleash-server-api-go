# GroupSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The group id | [optional] 
**Name** | **string** | The name of the group | 
**Description** | Pointer to **NullableString** | A custom description of the group | [optional] 
**MappingsSSO** | Pointer to **[]string** | A list of SSO groups that should map to this Unleash group | [optional] 
**RootRole** | Pointer to **NullableFloat32** | A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role. | [optional] 
**CreatedBy** | Pointer to **NullableString** | A user who created this group | [optional] 
**CreatedAt** | Pointer to **NullableTime** | When was this group created | [optional] 
**Users** | Pointer to [**[]GroupUserModelSchema**](GroupUserModelSchema.md) | A list of users belonging to this group | [optional] 
**Projects** | Pointer to **[]string** | A list of projects where this group is used | [optional] 
**UserCount** | Pointer to **int32** | The number of users that belong to this group | [optional] 

## Methods

### NewGroupSchema

`func NewGroupSchema(name string, ) *GroupSchema`

NewGroupSchema instantiates a new GroupSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSchemaWithDefaults

`func NewGroupSchemaWithDefaults() *GroupSchema`

NewGroupSchemaWithDefaults instantiates a new GroupSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupSchema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupSchema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupSchema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GroupSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GroupSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GroupSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GroupSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GroupSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMappingsSSO

`func (o *GroupSchema) GetMappingsSSO() []string`

GetMappingsSSO returns the MappingsSSO field if non-nil, zero value otherwise.

### GetMappingsSSOOk

`func (o *GroupSchema) GetMappingsSSOOk() (*[]string, bool)`

GetMappingsSSOOk returns a tuple with the MappingsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingsSSO

`func (o *GroupSchema) SetMappingsSSO(v []string)`

SetMappingsSSO sets MappingsSSO field to given value.

### HasMappingsSSO

`func (o *GroupSchema) HasMappingsSSO() bool`

HasMappingsSSO returns a boolean if a field has been set.

### GetRootRole

`func (o *GroupSchema) GetRootRole() float32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *GroupSchema) GetRootRoleOk() (*float32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *GroupSchema) SetRootRole(v float32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *GroupSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### SetRootRoleNil

`func (o *GroupSchema) SetRootRoleNil(b bool)`

 SetRootRoleNil sets the value for RootRole to be an explicit nil

### UnsetRootRole
`func (o *GroupSchema) UnsetRootRole()`

UnsetRootRole ensures that no value is present for RootRole, not even an explicit nil
### GetCreatedBy

`func (o *GroupSchema) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GroupSchema) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GroupSchema) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GroupSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GroupSchema) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GroupSchema) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *GroupSchema) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupSchema) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupSchema) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupSchema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupSchema) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupSchema) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUsers

`func (o *GroupSchema) GetUsers() []GroupUserModelSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupSchema) GetUsersOk() (*[]GroupUserModelSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupSchema) SetUsers(v []GroupUserModelSchema)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupSchema) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetProjects

`func (o *GroupSchema) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GroupSchema) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GroupSchema) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GroupSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetUserCount

`func (o *GroupSchema) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *GroupSchema) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *GroupSchema) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *GroupSchema) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


