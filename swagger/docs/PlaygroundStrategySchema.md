# PlaygroundStrategySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The strategy&#x27;s name. | [default to null]
**Title** | **string** | Description of the feature&#x27;s purpose. | [optional] [default to null]
**Id** | **string** | The strategy&#x27;s id. | [default to null]
**Result** | [***AnyOfplaygroundStrategySchemaResult**](AnyOfplaygroundStrategySchemaResult.md) | The strategy&#x27;s evaluation result. If the strategy is a custom strategy that Unleash can&#x27;t evaluate, &#x60;evaluationStatus&#x60; will be &#x60;unknown&#x60;. Otherwise, it will be &#x60;true&#x60; or &#x60;false&#x60; | [default to null]
**Disabled** | **bool** | The strategy&#x27;s status. Disabled strategies are not evaluated | [default to null]
**Segments** | [**[]PlaygroundSegmentSchema**](playgroundSegmentSchema.md) | The strategy&#x27;s segments and their evaluation results. | [default to null]
**Constraints** | [**[]PlaygroundConstraintSchema**](playgroundConstraintSchema.md) | The strategy&#x27;s constraints and their evaluation results. | [default to null]
**Parameters** | [***ModelMap**](map.md) |  | [default to null]
**Links** | [***PlaygroundStrategySchemaLinks**](playgroundStrategySchema_links.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

