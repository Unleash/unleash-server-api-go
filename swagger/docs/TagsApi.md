# {{classname}}

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagToFeatures**](TagsApi.md#AddTagToFeatures) | **Put** /api/admin/projects/{projectId}/tags | 
[**CreateTag**](TagsApi.md#CreateTag) | **Post** /api/admin/tags | Create a new tag.
[**CreateTagType**](TagsApi.md#CreateTagType) | **Post** /api/admin/tag-types | Create a tag type
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /api/admin/tags/{type}/{value} | Delete a tag.
[**DeleteTagType**](TagsApi.md#DeleteTagType) | **Delete** /api/admin/tag-types/{name} | Delete a tag type
[**GetTag**](TagsApi.md#GetTag) | **Get** /api/admin/tags/{type}/{value} | Get a tag by type and value.
[**GetTagType**](TagsApi.md#GetTagType) | **Get** /api/admin/tag-types/{name} | Get a tag type
[**GetTagTypes**](TagsApi.md#GetTagTypes) | **Get** /api/admin/tag-types | Get all tag types
[**GetTags**](TagsApi.md#GetTags) | **Get** /api/admin/tags | List all tags.
[**GetTagsByType**](TagsApi.md#GetTagsByType) | **Get** /api/admin/tags/{type} | List all tags of a given type.
[**UpdateTagType**](TagsApi.md#UpdateTagType) | **Put** /api/admin/tag-types/{name} | Update a tag type
[**ValidateTagType**](TagsApi.md#ValidateTagType) | **Post** /api/admin/tag-types/validate | Validate a tag type

# **AddTagToFeatures**
> AddTagToFeatures(ctx, body, projectId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsBulkAddSchema**](TagsBulkAddSchema.md)| tagsBulkAddSchema | 
  **projectId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTag**
> TagWithVersionSchema CreateTag(ctx, body)
Create a new tag.

Create a new tag with the specified data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagSchema**](TagSchema.md)| tagSchema | 

### Return type

[**TagWithVersionSchema**](tagWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTagType**
> TagTypeSchema CreateTagType(ctx, body)
Create a tag type

Create a new tag type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagTypeSchema**](TagTypeSchema.md)| tagTypeSchema | 

### Return type

[**TagTypeSchema**](tagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTag**
> DeleteTag(ctx, type_, value)
Delete a tag.

Delete a tag by type and value. When a tag is deleted all references to the tag are removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **value** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTagType**
> DeleteTagType(ctx, name)
Delete a tag type

Deletes a tag type. If any features have tags of this type, those tags will be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTag**
> TagWithVersionSchema GetTag(ctx, type_, value)
Get a tag by type and value.

Get a tag by type and value. Can be used to check whether a given tag already exists in Unleash or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 
  **value** | **string**|  | 

### Return type

[**TagWithVersionSchema**](tagWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagType**
> TagTypeSchema GetTagType(ctx, name)
Get a tag type

Get a tag type by name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**TagTypeSchema**](tagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagTypes**
> TagTypesSchema GetTagTypes(ctx, )
Get all tag types

Get a list of all available tag types.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TagTypesSchema**](tagTypesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> TagsSchema GetTags(ctx, )
List all tags.

List all tags available in Unleash.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TagsSchema**](tagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsByType**
> TagsSchema GetTagsByType(ctx, type_)
List all tags of a given type.

List all tags of a given type. If the tag type does not exist it returns an empty list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 

### Return type

[**TagsSchema**](tagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTagType**
> UpdateTagType(ctx, body, name)
Update a tag type

Update the configuration for the specified tag type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTagTypeSchema**](UpdateTagTypeSchema.md)| updateTagTypeSchema | 
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateTagType**
> ValidateTagTypeSchema ValidateTagType(ctx, body)
Validate a tag type

Validates whether if the body of the request is a valid tag and whether the a tag type with that name already exists or not. If a tag type with the same name exists, this operation will return a 409 status code.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagTypeSchema**](TagTypeSchema.md)| tagTypeSchema | 

### Return type

[**ValidateTagTypeSchema**](validateTagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

