# EdgeTokenSchema

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | **[]string** | The list of projects this token has access to. If the token has access to specific projects they will be listed here. If the token has access to all projects it will be represented as [&#x60;*&#x60;] | [default to null]
**Type_** | **string** | The [API token](https://docs.getunleash.io/reference/api-tokens-and-client-keys#api-tokens)&#x27;s **type**. Unleash supports three different types of API tokens ([ADMIN](https://docs.getunleash.io/reference/api-tokens-and-client-keys#admin-tokens), [CLIENT](https://docs.getunleash.io/reference/api-tokens-and-client-keys#client-tokens), [FRONTEND](https://docs.getunleash.io/reference/api-tokens-and-client-keys#front-end-tokens)). They all have varying access, so when validating a token it&#x27;s important to know what kind you&#x27;re dealing with | [default to null]
**Token** | **string** | The actual token value. [Unleash API tokens](https://docs.getunleash.io/reference/api-tokens-and-client-keys) are comprised of three parts. &lt;project(s)&gt;:&lt;environment&gt;.randomcharacters | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

