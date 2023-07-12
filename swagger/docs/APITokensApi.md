# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](APITokensApi.md#CreateApiToken) | **Post** /api/admin/api-tokens | Create API token
[**DeleteApiToken**](APITokensApi.md#DeleteApiToken) | **Delete** /api/admin/api-tokens/{token} | Delete API token
[**GetAllApiTokens**](APITokensApi.md#GetAllApiTokens) | **Get** /api/admin/api-tokens | Get API tokens
[**UpdateApiToken**](APITokensApi.md#UpdateApiToken) | **Put** /api/admin/api-tokens/{token} | Update API token

# **CreateApiToken**
> ApiTokenSchema CreateApiToken(ctx, body)
Create API token

Create an API token of a specific type: one of client, admin, frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateApiTokenSchema**](CreateApiTokenSchema.md)| createApiTokenSchema | 

### Return type

[**ApiTokenSchema**](apiTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiToken**
> DeleteApiToken(ctx, token)
Delete API token

Deletes an existing API token. The `token` path parameter is the token's `secret`. If the token does not exist, this endpoint returns a 200 OK, but does nothing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllApiTokens**
> ApiTokensSchema GetAllApiTokens(ctx, )
Get API tokens

Retrieves all API tokens that exist in the Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiTokensSchema**](apiTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiToken**
> UpdateApiToken(ctx, body, token)
Update API token

Updates an existing API token with a new expiry date. The `token` path parameter is the token's `secret`. If the token does not exist, this endpoint returns a 200 OK, but does nothing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateApiTokenSchema**](UpdateApiTokenSchema.md)| updateApiTokenSchema | 
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

