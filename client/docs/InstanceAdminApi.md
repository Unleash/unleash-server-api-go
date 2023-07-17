# \InstanceAdminApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceAdminStats**](InstanceAdminApi.md#GetInstanceAdminStats) | **Get** /api/admin/instance-admin/statistics | Instance usage statistics
[**GetInstanceAdminStatsCsv**](InstanceAdminApi.md#GetInstanceAdminStatsCsv) | **Get** /api/admin/instance-admin/statistics/csv | Instance usage statistics



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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceAdminApi.GetInstanceAdminStats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceAdminApi.GetInstanceAdminStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceAdminStats`: InstanceAdminStatsSchema
    fmt.Fprintf(os.Stdout, "Response from `InstanceAdminApi.GetInstanceAdminStats`: %v\n", resp)
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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InstanceAdminApi.GetInstanceAdminStatsCsv(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceAdminApi.GetInstanceAdminStatsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInstanceAdminStatsCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `InstanceAdminApi.GetInstanceAdminStatsCsv`: %v\n", resp)
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

