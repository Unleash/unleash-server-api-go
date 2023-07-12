# CreateFeatureStrategySchema

## Properties

| Name            | Type                                          | Description                                                                                                        | Notes                        |
| --------------- | --------------------------------------------- | ------------------------------------------------------------------------------------------------------------------ | ---------------------------- |
| **Name**        | **string**                                    | The name of the strategy type                                                                                      | [default to null]            |
| **Title**       | **string**                                    | A descriptive title for the strategy                                                                               | [optional] [default to null] |
| **Disabled**    | **bool**                                      | A toggle to disable the strategy. defaults to false. Disabled strategies are not evaluated or returned to the SDKs | [optional] [default to null] |
| **SortOrder**   | **float64**                                   | The order of the strategy in the list                                                                              | [optional] [default to null] |
| **Constraints** | [**[]ConstraintSchema**](constraintSchema.md) | A list of the constraints attached to the strategy. See https://docs.getunleash.io/reference/strategy-constraints  | [optional] [default to null] |
| **Parameters**  | [**\*interface{}**](map.md)                   |                                                                                                                    | [optional] [default to null] |
| **Segments**    | **[]float64**                                 | Ids of segments to use for this strategy                                                                           | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
