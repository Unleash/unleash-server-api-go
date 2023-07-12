# StrategySchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | An optional title for the strategy | [optional] [default to null]
**Name** | **string** | The name (type) of the strategy | [default to null]
**DisplayName** | **string** | A human friendly name for the strategy | [default to null]
**Description** | **string** | A short description of the strategy | [default to null]
**Editable** | **bool** | Whether the strategy can be edited or not. Strategies bundled with Unleash cannot be edited. | [default to null]
**Deprecated** | **bool** |  | [default to null]
**Parameters** | [**[]StrategySchemaParameters**](strategySchema_parameters.md) | A list of relevant parameters for each strategy | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

