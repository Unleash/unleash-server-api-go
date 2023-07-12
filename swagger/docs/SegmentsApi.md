# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](SegmentsApi.md#CreateSegment) | **Post** /api/admin/segments | Create a new segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /api/admin/segments/{id} | Get a segment
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /api/admin/segments | Get all segments
[**GetSegmentsByStrategyId**](SegmentsApi.md#GetSegmentsByStrategyId) | **Get** /api/admin/segments/strategies/{strategyId} | Get strategy segments
[**GetStrategiesBySegmentId**](SegmentsApi.md#GetStrategiesBySegmentId) | **Get** /api/admin/segments/{id}/strategies | Get strategies that reference segment
[**RemoveSegment**](SegmentsApi.md#RemoveSegment) | **Delete** /api/admin/segments/{id} | Deletes a segment by id
[**UpdateSegment**](SegmentsApi.md#UpdateSegment) | **Put** /api/admin/segments/{id} | Update segment by id
[**ValidateSegment**](SegmentsApi.md#ValidateSegment) | **Post** /api/admin/segments/validate | Validates if a segment name exists

# **CreateSegment**
> AdminSegmentSchema CreateSegment(ctx, body)
Create a new segment

Creates a new segment using the payload provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertSegmentSchema**](UpsertSegmentSchema.md)| upsertSegmentSchema | 

### Return type

[**AdminSegmentSchema**](adminSegmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegment**
> AdminSegmentSchema GetSegment(ctx, id)
Get a segment

Retrieves a segment based on its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AdminSegmentSchema**](adminSegmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegments**
> SegmentsSchema GetSegments(ctx, )
Get all segments

Retrieves all segments that exist in this Unleash instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SegmentsSchema**](segmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegmentsByStrategyId**
> SegmentsSchema GetSegmentsByStrategyId(ctx, strategyId)
Get strategy segments

Retrieve all segments that are referenced by the specified strategy. Returns an empty list of segments if the strategy ID doesn't exist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **strategyId** | **string**|  | 

### Return type

[**SegmentsSchema**](segmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStrategiesBySegmentId**
> AdminStrategiesSchema GetStrategiesBySegmentId(ctx, id)
Get strategies that reference segment

Retrieve all strategies that reference the specified segment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AdminStrategiesSchema**](adminStrategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveSegment**
> RemoveSegment(ctx, id)
Deletes a segment by id

Deletes a segment by its id, if not found returns a 409 error

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

# **UpdateSegment**
> UpdateSegment(ctx, id)
Update segment by id

Updates the content of the segment with the provided payload. Any fields not specified will be left untouched.

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

# **ValidateSegment**
> ValidateSegment(ctx, body)
Validates if a segment name exists

Uses the name provided in the body of the request to validate if the given name exists or not

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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

