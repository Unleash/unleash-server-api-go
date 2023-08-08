# \APITokensApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](APITokensApi.md#CreateApiToken) | **Post** /api/admin/api-tokens | Create API token
[**DeleteApiToken**](APITokensApi.md#DeleteApiToken) | **Delete** /api/admin/api-tokens/{token} | Delete API token
[**GetAllApiTokens**](APITokensApi.md#GetAllApiTokens) | **Get** /api/admin/api-tokens | Get API tokens
[**UpdateApiToken**](APITokensApi.md#UpdateApiToken) | **Put** /api/admin/api-tokens/{token} | Update API token



## CreateApiToken

> ApiTokenSchema CreateApiToken(ctx).CreateApiTokenSchema(createApiTokenSchema).Execute()

Create API token



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
    createApiTokenSchema := *openapiclient.NewCreateApiTokenSchema() // CreateApiTokenSchema | createApiTokenSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokensApi.CreateApiToken(context.Background()).CreateApiTokenSchema(createApiTokenSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.CreateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiToken`: ApiTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `APITokensApi.CreateApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiTokenSchema** | [**CreateApiTokenSchema**](CreateApiTokenSchema.md) | createApiTokenSchema | 

### Return type

[**ApiTokenSchema**](ApiTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiToken

> DeleteApiToken(ctx, token).Execute()

Delete API token



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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.APITokensApi.DeleteApiToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.DeleteApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


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


## GetAllApiTokens

> ApiTokensSchema GetAllApiTokens(ctx).Execute()

Get API tokens



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
    resp, r, err := apiClient.APITokensApi.GetAllApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.GetAllApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllApiTokens`: ApiTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `APITokensApi.GetAllApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApiTokensRequest struct via the builder pattern


### Return type

[**ApiTokensSchema**](ApiTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiToken

> UpdateApiToken(ctx, token).UpdateApiTokenSchema(updateApiTokenSchema).Execute()

Update API token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/Unleash/unleash-server-api-go"
)

func main() {
    token := "token_example" // string | 
    updateApiTokenSchema := *openapiclient.NewUpdateApiTokenSchema(time.Now()) // UpdateApiTokenSchema | updateApiTokenSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.APITokensApi.UpdateApiToken(context.Background(), token).UpdateApiTokenSchema(updateApiTokenSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.UpdateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApiTokenSchema** | [**UpdateApiTokenSchema**](UpdateApiTokenSchema.md) | updateApiTokenSchema | 

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

