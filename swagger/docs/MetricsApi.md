# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](MetricsApi.md#CreateApplication) | **Post** /api/admin/metrics/applications/{appName} | Create an application to connect reported metrics
[**DeleteApplication**](MetricsApi.md#DeleteApplication) | **Delete** /api/admin/metrics/applications/{appName} | Delete an application
[**GetApplication**](MetricsApi.md#GetApplication) | **Get** /api/admin/metrics/applications/{appName} | Get application data
[**GetApplications**](MetricsApi.md#GetApplications) | **Get** /api/admin/metrics/applications | Get all applications
[**GetFeatureUsageSummary**](MetricsApi.md#GetFeatureUsageSummary) | **Get** /api/admin/client-metrics/features/{name} | Last hour of usage and a list of applications that have reported seeing this feature toggle
[**GetRawFeatureMetrics**](MetricsApi.md#GetRawFeatureMetrics) | **Get** /api/admin/client-metrics/features/{name}/raw | Feature usage metrics for the last 48 hours, grouped by hour
[**GetRequestsPerSecond**](MetricsApi.md#GetRequestsPerSecond) | **Get** /api/admin/metrics/rps | Gets usage data

# **CreateApplication**
> CreateApplication(ctx, body, appName)
Create an application to connect reported metrics

Is used to report usage as well which sdk the application uses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| createApplicationSchema | 
  **appName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplication**
> DeleteApplication(ctx, appName)
Delete an application

Delete the application specified in the request URL. Returns 200 OK if the application was successfully deleted or if it didn't exist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplication**
> ApplicationSchema GetApplication(ctx, appName)
Get application data

Returns data about the specified application (`appName`). The data contains information on the name of the application, sdkVersion (which sdk reported these metrics, typically `unleash-client-node:3.4.1` or `unleash-client-java:7.1.0`), as well as data about how to display this application in a list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appName** | **string**|  | 

### Return type

[**ApplicationSchema**](applicationSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApplications**
> ApplicationsSchema GetApplications(ctx, )
Get all applications

Returns all applications registered with Unleash. Applications can be created via metrics reporting or manual creation

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApplicationsSchema**](applicationsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureUsageSummary**
> FeatureUsageSchema GetFeatureUsageSummary(ctx, name)
Last hour of usage and a list of applications that have reported seeing this feature toggle

Separate counts for yes (enabled), no (disabled), as well as how many times each variant was selected during the last hour

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**FeatureUsageSchema**](featureUsageSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRawFeatureMetrics**
> FeatureMetricsSchema GetRawFeatureMetrics(ctx, name)
Feature usage metrics for the last 48 hours, grouped by hour

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**FeatureMetricsSchema**](featureMetricsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestsPerSecond**
> RequestsPerSecondSegmentedSchema GetRequestsPerSecond(ctx, )
Gets usage data

Gets usage data per app/endpoint from a prometheus compatible metrics endpoint

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RequestsPerSecondSegmentedSchema**](requestsPerSecondSegmentedSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

