# PlaygroundFeatureSchemaStrategies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [***AnyOfplaygroundFeatureSchemaStrategiesResult**](AnyOfplaygroundFeatureSchemaStrategiesResult.md) | The cumulative results of all the feature&#x27;s strategies. Can be &#x60;true&#x60;,                                   &#x60;false&#x60;, or &#x60;unknown&#x60;.                                   This property will only be &#x60;unknown&#x60;                                   if one or more of the strategies can&#x27;t be fully evaluated and the rest of the strategies                                   all resolve to &#x60;false&#x60;. | [default to null]
**Data** | [**[]PlaygroundStrategySchema**](playgroundStrategySchema.md) | The strategies that apply to this feature. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

