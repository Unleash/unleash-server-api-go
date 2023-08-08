# \ServiceAccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccountToken**](ServiceAccountsApi.md#CreateServiceAccountToken) | **Post** /api/admin/service-account/{id}/token | Create a token for a service account.
[**DeleteServiceAccount**](ServiceAccountsApi.md#DeleteServiceAccount) | **Delete** /api/admin/service-account/{id} | Delete a service account.
[**DeleteServiceAccountToken**](ServiceAccountsApi.md#DeleteServiceAccountToken) | **Delete** /api/admin/service-account/{id}/token/{tokenId} | Delete a token for a service account.
[**GetServiceAccountTokens**](ServiceAccountsApi.md#GetServiceAccountTokens) | **Get** /api/admin/service-account/{id}/token | List all tokens for a service account.
[**UpdateServiceAccount**](ServiceAccountsApi.md#UpdateServiceAccount) | **Put** /api/admin/service-account/{id} | Update a service account.



## CreateServiceAccountToken

> PatSchema CreateServiceAccountToken(ctx, id).PatSchema(patSchema).Execute()

Create a token for a service account.



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
    patSchema := *openapiclient.NewPatSchema() // PatSchema | patSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.CreateServiceAccountToken(context.Background(), id).PatSchema(patSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccountToken`: PatSchema
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patSchema** | [**PatSchema**](PatSchema.md) | patSchema | 

### Return type

[**PatSchema**](PatSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> DeleteServiceAccount(ctx, id).Execute()

Delete a service account.



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
    r, err := apiClient.ServiceAccountsApi.DeleteServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


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


## DeleteServiceAccountToken

> DeleteServiceAccountToken(ctx, id, tokenId).Execute()

Delete a token for a service account.



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
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceAccountsApi.DeleteServiceAccountToken(context.Background(), id, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountTokenRequest struct via the builder pattern


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


## GetServiceAccountTokens

> PatsSchema GetServiceAccountTokens(ctx, id).Execute()

List all tokens for a service account.



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
    resp, r, err := apiClient.ServiceAccountsApi.GetServiceAccountTokens(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccountTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountTokens`: PatsSchema
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PatsSchema**](PatsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccount

> ServiceAccountSchema UpdateServiceAccount(ctx, id).RequestBody(requestBody).Execute()

Update a service account.



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | updateServiceAccountSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountsApi.UpdateServiceAccount(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAccount`: ServiceAccountSchema
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** | updateServiceAccountSchema | 

### Return type

[**ServiceAccountSchema**](ServiceAccountSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
