# \ArchiveApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeature**](ArchiveApi.md#DeleteFeature) | **Delete** /api/admin/archive/{featureName} | Archives a feature
[**DeleteFeatures**](ArchiveApi.md#DeleteFeatures) | **Post** /api/admin/projects/{projectId}/delete | Deletes a list of features
[**GetArchivedFeatures**](ArchiveApi.md#GetArchivedFeatures) | **Get** /api/admin/archive/features | 
[**GetArchivedFeaturesByProjectId**](ArchiveApi.md#GetArchivedFeaturesByProjectId) | **Get** /api/admin/archive/features/{projectId} | 
[**ReviveFeature**](ArchiveApi.md#ReviveFeature) | **Post** /api/admin/archive/revive/{featureName} | Revives a feature
[**ReviveFeatures**](ArchiveApi.md#ReviveFeatures) | **Post** /api/admin/projects/{projectId}/revive | Revives a list of features



## DeleteFeature

> DeleteFeature(ctx, featureName).Execute()

Archives a feature



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArchiveApi.DeleteFeature(context.Background(), featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.DeleteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


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


## DeleteFeatures

> DeleteFeatures(ctx, projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()

Deletes a list of features



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    projectId := "projectId_example" // string | 
    batchFeaturesSchema := *openapiclient.NewBatchFeaturesSchema([]string{"Features_example"}) // BatchFeaturesSchema | batchFeaturesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArchiveApi.DeleteFeatures(context.Background(), projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.DeleteFeatures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchFeaturesSchema** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md) | batchFeaturesSchema | 

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


## GetArchivedFeatures

> FeaturesSchema GetArchivedFeatures(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.GetArchivedFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.GetArchivedFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchivedFeatures`: FeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `ArchiveApi.GetArchivedFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedFeaturesRequest struct via the builder pattern


### Return type

[**FeaturesSchema**](FeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArchivedFeaturesByProjectId

> FeaturesSchema GetArchivedFeaturesByProjectId(ctx, projectId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArchiveApi.GetArchivedFeaturesByProjectId(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.GetArchivedFeaturesByProjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArchivedFeaturesByProjectId`: FeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `ArchiveApi.GetArchivedFeaturesByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArchivedFeaturesByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeaturesSchema**](FeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviveFeature

> ReviveFeature(ctx, featureName).Execute()

Revives a feature



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArchiveApi.ReviveFeature(context.Background(), featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.ReviveFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviveFeatureRequest struct via the builder pattern


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


## ReviveFeatures

> ReviveFeatures(ctx, projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()

Revives a list of features



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    projectId := "projectId_example" // string | 
    batchFeaturesSchema := *openapiclient.NewBatchFeaturesSchema([]string{"Features_example"}) // BatchFeaturesSchema | batchFeaturesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArchiveApi.ReviveFeatures(context.Background(), projectId).BatchFeaturesSchema(batchFeaturesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveApi.ReviveFeatures``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReviveFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchFeaturesSchema** | [**BatchFeaturesSchema**](BatchFeaturesSchema.md) | batchFeaturesSchema | 

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

