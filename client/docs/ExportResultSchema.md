# ExportResultSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | [**[]FeatureSchema**](FeatureSchema.md) |  | 
**FeatureStrategies** | [**[]FeatureStrategySchema**](FeatureStrategySchema.md) |  | 
**FeatureEnvironments** | Pointer to [**[]FeatureEnvironmentSchema**](FeatureEnvironmentSchema.md) |  | [optional] 
**ContextFields** | Pointer to [**[]ContextFieldSchema**](ContextFieldSchema.md) |  | [optional] 
**FeatureTags** | Pointer to [**[]FeatureTagSchema**](FeatureTagSchema.md) |  | [optional] 
**Segments** | Pointer to [**[]ExportResultSchemaSegmentsInner**](ExportResultSchemaSegmentsInner.md) |  | [optional] 
**TagTypes** | [**[]TagTypeSchema**](TagTypeSchema.md) |  | 

## Methods

### NewExportResultSchema

`func NewExportResultSchema(features []FeatureSchema, featureStrategies []FeatureStrategySchema, tagTypes []TagTypeSchema, ) *ExportResultSchema`

NewExportResultSchema instantiates a new ExportResultSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportResultSchemaWithDefaults

`func NewExportResultSchemaWithDefaults() *ExportResultSchema`

NewExportResultSchemaWithDefaults instantiates a new ExportResultSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *ExportResultSchema) GetFeatures() []FeatureSchema`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ExportResultSchema) GetFeaturesOk() (*[]FeatureSchema, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ExportResultSchema) SetFeatures(v []FeatureSchema)`

SetFeatures sets Features field to given value.


### GetFeatureStrategies

`func (o *ExportResultSchema) GetFeatureStrategies() []FeatureStrategySchema`

GetFeatureStrategies returns the FeatureStrategies field if non-nil, zero value otherwise.

### GetFeatureStrategiesOk

`func (o *ExportResultSchema) GetFeatureStrategiesOk() (*[]FeatureStrategySchema, bool)`

GetFeatureStrategiesOk returns a tuple with the FeatureStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureStrategies

`func (o *ExportResultSchema) SetFeatureStrategies(v []FeatureStrategySchema)`

SetFeatureStrategies sets FeatureStrategies field to given value.


### GetFeatureEnvironments

`func (o *ExportResultSchema) GetFeatureEnvironments() []FeatureEnvironmentSchema`

GetFeatureEnvironments returns the FeatureEnvironments field if non-nil, zero value otherwise.

### GetFeatureEnvironmentsOk

`func (o *ExportResultSchema) GetFeatureEnvironmentsOk() (*[]FeatureEnvironmentSchema, bool)`

GetFeatureEnvironmentsOk returns a tuple with the FeatureEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureEnvironments

`func (o *ExportResultSchema) SetFeatureEnvironments(v []FeatureEnvironmentSchema)`

SetFeatureEnvironments sets FeatureEnvironments field to given value.

### HasFeatureEnvironments

`func (o *ExportResultSchema) HasFeatureEnvironments() bool`

HasFeatureEnvironments returns a boolean if a field has been set.

### GetContextFields

`func (o *ExportResultSchema) GetContextFields() []ContextFieldSchema`

GetContextFields returns the ContextFields field if non-nil, zero value otherwise.

### GetContextFieldsOk

`func (o *ExportResultSchema) GetContextFieldsOk() (*[]ContextFieldSchema, bool)`

GetContextFieldsOk returns a tuple with the ContextFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextFields

`func (o *ExportResultSchema) SetContextFields(v []ContextFieldSchema)`

SetContextFields sets ContextFields field to given value.

### HasContextFields

`func (o *ExportResultSchema) HasContextFields() bool`

HasContextFields returns a boolean if a field has been set.

### GetFeatureTags

`func (o *ExportResultSchema) GetFeatureTags() []FeatureTagSchema`

GetFeatureTags returns the FeatureTags field if non-nil, zero value otherwise.

### GetFeatureTagsOk

`func (o *ExportResultSchema) GetFeatureTagsOk() (*[]FeatureTagSchema, bool)`

GetFeatureTagsOk returns a tuple with the FeatureTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureTags

`func (o *ExportResultSchema) SetFeatureTags(v []FeatureTagSchema)`

SetFeatureTags sets FeatureTags field to given value.

### HasFeatureTags

`func (o *ExportResultSchema) HasFeatureTags() bool`

HasFeatureTags returns a boolean if a field has been set.

### GetSegments

`func (o *ExportResultSchema) GetSegments() []ExportResultSchemaSegmentsInner`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *ExportResultSchema) GetSegmentsOk() (*[]ExportResultSchemaSegmentsInner, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *ExportResultSchema) SetSegments(v []ExportResultSchemaSegmentsInner)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *ExportResultSchema) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetTagTypes

`func (o *ExportResultSchema) GetTagTypes() []TagTypeSchema`

GetTagTypes returns the TagTypes field if non-nil, zero value otherwise.

### GetTagTypesOk

`func (o *ExportResultSchema) GetTagTypesOk() (*[]TagTypeSchema, bool)`

GetTagTypesOk returns a tuple with the TagTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagTypes

`func (o *ExportResultSchema) SetTagTypes(v []TagTypeSchema)`

SetTagTypes sets TagTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


