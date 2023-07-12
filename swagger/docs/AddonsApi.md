# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddon**](AddonsApi.md#CreateAddon) | **Post** /api/admin/addons | Create a new addon
[**DeleteAddon**](AddonsApi.md#DeleteAddon) | **Delete** /api/admin/addons/{id} | Delete an addon
[**GetAddon**](AddonsApi.md#GetAddon) | **Get** /api/admin/addons/{id} | Get a specific addon
[**GetAddons**](AddonsApi.md#GetAddons) | **Get** /api/admin/addons | Get all addons and providers
[**UpdateAddon**](AddonsApi.md#UpdateAddon) | **Put** /api/admin/addons/{id} | Update an addon

# **CreateAddon**
> AddonSchema CreateAddon(ctx, body)
Create a new addon

Create an addon instance. The addon must use one of the providers available on this Unleash instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddonCreateUpdateSchema**](AddonCreateUpdateSchema.md)| addonCreateUpdateSchema | 

### Return type

[**AddonSchema**](addonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAddon**
> DeleteAddon(ctx, id)
Delete an addon

Delete the addon specified by the ID in the request path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddon**
> AddonSchema GetAddon(ctx, id)
Get a specific addon

Retrieve information about the addon whose ID matches the ID in the request URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AddonSchema**](addonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAddons**
> AddonsSchema GetAddons(ctx, )
Get all addons and providers

Retrieve all addons and providers that are defined on this Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AddonsSchema**](addonsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAddon**
> AddonSchema UpdateAddon(ctx, body, id)
Update an addon

Update the addon with a specific ID. Any fields in the update object will be updated. Properties that are not included in the update object will not be affected. To empty a property, pass `null` as that property's value.  Note: passing `null` as a value for the description property will set it to an empty string.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddonCreateUpdateSchema**](AddonCreateUpdateSchema.md)| addonCreateUpdateSchema | 
  **id** | **string**|  | 

### Return type

[**AddonSchema**](addonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

