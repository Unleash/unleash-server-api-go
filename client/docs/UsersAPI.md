# \UsersAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /api/admin/user-admin | Create a new user
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /api/admin/user-admin/{id} | Delete a user
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/admin/user-admin/{id} | Get user
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /api/admin/user-admin/{id} | Update a user



## CreateUser

> CreateUserResponseSchema CreateUser(ctx).CreateUserSchema(createUserSchema).Execute()

Create a new user



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
    createUserSchema := *openapiclient.NewCreateUserSchema(openapiclient.createUserSchema_rootRole{Int32: new(int32)}) // CreateUserSchema | createUserSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.CreateUser(context.Background()).CreateUserSchema(createUserSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: CreateUserResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserSchema** | [**CreateUserSchema**](CreateUserSchema.md) | createUserSchema | 

### Return type

[**CreateUserResponseSchema**](CreateUserResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, id).Execute()

Delete a user



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
    r, err := apiClient.UsersAPI.DeleteUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## GetUser

> UserSchema GetUser(ctx, id).Execute()

Get user



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
    resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserSchema
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSchema**](UserSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> CreateUserResponseSchema UpdateUser(ctx, id).UpdateUserSchema(updateUserSchema).Execute()

Update a user



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
    updateUserSchema := *openapiclient.NewUpdateUserSchema() // UpdateUserSchema | updateUserSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), id).UpdateUserSchema(updateUserSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: CreateUserResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserSchema** | [**UpdateUserSchema**](UpdateUserSchema.md) | updateUserSchema | 

### Return type

[**CreateUserResponseSchema**](CreateUserResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

