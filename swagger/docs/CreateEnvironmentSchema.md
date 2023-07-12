# CreateEnvironmentSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the environment. Must be a URL-friendly string according to [RFC 3968, section 2.3](https://www.rfc-editor.org/rfc/rfc3986#section-2.3) | [default to null]
**Type_** | **string** | The [type of environment](https://docs.getunleash.io/reference/environments#environment-types) you would like to create. Unleash officially recognizes the following values: - &#x60;development&#x60; - &#x60;test&#x60; - &#x60;preproduction&#x60; - &#x60;production&#x60;  If you pass a string that is not one of the recognized values, Unleash will accept it, but it will carry no special semantics. | [default to null]
**Enabled** | **bool** | Newly created environments are enabled by default. Set this property to &#x60;false&#x60; to create the environment in a disabled state. | [optional] [default to null]
**SortOrder** | **int32** | Defines where in the list of environments to place this environment. The list uses an ascending sort, so lower numbers are shown first. You can change this value later. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

