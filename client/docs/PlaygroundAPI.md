# \PlaygroundAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedPlayground**](PlaygroundAPI.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**GetPlayground**](PlaygroundAPI.md#GetPlayground) | **Post** /api/admin/playground | Evaluate an Unleash context against a set of environments and projects.



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
    resp, r, err := apiClient.PlaygroundAPI.GetAdvancedPlayground(context.Background()).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.GetAdvancedPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdvancedPlayground`: AdvancedPlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.GetAdvancedPlayground`: %v\n", resp)
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


## GetPlayground

> PlaygroundResponseSchema GetPlayground(ctx).PlaygroundRequestSchema(playgroundRequestSchema).Execute()

Evaluate an Unleash context against a set of environments and projects.



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
    playgroundRequestSchema := *openapiclient.NewPlaygroundRequestSchema("development", *openapiclient.NewSdkContextSchema("My cool application.")) // PlaygroundRequestSchema | playgroundRequestSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaygroundAPI.GetPlayground(context.Background()).PlaygroundRequestSchema(playgroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.GetPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayground`: PlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.GetPlayground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaygroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playgroundRequestSchema** | [**PlaygroundRequestSchema**](PlaygroundRequestSchema.md) | playgroundRequestSchema | 

### Return type

[**PlaygroundResponseSchema**](PlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

