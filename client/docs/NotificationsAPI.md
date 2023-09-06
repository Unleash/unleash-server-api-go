# \NotificationsAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotifications**](NotificationsAPI.md#GetNotifications) | **Get** /api/admin/notifications | Retrieves a list of notifications
[**MarkNotificationsAsRead**](NotificationsAPI.md#MarkNotificationsAsRead) | **Post** /api/admin/notifications/read | Mark notifications as read



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
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetNotifications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: []NotificationsSchemaInner
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotifications`: %v\n", resp)
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

> MarkNotificationsAsRead(ctx).MarkNotificationsAsReadSchema(markNotificationsAsReadSchema).Execute()

Mark notifications as read



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
    markNotificationsAsReadSchema := *openapiclient.NewMarkNotificationsAsReadSchema([]int32{int32(5)}) // MarkNotificationsAsReadSchema | markNotificationsAsReadSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsAPI.MarkNotificationsAsRead(context.Background()).MarkNotificationsAsReadSchema(markNotificationsAsReadSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.MarkNotificationsAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markNotificationsAsReadSchema** | [**MarkNotificationsAsReadSchema**](MarkNotificationsAsReadSchema.md) | markNotificationsAsReadSchema | 

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

