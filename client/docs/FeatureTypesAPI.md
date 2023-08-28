# \FeatureTypesAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllFeatureTypes**](FeatureTypesAPI.md#GetAllFeatureTypes) | **Get** /api/admin/feature-types | Get all feature types
[**UpdateFeatureTypeLifetime**](FeatureTypesAPI.md#UpdateFeatureTypeLifetime) | **Put** /api/admin/feature-types/{id}/lifetime | Update feature type lifetime



## GetAllFeatureTypes

> FeatureTypesSchema GetAllFeatureTypes(ctx).Execute()

Get all feature types



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
    resp, r, err := apiClient.FeatureTypesAPI.GetAllFeatureTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureTypesAPI.GetAllFeatureTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFeatureTypes`: FeatureTypesSchema
    fmt.Fprintf(os.Stdout, "Response from `FeatureTypesAPI.GetAllFeatureTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFeatureTypesRequest struct via the builder pattern


### Return type

[**FeatureTypesSchema**](FeatureTypesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := apiClient.FeatureTypesAPI.UpdateFeatureTypeLifetime(context.Background(), id).UpdateFeatureTypeLifetimeSchema(updateFeatureTypeLifetimeSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureTypesAPI.UpdateFeatureTypeLifetime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureTypeLifetime`: FeatureTypeSchema
    fmt.Fprintf(os.Stdout, "Response from `FeatureTypesAPI.UpdateFeatureTypeLifetime`: %v\n", resp)
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

