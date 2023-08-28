# ProjectAddRoleAccessSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]ProjectAddRoleAccessSchemaGroupsInner**](ProjectAddRoleAccessSchemaGroupsInner.md) | A list of groups IDs | 
**Users** | [**[]ProjectAddRoleAccessSchemaUsersInner**](ProjectAddRoleAccessSchemaUsersInner.md) | A list of user IDs | 

## Methods

### NewProjectAddRoleAccessSchema

`func NewProjectAddRoleAccessSchema(groups []ProjectAddRoleAccessSchemaGroupsInner, users []ProjectAddRoleAccessSchemaUsersInner, ) *ProjectAddRoleAccessSchema`

NewProjectAddRoleAccessSchema instantiates a new ProjectAddRoleAccessSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAddRoleAccessSchemaWithDefaults

`func NewProjectAddRoleAccessSchemaWithDefaults() *ProjectAddRoleAccessSchema`

NewProjectAddRoleAccessSchemaWithDefaults instantiates a new ProjectAddRoleAccessSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ProjectAddRoleAccessSchema) GetGroups() []ProjectAddRoleAccessSchemaGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectAddRoleAccessSchema) GetGroupsOk() (*[]ProjectAddRoleAccessSchemaGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectAddRoleAccessSchema) SetGroups(v []ProjectAddRoleAccessSchemaGroupsInner)`

SetGroups sets Groups field to given value.


### GetUsers

`func (o *ProjectAddRoleAccessSchema) GetUsers() []ProjectAddRoleAccessSchemaUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectAddRoleAccessSchema) GetUsersOk() (*[]ProjectAddRoleAccessSchemaUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectAddRoleAccessSchema) SetUsers(v []ProjectAddRoleAccessSchemaUsersInner)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


