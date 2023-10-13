# ProjectAccessConfigurationSchemaRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the role. | [optional] 
**Groups** | Pointer to **[]int32** | A list of group ids that will be assigned this role | [optional] 
**Users** | Pointer to **[]int32** | A list of user ids that will be assigned this role | [optional] 

## Methods

### NewProjectAccessConfigurationSchemaRolesInner

`func NewProjectAccessConfigurationSchemaRolesInner() *ProjectAccessConfigurationSchemaRolesInner`

NewProjectAccessConfigurationSchemaRolesInner instantiates a new ProjectAccessConfigurationSchemaRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectAccessConfigurationSchemaRolesInnerWithDefaults

`func NewProjectAccessConfigurationSchemaRolesInnerWithDefaults() *ProjectAccessConfigurationSchemaRolesInner`

NewProjectAccessConfigurationSchemaRolesInnerWithDefaults instantiates a new ProjectAccessConfigurationSchemaRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectAccessConfigurationSchemaRolesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectAccessConfigurationSchemaRolesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroups

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProjectAccessConfigurationSchemaRolesInner) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ProjectAccessConfigurationSchemaRolesInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectAccessConfigurationSchemaRolesInner) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectAccessConfigurationSchemaRolesInner) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ProjectAccessConfigurationSchemaRolesInner) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


