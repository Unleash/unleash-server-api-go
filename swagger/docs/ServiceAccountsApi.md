# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceAccountsApi.md#CreateServiceAccount) | **Post** /api/admin/service-account | Create a service account.
[**CreateServiceAccountToken**](ServiceAccountsApi.md#CreateServiceAccountToken) | **Post** /api/admin/service-account/{id}/token | Create a token for a service account.
[**DeleteServiceAccount**](ServiceAccountsApi.md#DeleteServiceAccount) | **Delete** /api/admin/service-account/{id} | Delete a service account.
[**DeleteServiceAccountToken**](ServiceAccountsApi.md#DeleteServiceAccountToken) | **Delete** /api/admin/service-account/{id}/token/{tokenId} | Delete a token for a service account.
[**GetServiceAccountTokens**](ServiceAccountsApi.md#GetServiceAccountTokens) | **Get** /api/admin/service-account/{id}/token | List all tokens for a service account.
[**GetServiceAccounts**](ServiceAccountsApi.md#GetServiceAccounts) | **Get** /api/admin/service-account | List service accounts.
[**UpdateServiceAccount**](ServiceAccountsApi.md#UpdateServiceAccount) | **Put** /api/admin/service-account/{id} | Update a service account.

# **CreateServiceAccount**
> ServiceAccountSchema CreateServiceAccount(ctx, body)
Create a service account.

Creates a new service account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateServiceAccountSchema**](CreateServiceAccountSchema.md)| createServiceAccountSchema | 

### Return type

[**ServiceAccountSchema**](serviceAccountSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceAccountToken**
> PatSchema CreateServiceAccountToken(ctx, body, id)
Create a token for a service account.

Creates a new token for the service account identified by the id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatSchema**](PatSchema.md)| patSchema | 
  **id** | **string**|  | 

### Return type

[**PatSchema**](patSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceAccount**
> DeleteServiceAccount(ctx, id)
Delete a service account.

Deletes an existing service account identified by its id.

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

# **DeleteServiceAccountToken**
> DeleteServiceAccountToken(ctx, id, tokenId)
Delete a token for a service account.

Deletes a token for the service account identified both by the service account's id and the token's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **tokenId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAccountTokens**
> PatsSchema GetServiceAccountTokens(ctx, id)
List all tokens for a service account.

Returns the list of all tokens for a service account identified by the id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**PatsSchema**](patsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAccounts**
> ServiceAccountsSchema GetServiceAccounts(ctx, )
List service accounts.

Returns the list of all service accounts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceAccountsSchema**](serviceAccountsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceAccount**
> ServiceAccountSchema UpdateServiceAccount(ctx, body, id)
Update a service account.

Updates an existing service account identified by its id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| updateServiceAccountSchema | 
  **id** | **string**|  | 

### Return type

[**ServiceAccountSchema**](serviceAccountSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

