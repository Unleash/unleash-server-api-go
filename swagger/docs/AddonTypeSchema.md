# AddonTypeSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the addon type. When creating new addons, this goes in the payload&#x27;s &#x60;type&#x60; field. | [default to null]
**DisplayName** | **string** | The addon type&#x27;s name as it should be displayed in the admin UI. | [default to null]
**DocumentationUrl** | **string** | A URL to where you can find more information about using this addon type. | [default to null]
**Description** | **string** | A description of the addon type. | [default to null]
**TagTypes** | [**[]TagTypeSchema**](tagTypeSchema.md) | A list of [Unleash tag types](https://docs.getunleash.io/reference/tags#tag-types) that this addon uses. These tags will be added to the Unleash instance when an addon of this type is created. | [optional] [default to null]
**Parameters** | [**[]AddonParameterSchema**](addonParameterSchema.md) | The addon provider&#x27;s parameters. Use these to configure an addon of this provider type. Items with &#x60;required: true&#x60; must be provided. | [optional] [default to null]
**Events** | **[]string** | All the [event types](https://docs.getunleash.io/reference/api/legacy/unleash/admin/events#feature-toggle-events) that are available for this addon provider. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

