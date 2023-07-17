# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContextField**](ContextApi.md#CreateContextField) | **Post** /api/admin/context | Create a context field
[**DeleteContextField**](ContextApi.md#DeleteContextField) | **Delete** /api/admin/context/{contextField} | Delete an existing context field
[**GetContextField**](ContextApi.md#GetContextField) | **Get** /api/admin/context/{contextField} | Gets context field
[**GetContextFields**](ContextApi.md#GetContextFields) | **Get** /api/admin/context | Gets configured context fields
[**UpdateContextField**](ContextApi.md#UpdateContextField) | **Put** /api/admin/context/{contextField} | Update an existing context field
[**Validate**](ContextApi.md#Validate) | **Post** /api/admin/context/validate | Validate a context field

# **CreateContextField**
> ContextFieldSchema CreateContextField(ctx, body)
Create a context field

Endpoint that allows creation of [custom context fields](https://docs.getunleash.io/reference/unleash-context#custom-context-fields)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertContextFieldSchema**](UpsertContextFieldSchema.md)| upsertContextFieldSchema | 

### Return type

[**ContextFieldSchema**](contextFieldSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContextField**
> DeleteContextField(ctx, contextField)
Delete an existing context field

Endpoint that allows deletion of a custom context field. Does not validate that context field is not in use, but since context field configuration is stored in a json blob for the strategy, existing strategies are safe.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextField** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextField**
> ContextFieldSchema GetContextField(ctx, contextField)
Gets context field

Returns specific [context field](https://docs.getunleash.io/reference/unleash-context) identified by the name in the path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextField** | **string**|  | 

### Return type

[**ContextFieldSchema**](contextFieldSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextFields**
> Array GetContextFields(ctx, )
Gets configured context fields

Returns all configured [Context fields](https://docs.getunleash.io/how-to/how-to-define-custom-context-fields) that have been created.

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

# **UpdateContextField**
> UpdateContextField(ctx, body, contextField)
Update an existing context field

Endpoint that allows updating a custom context field. Used to toggle stickiness and add/remove legal values for this context field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertContextFieldSchema**](UpsertContextFieldSchema.md)| upsertContextFieldSchema | 
  **contextField** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Validate**
> Validate(ctx, body)
Validate a context field

Check whether the provided data can be used to create a context field. If the data is not valid, ...?

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NameSchema**](NameSchema.md)| nameSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

