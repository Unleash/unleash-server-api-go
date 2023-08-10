# \SegmentsApi

All URIs are relative to *https://app.unleash-hosted.com/hosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegment**](SegmentsApi.md#CreateSegment) | **Post** /api/admin/segments | Create a new segment
[**GetSegment**](SegmentsApi.md#GetSegment) | **Get** /api/admin/segments/{id} | Get a segment
[**GetSegments**](SegmentsApi.md#GetSegments) | **Get** /api/admin/segments | Get all segments
[**GetSegmentsByStrategyId**](SegmentsApi.md#GetSegmentsByStrategyId) | **Get** /api/admin/segments/strategies/{strategyId} | Get strategy segments
[**GetStrategiesBySegmentId**](SegmentsApi.md#GetStrategiesBySegmentId) | **Get** /api/admin/segments/{id}/strategies | Get strategies that reference segment
[**RemoveSegment**](SegmentsApi.md#RemoveSegment) | **Delete** /api/admin/segments/{id} | Deletes a segment by id
[**UpdateSegment**](SegmentsApi.md#UpdateSegment) | **Put** /api/admin/segments/{id} | Update segment by id
[**ValidateSegment**](SegmentsApi.md#ValidateSegment) | **Post** /api/admin/segments/validate | Validates if a segment name exists



## CreateSegment

> AdminSegmentSchema CreateSegment(ctx).UpsertSegmentSchema(upsertSegmentSchema).Execute()

Create a new segment



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
    upsertSegmentSchema := *openapiclient.NewUpsertSegmentSchema("Name_example", []openapiclient.ConstraintSchema{*openapiclient.NewConstraintSchema("appName", "IN")}) // UpsertSegmentSchema | upsertSegmentSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.CreateSegment(context.Background()).UpsertSegmentSchema(upsertSegmentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.CreateSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSegment`: AdminSegmentSchema
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.CreateSegment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertSegmentSchema** | [**UpsertSegmentSchema**](UpsertSegmentSchema.md) | upsertSegmentSchema | 

### Return type

[**AdminSegmentSchema**](AdminSegmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegment

> AdminSegmentSchema GetSegment(ctx, id).Execute()

Get a segment



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
    resp, r, err := apiClient.SegmentsApi.GetSegment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegment`: AdminSegmentSchema
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminSegmentSchema**](AdminSegmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegments

> SegmentsSchema GetSegments(ctx).Execute()

Get all segments



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
    resp, r, err := apiClient.SegmentsApi.GetSegments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegments`: SegmentsSchema
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsRequest struct via the builder pattern


### Return type

[**SegmentsSchema**](SegmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegmentsByStrategyId

> SegmentsSchema GetSegmentsByStrategyId(ctx, strategyId).Execute()

Get strategy segments



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
    strategyId := "strategyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SegmentsApi.GetSegmentsByStrategyId(context.Background(), strategyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetSegmentsByStrategyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegmentsByStrategyId`: SegmentsSchema
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetSegmentsByStrategyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsByStrategyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SegmentsSchema**](SegmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategiesBySegmentId

> AdminStrategiesSchema GetStrategiesBySegmentId(ctx, id).Execute()

Get strategies that reference segment



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
    resp, r, err := apiClient.SegmentsApi.GetStrategiesBySegmentId(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.GetStrategiesBySegmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStrategiesBySegmentId`: AdminStrategiesSchema
    fmt.Fprintf(os.Stdout, "Response from `SegmentsApi.GetStrategiesBySegmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategiesBySegmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminStrategiesSchema**](AdminStrategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSegment

> RemoveSegment(ctx, id).Execute()

Deletes a segment by id



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
    r, err := apiClient.SegmentsApi.RemoveSegment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.RemoveSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSegmentRequest struct via the builder pattern


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


## UpdateSegment

> UpdateSegment(ctx, id).Execute()

Update segment by id



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
    r, err := apiClient.SegmentsApi.UpdateSegment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.UpdateSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSegmentRequest struct via the builder pattern


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


## ValidateSegment

> ValidateSegment(ctx).NameSchema(nameSchema).Execute()

Validates if a segment name exists



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
    nameSchema := *openapiclient.NewNameSchema("betaUser") // NameSchema | nameSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SegmentsApi.ValidateSegment(context.Background()).NameSchema(nameSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SegmentsApi.ValidateSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameSchema** | [**NameSchema**](NameSchema.md) | nameSchema | 

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

