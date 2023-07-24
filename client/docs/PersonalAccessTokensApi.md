# \PersonalAccessTokensApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePat**](PersonalAccessTokensApi.md#CreatePat) | **Post** /api/admin/user/tokens | Create a new Personal Access Token.
[**DeletePat**](PersonalAccessTokensApi.md#DeletePat) | **Delete** /api/admin/user/tokens/{id} | Delete a Personal Access Token.
[**GetPats**](PersonalAccessTokensApi.md#GetPats) | **Get** /api/admin/user/tokens | Get all Personal Access Tokens for the current user.



## CreatePat

> PatSchema CreatePat(ctx).PatSchema(patSchema).Execute()

Create a new Personal Access Token.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    patSchema := *openapiclient.NewPatSchema() // PatSchema | patSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.CreatePat(context.Background()).PatSchema(patSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.CreatePat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePat`: PatSchema
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.CreatePat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePatRequest struct via the builder pattern


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


## DeletePat

> DeletePat(ctx, id).Execute()

Delete a Personal Access Token.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PersonalAccessTokensApi.DeletePat(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.DeletePat``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePatRequest struct via the builder pattern


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


## GetPats

> PatsSchema GetPats(ctx).Execute()

Get all Personal Access Tokens for the current user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.GetPats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.GetPats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPats`: PatsSchema
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.GetPats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatsRequest struct via the builder pattern


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
