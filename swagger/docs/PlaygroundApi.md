# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayground**](PlaygroundApi.md#GetPlayground) | **Post** /api/admin/playground | Evaluate an Unleash context against a set of environments and projects.

# **GetPlayground**
> PlaygroundResponseSchema GetPlayground(ctx, body)
Evaluate an Unleash context against a set of environments and projects.

Use the provided `context`, `environment`, and `projects` to evaluate toggles on this Unleash instance. Returns a list of all toggles that match the parameters and what they evaluate to. The response also contains the input parameters that were provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlaygroundRequestSchema**](PlaygroundRequestSchema.md)| playgroundRequestSchema | 

### Return type

[**PlaygroundResponseSchema**](playgroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

