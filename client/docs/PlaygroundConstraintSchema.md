# PlaygroundConstraintSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextName** | **string** | The name of the context field that this constraint should apply to. | 
**Operator** | **string** | The operator to use when evaluating this constraint. For more information about the various operators, refer to [the strategy constraint operator documentation](https://docs.getunleash.io/reference/strategy-constraints#strategy-constraint-operators). | 
**CaseInsensitive** | Pointer to **bool** | Whether the operator should be case sensitive or not. Defaults to &#x60;false&#x60; (being case sensitive). | [optional] [default to false]
**Inverted** | Pointer to **bool** | Whether the result should be negated or not. If &#x60;true&#x60;, will turn a &#x60;true&#x60; result into a &#x60;false&#x60; result and vice versa. | [optional] [default to false]
**Values** | Pointer to **[]string** | The context values that should be used for constraint evaluation. Use this property instead of &#x60;value&#x60; for properties that accept multiple values. | [optional] 
**Value** | Pointer to **string** | The context value that should be used for constraint evaluation. Use this property instead of &#x60;values&#x60; for properties that only accept single values. | [optional] 
**Result** | **bool** | Whether this was evaluated as true or false. | 

## Methods

### NewPlaygroundConstraintSchema

`func NewPlaygroundConstraintSchema(contextName string, operator string, result bool, ) *PlaygroundConstraintSchema`

NewPlaygroundConstraintSchema instantiates a new PlaygroundConstraintSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundConstraintSchemaWithDefaults

`func NewPlaygroundConstraintSchemaWithDefaults() *PlaygroundConstraintSchema`

NewPlaygroundConstraintSchemaWithDefaults instantiates a new PlaygroundConstraintSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextName

`func (o *PlaygroundConstraintSchema) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *PlaygroundConstraintSchema) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *PlaygroundConstraintSchema) SetContextName(v string)`

SetContextName sets ContextName field to given value.


### GetOperator

`func (o *PlaygroundConstraintSchema) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PlaygroundConstraintSchema) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PlaygroundConstraintSchema) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetCaseInsensitive

`func (o *PlaygroundConstraintSchema) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *PlaygroundConstraintSchema) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *PlaygroundConstraintSchema) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *PlaygroundConstraintSchema) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetInverted

`func (o *PlaygroundConstraintSchema) GetInverted() bool`

GetInverted returns the Inverted field if non-nil, zero value otherwise.

### GetInvertedOk

`func (o *PlaygroundConstraintSchema) GetInvertedOk() (*bool, bool)`

GetInvertedOk returns a tuple with the Inverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverted

`func (o *PlaygroundConstraintSchema) SetInverted(v bool)`

SetInverted sets Inverted field to given value.

### HasInverted

`func (o *PlaygroundConstraintSchema) HasInverted() bool`

HasInverted returns a boolean if a field has been set.

### GetValues

`func (o *PlaygroundConstraintSchema) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PlaygroundConstraintSchema) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PlaygroundConstraintSchema) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *PlaygroundConstraintSchema) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetValue

`func (o *PlaygroundConstraintSchema) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlaygroundConstraintSchema) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlaygroundConstraintSchema) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PlaygroundConstraintSchema) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetResult

`func (o *PlaygroundConstraintSchema) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PlaygroundConstraintSchema) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PlaygroundConstraintSchema) SetResult(v bool)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


