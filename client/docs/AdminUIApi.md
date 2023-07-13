# \AdminUIApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeedback**](AdminUIApi.md#CreateFeedback) | **Post** /api/admin/feedback | 
[**GetUiConfig**](AdminUIApi.md#GetUiConfig) | **Get** /api/admin/ui-config | 
[**SetUiConfig**](AdminUIApi.md#SetUiConfig) | **Post** /api/admin/ui-config | 
[**UpdateFeedback**](AdminUIApi.md#UpdateFeedback) | **Put** /api/admin/feedback/{id} | 
[**UpdateSplashSettings**](AdminUIApi.md#UpdateSplashSettings) | **Post** /api/admin/splash/{id} | 



## CreateFeedback

> FeedbackSchema CreateFeedback(ctx).FeedbackSchema(feedbackSchema).Execute()



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
    feedbackSchema := *openapiclient.NewFeedbackSchema() // FeedbackSchema | feedbackSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.CreateFeedback(context.Background()).FeedbackSchema(feedbackSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.CreateFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeedback`: FeedbackSchema
    fmt.Fprintf(os.Stdout, "Response from `AdminUIApi.CreateFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackSchema** | [**FeedbackSchema**](FeedbackSchema.md) | feedbackSchema | 

### Return type

[**FeedbackSchema**](FeedbackSchema.md)

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

> FeedbackSchema UpdateFeedback(ctx, id).FeedbackSchema(feedbackSchema).Execute()



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
    id := "id_example" // string | 
    feedbackSchema := *openapiclient.NewFeedbackSchema() // FeedbackSchema | feedbackSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.UpdateFeedback(context.Background(), id).FeedbackSchema(feedbackSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.UpdateFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeedback`: FeedbackSchema
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

 **feedbackSchema** | [**FeedbackSchema**](FeedbackSchema.md) | feedbackSchema | 

### Return type

[**FeedbackSchema**](FeedbackSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSplashSettings

> SplashSchema UpdateSplashSettings(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminUIApi.UpdateSplashSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUIApi.UpdateSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSplashSettings`: SplashSchema
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

[**SplashSchema**](SplashSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

