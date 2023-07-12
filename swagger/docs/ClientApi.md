# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllClientFeatures**](ClientApi.md#GetAllClientFeatures) | **Get** /api/client/features | Get all toggles (SDK)
[**GetClientFeature**](ClientApi.md#GetClientFeature) | **Get** /api/client/features/{featureName} | Get a single feature toggle
[**RegisterClientApplication**](ClientApi.md#RegisterClientApplication) | **Post** /api/client/register | Register a client SDK
[**RegisterClientMetrics**](ClientApi.md#RegisterClientMetrics) | **Post** /api/client/metrics | Register client usage metrics

# **GetAllClientFeatures**
> ClientFeaturesSchema GetAllClientFeatures(ctx, )
Get all toggles (SDK)

Returns the SDK configuration for all feature toggles that are available to the provided API key. Used by SDKs to configure local evaluation

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClientFeaturesSchema**](clientFeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientFeature**
> ClientFeatureSchema GetClientFeature(ctx, featureName)
Get a single feature toggle

Gets all the client data for a single toggle. Contains the exact same information about a toggle as the `/api/client/features` endpoint does, but only contains data about the specified toggle. All SDKs should use `/api/client/features`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 

### Return type

[**ClientFeatureSchema**](clientFeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterClientApplication**
> RegisterClientApplication(ctx, body)
Register a client SDK

Register a client SDK with Unleash. SDKs call this endpoint on startup to tell Unleash about their existence. Used to track custom strategies in use as well as SDK versions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientApplicationSchema**](ClientApplicationSchema.md)| clientApplicationSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterClientMetrics**
> RegisterClientMetrics(ctx, body)
Register client usage metrics

Registers usage metrics. Stores information about how many times each toggle was evaluated to enabled and disabled within a time frame. If provided, this operation will also store data on how many times each feature toggle's variants were displayed to the end user.

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

