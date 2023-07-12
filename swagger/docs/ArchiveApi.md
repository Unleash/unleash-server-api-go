# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeature**](ArchiveApi.md#DeleteFeature) | **Delete** /api/admin/archive/{featureName} | Archives a feature
[**DeleteFeatures**](ArchiveApi.md#DeleteFeatures) | **Post** /api/admin/projects/{projectId}/delete | Deletes a list of features
[**GetArchivedFeatures**](ArchiveApi.md#GetArchivedFeatures) | **Get** /api/admin/archive/features | 
[**GetArchivedFeaturesByProjectId**](ArchiveApi.md#GetArchivedFeaturesByProjectId) | **Get** /api/admin/archive/features/{projectId} | 
[**ReviveFeature**](ArchiveApi.md#ReviveFeature) | **Post** /api/admin/archive/revive/{featureName} | Revives a feature
[**ReviveFeatures**](ArchiveApi.md#ReviveFeatures) | **Post** /api/admin/projects/{projectId}/revive | Revives a list of features

# **DeleteFeature**
> DeleteFeature(ctx, featureName)
Archives a feature

This endpoint archives the specified feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatures**
> DeleteFeatures(ctx, body, projectId)
Deletes a list of features

This endpoint deletes the specified features, that are in archive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md)| batchFeaturesSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchivedFeatures**
> FeaturesSchema GetArchivedFeatures(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FeaturesSchema**](featuresSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetArchivedFeaturesByProjectId**
> FeaturesSchema GetArchivedFeaturesByProjectId(ctx, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**|  | 

### Return type

[**FeaturesSchema**](featuresSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviveFeature**
> ReviveFeature(ctx, featureName)
Revives a feature

This endpoint revives the specified feature from archive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureName** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviveFeatures**
> ReviveFeatures(ctx, body, projectId)
Revives a list of features

This endpoint revives the specified features.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md)| batchFeaturesSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

