# ProjectRoleSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | The id of the project user and group count are counted for. | 
**Role** | Pointer to **int32** | Id of the role the user and group count are counted for. | [optional] 
**UserCount** | Pointer to **int32** | Number of users mapped to this project. | [optional] 
**ServiceAccountCount** | Pointer to **int32** | Number of service accounts mapped to this project. | [optional] 
**GroupCount** | Pointer to **int32** | Number of groups mapped to this project. | [optional] 

## Methods

### NewProjectRoleSchema

`func NewProjectRoleSchema(project string, ) *ProjectRoleSchema`

NewProjectRoleSchema instantiates a new ProjectRoleSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRoleSchemaWithDefaults

`func NewProjectRoleSchemaWithDefaults() *ProjectRoleSchema`

NewProjectRoleSchemaWithDefaults instantiates a new ProjectRoleSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectRoleSchema) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectRoleSchema) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectRoleSchema) SetProject(v string)`

SetProject sets Project field to given value.


### GetRole

`func (o *ProjectRoleSchema) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProjectRoleSchema) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProjectRoleSchema) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProjectRoleSchema) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUserCount

`func (o *ProjectRoleSchema) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *ProjectRoleSchema) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *ProjectRoleSchema) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *ProjectRoleSchema) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetServiceAccountCount

`func (o *ProjectRoleSchema) GetServiceAccountCount() int32`

GetServiceAccountCount returns the ServiceAccountCount field if non-nil, zero value otherwise.

### GetServiceAccountCountOk

`func (o *ProjectRoleSchema) GetServiceAccountCountOk() (*int32, bool)`

GetServiceAccountCountOk returns a tuple with the ServiceAccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountCount

`func (o *ProjectRoleSchema) SetServiceAccountCount(v int32)`

SetServiceAccountCount sets ServiceAccountCount field to given value.

### HasServiceAccountCount

`func (o *ProjectRoleSchema) HasServiceAccountCount() bool`

HasServiceAccountCount returns a boolean if a field has been set.

### GetGroupCount

`func (o *ProjectRoleSchema) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *ProjectRoleSchema) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *ProjectRoleSchema) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *ProjectRoleSchema) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


