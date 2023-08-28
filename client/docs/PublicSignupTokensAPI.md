# \PublicSignupTokensAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPublicSignupTokenUser**](PublicSignupTokensAPI.md#AddPublicSignupTokenUser) | **Post** /invite/{token}/signup | Add a user via a signup token
[**CreatePublicSignupToken**](PublicSignupTokensAPI.md#CreatePublicSignupToken) | **Post** /api/admin/invite-link/tokens | Create a public signup token
[**GetAllPublicSignupTokens**](PublicSignupTokensAPI.md#GetAllPublicSignupTokens) | **Get** /api/admin/invite-link/tokens | Get public signup tokens
[**GetPublicSignupToken**](PublicSignupTokensAPI.md#GetPublicSignupToken) | **Get** /api/admin/invite-link/tokens/{token} | Retrieve a token
[**UpdatePublicSignupToken**](PublicSignupTokensAPI.md#UpdatePublicSignupToken) | **Put** /api/admin/invite-link/tokens/{token} | Update a public signup token
[**ValidatePublicSignupToken**](PublicSignupTokensAPI.md#ValidatePublicSignupToken) | **Get** /invite/{token}/validate | Validate signup token



## AddPublicSignupTokenUser

> UserSchema AddPublicSignupTokenUser(ctx, token).CreateInvitedUserSchema(createInvitedUserSchema).Execute()

Add a user via a signup token



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
    createInvitedUserSchema := *openapiclient.NewCreateInvitedUserSchema("hunter@example.com", "Hunter Burgan", "hunter2") // CreateInvitedUserSchema | createInvitedUserSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensAPI.AddPublicSignupTokenUser(context.Background(), token).CreateInvitedUserSchema(createInvitedUserSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.AddPublicSignupTokenUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPublicSignupTokenUser`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensAPI.AddPublicSignupTokenUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPublicSignupTokenUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInvitedUserSchema** | [**CreateInvitedUserSchema**](CreateInvitedUserSchema.md) | createInvitedUserSchema | 

### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublicSignupToken

> PublicSignupTokenSchema CreatePublicSignupToken(ctx).PublicSignupTokenCreateSchema(publicSignupTokenCreateSchema).Execute()

Create a public signup token



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
    publicSignupTokenCreateSchema := *openapiclient.NewPublicSignupTokenCreateSchema("Name_example", time.Now()) // PublicSignupTokenCreateSchema | publicSignupTokenCreateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensAPI.CreatePublicSignupToken(context.Background()).PublicSignupTokenCreateSchema(publicSignupTokenCreateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.CreatePublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensAPI.CreatePublicSignupToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicSignupTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicSignupTokenCreateSchema** | [**PublicSignupTokenCreateSchema**](PublicSignupTokenCreateSchema.md) | publicSignupTokenCreateSchema | 

### Return type

[**PublicSignupTokenSchema**](PublicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPublicSignupTokens

> PublicSignupTokensSchema GetAllPublicSignupTokens(ctx).Execute()

Get public signup tokens



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
    resp, r, err := apiClient.PublicSignupTokensAPI.GetAllPublicSignupTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.GetAllPublicSignupTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPublicSignupTokens`: PublicSignupTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensAPI.GetAllPublicSignupTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPublicSignupTokensRequest struct via the builder pattern


### Return type

[**PublicSignupTokensSchema**](PublicSignupTokensSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicSignupToken

> PublicSignupTokenSchema GetPublicSignupToken(ctx, token).Execute()

Retrieve a token



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
    resp, r, err := apiClient.PublicSignupTokensAPI.GetPublicSignupToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.GetPublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensAPI.GetPublicSignupToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicSignupTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicSignupTokenSchema**](PublicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicSignupToken

> PublicSignupTokenSchema UpdatePublicSignupToken(ctx, token).PublicSignupTokenUpdateSchema(publicSignupTokenUpdateSchema).Execute()

Update a public signup token



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
    publicSignupTokenUpdateSchema := *openapiclient.NewPublicSignupTokenUpdateSchema() // PublicSignupTokenUpdateSchema | publicSignupTokenUpdateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensAPI.UpdatePublicSignupToken(context.Background(), token).PublicSignupTokenUpdateSchema(publicSignupTokenUpdateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.UpdatePublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensAPI.UpdatePublicSignupToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicSignupTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicSignupTokenUpdateSchema** | [**PublicSignupTokenUpdateSchema**](PublicSignupTokenUpdateSchema.md) | publicSignupTokenUpdateSchema | 

### Return type

[**PublicSignupTokenSchema**](PublicSignupTokenSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePublicSignupToken

> ValidatePublicSignupToken(ctx, token).Execute()

Validate signup token



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
    r, err := apiClient.PublicSignupTokensAPI.ValidatePublicSignupToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensAPI.ValidatePublicSignupToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiValidatePublicSignupTokenRequest struct via the builder pattern


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

