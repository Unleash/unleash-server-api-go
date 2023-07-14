# \UnstableApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedPlayground**](UnstableApi.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**GetLoginHistory**](UnstableApi.md#GetLoginHistory) | **Get** /api/admin/logins | Get all login events.
[**GetNotifications**](UnstableApi.md#GetNotifications) | **Get** /api/admin/notifications | Retrieves a list of notifications
[**MarkNotificationsAsRead**](UnstableApi.md#MarkNotificationsAsRead) | **Post** /api/admin/notifications/read | Mark notifications as read



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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
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


## GetLoginHistory

> LoginHistorySchema GetLoginHistory(ctx).Execute()

Get all login events.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetLoginHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetLoginHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginHistory`: LoginHistorySchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetLoginHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginHistoryRequest struct via the builder pattern


### Return type

[**LoginHistorySchema**](LoginHistorySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> []NotificationsSchemaInner GetNotifications(ctx).Execute()

Retrieves a list of notifications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: []NotificationsSchemaInner
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


### Return type

[**[]NotificationsSchemaInner**](NotificationsSchemaInner.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationsAsRead

> MarkNotificationsAsRead(ctx).RequestBody(requestBody).Execute()

Mark notifications as read



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | markNotificationsAsReadSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.MarkNotificationsAsRead(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.MarkNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | markNotificationsAsReadSchema | 

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

