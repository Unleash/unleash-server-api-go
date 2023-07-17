# ProjectSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this project | [default to null]
**Name** | **string** | The name of this project | [default to null]
**Description** | **string** | Additional information about the project | [optional] [default to null]
**Health** | **float64** | An indicator of the [project&#x27;s health](https://docs.getunleash.io/reference/technical-debt#health-rating) on a scale from 0 to 100 | [optional] [default to null]
**FeatureCount** | **float64** | The number of features this project has | [optional] [default to null]
**MemberCount** | **float64** | The number of members this project has | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Favorite** | **bool** | &#x60;true&#x60; if the project was favorited, otherwise &#x60;false&#x60;. | [optional] [default to null]
**Mode** | **string** | The project&#x27;s [collaboration mode](https://docs.getunleash.io/reference/project-collaboration-mode). Determines whether non-project members can submit change requests or not. | [optional] [default to null]
**DefaultStickiness** | **string** | A default stickiness for the project affecting the default stickiness value for variants and Gradual Rollout strategy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

