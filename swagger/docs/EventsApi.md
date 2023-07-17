# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /api/admin/events | Get the most recent events from the Unleash instance or all events related to a project.
[**GetEventsForToggle**](EventsApi.md#GetEventsForToggle) | **Get** /api/admin/events/{featureName} | Get all events related to a specific feature toggle.
[**SearchEvents**](EventsApi.md#SearchEvents) | **Post** /api/admin/events/search | Search for events

# **GetEvents**
> EventsSchema GetEvents(ctx, optional)
Get the most recent events from the Unleash instance or all events related to a project.

Returns **the last 100** events from the Unleash instance when called without a query parameter. When called with a `project` parameter, returns **all events** for the specified project.  If the provided project does not exist, the list of events will be empty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **optional.String**| The name of the project whose events you want to retrieve | 

### Return type

[**EventsSchema**](eventsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsForToggle**
> FeatureEventsSchema GetEventsForToggle(ctx, featureName)
Get all events related to a specific feature toggle.

Returns all events related to the specified feature toggle. If the feature toggle does not exist, the list of events will be empty.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 

### Return type

[**FeatureEventsSchema**](featureEventsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchEvents**
> EventsSchema SearchEvents(ctx, body)
Search for events

Allows searching for events matching the search criteria in the request body

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchEventsSchema**](SearchEventsSchema.md)| searchEventsSchema | 

### Return type

[**EventsSchema**](eventsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

