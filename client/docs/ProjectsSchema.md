# ProjectsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  | 
**Projects** | [**[]ProjectSchema**](ProjectSchema.md) | A list of projects in the Unleash instance | 

## Methods

### NewProjectsSchema

`func NewProjectsSchema(version int32, projects []ProjectSchema, ) *ProjectsSchema`

NewProjectsSchema instantiates a new ProjectsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsSchemaWithDefaults

`func NewProjectsSchemaWithDefaults() *ProjectsSchema`

NewProjectsSchemaWithDefaults instantiates a new ProjectsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ProjectsSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectsSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectsSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetProjects

`func (o *ProjectsSchema) GetProjects() []ProjectSchema`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectsSchema) GetProjectsOk() (*[]ProjectSchema, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectsSchema) SetProjects(v []ProjectSchema)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


