# \ClientAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllClientFeatures**](ClientAPI.md#GetAllClientFeatures) | **Get** /api/client/features | Get all toggles (SDK)
[**GetClientFeature**](ClientAPI.md#GetClientFeature) | **Get** /api/client/features/{featureName} | Get a single feature toggle
[**RegisterClientApplication**](ClientAPI.md#RegisterClientApplication) | **Post** /api/client/register | Register a client SDK
[**RegisterClientMetrics**](ClientAPI.md#RegisterClientMetrics) | **Post** /api/client/metrics | Register client usage metrics



## GetAllClientFeatures

> ClientFeaturesSchema GetAllClientFeatures(ctx).Execute()

Get all toggles (SDK)



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
    resp, r, err := apiClient.ClientAPI.GetAllClientFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetAllClientFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllClientFeatures`: ClientFeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetAllClientFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllClientFeaturesRequest struct via the builder pattern


### Return type

[**ClientFeaturesSchema**](ClientFeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientFeature

> ClientFeatureSchema GetClientFeature(ctx, featureName).Execute()

Get a single feature toggle



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
    featureName := "featureName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClientAPI.GetClientFeature(context.Background(), featureName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.GetClientFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientFeature`: ClientFeatureSchema
    fmt.Fprintf(os.Stdout, "Response from `ClientAPI.GetClientFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientFeatureSchema**](ClientFeatureSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterClientApplication

> RegisterClientApplication(ctx).ClientApplicationSchema(clientApplicationSchema).Execute()

Register a client SDK



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
    clientApplicationSchema := *openapiclient.NewClientApplicationSchema("example-app", float32(10), openapiclient.clientApplicationSchema_started{Float32: new(float32)}, []string{"Strategies_example"}) // ClientApplicationSchema | clientApplicationSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientAPI.RegisterClientApplication(context.Background()).ClientApplicationSchema(clientApplicationSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RegisterClientApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterClientApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApplicationSchema** | [**ClientApplicationSchema**](ClientApplicationSchema.md) | clientApplicationSchema | 

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


## RegisterClientMetrics

> RegisterClientMetrics(ctx).ClientMetricsSchema(clientMetricsSchema).Execute()

Register client usage metrics



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
    clientMetricsSchema := *openapiclient.NewClientMetricsSchema("insurance-selector", *openapiclient.NewClientMetricsSchemaBucket(openapiclient.dateSchema{Int32: new(int32)}, openapiclient.dateSchema{Int32: new(int32)}, map[string]ClientMetricsSchemaBucketTogglesValue{"key": *openapiclient.NewClientMetricsSchemaBucketTogglesValue()})) // ClientMetricsSchema | clientMetricsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ClientAPI.RegisterClientMetrics(context.Background()).ClientMetricsSchema(clientMetricsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientAPI.RegisterClientMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterClientMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientMetricsSchema** | [**ClientMetricsSchema**](ClientMetricsSchema.md) | clientMetricsSchema | 

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

