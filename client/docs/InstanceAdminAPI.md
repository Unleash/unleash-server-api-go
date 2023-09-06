# \InstanceAdminAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceAdminStats**](InstanceAdminAPI.md#GetInstanceAdminStats) | **Get** /api/admin/instance-admin/statistics | Instance usage statistics
[**GetInstanceAdminStatsCsv**](InstanceAdminAPI.md#GetInstanceAdminStatsCsv) | **Get** /api/admin/instance-admin/statistics/csv | Instance usage statistics
[**GetLoginHistory**](InstanceAdminAPI.md#GetLoginHistory) | **Get** /api/admin/logins | Get all login events.



## GetInstanceAdminStats

> InstanceAdminStatsSchema GetInstanceAdminStats(ctx).Execute()

Instance usage statistics



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
    resp, r, err := apiClient.InstanceAdminAPI.GetInstanceAdminStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceAdminAPI.GetInstanceAdminStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceAdminStats`: InstanceAdminStatsSchema
    fmt.Fprintf(os.Stdout, "Response from `InstanceAdminAPI.GetInstanceAdminStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceAdminStatsRequest struct via the builder pattern


### Return type

[**InstanceAdminStatsSchema**](InstanceAdminStatsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceAdminStatsCsv

> string GetInstanceAdminStatsCsv(ctx).Execute()

Instance usage statistics



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
    resp, r, err := apiClient.InstanceAdminAPI.GetInstanceAdminStatsCsv(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceAdminAPI.GetInstanceAdminStatsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceAdminStatsCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `InstanceAdminAPI.GetInstanceAdminStatsCsv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceAdminStatsCsvRequest struct via the builder pattern


### Return type

**string**

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

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
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceAdminAPI.GetLoginHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceAdminAPI.GetLoginHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoginHistory`: LoginHistorySchema
    fmt.Fprintf(os.Stdout, "Response from `InstanceAdminAPI.GetLoginHistory`: %v\n", resp)
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

