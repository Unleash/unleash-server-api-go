# AddonCreateUpdateSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | The addon provider, such as \&quot;webhook\&quot; or \&quot;slack\&quot;. This string is **case sensitive** and maps to the provider&#x27;s &#x60;name&#x60; property.  The list of all supported providers and their parameters for a specific Unleash instance can be found by making a GET request to the &#x60;api/admin/addons&#x60; endpoint: the &#x60;providers&#x60; property of that response will contain all available providers.  The default set of providers can be found in the [addons reference documentation](https://docs.getunleash.io/reference/addons). The default supported options are: - &#x60;datadog&#x60; for [Datadog](https://docs.getunleash.io/reference/addons/datadog) - &#x60;slack&#x60; for [Slack](https://docs.getunleash.io/reference/addons/slack) - &#x60;teams&#x60; for [Microsoft Teams](https://docs.getunleash.io/reference/addons/teams) - &#x60;webhook&#x60; for [webhooks](https://docs.getunleash.io/reference/addons/webhook)  The provider you choose for your addon dictates what properties the &#x60;parameters&#x60; object needs. Refer to the documentation for each provider for more information.  | [default to null]
**Description** | **string** | A description of the addon. | [optional] [default to null]
**Enabled** | **bool** | Whether the addon should be enabled or not. | [default to null]
**Parameters** | [**map[string]Object**](.md) | Parameters for the addon provider. This object has different required and optional properties depending on the provider you choose. Consult the documentation for details. | [default to null]
**Events** | **[]string** | The event types that will trigger this specific addon. | [default to null]
**Projects** | **[]string** | The projects that this addon will listen to events from. An empty list means it will listen to events from **all** projects. | [optional] [default to null]
**Environments** | **[]string** | The list of environments that this addon will listen to events from. An empty list means it will listen to events from **all** environments. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

