# \EdgeAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkMetrics**](EdgeAPI.md#BulkMetrics) | **Post** /edge/metrics | Send metrics from Edge
[**GetValidTokens**](EdgeAPI.md#GetValidTokens) | **Post** /edge/validate | Check which tokens are valid



## BulkMetrics

> BulkMetrics(ctx).BulkMetricsSchema(bulkMetricsSchema).Execute()

Send metrics from Edge



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
    bulkMetricsSchema := *openapiclient.NewBulkMetricsSchema([]openapiclient.BulkRegistrationSchema{*openapiclient.NewBulkRegistrationSchema("Ingress load balancer", "development", "application-name-dacb1234")}, []openapiclient.ClientMetricsEnvSchema{*openapiclient.NewClientMetricsEnvSchema("my.special.feature", "accounting", "development")}) // BulkMetricsSchema | bulkMetricsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EdgeAPI.BulkMetrics(context.Background()).BulkMetricsSchema(bulkMetricsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.BulkMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkMetricsSchema** | [**BulkMetricsSchema**](BulkMetricsSchema.md) | bulkMetricsSchema | 

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


## GetValidTokens

> ValidatedEdgeTokensSchema GetValidTokens(ctx).TokenStringListSchema(tokenStringListSchema).Execute()

Check which tokens are valid



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
    tokenStringListSchema := *openapiclient.NewTokenStringListSchema([]string{"Tokens_example"}) // TokenStringListSchema | tokenStringListSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeAPI.GetValidTokens(context.Background()).TokenStringListSchema(tokenStringListSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.GetValidTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidTokens`: ValidatedEdgeTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `EdgeAPI.GetValidTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenStringListSchema** | [**TokenStringListSchema**](TokenStringListSchema.md) | tokenStringListSchema | 

### Return type

[**ValidatedEdgeTokensSchema**](ValidatedEdgeTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

