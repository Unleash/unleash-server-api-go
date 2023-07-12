# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedPlayground**](UnstableApi.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**GetLoginHistory**](UnstableApi.md#GetLoginHistory) | **Get** /api/admin/logins | Get all login events.
[**GetNotifications**](UnstableApi.md#GetNotifications) | **Get** /api/admin/notifications | Retrieves a list of notifications
[**MarkNotificationsAsRead**](UnstableApi.md#MarkNotificationsAsRead) | **Post** /api/admin/notifications/read | Mark notifications as read

# **GetAdvancedPlayground**
> AdvancedPlaygroundResponseSchema GetAdvancedPlayground(ctx, body)
Batch evaluate an Unleash context against a set of environments and projects.

Use the provided `context`, `environments`, and `projects` to evaluate toggles on this Unleash instance. You can use comma-separated values to provide multiple values to each context field. Returns a combinatorial list of all toggles that match the parameters and what they evaluate to. The response also contains the input parameters that were provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md)| advancedPlaygroundRequestSchema | 

### Return type

[**AdvancedPlaygroundResponseSchema**](advancedPlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoginHistory**
> LoginHistorySchema GetLoginHistory(ctx, )
Get all login events.

Returns **all** login events in the Unleash system. You can optionally get them in CSV format by specifying the `Accept` header as `text/csv`.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LoginHistorySchema**](loginHistorySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotifications**
> Array GetNotifications(ctx, )
Retrieves a list of notifications

A user may get relevant notifications from the projects they are a member of

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Array**](array.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkNotificationsAsRead**
> MarkNotificationsAsRead(ctx, body)
Mark notifications as read

Allow to select which notifications were read and saving a read date

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| markNotificationsAsReadSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

