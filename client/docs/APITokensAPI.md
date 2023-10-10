# \APITokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](APITokensAPI.md#CreateApiToken) | **Post** /api/admin/api-tokens | Create API token
[**DeleteApiToken**](APITokensAPI.md#DeleteApiToken) | **Delete** /api/admin/api-tokens/{token} | Delete API token
[**GetAllApiTokens**](APITokensAPI.md#GetAllApiTokens) | **Get** /api/admin/api-tokens | Get API tokens
[**GetApiTokensByName**](APITokensAPI.md#GetApiTokensByName) | **Get** /api/admin/api-tokens/{name} | Get API tokens by name
[**UpdateApiToken**](APITokensAPI.md#UpdateApiToken) | **Put** /api/admin/api-tokens/{token} | Update API token



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
    createApiTokenSchema := openapiclient.createApiTokenSchema{CreateApiTokenSchemaOneOf: openapiclient.NewCreateApiTokenSchemaOneOf("admin", "token-64522")} // CreateApiTokenSchema | createApiTokenSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokensAPI.CreateApiToken(context.Background()).CreateApiTokenSchema(createApiTokenSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensAPI.CreateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiToken`: ApiTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `APITokensAPI.CreateApiToken`: %v\n", resp)
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
    r, err := apiClient.APITokensAPI.DeleteApiToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensAPI.DeleteApiToken``: %v\n", err)
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
    resp, r, err := apiClient.APITokensAPI.GetAllApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensAPI.GetAllApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllApiTokens`: ApiTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `APITokensAPI.GetAllApiTokens`: %v\n", resp)
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


## GetApiTokensByName

> ApiTokensSchema GetApiTokensByName(ctx, name).Execute()

Get API tokens by name



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
    resp, r, err := apiClient.APITokensAPI.GetApiTokensByName(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensAPI.GetApiTokensByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTokensByName`: ApiTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `APITokensAPI.GetApiTokensByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokensByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
    r, err := apiClient.APITokensAPI.UpdateApiToken(context.Background(), token).UpdateApiTokenSchema(updateApiTokenSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensAPI.UpdateApiToken``: %v\n", err)
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

