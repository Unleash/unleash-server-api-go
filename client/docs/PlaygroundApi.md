# \PlaygroundApi

All URIs are relative to *https://app.unleash-hosted.com/hosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlayground**](PlaygroundApi.md#GetPlayground) | **Post** /api/admin/playground | Evaluate an Unleash context against a set of environments and projects.



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
    resp, r, err := apiClient.PlaygroundApi.GetPlayground(context.Background()).PlaygroundRequestSchema(playgroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundApi.GetPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayground`: PlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `PlaygroundApi.GetPlayground`: %v\n", resp)
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

