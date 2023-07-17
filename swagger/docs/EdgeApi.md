# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkMetrics**](EdgeApi.md#BulkMetrics) | **Post** /edge/metrics | Send metrics from Edge
[**GetValidTokens**](EdgeApi.md#GetValidTokens) | **Post** /edge/validate | Check which tokens are valid

# **BulkMetrics**
> BulkMetrics(ctx, body)
Send metrics from Edge

This operation accepts batched metrics from Edge. Metrics will be inserted into Unleash's metrics storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkMetricsSchema**](BulkMetricsSchema.md)| bulkMetricsSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetValidTokens**
> ValidatedEdgeTokensSchema GetValidTokens(ctx, body)
Check which tokens are valid

This operation accepts a list of tokens to validate. Unleash will validate each token you provide. For each valid token you provide, Unleash will return the token along with its type and which projects it has access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)| tokenStringListSchema | 

### Return type

[**ValidatedEdgeTokensSchema**](validatedEdgeTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

