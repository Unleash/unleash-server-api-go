# \StrategiesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStrategy**](StrategiesApi.md#CreateStrategy) | **Post** /api/admin/strategies | Create a strategy
[**DeprecateStrategy**](StrategiesApi.md#DeprecateStrategy) | **Post** /api/admin/strategies/{strategyName}/deprecate | Deprecate a strategy
[**GetAllStrategies**](StrategiesApi.md#GetAllStrategies) | **Get** /api/admin/strategies | Get all strategies
[**GetStrategiesByContextField**](StrategiesApi.md#GetStrategiesByContextField) | **Get** /api/admin/context/{contextField}/strategies | Get strategies that use a context field
[**GetStrategy**](StrategiesApi.md#GetStrategy) | **Get** /api/admin/strategies/{name} | Get a strategy definition
[**ReactivateStrategy**](StrategiesApi.md#ReactivateStrategy) | **Post** /api/admin/strategies/{strategyName}/reactivate | Reactivate a strategy
[**RemoveStrategy**](StrategiesApi.md#RemoveStrategy) | **Delete** /api/admin/strategies/{name} | Delete a strategy
[**UpdateFeatureStrategySegments**](StrategiesApi.md#UpdateFeatureStrategySegments) | **Post** /api/admin/segments/strategies | Update strategy segments
[**UpdateStrategy**](StrategiesApi.md#UpdateStrategy) | **Put** /api/admin/strategies/{strategyName} | Update a strategy type



## CreateStrategy

> StrategySchema CreateStrategy(ctx).CreateStrategySchema(createStrategySchema).Execute()

Create a strategy



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
    createStrategySchema := *openapiclient.NewCreateStrategySchema("my-custom-strategy", []openapiclient.CreateStrategySchemaParametersInner{*openapiclient.NewCreateStrategySchemaParametersInner("Rollout percentage", "percentage")}) // CreateStrategySchema | createStrategySchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StrategiesApi.CreateStrategy(context.Background()).CreateStrategySchema(createStrategySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.CreateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStrategy`: StrategySchema
    fmt.Fprintf(os.Stdout, "Response from `StrategiesApi.CreateStrategy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStrategySchema** | [**CreateStrategySchema**](CreateStrategySchema.md) | createStrategySchema | 

### Return type

[**StrategySchema**](StrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeprecateStrategy

> DeprecateStrategy(ctx, strategyName).Execute()

Deprecate a strategy



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
    strategyName := "strategyName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StrategiesApi.DeprecateStrategy(context.Background(), strategyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.DeprecateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecateStrategyRequest struct via the builder pattern


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


## GetAllStrategies

> StrategiesSchema GetAllStrategies(ctx).Execute()

Get all strategies



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
    resp, r, err := apiClient.StrategiesApi.GetAllStrategies(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.GetAllStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllStrategies`: StrategiesSchema
    fmt.Fprintf(os.Stdout, "Response from `StrategiesApi.GetAllStrategies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllStrategiesRequest struct via the builder pattern


### Return type

[**StrategiesSchema**](StrategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategiesByContextField

> ContextFieldStrategiesSchema GetStrategiesByContextField(ctx, contextField).Execute()

Get strategies that use a context field



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
    contextField := "contextField_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StrategiesApi.GetStrategiesByContextField(context.Background(), contextField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.GetStrategiesByContextField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStrategiesByContextField`: ContextFieldStrategiesSchema
    fmt.Fprintf(os.Stdout, "Response from `StrategiesApi.GetStrategiesByContextField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategiesByContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextFieldStrategiesSchema**](ContextFieldStrategiesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategy

> StrategySchema GetStrategy(ctx, name).Execute()

Get a strategy definition



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
    resp, r, err := apiClient.StrategiesApi.GetStrategy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.GetStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStrategy`: StrategySchema
    fmt.Fprintf(os.Stdout, "Response from `StrategiesApi.GetStrategy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StrategySchema**](StrategySchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReactivateStrategy

> ReactivateStrategy(ctx, strategyName).Execute()

Reactivate a strategy



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
    strategyName := "strategyName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StrategiesApi.ReactivateStrategy(context.Background(), strategyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.ReactivateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateStrategyRequest struct via the builder pattern


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


## RemoveStrategy

> RemoveStrategy(ctx, name).Execute()

Delete a strategy



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
    r, err := apiClient.StrategiesApi.RemoveStrategy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.RemoveStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveStrategyRequest struct via the builder pattern


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


## UpdateFeatureStrategySegments

> UpdateFeatureStrategySegmentsSchema UpdateFeatureStrategySegments(ctx).UpdateFeatureStrategySegmentsSchema(updateFeatureStrategySegmentsSchema).Execute()

Update strategy segments



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
    updateFeatureStrategySegmentsSchema := *openapiclient.NewUpdateFeatureStrategySegmentsSchema("red-vista", "15d1e20b-6310-4e17-a0c2-9fb84de3053a", "development", []int32{int32(123)}) // UpdateFeatureStrategySegmentsSchema | updateFeatureStrategySegmentsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StrategiesApi.UpdateFeatureStrategySegments(context.Background()).UpdateFeatureStrategySegmentsSchema(updateFeatureStrategySegmentsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.UpdateFeatureStrategySegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureStrategySegments`: UpdateFeatureStrategySegmentsSchema
    fmt.Fprintf(os.Stdout, "Response from `StrategiesApi.UpdateFeatureStrategySegments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureStrategySegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFeatureStrategySegmentsSchema** | [**UpdateFeatureStrategySegmentsSchema**](UpdateFeatureStrategySegmentsSchema.md) | updateFeatureStrategySegmentsSchema | 

### Return type

[**UpdateFeatureStrategySegmentsSchema**](UpdateFeatureStrategySegmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStrategy

> UpdateStrategy(ctx, strategyName).UpdateStrategySchema(updateStrategySchema).Execute()

Update a strategy type



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
    strategyName := "strategyName_example" // string | 
    updateStrategySchema := *openapiclient.NewUpdateStrategySchema([]openapiclient.CreateStrategySchemaParametersInner{*openapiclient.NewCreateStrategySchemaParametersInner("Rollout percentage", "percentage")}) // UpdateStrategySchema | updateStrategySchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StrategiesApi.UpdateStrategy(context.Background(), strategyName).UpdateStrategySchema(updateStrategySchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StrategiesApi.UpdateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**strategyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStrategySchema** | [**UpdateStrategySchema**](UpdateStrategySchema.md) | updateStrategySchema | 

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

