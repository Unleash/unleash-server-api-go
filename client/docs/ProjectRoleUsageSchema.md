# ProjectRoleUsageSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]ProjectRoleSchema**](ProjectRoleSchema.md) | A collection of projects with counts of users and groups mapped to them with specified roles. | [optional] 

## Methods

### NewProjectRoleUsageSchema

`func NewProjectRoleUsageSchema() *ProjectRoleUsageSchema`

NewProjectRoleUsageSchema instantiates a new ProjectRoleUsageSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleUsageSchemaWithDefaults

`func NewProjectRoleUsageSchemaWithDefaults() *ProjectRoleUsageSchema`

NewProjectRoleUsageSchemaWithDefaults instantiates a new ProjectRoleUsageSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *ProjectRoleUsageSchema) GetProjects() []ProjectRoleSchema`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectRoleUsageSchema) GetProjectsOk() (*[]ProjectRoleSchema, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectRoleUsageSchema) SetProjects(v []ProjectRoleSchema)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectRoleUsageSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


