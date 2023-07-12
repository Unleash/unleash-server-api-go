# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePat**](PersonalAccessTokensApi.md#CreatePat) | **Post** /api/admin/user/tokens | Create a new Personal Access Token.
[**DeletePat**](PersonalAccessTokensApi.md#DeletePat) | **Delete** /api/admin/user/tokens/{id} | Delete a Personal Access Token.
[**GetPats**](PersonalAccessTokensApi.md#GetPats) | **Get** /api/admin/user/tokens | Get all Personal Access Tokens for the current user.

# **CreatePat**
> PatSchema CreatePat(ctx, body)
Create a new Personal Access Token.

Creates a new [Personal Access Token](https://docs.getunleash.io/how-to/how-to-create-personal-access-tokens) for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatSchema**](PatSchema.md)| patSchema | 

### Return type

[**PatSchema**](patSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePat**
> DeletePat(ctx, id)
Delete a Personal Access Token.

This endpoint allows for deleting a [Personal Access Token](https://docs.getunleash.io/how-to/how-to-create-personal-access-tokens) belonging to the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPats**
> PatsSchema GetPats(ctx, )
Get all Personal Access Tokens for the current user.

Returns all of the [Personal Access Tokens](https://docs.getunleash.io/how-to/how-to-create-personal-access-tokens) belonging to the current user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PatsSchema**](patsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

