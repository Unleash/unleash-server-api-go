# ProjectStatsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgTimeToProdCurrentWindow** | **float32** | The average time from when a feature was created to when it was enabled in the \&quot;production\&quot; environment during the current window | 
**CreatedCurrentWindow** | **float32** | The number of feature flags created during the current window | 
**CreatedPastWindow** | **float32** | The number of feature flags created during the previous window | 
**ArchivedCurrentWindow** | **float32** | The number of feature flags that were archived during the current window | 
**ArchivedPastWindow** | **float32** | The number of feature flags that were archived during the previous window | 
**ProjectActivityCurrentWindow** | **float32** | The number of project events that occurred during the current window | 
**ProjectActivityPastWindow** | **float32** | The number of project events that occurred during the previous window | 
**ProjectMembersAddedCurrentWindow** | **float32** | The number of members that were added to the project during the current window | 

## Methods

### NewProjectStatsSchema

`func NewProjectStatsSchema(avgTimeToProdCurrentWindow float32, createdCurrentWindow float32, createdPastWindow float32, archivedCurrentWindow float32, archivedPastWindow float32, projectActivityCurrentWindow float32, projectActivityPastWindow float32, projectMembersAddedCurrentWindow float32, ) *ProjectStatsSchema`

NewProjectStatsSchema instantiates a new ProjectStatsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectStatsSchemaWithDefaults

`func NewProjectStatsSchemaWithDefaults() *ProjectStatsSchema`

NewProjectStatsSchemaWithDefaults instantiates a new ProjectStatsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgTimeToProdCurrentWindow

`func (o *ProjectStatsSchema) GetAvgTimeToProdCurrentWindow() float32`

GetAvgTimeToProdCurrentWindow returns the AvgTimeToProdCurrentWindow field if non-nil, zero value otherwise.

### GetAvgTimeToProdCurrentWindowOk

`func (o *ProjectStatsSchema) GetAvgTimeToProdCurrentWindowOk() (*float32, bool)`

GetAvgTimeToProdCurrentWindowOk returns a tuple with the AvgTimeToProdCurrentWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgTimeToProdCurrentWindow

`func (o *ProjectStatsSchema) SetAvgTimeToProdCurrentWindow(v float32)`

SetAvgTimeToProdCurrentWindow sets AvgTimeToProdCurrentWindow field to given value.


### GetCreatedCurrentWindow

`func (o *ProjectStatsSchema) GetCreatedCurrentWindow() float32`

GetCreatedCurrentWindow returns the CreatedCurrentWindow field if non-nil, zero value otherwise.

### GetCreatedCurrentWindowOk

`func (o *ProjectStatsSchema) GetCreatedCurrentWindowOk() (*float32, bool)`

GetCreatedCurrentWindowOk returns a tuple with the CreatedCurrentWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedCurrentWindow

`func (o *ProjectStatsSchema) SetCreatedCurrentWindow(v float32)`

SetCreatedCurrentWindow sets CreatedCurrentWindow field to given value.


### GetCreatedPastWindow

`func (o *ProjectStatsSchema) GetCreatedPastWindow() float32`

GetCreatedPastWindow returns the CreatedPastWindow field if non-nil, zero value otherwise.

### GetCreatedPastWindowOk

`func (o *ProjectStatsSchema) GetCreatedPastWindowOk() (*float32, bool)`

GetCreatedPastWindowOk returns a tuple with the CreatedPastWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedPastWindow

`func (o *ProjectStatsSchema) SetCreatedPastWindow(v float32)`

SetCreatedPastWindow sets CreatedPastWindow field to given value.


### GetArchivedCurrentWindow

`func (o *ProjectStatsSchema) GetArchivedCurrentWindow() float32`

GetArchivedCurrentWindow returns the ArchivedCurrentWindow field if non-nil, zero value otherwise.

### GetArchivedCurrentWindowOk

`func (o *ProjectStatsSchema) GetArchivedCurrentWindowOk() (*float32, bool)`

GetArchivedCurrentWindowOk returns a tuple with the ArchivedCurrentWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedCurrentWindow

`func (o *ProjectStatsSchema) SetArchivedCurrentWindow(v float32)`

SetArchivedCurrentWindow sets ArchivedCurrentWindow field to given value.


### GetArchivedPastWindow

`func (o *ProjectStatsSchema) GetArchivedPastWindow() float32`

GetArchivedPastWindow returns the ArchivedPastWindow field if non-nil, zero value otherwise.

### GetArchivedPastWindowOk

`func (o *ProjectStatsSchema) GetArchivedPastWindowOk() (*float32, bool)`

GetArchivedPastWindowOk returns a tuple with the ArchivedPastWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedPastWindow

`func (o *ProjectStatsSchema) SetArchivedPastWindow(v float32)`

SetArchivedPastWindow sets ArchivedPastWindow field to given value.


### GetProjectActivityCurrentWindow

`func (o *ProjectStatsSchema) GetProjectActivityCurrentWindow() float32`

GetProjectActivityCurrentWindow returns the ProjectActivityCurrentWindow field if non-nil, zero value otherwise.

### GetProjectActivityCurrentWindowOk

`func (o *ProjectStatsSchema) GetProjectActivityCurrentWindowOk() (*float32, bool)`

GetProjectActivityCurrentWindowOk returns a tuple with the ProjectActivityCurrentWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectActivityCurrentWindow

`func (o *ProjectStatsSchema) SetProjectActivityCurrentWindow(v float32)`

SetProjectActivityCurrentWindow sets ProjectActivityCurrentWindow field to given value.


### GetProjectActivityPastWindow

`func (o *ProjectStatsSchema) GetProjectActivityPastWindow() float32`

GetProjectActivityPastWindow returns the ProjectActivityPastWindow field if non-nil, zero value otherwise.

### GetProjectActivityPastWindowOk

`func (o *ProjectStatsSchema) GetProjectActivityPastWindowOk() (*float32, bool)`

GetProjectActivityPastWindowOk returns a tuple with the ProjectActivityPastWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectActivityPastWindow

`func (o *ProjectStatsSchema) SetProjectActivityPastWindow(v float32)`

SetProjectActivityPastWindow sets ProjectActivityPastWindow field to given value.


### GetProjectMembersAddedCurrentWindow

`func (o *ProjectStatsSchema) GetProjectMembersAddedCurrentWindow() float32`

GetProjectMembersAddedCurrentWindow returns the ProjectMembersAddedCurrentWindow field if non-nil, zero value otherwise.

### GetProjectMembersAddedCurrentWindowOk

`func (o *ProjectStatsSchema) GetProjectMembersAddedCurrentWindowOk() (*float32, bool)`

GetProjectMembersAddedCurrentWindowOk returns a tuple with the ProjectMembersAddedCurrentWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMembersAddedCurrentWindow

`func (o *ProjectStatsSchema) SetProjectMembersAddedCurrentWindow(v float32)`

SetProjectMembersAddedCurrentWindow sets ProjectMembersAddedCurrentWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


