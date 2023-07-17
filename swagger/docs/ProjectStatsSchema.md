# ProjectStatsSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgTimeToProdCurrentWindow** | **float64** | The average time from when a feature was created to when it was enabled in the \&quot;production\&quot; environment during the current window | [default to null]
**CreatedCurrentWindow** | **float64** | The number of feature toggles created during the current window | [default to null]
**CreatedPastWindow** | **float64** | The number of feature toggles created during the previous window | [default to null]
**ArchivedCurrentWindow** | **float64** | The number of feature toggles that were archived during the current window | [default to null]
**ArchivedPastWindow** | **float64** | The number of feature toggles that were archived during the previous window | [default to null]
**ProjectActivityCurrentWindow** | **float64** | The number of project events that occurred during the current window | [default to null]
**ProjectActivityPastWindow** | **float64** | The number of project events that occurred during the previous window | [default to null]
**ProjectMembersAddedCurrentWindow** | **float64** | The number of members that were added to the project during the current window | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

