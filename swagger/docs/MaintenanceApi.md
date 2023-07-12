# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMaintenance**](MaintenanceApi.md#GetMaintenance) | **Get** /api/admin/maintenance | Get maintenance mode status
[**ToggleMaintenance**](MaintenanceApi.md#ToggleMaintenance) | **Post** /api/admin/maintenance | Enabled/disabled maintenance mode

# **GetMaintenance**
> MaintenanceSchema GetMaintenance(ctx, )
Get maintenance mode status

Tells you whether maintenance mode is enabled or disabled

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MaintenanceSchema**](maintenanceSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleMaintenance**
> ToggleMaintenance(ctx, body)
Enabled/disabled maintenance mode

Lets administrators put Unleash into a mostly read-only mode. While Unleash is in maintenance mode, users can not change any configuration settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ToggleMaintenanceSchema**](ToggleMaintenanceSchema.md)| toggleMaintenanceSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

