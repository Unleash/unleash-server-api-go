# \FrontendAPIApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFrontendFeatures**](FrontendAPIApi.md#GetFrontendFeatures) | **Get** /api/frontend | Retrieve enabled feature toggles for the provided context.
[**RegisterFrontendClient**](FrontendAPIApi.md#RegisterFrontendClient) | **Post** /api/frontend/client/register | Register a client SDK
[**RegisterFrontendMetrics**](FrontendAPIApi.md#RegisterFrontendMetrics) | **Post** /api/frontend/client/metrics | Register client usage metrics



## GetFrontendFeatures

> ProxyFeaturesSchema GetFrontendFeatures(ctx).Execute()

Retrieve enabled feature toggles for the provided context.



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
    resp, r, err := apiClient.FrontendAPIApi.GetFrontendFeatures(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrontendAPIApi.GetFrontendFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrontendFeatures`: ProxyFeaturesSchema
    fmt.Fprintf(os.Stdout, "Response from `FrontendAPIApi.GetFrontendFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrontendFeaturesRequest struct via the builder pattern


### Return type

[**ProxyFeaturesSchema**](ProxyFeaturesSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterFrontendClient

> RegisterFrontendClient(ctx).ProxyClientSchema(proxyClientSchema).Execute()

Register a client SDK



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
    proxyClientSchema := *openapiclient.NewProxyClientSchema("AppName_example", float32(123), openapiclient.proxyClientSchema_started{Float32: new(float32)}, []string{"Strategies_example"}) // ProxyClientSchema | proxyClientSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FrontendAPIApi.RegisterFrontendClient(context.Background()).ProxyClientSchema(proxyClientSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrontendAPIApi.RegisterFrontendClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterFrontendClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyClientSchema** | [**ProxyClientSchema**](ProxyClientSchema.md) | proxyClientSchema | 

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


## RegisterFrontendMetrics

> RegisterFrontendMetrics(ctx).ClientMetricsSchema(clientMetricsSchema).Execute()

Register client usage metrics



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
    clientMetricsSchema := *openapiclient.NewClientMetricsSchema("insurance-selector", *openapiclient.NewClientMetricsSchemaBucket(openapiclient.dateSchema{Float32: new(float32)}, openapiclient.dateSchema{Float32: new(float32)}, map[string]ClientMetricsSchemaBucketTogglesValue{"key": *openapiclient.NewClientMetricsSchemaBucketTogglesValue()})) // ClientMetricsSchema | clientMetricsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FrontendAPIApi.RegisterFrontendMetrics(context.Background()).ClientMetricsSchema(clientMetricsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FrontendAPIApi.RegisterFrontendMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterFrontendMetricsRequest struct via the builder pattern


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

