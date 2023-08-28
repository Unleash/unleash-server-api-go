# StateSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of the schema used to describe the state | 
**Features** | Pointer to [**[]FeatureSchema**](FeatureSchema.md) | A list of features | [optional] 
**Strategies** | Pointer to [**[]StrategySchema**](StrategySchema.md) | A list of strategies | [optional] 
**Tags** | Pointer to [**[]TagSchema**](TagSchema.md) | A list of tags | [optional] 
**TagTypes** | Pointer to [**[]TagTypeSchema**](TagTypeSchema.md) | A list of tag types | [optional] 
**FeatureTags** | Pointer to [**[]FeatureTagSchema**](FeatureTagSchema.md) | A list of tags applied to features | [optional] 
**Projects** | Pointer to [**[]ProjectSchema**](ProjectSchema.md) | A list of projects | [optional] 
**FeatureStrategies** | Pointer to [**[]FeatureStrategySchema**](FeatureStrategySchema.md) | A list of feature strategies as applied to features | [optional] 
**FeatureEnvironments** | Pointer to [**[]FeatureEnvironmentSchema**](FeatureEnvironmentSchema.md) | A list of feature environment configurations | [optional] 
**Environments** | Pointer to [**[]EnvironmentSchema**](EnvironmentSchema.md) | A list of environments | [optional] 
**Segments** | Pointer to [**[]SegmentSchema**](SegmentSchema.md) | A list of segments | [optional] 
**FeatureStrategySegments** | Pointer to [**[]FeatureStrategySegmentSchema**](FeatureStrategySegmentSchema.md) | A list of segment/strategy pairings | [optional] 

## Methods

### NewStateSchema

`func NewStateSchema(version int32, ) *StateSchema`

NewStateSchema instantiates a new StateSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateSchemaWithDefaults

`func NewStateSchemaWithDefaults() *StateSchema`

NewStateSchemaWithDefaults instantiates a new StateSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *StateSchema) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StateSchema) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StateSchema) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFeatures

`func (o *StateSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *StateSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *StateSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *StateSchema) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetStrategies

`func (o *StateSchema) GetStrategies() []StrategySchema`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *StateSchema) GetStrategiesOk() (*[]StrategySchema, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *StateSchema) SetStrategies(v []StrategySchema)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *StateSchema) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetTags

`func (o *StateSchema) GetTags() []TagSchema`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StateSchema) GetTagsOk() (*[]TagSchema, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StateSchema) SetTags(v []TagSchema)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StateSchema) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagTypes

`func (o *StateSchema) GetTagTypes() []TagTypeSchema`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *StateSchema) GetTagTypesOk() (*[]TagTypeSchema, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *StateSchema) SetTagTypes(v []TagTypeSchema)`

SetTagTypes sets TagTypes field to given value.

### HasTagTypes

`func (o *StateSchema) HasTagTypes() bool`

HasTagTypes returns a boolean if a field has been set.

### GetFeatureTags

`func (o *StateSchema) GetFeatureTags() []FeatureTagSchema`

GetFeatureTags returns the FeatureTags field if non-nil, zero value otherwise.

### GetFeatureTagsOk

`func (o *StateSchema) GetFeatureTagsOk() (*[]FeatureTagSchema, bool)`

GetFeatureTagsOk returns a tuple with the FeatureTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTags

`func (o *StateSchema) SetFeatureTags(v []FeatureTagSchema)`

SetFeatureTags sets FeatureTags field to given value.

### HasFeatureTags

`func (o *StateSchema) HasFeatureTags() bool`

HasFeatureTags returns a boolean if a field has been set.

### GetProjects

`func (o *StateSchema) GetProjects() []ProjectSchema`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *StateSchema) GetProjectsOk() (*[]ProjectSchema, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *StateSchema) SetProjects(v []ProjectSchema)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *StateSchema) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetFeatureStrategies

`func (o *StateSchema) GetFeatureStrategies() []FeatureStrategySchema`

GetFeatureStrategies returns the FeatureStrategies field if non-nil, zero value otherwise.

### GetFeatureStrategiesOk

`func (o *StateSchema) GetFeatureStrategiesOk() (*[]FeatureStrategySchema, bool)`

GetFeatureStrategiesOk returns a tuple with the FeatureStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStrategies

`func (o *StateSchema) SetFeatureStrategies(v []FeatureStrategySchema)`

SetFeatureStrategies sets FeatureStrategies field to given value.

### HasFeatureStrategies

`func (o *StateSchema) HasFeatureStrategies() bool`

HasFeatureStrategies returns a boolean if a field has been set.

### GetFeatureEnvironments

`func (o *StateSchema) GetFeatureEnvironments() []FeatureEnvironmentSchema`

GetFeatureEnvironments returns the FeatureEnvironments field if non-nil, zero value otherwise.

### GetFeatureEnvironmentsOk

`func (o *StateSchema) GetFeatureEnvironmentsOk() (*[]FeatureEnvironmentSchema, bool)`

GetFeatureEnvironmentsOk returns a tuple with the FeatureEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureEnvironments

`func (o *StateSchema) SetFeatureEnvironments(v []FeatureEnvironmentSchema)`

SetFeatureEnvironments sets FeatureEnvironments field to given value.

### HasFeatureEnvironments

`func (o *StateSchema) HasFeatureEnvironments() bool`

HasFeatureEnvironments returns a boolean if a field has been set.

### GetEnvironments

`func (o *StateSchema) GetEnvironments() []EnvironmentSchema`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *StateSchema) GetEnvironmentsOk() (*[]EnvironmentSchema, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *StateSchema) SetEnvironments(v []EnvironmentSchema)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *StateSchema) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetSegments

`func (o *StateSchema) GetSegments() []SegmentSchema`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *StateSchema) GetSegmentsOk() (*[]SegmentSchema, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *StateSchema) SetSegments(v []SegmentSchema)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *StateSchema) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetFeatureStrategySegments

`func (o *StateSchema) GetFeatureStrategySegments() []FeatureStrategySegmentSchema`

GetFeatureStrategySegments returns the FeatureStrategySegments field if non-nil, zero value otherwise.

### GetFeatureStrategySegmentsOk

`func (o *StateSchema) GetFeatureStrategySegmentsOk() (*[]FeatureStrategySegmentSchema, bool)`

GetFeatureStrategySegmentsOk returns a tuple with the FeatureStrategySegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStrategySegments

`func (o *StateSchema) SetFeatureStrategySegments(v []FeatureStrategySegmentSchema)`

SetFeatureStrategySegments sets FeatureStrategySegments field to given value.

### HasFeatureStrategySegments

`func (o *StateSchema) HasFeatureStrategySegments() bool`

HasFeatureStrategySegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


