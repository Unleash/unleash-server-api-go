# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStrategy**](StrategiesApi.md#CreateStrategy) | **Post** /api/admin/strategies | Create a strategy
[**DeprecateStrategy**](StrategiesApi.md#DeprecateStrategy) | **Post** /api/admin/strategies/{strategyName}/deprecate | Deprecate a strategy
[**GetAllStrategies**](StrategiesApi.md#GetAllStrategies) | **Get** /api/admin/strategies | Get all strategies
[**GetStrategiesByContextField**](StrategiesApi.md#GetStrategiesByContextField) | **Get** /api/admin/context/{contextField}/strategies | Get strategies that use a context field
[**GetStrategy**](StrategiesApi.md#GetStrategy) | **Get** /api/admin/strategies/{name} | Get a strategy definition
[**ReactivateStrategy**](StrategiesApi.md#ReactivateStrategy) | **Post** /api/admin/strategies/{strategyName}/reactivate | Reactivate a strategy
[**RemoveStrategy**](StrategiesApi.md#RemoveStrategy) | **Delete** /api/admin/strategies/{name} | Delete a strategy
[**UpdateFeatureStrategySegments**](StrategiesApi.md#UpdateFeatureStrategySegments) | **Post** /api/admin/segments/strategies | Update strategy segments
[**UpdateStrategy**](StrategiesApi.md#UpdateStrategy) | **Put** /api/admin/strategies/{strategyName} | Update a strategy type

# **CreateStrategy**
> StrategySchema CreateStrategy(ctx, body)
Create a strategy

Creates a strategy type based on the supplied data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStrategySchema**](CreateStrategySchema.md)| createStrategySchema | 

### Return type

[**StrategySchema**](strategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeprecateStrategy**
> DeprecateStrategy(ctx, strategyName)
Deprecate a strategy

Marks the specified strategy as deprecated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **strategyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStrategies**
> StrategiesSchema GetAllStrategies(ctx, )
Get all strategies

Retrieves all strategy types ([predefined](https://docs.getunleash.io/reference/activation-strategies \"predefined strategies\") and [custom strategies](https://docs.getunleash.io/reference/custom-activation-strategies)) that are defined on this Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StrategiesSchema**](strategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrategiesByContextField**
> ContextFieldStrategiesSchema GetStrategiesByContextField(ctx, contextField)
Get strategies that use a context field

Retrieves a list of all strategies that use the specified context field. If the context field doesn't exist, returns an empty list of strategies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextField** | **string**|  | 

### Return type

[**ContextFieldStrategiesSchema**](contextFieldStrategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrategy**
> StrategySchema GetStrategy(ctx, name)
Get a strategy definition

Retrieves the definition of the strategy specified in the URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**StrategySchema**](strategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactivateStrategy**
> ReactivateStrategy(ctx, strategyName)
Reactivate a strategy

Marks the specified strategy as not deprecated. If the strategy wasn't already deprecated, nothing changes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **strategyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveStrategy**
> RemoveStrategy(ctx, name)
Delete a strategy

Deletes the specified strategy definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatureStrategySegments**
> UpdateFeatureStrategySegmentsSchema UpdateFeatureStrategySegments(ctx, body)
Update strategy segments

Sets the segments of the strategy specified to be exactly the ones passed in the payload. Any segments that were used by the strategy before will be removed if they are not in the provided list of segments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFeatureStrategySegmentsSchema**](UpdateFeatureStrategySegmentsSchema.md)| updateFeatureStrategySegmentsSchema | 

### Return type

[**UpdateFeatureStrategySegmentsSchema**](updateFeatureStrategySegmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStrategy**
> UpdateStrategy(ctx, body, strategyName)
Update a strategy type

Updates the specified strategy type. Any properties not specified in the request body are left untouched.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateStrategySchema**](UpdateStrategySchema.md)| updateStrategySchema | 
  **strategyName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

