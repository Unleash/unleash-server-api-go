# FeatureStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A uuid for the feature strategy | [optional] 
**Name** | **string** | The name or type of strategy | 
**Title** | Pointer to **NullableString** | A descriptive title for the strategy | [optional] 
**Disabled** | Pointer to **NullableBool** | A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs | [optional] 
**FeatureName** | Pointer to **string** | The name or feature the strategy is attached to | [optional] 
**SortOrder** | Pointer to **float32** | The order of the strategy in the list | [optional] 
**Segments** | Pointer to **[]float32** | A list of segment ids attached to the strategy | [optional] 
**Constraints** | Pointer to [**[]ConstraintSchema**](ConstraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints | [optional] 
**Parameters** | Pointer to **map[string]string** | A list of parameters for a strategy | [optional] 

## Methods

### NewFeatureStrategySchema

`func NewFeatureStrategySchema(name string, ) *FeatureStrategySchema`

NewFeatureStrategySchema instantiates a new FeatureStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureStrategySchemaWithDefaults

`func NewFeatureStrategySchemaWithDefaults() *FeatureStrategySchema`

NewFeatureStrategySchemaWithDefaults instantiates a new FeatureStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureStrategySchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureStrategySchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureStrategySchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureStrategySchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FeatureStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureStrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FeatureStrategySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FeatureStrategySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FeatureStrategySchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FeatureStrategySchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FeatureStrategySchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FeatureStrategySchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDisabled

`func (o *FeatureStrategySchema) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FeatureStrategySchema) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FeatureStrategySchema) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FeatureStrategySchema) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *FeatureStrategySchema) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *FeatureStrategySchema) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetFeatureName

`func (o *FeatureStrategySchema) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *FeatureStrategySchema) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *FeatureStrategySchema) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.

### HasFeatureName

`func (o *FeatureStrategySchema) HasFeatureName() bool`

HasFeatureName returns a boolean if a field has been set.

### GetSortOrder

`func (o *FeatureStrategySchema) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FeatureStrategySchema) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FeatureStrategySchema) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FeatureStrategySchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSegments

`func (o *FeatureStrategySchema) GetSegments() []float32`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *FeatureStrategySchema) GetSegmentsOk() (*[]float32, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *FeatureStrategySchema) SetSegments(v []float32)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *FeatureStrategySchema) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetConstraints

`func (o *FeatureStrategySchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *FeatureStrategySchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *FeatureStrategySchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *FeatureStrategySchema) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetParameters

`func (o *FeatureStrategySchema) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FeatureStrategySchema) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FeatureStrategySchema) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *FeatureStrategySchema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


