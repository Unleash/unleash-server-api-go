# AddonParameterSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the parameter as it is used in code. References to this parameter should use this value. | [default to null]
**DisplayName** | **string** | The name of the parameter as it is shown to the end user in the Admin UI. | [default to null]
**Type_** | **string** | The type of the parameter. Corresponds roughly to [HTML &#x60;input&#x60; field types](https://developer.mozilla.org/docs/Web/HTML/Element/Input#input_types). Multi-line inut fields are indicated as &#x60;textfield&#x60; (equivalent to the HTML &#x60;textarea&#x60; tag). | [default to null]
**Description** | **string** | A description of the parameter. This should explain to the end user what the parameter is used for. | [optional] [default to null]
**Placeholder** | **string** | The default value for this parameter. This value is used if no other value is provided. | [optional] [default to null]
**Required** | **bool** | Whether this parameter is required or not. If a parameter is required, you must give it a value when you create the addon. If it is not required it can be left out. It may receive a default value in those cases. | [default to null]
**Sensitive** | **bool** | Indicates whether this parameter is **sensitive** or not. Unleash will not return sensitive parameters to API requests. It will instead use a number of asterisks to indicate that a value is set, e.g. \&quot;******\&quot;. The number of asterisks does not correlate to the parameter&#x27;s value. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

