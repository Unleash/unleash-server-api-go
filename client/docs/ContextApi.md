# \ContextApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContextField**](ContextApi.md#CreateContextField) | **Post** /api/admin/context | 
[**DeleteContextField**](ContextApi.md#DeleteContextField) | **Delete** /api/admin/context/{contextField} | 
[**GetContextField**](ContextApi.md#GetContextField) | **Get** /api/admin/context/{contextField} | 
[**GetContextFields**](ContextApi.md#GetContextFields) | **Get** /api/admin/context | 
[**UpdateContextField**](ContextApi.md#UpdateContextField) | **Put** /api/admin/context/{contextField} | 
[**Validate**](ContextApi.md#Validate) | **Post** /api/admin/context/validate | 



## CreateContextField

> ContextFieldSchema CreateContextField(ctx).UpsertContextFieldSchema(upsertContextFieldSchema).Execute()



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
    upsertContextFieldSchema := *openapiclient.NewUpsertContextFieldSchema("Name_example") // UpsertContextFieldSchema | upsertContextFieldSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.CreateContextField(context.Background()).UpsertContextFieldSchema(upsertContextFieldSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.CreateContextField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContextField`: ContextFieldSchema
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.CreateContextField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upsertContextFieldSchema** | [**UpsertContextFieldSchema**](UpsertContextFieldSchema.md) | upsertContextFieldSchema | 

### Return type

[**ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContextField

> DeleteContextField(ctx, contextField).Execute()



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
    contextField := "contextField_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContextApi.DeleteContextField(context.Background(), contextField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.DeleteContextField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextField

> ContextFieldSchema GetContextField(ctx, contextField).Execute()



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
    contextField := "contextField_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContextApi.GetContextField(context.Background(), contextField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.GetContextField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextField`: ContextFieldSchema
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.GetContextField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextFields

> []ContextFieldSchema GetContextFields(ctx).Execute()



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
    resp, r, err := apiClient.ContextApi.GetContextFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.GetContextFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContextFields`: []ContextFieldSchema
    fmt.Fprintf(os.Stdout, "Response from `ContextApi.GetContextFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextFieldsRequest struct via the builder pattern


### Return type

[**[]ContextFieldSchema**](ContextFieldSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContextField

> UpdateContextField(ctx, contextField).UpsertContextFieldSchema(upsertContextFieldSchema).Execute()



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
    contextField := "contextField_example" // string | 
    upsertContextFieldSchema := *openapiclient.NewUpsertContextFieldSchema("Name_example") // UpsertContextFieldSchema | upsertContextFieldSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContextApi.UpdateContextField(context.Background(), contextField).UpsertContextFieldSchema(upsertContextFieldSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.UpdateContextField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextField** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContextFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertContextFieldSchema** | [**UpsertContextFieldSchema**](UpsertContextFieldSchema.md) | upsertContextFieldSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> Validate(ctx).NameSchema(nameSchema).Execute()



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
    nameSchema := *openapiclient.NewNameSchema("Name_example") // NameSchema | nameSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContextApi.Validate(context.Background()).NameSchema(nameSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContextApi.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameSchema** | [**NameSchema**](NameSchema.md) | nameSchema | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
