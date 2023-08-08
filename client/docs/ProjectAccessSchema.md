# ProjectAccessSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]GroupSchema**](GroupSchema.md) |  | 
**Users** | [**[]UserWithProjectRoleSchema**](UserWithProjectRoleSchema.md) |  | 
**Roles** | [**[]RoleSchema**](RoleSchema.md) |  | 

## Methods

### NewProjectAccessSchema

`func NewProjectAccessSchema(groups []GroupSchema, users []UserWithProjectRoleSchema, roles []RoleSchema, ) *ProjectAccessSchema`

NewProjectAccessSchema instantiates a new ProjectAccessSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAccessSchemaWithDefaults

`func NewProjectAccessSchemaWithDefaults() *ProjectAccessSchema`

NewProjectAccessSchemaWithDefaults instantiates a new ProjectAccessSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ProjectAccessSchema) GetGroups() []GroupSchema`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectAccessSchema) GetGroupsOk() (*[]GroupSchema, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectAccessSchema) SetGroups(v []GroupSchema)`

SetGroups sets Groups field to given value.


### GetUsers

`func (o *ProjectAccessSchema) GetUsers() []UserWithProjectRoleSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectAccessSchema) GetUsersOk() (*[]UserWithProjectRoleSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectAccessSchema) SetUsers(v []UserWithProjectRoleSchema)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *ProjectAccessSchema) GetRoles() []RoleSchema`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ProjectAccessSchema) GetRolesOk() (*[]RoleSchema, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ProjectAccessSchema) SetRoles(v []RoleSchema)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


