# HealthReportSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The project overview version. | [default to null]
**Name** | **string** | The project&#x27;s name | [default to null]
**Description** | **string** | The project&#x27;s description | [optional] [default to null]
**DefaultStickiness** | **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [default to null]
**Mode** | **string** | The project&#x27;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | [default to null]
**Members** | **int32** | The number of users/members in the project. | [default to null]
**Health** | **int32** | The overall [health rating](https://docs.getunleash.io/reference/technical-debt#health-rating) of the project. | [default to null]
**Environments** | [**[]ProjectEnvironmentSchema**](projectEnvironmentSchema.md) | An array containing the names of all the environments configured for the project. | [default to null]
**Features** | [**[]FeatureSchema**](featureSchema.md) | An array containing an overview of all the features of the project and their individual status | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | When the project was last updated. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | When the project was last updated. | [optional] [default to null]
**Favorite** | **bool** | Indicates if the project has been marked as a favorite by the current user requesting the project health overview. | [optional] [default to null]
**Stats** | [***ProjectStatsSchema**](projectStatsSchema.md) |  | [optional] [default to null]
**PotentiallyStaleCount** | **float64** | The number of potentially stale feature toggles. | [default to null]
**ActiveCount** | **float64** | The number of active feature toggles. | [default to null]
**StaleCount** | **float64** | The number of stale feature toggles. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

