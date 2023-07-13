# ProjectsOverviewSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureCount** | **float32** |  | 
**MemberCount** | **float32** |  | 
**Projects** | [**[]ProjectSchema**](ProjectSchema.md) |  | 

## Methods

### NewProjectsOverviewSchema

`func NewProjectsOverviewSchema(featureCount float32, memberCount float32, projects []ProjectSchema, ) *ProjectsOverviewSchema`

NewProjectsOverviewSchema instantiates a new ProjectsOverviewSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsOverviewSchemaWithDefaults

`func NewProjectsOverviewSchemaWithDefaults() *ProjectsOverviewSchema`

NewProjectsOverviewSchemaWithDefaults instantiates a new ProjectsOverviewSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureCount

`func (o *ProjectsOverviewSchema) GetFeatureCount() float32`

GetFeatureCount returns the FeatureCount field if non-nil, zero value otherwise.

### GetFeatureCountOk

`func (o *ProjectsOverviewSchema) GetFeatureCountOk() (*float32, bool)`

GetFeatureCountOk returns a tuple with the FeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureCount

`func (o *ProjectsOverviewSchema) SetFeatureCount(v float32)`

SetFeatureCount sets FeatureCount field to given value.


### GetMemberCount

`func (o *ProjectsOverviewSchema) GetMemberCount() float32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *ProjectsOverviewSchema) GetMemberCountOk() (*float32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *ProjectsOverviewSchema) SetMemberCount(v float32)`

SetMemberCount sets MemberCount field to given value.


### GetProjects

`func (o *ProjectsOverviewSchema) GetProjects() []ProjectSchema`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectsOverviewSchema) GetProjectsOk() (*[]ProjectSchema, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectsOverviewSchema) SetProjects(v []ProjectSchema)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


