# \UnstableApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedPlayground**](UnstableApi.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**UpdateFeatureTypeLifetime**](UnstableApi.md#UpdateFeatureTypeLifetime) | **Put** /api/admin/feature-types/{id}/lifetime | Update feature type lifetime



## GetAdvancedPlayground

> AdvancedPlaygroundResponseSchema GetAdvancedPlayground(ctx).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()

Batch evaluate an Unleash context against a set of environments and projects.



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
    advancedPlaygroundRequestSchema := *openapiclient.NewAdvancedPlaygroundRequestSchema([]string{"Environments_example"}, *openapiclient.NewSdkContextSchema("My cool application.")) // AdvancedPlaygroundRequestSchema | advancedPlaygroundRequestSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetAdvancedPlayground(context.Background()).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetAdvancedPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdvancedPlayground`: AdvancedPlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetAdvancedPlayground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedPlaygroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedPlaygroundRequestSchema** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md) | advancedPlaygroundRequestSchema | 

### Return type

[**AdvancedPlaygroundResponseSchema**](AdvancedPlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureTypeLifetime

> FeatureTypeSchema UpdateFeatureTypeLifetime(ctx, id).UpdateFeatureTypeLifetimeSchema(updateFeatureTypeLifetimeSchema).Execute()

Update feature type lifetime



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
    id := "id_example" // string | 
    updateFeatureTypeLifetimeSchema := *openapiclient.NewUpdateFeatureTypeLifetimeSchema(NullableInt32(7)) // UpdateFeatureTypeLifetimeSchema | updateFeatureTypeLifetimeSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.UpdateFeatureTypeLifetime(context.Background(), id).UpdateFeatureTypeLifetimeSchema(updateFeatureTypeLifetimeSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.UpdateFeatureTypeLifetime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureTypeLifetime`: FeatureTypeSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.UpdateFeatureTypeLifetime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureTypeLifetimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFeatureTypeLifetimeSchema** | [**UpdateFeatureTypeLifetimeSchema**](UpdateFeatureTypeLifetimeSchema.md) | updateFeatureTypeLifetimeSchema | 

### Return type

[**FeatureTypeSchema**](FeatureTypeSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

