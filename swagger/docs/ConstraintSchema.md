# ConstraintSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextName** | **string** | The name of the context field that this constraint should apply to. | [default to null]
**Operator** | **string** | The operator to use when evaluating this constraint. For more information about the various operators, refer to [the strategy constraint operator documentation](https://docs.getunleash.io/reference/strategy-constraints#strategy-constraint-operators). | [default to null]
**CaseInsensitive** | **bool** | Whether the operator should be case sensitive or not. Defaults to &#x60;false&#x60; (being case sensitive). | [optional] [default to false]
**Inverted** | **bool** | Whether the result should be negated or not. If &#x60;true&#x60;, will turn a &#x60;true&#x60; result into a &#x60;false&#x60; result and vice versa. | [optional] [default to false]
**Values** | **[]string** | The context values that should be used for constraint evaluation. Use this property instead of &#x60;value&#x60; for properties that accept multiple values. | [optional] [default to null]
**Value** | **string** | The context value that should be used for constraint evaluation. Use this property instead of &#x60;values&#x60; for properties that only accept single values. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

