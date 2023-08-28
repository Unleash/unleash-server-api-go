# ProjectAddAccessSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]int32** | A list of role IDs | 
**Groups** | **[]int32** | A list of group IDs | 
**Users** | **[]int32** | A list of user IDs | 

## Methods

### NewProjectAddAccessSchema

`func NewProjectAddAccessSchema(roles []int32, groups []int32, users []int32, ) *ProjectAddAccessSchema`

NewProjectAddAccessSchema instantiates a new ProjectAddAccessSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAddAccessSchemaWithDefaults

`func NewProjectAddAccessSchemaWithDefaults() *ProjectAddAccessSchema`

NewProjectAddAccessSchemaWithDefaults instantiates a new ProjectAddAccessSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ProjectAddAccessSchema) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ProjectAddAccessSchema) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ProjectAddAccessSchema) SetRoles(v []int32)`

SetRoles sets Roles field to given value.


### GetGroups

`func (o *ProjectAddAccessSchema) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectAddAccessSchema) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectAddAccessSchema) SetGroups(v []int32)`

SetGroups sets Groups field to given value.


### GetUsers

`func (o *ProjectAddAccessSchema) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectAddAccessSchema) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectAddAccessSchema) SetUsers(v []int32)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


