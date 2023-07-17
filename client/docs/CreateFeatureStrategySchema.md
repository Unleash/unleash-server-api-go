# CreateFeatureStrategySchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the strategy type | 
**Title** | Pointer to **NullableString** | A descriptive title for the strategy | [optional] 
**Disabled** | Pointer to **NullableBool** | A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs | [optional] 
**SortOrder** | Pointer to **float32** | The order of the strategy in the list | [optional] 
**Constraints** | Pointer to [**[]ConstraintSchema**](ConstraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints | [optional] 
**Variants** | Pointer to [**[]CreateStrategyVariantSchema**](CreateStrategyVariantSchema.md) | Strategy level variants | [optional] 
**Parameters** | Pointer to **map[string]string** | A list of parameters for a strategy | [optional] 
**Segments** | Pointer to **[]float32** | Ids of segments to use for this strategy | [optional] 

## Methods

### NewCreateFeatureStrategySchema

`func NewCreateFeatureStrategySchema(name string, ) *CreateFeatureStrategySchema`

NewCreateFeatureStrategySchema instantiates a new CreateFeatureStrategySchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFeatureStrategySchemaWithDefaults

`func NewCreateFeatureStrategySchemaWithDefaults() *CreateFeatureStrategySchema`

NewCreateFeatureStrategySchemaWithDefaults instantiates a new CreateFeatureStrategySchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateFeatureStrategySchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFeatureStrategySchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFeatureStrategySchema) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateFeatureStrategySchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFeatureStrategySchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFeatureStrategySchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateFeatureStrategySchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateFeatureStrategySchema) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateFeatureStrategySchema) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDisabled

`func (o *CreateFeatureStrategySchema) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateFeatureStrategySchema) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateFeatureStrategySchema) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateFeatureStrategySchema) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *CreateFeatureStrategySchema) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *CreateFeatureStrategySchema) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetSortOrder

`func (o *CreateFeatureStrategySchema) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CreateFeatureStrategySchema) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CreateFeatureStrategySchema) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CreateFeatureStrategySchema) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetConstraints

`func (o *CreateFeatureStrategySchema) GetConstraints() []ConstraintSchema`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CreateFeatureStrategySchema) GetConstraintsOk() (*[]ConstraintSchema, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CreateFeatureStrategySchema) SetConstraints(v []ConstraintSchema)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *CreateFeatureStrategySchema) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetVariants

`func (o *CreateFeatureStrategySchema) GetVariants() []CreateStrategyVariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *CreateFeatureStrategySchema) GetVariantsOk() (*[]CreateStrategyVariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *CreateFeatureStrategySchema) SetVariants(v []CreateStrategyVariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *CreateFeatureStrategySchema) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetParameters

`func (o *CreateFeatureStrategySchema) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateFeatureStrategySchema) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateFeatureStrategySchema) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateFeatureStrategySchema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSegments

`func (o *CreateFeatureStrategySchema) GetSegments() []float32`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *CreateFeatureStrategySchema) GetSegmentsOk() (*[]float32, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *CreateFeatureStrategySchema) SetSegments(v []float32)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *CreateFeatureStrategySchema) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


