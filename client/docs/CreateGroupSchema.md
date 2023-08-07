# CreateGroupSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the group | 
**Description** | Pointer to **NullableString** | A custom description of the group | [optional] 
**MappingsSSO** | Pointer to **[]string** | A list of SSO groups that should map to this Unleash group | [optional] 
**RootRole** | Pointer to **NullableFloat32** | A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role. | [optional] 
**Users** | Pointer to [**[]CreateGroupSchemaUsersInner**](CreateGroupSchemaUsersInner.md) | A list of users belonging to this group | [optional] 

## Methods

### NewCreateGroupSchema

`func NewCreateGroupSchema(name string, ) *CreateGroupSchema`

NewCreateGroupSchema instantiates a new CreateGroupSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupSchemaWithDefaults

`func NewCreateGroupSchemaWithDefaults() *CreateGroupSchema`

NewCreateGroupSchemaWithDefaults instantiates a new CreateGroupSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGroupSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupSchema) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGroupSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGroupSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGroupSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGroupSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateGroupSchema) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateGroupSchema) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMappingsSSO

`func (o *CreateGroupSchema) GetMappingsSSO() []string`

GetMappingsSSO returns the MappingsSSO field if non-nil, zero value otherwise.

### GetMappingsSSOOk

`func (o *CreateGroupSchema) GetMappingsSSOOk() (*[]string, bool)`

GetMappingsSSOOk returns a tuple with the MappingsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingsSSO

`func (o *CreateGroupSchema) SetMappingsSSO(v []string)`

SetMappingsSSO sets MappingsSSO field to given value.

### HasMappingsSSO

`func (o *CreateGroupSchema) HasMappingsSSO() bool`

HasMappingsSSO returns a boolean if a field has been set.

### GetRootRole

`func (o *CreateGroupSchema) GetRootRole() float32`

GetRootRole returns the RootRole field if non-nil, zero value otherwise.

### GetRootRoleOk

`func (o *CreateGroupSchema) GetRootRoleOk() (*float32, bool)`

GetRootRoleOk returns a tuple with the RootRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootRole

`func (o *CreateGroupSchema) SetRootRole(v float32)`

SetRootRole sets RootRole field to given value.

### HasRootRole

`func (o *CreateGroupSchema) HasRootRole() bool`

HasRootRole returns a boolean if a field has been set.

### SetRootRoleNil

`func (o *CreateGroupSchema) SetRootRoleNil(b bool)`

 SetRootRoleNil sets the value for RootRole to be an explicit nil

### UnsetRootRole
`func (o *CreateGroupSchema) UnsetRootRole()`

UnsetRootRole ensures that no value is present for RootRole, not even an explicit nil
### GetUsers

`func (o *CreateGroupSchema) GetUsers() []CreateGroupSchemaUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CreateGroupSchema) GetUsersOk() (*[]CreateGroupSchemaUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CreateGroupSchema) SetUsers(v []CreateGroupSchemaUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *CreateGroupSchema) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


