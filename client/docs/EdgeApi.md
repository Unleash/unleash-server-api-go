# \EdgeApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkMetrics**](EdgeApi.md#BulkMetrics) | **Post** /edge/metrics | Send metrics from Edge
[**GetValidTokens**](EdgeApi.md#GetValidTokens) | **Post** /edge/validate | Check which tokens are valid



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/openapi"
)

func main() {
    bulkMetricsSchema := *openapiclient.NewBulkMetricsSchema([]openapiclient.BulkRegistrationSchema{*openapiclient.NewBulkRegistrationSchema("Ingress load balancer", "development", "application-name-dacb1234")}, []openapiclient.ClientMetricsEnvSchema{*openapiclient.NewClientMetricsEnvSchema("my.special.feature", "accounting", "development")}) // BulkMetricsSchema | bulkMetricsSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EdgeApi.BulkMetrics(context.Background()).BulkMetricsSchema(bulkMetricsSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApi.BulkMetrics``: %v\n", err)
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

> ValidatedEdgeTokensSchema GetValidTokens(ctx).RequestBody(requestBody).Execute()

Check which tokens are valid



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | tokenStringListSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApi.GetValidTokens(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApi.GetValidTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidTokens`: ValidatedEdgeTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `EdgeApi.GetValidTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | tokenStringListSchema | 

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

