# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedback**](AdminUIApi.md#CreateFeedback) | **Post** /api/admin/feedback | 
[**GetUiConfig**](AdminUIApi.md#GetUiConfig) | **Get** /api/admin/ui-config | 
[**SetUiConfig**](AdminUIApi.md#SetUiConfig) | **Post** /api/admin/ui-config | 
[**UpdateFeedback**](AdminUIApi.md#UpdateFeedback) | **Put** /api/admin/feedback/{id} | 
[**UpdateSplashSettings**](AdminUIApi.md#UpdateSplashSettings) | **Post** /api/admin/splash/{id} | 

# **CreateFeedback**
> FeedbackSchema CreateFeedback(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FeedbackSchema**](FeedbackSchema.md)| feedbackSchema | 

### Return type

[**FeedbackSchema**](feedbackSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUiConfig**
> UiConfigSchema GetUiConfig(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UiConfigSchema**](uiConfigSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetUiConfig**
> SetUiConfig(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetUiConfigSchema**](SetUiConfigSchema.md)| setUiConfigSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeedback**
> FeedbackSchema UpdateFeedback(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FeedbackSchema**](FeedbackSchema.md)| feedbackSchema | 
  **id** | **string**|  | 

### Return type

[**FeedbackSchema**](feedbackSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSplashSettings**
> SplashSchema UpdateSplashSettings(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SplashSchema**](splashSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

