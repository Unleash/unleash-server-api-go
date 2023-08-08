# PlaygroundStrategySchemaResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationStatus** | **string** | Signals that this strategy was evaluated successfully. | 
**Enabled** | **bool** | Whether this strategy evaluates to true or not. | 
**Variant** | Pointer to [**NullablePlaygroundStrategySchemaResultAnyOf1Variant**](PlaygroundStrategySchemaResultAnyOf1Variant.md) |  | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | The feature variants. | [optional] 

## Methods

### NewPlaygroundStrategySchemaResult

`func NewPlaygroundStrategySchemaResult(evaluationStatus string, enabled bool, ) *PlaygroundStrategySchemaResult`

NewPlaygroundStrategySchemaResult instantiates a new PlaygroundStrategySchemaResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundStrategySchemaResultWithDefaults

`func NewPlaygroundStrategySchemaResultWithDefaults() *PlaygroundStrategySchemaResult`

NewPlaygroundStrategySchemaResultWithDefaults instantiates a new PlaygroundStrategySchemaResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationStatus

`func (o *PlaygroundStrategySchemaResult) GetEvaluationStatus() string`

GetEvaluationStatus returns the EvaluationStatus field if non-nil, zero value otherwise.

### GetEvaluationStatusOk

`func (o *PlaygroundStrategySchemaResult) GetEvaluationStatusOk() (*string, bool)`

GetEvaluationStatusOk returns a tuple with the EvaluationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationStatus

`func (o *PlaygroundStrategySchemaResult) SetEvaluationStatus(v string)`

SetEvaluationStatus sets EvaluationStatus field to given value.


### GetEnabled

`func (o *PlaygroundStrategySchemaResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlaygroundStrategySchemaResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlaygroundStrategySchemaResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVariant

`func (o *PlaygroundStrategySchemaResult) GetVariant() PlaygroundStrategySchemaResultAnyOf1Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PlaygroundStrategySchemaResult) GetVariantOk() (*PlaygroundStrategySchemaResultAnyOf1Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PlaygroundStrategySchemaResult) SetVariant(v PlaygroundStrategySchemaResultAnyOf1Variant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *PlaygroundStrategySchemaResult) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### SetVariantNil

`func (o *PlaygroundStrategySchemaResult) SetVariantNil(b bool)`

 SetVariantNil sets the value for Variant to be an explicit nil

### UnsetVariant
`func (o *PlaygroundStrategySchemaResult) UnsetVariant()`

UnsetVariant ensures that no value is present for Variant, not even an explicit nil
### GetVariants

`func (o *PlaygroundStrategySchemaResult) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PlaygroundStrategySchemaResult) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PlaygroundStrategySchemaResult) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *PlaygroundStrategySchemaResult) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


