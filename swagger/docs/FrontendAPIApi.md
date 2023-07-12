# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFrontendFeatures**](FrontendAPIApi.md#GetFrontendFeatures) | **Get** /api/frontend | Retrieve enabled feature toggles for the provided context.
[**RegisterFrontendClient**](FrontendAPIApi.md#RegisterFrontendClient) | **Post** /api/frontend/client/register | Register a client SDK
[**RegisterFrontendMetrics**](FrontendAPIApi.md#RegisterFrontendMetrics) | **Post** /api/frontend/client/metrics | Register client usage metrics

# **GetFrontendFeatures**
> ProxyFeaturesSchema GetFrontendFeatures(ctx, )
Retrieve enabled feature toggles for the provided context.

This endpoint returns the list of feature toggles that the proxy evaluates to enabled for the given context. Context values are provided as query parameters. If the Frontend API is disabled 404 is returned.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProxyFeaturesSchema**](proxyFeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterFrontendClient**
> RegisterFrontendClient(ctx, body)
Register a client SDK

This is for future use. Currently Frontend client registration is not supported. Returning 200 for clients that expect this status code. If the Frontend API is disabled 404 is returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProxyClientSchema**](ProxyClientSchema.md)| proxyClientSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterFrontendMetrics**
> RegisterFrontendMetrics(ctx, body)
Register client usage metrics

Registers usage metrics. Stores information about how many times each toggle was evaluated to enabled and disabled within a time frame. If provided, this operation will also store data on how many times each feature toggle's variants were displayed to the end user. If the Frontend API is disabled 404 is returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientMetricsSchema**](ClientMetricsSchema.md)| clientMetricsSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

