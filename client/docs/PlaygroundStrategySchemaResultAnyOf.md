# PlaygroundStrategySchemaResultAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluationStatus** | **string** | Signals that this strategy could not be evaluated. This is most likely because you&#39;re using a custom strategy that Unleash doesn&#39;t know about. | 
**Enabled** | [**PlaygroundStrategySchemaResultAnyOfEnabled**](PlaygroundStrategySchemaResultAnyOfEnabled.md) |  | 

## Methods

### NewPlaygroundStrategySchemaResultAnyOf

`func NewPlaygroundStrategySchemaResultAnyOf(evaluationStatus string, enabled PlaygroundStrategySchemaResultAnyOfEnabled, ) *PlaygroundStrategySchemaResultAnyOf`

NewPlaygroundStrategySchemaResultAnyOf instantiates a new PlaygroundStrategySchemaResultAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundStrategySchemaResultAnyOfWithDefaults

`func NewPlaygroundStrategySchemaResultAnyOfWithDefaults() *PlaygroundStrategySchemaResultAnyOf`

NewPlaygroundStrategySchemaResultAnyOfWithDefaults instantiates a new PlaygroundStrategySchemaResultAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluationStatus

`func (o *PlaygroundStrategySchemaResultAnyOf) GetEvaluationStatus() string`

GetEvaluationStatus returns the EvaluationStatus field if non-nil, zero value otherwise.

### GetEvaluationStatusOk

`func (o *PlaygroundStrategySchemaResultAnyOf) GetEvaluationStatusOk() (*string, bool)`

GetEvaluationStatusOk returns a tuple with the EvaluationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationStatus

`func (o *PlaygroundStrategySchemaResultAnyOf) SetEvaluationStatus(v string)`

SetEvaluationStatus sets EvaluationStatus field to given value.


### GetEnabled

`func (o *PlaygroundStrategySchemaResultAnyOf) GetEnabled() PlaygroundStrategySchemaResultAnyOfEnabled`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlaygroundStrategySchemaResultAnyOf) GetEnabledOk() (*PlaygroundStrategySchemaResultAnyOfEnabled, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlaygroundStrategySchemaResultAnyOf) SetEnabled(v PlaygroundStrategySchemaResultAnyOfEnabled)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


