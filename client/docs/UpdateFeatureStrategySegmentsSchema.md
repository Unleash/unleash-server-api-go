# UpdateFeatureStrategySegmentsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | The ID of the project that the strategy belongs to. | 
**StrategyId** | **string** | The ID of the strategy to update segments for. | 
**EnvironmentId** | **string** | The ID of the strategy environment. | 
**SegmentIds** | **[]int32** | The new list of segments (IDs) to use for this strategy. Any segments not in this list will be removed from the strategy. | 

## Methods

### NewUpdateFeatureStrategySegmentsSchema

`func NewUpdateFeatureStrategySegmentsSchema(projectId string, strategyId string, environmentId string, segmentIds []int32, ) *UpdateFeatureStrategySegmentsSchema`

NewUpdateFeatureStrategySegmentsSchema instantiates a new UpdateFeatureStrategySegmentsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFeatureStrategySegmentsSchemaWithDefaults

`func NewUpdateFeatureStrategySegmentsSchemaWithDefaults() *UpdateFeatureStrategySegmentsSchema`

NewUpdateFeatureStrategySegmentsSchemaWithDefaults instantiates a new UpdateFeatureStrategySegmentsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *UpdateFeatureStrategySegmentsSchema) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateFeatureStrategySegmentsSchema) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateFeatureStrategySegmentsSchema) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetStrategyId

`func (o *UpdateFeatureStrategySegmentsSchema) GetStrategyId() string`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *UpdateFeatureStrategySegmentsSchema) GetStrategyIdOk() (*string, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *UpdateFeatureStrategySegmentsSchema) SetStrategyId(v string)`

SetStrategyId sets StrategyId field to given value.


### GetEnvironmentId

`func (o *UpdateFeatureStrategySegmentsSchema) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *UpdateFeatureStrategySegmentsSchema) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *UpdateFeatureStrategySegmentsSchema) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetSegmentIds

`func (o *UpdateFeatureStrategySegmentsSchema) GetSegmentIds() []int32`

GetSegmentIds returns the SegmentIds field if non-nil, zero value otherwise.

### GetSegmentIdsOk

`func (o *UpdateFeatureStrategySegmentsSchema) GetSegmentIdsOk() (*[]int32, bool)`

GetSegmentIdsOk returns a tuple with the SegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentIds

`func (o *UpdateFeatureStrategySegmentsSchema) SetSegmentIds(v []int32)`

SetSegmentIds sets SegmentIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


