# PlaygroundStrategySchemaResultAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationStatus** | **string** | Signals that this strategy was evaluated successfully. | 
**Enabled** | **bool** | Whether this strategy evaluates to true or not. | 
**Variant** | Pointer to [**NullablePlaygroundStrategySchemaResultAnyOf1Variant**](PlaygroundStrategySchemaResultAnyOf1Variant.md) |  | [optional] 
**Variants** | Pointer to [**[]VariantSchema**](VariantSchema.md) | The feature variants. | [optional] 

## Methods

### NewPlaygroundStrategySchemaResultAnyOf1

`func NewPlaygroundStrategySchemaResultAnyOf1(evaluationStatus string, enabled bool, ) *PlaygroundStrategySchemaResultAnyOf1`

NewPlaygroundStrategySchemaResultAnyOf1 instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundStrategySchemaResultAnyOf1WithDefaults

`func NewPlaygroundStrategySchemaResultAnyOf1WithDefaults() *PlaygroundStrategySchemaResultAnyOf1`

NewPlaygroundStrategySchemaResultAnyOf1WithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationStatus

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatus() string`

GetEvaluationStatus returns the EvaluationStatus field if non-nil, zero value otherwise.

### GetEvaluationStatusOk

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetEvaluationStatusOk() (*string, bool)`

GetEvaluationStatusOk returns a tuple with the EvaluationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationStatus

`func (o *PlaygroundStrategySchemaResultAnyOf1) SetEvaluationStatus(v string)`

SetEvaluationStatus sets EvaluationStatus field to given value.


### GetEnabled

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlaygroundStrategySchemaResultAnyOf1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetVariant

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariant() PlaygroundStrategySchemaResultAnyOf1Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariantOk() (*PlaygroundStrategySchemaResultAnyOf1Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariant(v PlaygroundStrategySchemaResultAnyOf1Variant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *PlaygroundStrategySchemaResultAnyOf1) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### SetVariantNil

`func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariantNil(b bool)`

 SetVariantNil sets the value for Variant to be an explicit nil

### UnsetVariant
`func (o *PlaygroundStrategySchemaResultAnyOf1) UnsetVariant()`

UnsetVariant ensures that no value is present for Variant, not even an explicit nil
### GetVariants

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariants() []VariantSchema`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PlaygroundStrategySchemaResultAnyOf1) GetVariantsOk() (*[]VariantSchema, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PlaygroundStrategySchemaResultAnyOf1) SetVariants(v []VariantSchema)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *PlaygroundStrategySchemaResultAnyOf1) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


