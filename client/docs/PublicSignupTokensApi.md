# \PublicSignupTokensApi

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPublicSignupTokenUser**](PublicSignupTokensApi.md#AddPublicSignupTokenUser) | **Post** /invite/{token}/signup | Add a user via a signup token
[**CreatePublicSignupToken**](PublicSignupTokensApi.md#CreatePublicSignupToken) | **Post** /api/admin/invite-link/tokens | Create a public signup token
[**GetAllPublicSignupTokens**](PublicSignupTokensApi.md#GetAllPublicSignupTokens) | **Get** /api/admin/invite-link/tokens | Retrieve all existing public signup tokens
[**GetPublicSignupToken**](PublicSignupTokensApi.md#GetPublicSignupToken) | **Get** /api/admin/invite-link/tokens/{token} | Retrieve a token
[**UpdatePublicSignupToken**](PublicSignupTokensApi.md#UpdatePublicSignupToken) | **Put** /api/admin/invite-link/tokens/{token} | Update a public signup token
[**ValidatePublicSignupToken**](PublicSignupTokensApi.md#ValidatePublicSignupToken) | **Get** /invite/{token}/validate | Check whether a public sign-up token exists, has not expired and is enabled



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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    token := "token_example" // string | 
    createInvitedUserSchema := *openapiclient.NewCreateInvitedUserSchema("Email_example", "Name_example", "Password_example") // CreateInvitedUserSchema | createInvitedUserSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensApi.AddPublicSignupTokenUser(context.Background(), token).CreateInvitedUserSchema(createInvitedUserSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.AddPublicSignupTokenUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPublicSignupTokenUser`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensApi.AddPublicSignupTokenUser`: %v\n", resp)
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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    publicSignupTokenCreateSchema := *openapiclient.NewPublicSignupTokenCreateSchema("Name_example", time.Now()) // PublicSignupTokenCreateSchema | publicSignupTokenCreateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensApi.CreatePublicSignupToken(context.Background()).PublicSignupTokenCreateSchema(publicSignupTokenCreateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.CreatePublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensApi.CreatePublicSignupToken`: %v\n", resp)
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

Retrieve all existing public signup tokens

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
    resp, r, err := apiClient.PublicSignupTokensApi.GetAllPublicSignupTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.GetAllPublicSignupTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPublicSignupTokens`: PublicSignupTokensSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensApi.GetAllPublicSignupTokens`: %v\n", resp)
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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensApi.GetPublicSignupToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.GetPublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensApi.GetPublicSignupToken`: %v\n", resp)
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
    openapiclient "github.com/Unleash/unleash-server-api-go/client"
)

func main() {
    token := "token_example" // string | 
    publicSignupTokenUpdateSchema := *openapiclient.NewPublicSignupTokenUpdateSchema() // PublicSignupTokenUpdateSchema | publicSignupTokenUpdateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicSignupTokensApi.UpdatePublicSignupToken(context.Background(), token).PublicSignupTokenUpdateSchema(publicSignupTokenUpdateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.UpdatePublicSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublicSignupToken`: PublicSignupTokenSchema
    fmt.Fprintf(os.Stdout, "Response from `PublicSignupTokensApi.UpdatePublicSignupToken`: %v\n", resp)
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

Check whether a public sign-up token exists, has not expired and is enabled

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
    token := "token_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PublicSignupTokensApi.ValidatePublicSignupToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicSignupTokensApi.ValidatePublicSignupToken``: %v\n", err)
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

