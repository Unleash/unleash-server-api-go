# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallImport**](ImportExportApi.md#CallImport) | **Post** /api/admin/state/import | Import state (deprecated)
[**Export**](ImportExportApi.md#Export) | **Get** /api/admin/state/export | Export state (deprecated)
[**ExportFeatures**](ImportExportApi.md#ExportFeatures) | **Post** /api/admin/features-batch/export | Export feature toggles from an environment
[**ImportToggles**](ImportExportApi.md#ImportToggles) | **Post** /api/admin/features-batch/import | Import feature toggles
[**ValidateImport**](ImportExportApi.md#ValidateImport) | **Post** /api/admin/features-batch/validate | Validate feature import data

# **CallImport**
> CallImport(ctx, body)
Import state (deprecated)

Imports state into the system. Deprecated in favor of /api/admin/features-batch/import

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| stateSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Export**
> ModelMap Export(ctx, optional)
Export state (deprecated)

Exports the current state of the system. Deprecated in favor of /api/admin/features-batch/export

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportExportApiExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportExportApiExportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| Desired export format. Must be either &#x60;json&#x60; or &#x60;yaml&#x60;. | [default to json]
 **download** | [**optional.Interface of Download**](.md)| Whether exported data should be downloaded as a file. | 
 **strategies** | [**optional.Interface of Strategies**](.md)| Whether strategies should be included in the exported data. | 
 **featureToggles** | [**optional.Interface of FeatureToggles**](.md)| Whether feature toggles should be included in the exported data. | 
 **projects** | [**optional.Interface of Projects**](.md)| Whether projects should be included in the exported data. | 
 **tags** | [**optional.Interface of Tags**](.md)| Whether tag types, tags, and feature_tags should be included in the exported data. | 
 **environments** | [**optional.Interface of Environments**](.md)| Whether environments should be included in the exported data. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportFeatures**
> ExportResultSchema ExportFeatures(ctx, body)
Export feature toggles from an environment

Exports all features listed in the `features` property from the environment specified in the request body. If set to `true`, the `downloadFile` property will let you download a file with the exported data. Otherwise, the export data is returned directly as JSON. Refer to the documentation for more information about [Unleash's export functionality](https://docs.getunleash.io/reference/deploy/environment-import-export#export).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportQuerySchema**](ExportQuerySchema.md)| exportQuerySchema | 

### Return type

[**ExportResultSchema**](exportResultSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportToggles**
> ImportToggles(ctx, body)
Import feature toggles

[Import feature toggles](https://docs.getunleash.io/reference/deploy/environment-import-export#import) into a specific project and environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ImportTogglesSchema**](ImportTogglesSchema.md)| importTogglesSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateImport**
> ImportTogglesValidateSchema ValidateImport(ctx, body)
Validate feature import data

Validates a feature toggle data set. Checks whether the data can be imported into the specified project and environment. The returned value is an object that contains errors, warnings, and permissions required to perform the import, as described in the [import documentation](https://docs.getunleash.io/reference/deploy/environment-import-export#import).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ImportTogglesSchema**](ImportTogglesSchema.md)| importTogglesSchema | 

### Return type

[**ImportTogglesValidateSchema**](importTogglesValidateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

