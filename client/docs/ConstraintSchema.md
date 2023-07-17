# ConstraintSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextName** | **string** | The name of the context field that this constraint should apply to. | 
**Operator** | **string** | The operator to use when evaluating this constraint. For more information about the various operators, refer to [the strategy constraint operator documentation](https://docs.getunleash.io/reference/strategy-constraints#strategy-constraint-operators). | 
**CaseInsensitive** | Pointer to **bool** | Whether the operator should be case sensitive or not. Defaults to &#x60;false&#x60; (being case sensitive). | [optional] [default to false]
**Inverted** | Pointer to **bool** | Whether the result should be negated or not. If &#x60;true&#x60;, will turn a &#x60;true&#x60; result into a &#x60;false&#x60; result and vice versa. | [optional] [default to false]
**Values** | Pointer to **[]string** | The context values that should be used for constraint evaluation. Use this property instead of &#x60;value&#x60; for properties that accept multiple values. | [optional] 
**Value** | Pointer to **string** | The context value that should be used for constraint evaluation. Use this property instead of &#x60;values&#x60; for properties that only accept single values. | [optional] 

## Methods

### NewConstraintSchema

`func NewConstraintSchema(contextName string, operator string, ) *ConstraintSchema`

NewConstraintSchema instantiates a new ConstraintSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintSchemaWithDefaults

`func NewConstraintSchemaWithDefaults() *ConstraintSchema`

NewConstraintSchemaWithDefaults instantiates a new ConstraintSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextName

`func (o *ConstraintSchema) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *ConstraintSchema) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *ConstraintSchema) SetContextName(v string)`

SetContextName sets ContextName field to given value.


### GetOperator

`func (o *ConstraintSchema) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ConstraintSchema) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ConstraintSchema) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetCaseInsensitive

`func (o *ConstraintSchema) GetCaseInsensitive() bool`

GetCaseInsensitive returns the CaseInsensitive field if non-nil, zero value otherwise.

### GetCaseInsensitiveOk

`func (o *ConstraintSchema) GetCaseInsensitiveOk() (*bool, bool)`

GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseInsensitive

`func (o *ConstraintSchema) SetCaseInsensitive(v bool)`

SetCaseInsensitive sets CaseInsensitive field to given value.

### HasCaseInsensitive

`func (o *ConstraintSchema) HasCaseInsensitive() bool`

HasCaseInsensitive returns a boolean if a field has been set.

### GetInverted

`func (o *ConstraintSchema) GetInverted() bool`

GetInverted returns the Inverted field if non-nil, zero value otherwise.

### GetInvertedOk

`func (o *ConstraintSchema) GetInvertedOk() (*bool, bool)`

GetInvertedOk returns a tuple with the Inverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverted

`func (o *ConstraintSchema) SetInverted(v bool)`

SetInverted sets Inverted field to given value.

### HasInverted

`func (o *ConstraintSchema) HasInverted() bool`

HasInverted returns a boolean if a field has been set.

### GetValues

`func (o *ConstraintSchema) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConstraintSchema) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConstraintSchema) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ConstraintSchema) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetValue

`func (o *ConstraintSchema) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConstraintSchema) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConstraintSchema) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConstraintSchema) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


