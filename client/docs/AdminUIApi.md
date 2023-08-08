# \AdminUIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedback**](AdminUIApi.md#CreateFeedback) | **Post** /api/admin/feedback | Send Unleash feedback
[**GetUiConfig**](AdminUIApi.md#GetUiConfig) | **Get** /api/admin/ui-config | Get UI configuration
[**SetUiConfig**](AdminUIApi.md#SetUiConfig) | **Post** /api/admin/ui-config | Set UI configuration
[**UpdateFeedback**](AdminUIApi.md#UpdateFeedback) | **Put** /api/admin/feedback/{id} | Update Unleash feedback
[**UpdateSplashSettings**](AdminUIApi.md#UpdateSplashSettings) | **Post** /api/admin/splash/{id} | Update splash settings



## CreateFeedback

> FeedbackResponseSchema CreateFeedback(ctx).FeedbackCreateSchema(feedbackCreateSchema).Execute()

Send Unleash feedback



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
    feedbackCreateSchema := *openapiclient.NewFeedbackCreateSchema("pnps") // FeedbackCreateSchema | feedbackCreateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.CreateFeedback(context.Background()).FeedbackCreateSchema(feedbackCreateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.CreateFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeedback`: FeedbackResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `AdminUIApi.CreateFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackCreateSchema** | [**FeedbackCreateSchema**](FeedbackCreateSchema.md) | feedbackCreateSchema | 

### Return type

[**FeedbackResponseSchema**](FeedbackResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUiConfig

> UiConfigSchema GetUiConfig(ctx).Execute()

Get UI configuration



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
    resp, r, err := apiClient.AdminUIApi.GetUiConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.GetUiConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUiConfig`: UiConfigSchema
    fmt.Fprintf(os.Stdout, "Response from `AdminUIApi.GetUiConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiConfigRequest struct via the builder pattern


### Return type

[**UiConfigSchema**](UiConfigSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetUiConfig

> SetUiConfig(ctx).SetUiConfigSchema(setUiConfigSchema).Execute()

Set UI configuration



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
    setUiConfigSchema := *openapiclient.NewSetUiConfigSchema() // SetUiConfigSchema | setUiConfigSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminUIApi.SetUiConfig(context.Background()).SetUiConfigSchema(setUiConfigSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.SetUiConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetUiConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setUiConfigSchema** | [**SetUiConfigSchema**](SetUiConfigSchema.md) | setUiConfigSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeedback

> FeedbackResponseSchema UpdateFeedback(ctx, id).FeedbackUpdateSchema(feedbackUpdateSchema).Execute()

Update Unleash feedback



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
    feedbackUpdateSchema := *openapiclient.NewFeedbackUpdateSchema() // FeedbackUpdateSchema | feedbackUpdateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.UpdateFeedback(context.Background(), id).FeedbackUpdateSchema(feedbackUpdateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.UpdateFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeedback`: FeedbackResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `AdminUIApi.UpdateFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feedbackUpdateSchema** | [**FeedbackUpdateSchema**](FeedbackUpdateSchema.md) | feedbackUpdateSchema | 

### Return type

[**FeedbackResponseSchema**](FeedbackResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSplashSettings

> SplashResponseSchema UpdateSplashSettings(ctx, id).Execute()

Update splash settings



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.UpdateSplashSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.UpdateSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSplashSettings`: SplashResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `AdminUIApi.UpdateSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SplashResponseSchema**](SplashResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

