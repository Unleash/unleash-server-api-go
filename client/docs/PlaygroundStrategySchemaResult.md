# PlaygroundStrategySchemaResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationStatus** | **string** | Signals that this strategy was evaluated successfully. | 
**Enabled** | **bool** | Whether this strategy evaluates to true or not. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


