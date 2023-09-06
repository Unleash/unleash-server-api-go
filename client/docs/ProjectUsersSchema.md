# ProjectUsersSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserWithProjectRoleSchema**](UserWithProjectRoleSchema.md) | A list of users with access to this project and their role within it. | 
**Roles** | [**[]RoleSchema**](RoleSchema.md) | A list of roles that are available for this project | 

## Methods

### NewProjectUsersSchema

`func NewProjectUsersSchema(users []UserWithProjectRoleSchema, roles []RoleSchema, ) *ProjectUsersSchema`

NewProjectUsersSchema instantiates a new ProjectUsersSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUsersSchemaWithDefaults

`func NewProjectUsersSchemaWithDefaults() *ProjectUsersSchema`

NewProjectUsersSchemaWithDefaults instantiates a new ProjectUsersSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ProjectUsersSchema) GetUsers() []UserWithProjectRoleSchema`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectUsersSchema) GetUsersOk() (*[]UserWithProjectRoleSchema, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectUsersSchema) SetUsers(v []UserWithProjectRoleSchema)`

SetUsers sets Users field to given value.


### GetRoles

`func (o *ProjectUsersSchema) GetRoles() []RoleSchema`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ProjectUsersSchema) GetRolesOk() (*[]RoleSchema, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ProjectUsersSchema) SetRoles(v []RoleSchema)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


