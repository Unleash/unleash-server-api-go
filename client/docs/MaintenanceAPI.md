# \MaintenanceAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMaintenance**](MaintenanceAPI.md#GetMaintenance) | **Get** /api/admin/maintenance | Get maintenance mode status
[**ToggleMaintenance**](MaintenanceAPI.md#ToggleMaintenance) | **Post** /api/admin/maintenance | Enabled/disabled maintenance mode



## GetMaintenance

> MaintenanceSchema GetMaintenance(ctx).Execute()

Get maintenance mode status



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
    resp, r, err := apiClient.MaintenanceAPI.GetMaintenance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.GetMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenance`: MaintenanceSchema
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.GetMaintenance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceRequest struct via the builder pattern


### Return type

[**MaintenanceSchema**](MaintenanceSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleMaintenance

> ToggleMaintenance(ctx).ToggleMaintenanceSchema(toggleMaintenanceSchema).Execute()

Enabled/disabled maintenance mode



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
    toggleMaintenanceSchema := *openapiclient.NewToggleMaintenanceSchema(true) // ToggleMaintenanceSchema | toggleMaintenanceSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceAPI.ToggleMaintenance(context.Background()).ToggleMaintenanceSchema(toggleMaintenanceSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.ToggleMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toggleMaintenanceSchema** | [**ToggleMaintenanceSchema**](ToggleMaintenanceSchema.md) | toggleMaintenanceSchema | 

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

