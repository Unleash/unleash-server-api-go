# \TagsAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTagToFeatures**](TagsAPI.md#AddTagToFeatures) | **Put** /api/admin/projects/{projectId}/tags | Adds a tag to the specified features
[**CreateTag**](TagsAPI.md#CreateTag) | **Post** /api/admin/tags | Create a new tag.
[**CreateTagType**](TagsAPI.md#CreateTagType) | **Post** /api/admin/tag-types | Create a tag type
[**DeleteTag**](TagsAPI.md#DeleteTag) | **Delete** /api/admin/tags/{type}/{value} | Delete a tag.
[**DeleteTagType**](TagsAPI.md#DeleteTagType) | **Delete** /api/admin/tag-types/{name} | Delete a tag type
[**GetTag**](TagsAPI.md#GetTag) | **Get** /api/admin/tags/{type}/{value} | Get a tag by type and value.
[**GetTagType**](TagsAPI.md#GetTagType) | **Get** /api/admin/tag-types/{name} | Get a tag type
[**GetTagTypes**](TagsAPI.md#GetTagTypes) | **Get** /api/admin/tag-types | Get all tag types
[**GetTags**](TagsAPI.md#GetTags) | **Get** /api/admin/tags | List all tags.
[**GetTagsByType**](TagsAPI.md#GetTagsByType) | **Get** /api/admin/tags/{type} | List all tags of a given type.
[**UpdateTagType**](TagsAPI.md#UpdateTagType) | **Put** /api/admin/tag-types/{name} | Update a tag type
[**ValidateTagType**](TagsAPI.md#ValidateTagType) | **Post** /api/admin/tag-types/validate | Validate a tag type



## AddTagToFeatures

> AddTagToFeatures(ctx, projectId).TagsBulkAddSchema(tagsBulkAddSchema).Execute()

Adds a tag to the specified features



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    projectId := "projectId_example" // string | 
    tagsBulkAddSchema := *openapiclient.NewTagsBulkAddSchema([]string{"Features_example"}, *openapiclient.NewUpdateTagsSchema([]openapiclient.TagSchema{*openapiclient.NewTagSchema("a-tag-value", "simple")}, []openapiclient.TagSchema{*openapiclient.NewTagSchema("a-tag-value", "simple")})) // TagsBulkAddSchema | tagsBulkAddSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsAPI.AddTagToFeatures(context.Background(), projectId).TagsBulkAddSchema(tagsBulkAddSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.AddTagToFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagToFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagsBulkAddSchema** | [**TagsBulkAddSchema**](TagsBulkAddSchema.md) | tagsBulkAddSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> TagWithVersionSchema CreateTag(ctx).TagSchema(tagSchema).Execute()

Create a new tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    tagSchema := *openapiclient.NewTagSchema("a-tag-value", "simple") // TagSchema | tagSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.CreateTag(context.Background()).TagSchema(tagSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTag`: TagWithVersionSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagSchema** | [**TagSchema**](TagSchema.md) | tagSchema | 

### Return type

[**TagWithVersionSchema**](TagWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTagType

> TagTypeSchema CreateTagType(ctx).TagTypeSchema(tagTypeSchema).Execute()

Create a tag type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    tagTypeSchema := *openapiclient.NewTagTypeSchema("color") // TagTypeSchema | tagTypeSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.CreateTagType(context.Background()).TagTypeSchema(tagTypeSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTagType`: TagTypeSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTagType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagTypeSchema** | [**TagTypeSchema**](TagTypeSchema.md) | tagTypeSchema | 

### Return type

[**TagTypeSchema**](TagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, type_, value).Execute()

Delete a tag.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    type_ := "type__example" // string | 
    value := "value_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsAPI.DeleteTag(context.Background(), type_, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**value** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTagType

> DeleteTagType(ctx, name).Execute()

Delete a tag type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsAPI.DeleteTagType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> TagWithVersionSchema GetTag(ctx, type_, value).Execute()

Get a tag by type and value.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    type_ := "type__example" // string | 
    value := "value_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.GetTag(context.Background(), type_, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: TagWithVersionSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 
**value** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TagWithVersionSchema**](TagWithVersionSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagType

> TagTypeSchema GetTagType(ctx, name).Execute()

Get a tag type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.GetTagType(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagType`: TagTypeSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagTypeSchema**](TagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagTypes

> TagTypesSchema GetTagTypes(ctx).Execute()

Get all tag types



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.GetTagTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagTypes`: TagTypesSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagTypesRequest struct via the builder pattern


### Return type

[**TagTypesSchema**](TagTypesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> TagsSchema GetTags(ctx).Execute()

List all tags.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.GetTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: TagsSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


### Return type

[**TagsSchema**](TagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsByType

> TagsSchema GetTagsByType(ctx, type_).Execute()

List all tags of a given type.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    type_ := "type__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.GetTagsByType(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTagsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsByType`: TagsSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTagsByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsSchema**](TagsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTagType

> UpdateTagType(ctx, name).UpdateTagTypeSchema(updateTagTypeSchema).Execute()

Update a tag type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    name := "name_example" // string | 
    updateTagTypeSchema := *openapiclient.NewUpdateTagTypeSchema() // UpdateTagTypeSchema | updateTagTypeSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagsAPI.UpdateTagType(context.Background(), name).UpdateTagTypeSchema(updateTagTypeSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagTypeSchema** | [**UpdateTagTypeSchema**](UpdateTagTypeSchema.md) | updateTagTypeSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateTagType

> ValidateTagTypeSchema ValidateTagType(ctx).TagTypeSchema(tagTypeSchema).Execute()

Validate a tag type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    tagTypeSchema := *openapiclient.NewTagTypeSchema("color") // TagTypeSchema | tagTypeSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TagsAPI.ValidateTagType(context.Background()).TagTypeSchema(tagTypeSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ValidateTagType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateTagType`: ValidateTagTypeSchema
    fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ValidateTagType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateTagTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagTypeSchema** | [**TagTypeSchema**](TagTypeSchema.md) | tagTypeSchema | 

### Return type

[**ValidateTagTypeSchema**](ValidateTagTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

