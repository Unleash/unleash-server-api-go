# \MetricsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](MetricsApi.md#CreateApplication) | **Post** /api/admin/metrics/applications/{appName} | Create an application to connect reported metrics
[**DeleteApplication**](MetricsApi.md#DeleteApplication) | **Delete** /api/admin/metrics/applications/{appName} | Delete an application
[**GetApplication**](MetricsApi.md#GetApplication) | **Get** /api/admin/metrics/applications/{appName} | Get application data
[**GetApplications**](MetricsApi.md#GetApplications) | **Get** /api/admin/metrics/applications | Get all applications
[**GetFeatureUsageSummary**](MetricsApi.md#GetFeatureUsageSummary) | **Get** /api/admin/client-metrics/features/{name} | Last hour of usage and a list of applications that have reported seeing this feature toggle
[**GetRawFeatureMetrics**](MetricsApi.md#GetRawFeatureMetrics) | **Get** /api/admin/client-metrics/features/{name}/raw | Feature usage metrics for the last 48 hours, grouped by hour
[**GetRequestsPerSecond**](MetricsApi.md#GetRequestsPerSecond) | **Get** /api/admin/metrics/rps | Gets usage data



## CreateApplication

> CreateApplication(ctx, appName).RequestBody(requestBody).Execute()

Create an application to connect reported metrics



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
    appName := "appName_example" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | createApplicationSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetricsApi.CreateApplication(context.Background(), appName).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** | createApplicationSchema | 

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


## DeleteApplication

> DeleteApplication(ctx, appName).Execute()

Delete an application



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetricsApi.DeleteApplication(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> ApplicationSchema GetApplication(ctx, appName).Execute()

Get application data



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
    appName := "appName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.GetApplication(context.Background(), appName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ApplicationSchema
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationSchema**](ApplicationSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsSchema GetApplications(ctx).Execute()

Get all applications



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
    resp, r, err := apiClient.MetricsApi.GetApplications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: ApplicationsSchema
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


### Return type

[**ApplicationsSchema**](ApplicationsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureUsageSummary

> FeatureUsageSchema GetFeatureUsageSummary(ctx, name).Execute()

Last hour of usage and a list of applications that have reported seeing this feature toggle



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.GetFeatureUsageSummary(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetFeatureUsageSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureUsageSummary`: FeatureUsageSchema
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetFeatureUsageSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureUsageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureUsageSchema**](FeatureUsageSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawFeatureMetrics

> FeatureMetricsSchema GetRawFeatureMetrics(ctx, name).Execute()

Feature usage metrics for the last 48 hours, grouped by hour

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.GetRawFeatureMetrics(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetRawFeatureMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawFeatureMetrics`: FeatureMetricsSchema
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetRawFeatureMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawFeatureMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FeatureMetricsSchema**](FeatureMetricsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestsPerSecond

> RequestsPerSecondSegmentedSchema GetRequestsPerSecond(ctx).Execute()

Gets usage data



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
    resp, r, err := apiClient.MetricsApi.GetRequestsPerSecond(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetRequestsPerSecond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestsPerSecond`: RequestsPerSecondSegmentedSchema
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetRequestsPerSecond`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsPerSecondRequest struct via the builder pattern


### Return type

[**RequestsPerSecondSegmentedSchema**](RequestsPerSecondSegmentedSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

