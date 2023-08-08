# \EnvironmentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneEnvironment**](EnvironmentsApi.md#CloneEnvironment) | **Post** /api/admin/environments/{name}/clone | Clones an environment
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /api/admin/environments | Creates a new environment
[**GetAllEnvironments**](EnvironmentsApi.md#GetAllEnvironments) | **Get** /api/admin/environments | Get all environments
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /api/admin/environments/{name} | Get the environment with &#x60;name&#x60;
[**GetProjectEnvironments**](EnvironmentsApi.md#GetProjectEnvironments) | **Get** /api/admin/environments/project/{projectId} | Get the environments available to a project
[**RemoveEnvironment**](EnvironmentsApi.md#RemoveEnvironment) | **Delete** /api/admin/environments/{name} | Deletes an environment by name
[**ToggleEnvironmentOff**](EnvironmentsApi.md#ToggleEnvironmentOff) | **Post** /api/admin/environments/{name}/off | Toggle the environment with &#x60;name&#x60; off
[**ToggleEnvironmentOn**](EnvironmentsApi.md#ToggleEnvironmentOn) | **Post** /api/admin/environments/{name}/on | Toggle the environment with &#x60;name&#x60; on
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Put** /api/admin/environments/update/{name} | Updates an environment by name
[**UpdateSortOrder**](EnvironmentsApi.md#UpdateSortOrder) | **Put** /api/admin/environments/sort-order | Update environment sort orders
[**ValidateEnvironmentName**](EnvironmentsApi.md#ValidateEnvironmentName) | **Post** /api/admin/environments/validate | Validates if an environment name exists



## CloneEnvironment

> CloneEnvironment(ctx, name).Execute()

Clones an environment



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
    r, err := apiClient.EnvironmentsApi.CloneEnvironment(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CloneEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneEnvironmentRequest struct via the builder pattern


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


## CreateEnvironment

> EnvironmentSchema CreateEnvironment(ctx).CreateEnvironmentSchema(createEnvironmentSchema).Execute()

Creates a new environment



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
    createEnvironmentSchema := *openapiclient.NewCreateEnvironmentSchema("Name_example", "Type_example") // CreateEnvironmentSchema | createEnvironmentSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironment(context.Background()).CreateEnvironmentSchema(createEnvironmentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironment`: EnvironmentSchema
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEnvironmentSchema** | [**CreateEnvironmentSchema**](CreateEnvironmentSchema.md) | createEnvironmentSchema | 

### Return type

[**EnvironmentSchema**](EnvironmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllEnvironments

> EnvironmentsSchema GetAllEnvironments(ctx).Execute()

Get all environments



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
    resp, r, err := apiClient.EnvironmentsApi.GetAllEnvironments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetAllEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllEnvironments`: EnvironmentsSchema
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetAllEnvironments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllEnvironmentsRequest struct via the builder pattern


### Return type

[**EnvironmentsSchema**](EnvironmentsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironment

> EnvironmentSchema GetEnvironment(ctx, name).Execute()

Get the environment with `name`



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
    resp, r, err := apiClient.EnvironmentsApi.GetEnvironment(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironment`: EnvironmentSchema
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentSchema**](EnvironmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectEnvironments

> EnvironmentsProjectSchema GetProjectEnvironments(ctx, projectId).Execute()

Get the environments available to a project



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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.GetProjectEnvironments(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetProjectEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjectEnvironments`: EnvironmentsProjectSchema
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetProjectEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentsProjectSchema**](EnvironmentsProjectSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEnvironment

> RemoveEnvironment(ctx, name).Execute()

Deletes an environment by name



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
    r, err := apiClient.EnvironmentsApi.RemoveEnvironment(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.RemoveEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEnvironmentRequest struct via the builder pattern


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


## ToggleEnvironmentOff

> ToggleEnvironmentOff(ctx, name).Execute()

Toggle the environment with `name` off



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
    r, err := apiClient.EnvironmentsApi.ToggleEnvironmentOff(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ToggleEnvironmentOff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleEnvironmentOffRequest struct via the builder pattern


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


## ToggleEnvironmentOn

> ToggleEnvironmentOn(ctx, name).Execute()

Toggle the environment with `name` on



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
    r, err := apiClient.EnvironmentsApi.ToggleEnvironmentOn(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ToggleEnvironmentOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleEnvironmentOnRequest struct via the builder pattern


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


## UpdateEnvironment

> EnvironmentSchema UpdateEnvironment(ctx, name).UpdateEnvironmentSchema(updateEnvironmentSchema).Execute()

Updates an environment by name



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
    updateEnvironmentSchema := *openapiclient.NewUpdateEnvironmentSchema() // UpdateEnvironmentSchema | updateEnvironmentSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateEnvironment(context.Background(), name).UpdateEnvironmentSchema(updateEnvironmentSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: EnvironmentSchema
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvironmentSchema** | [**UpdateEnvironmentSchema**](UpdateEnvironmentSchema.md) | updateEnvironmentSchema | 

### Return type

[**EnvironmentSchema**](EnvironmentSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSortOrder

> UpdateSortOrder(ctx).RequestBody(requestBody).Execute()

Update environment sort orders



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
    requestBody := map[string]float32{"key": float32(123)} // map[string]float32 | sortOrderSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentsApi.UpdateSortOrder(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateSortOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSortOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]float32** | sortOrderSchema | 

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


## ValidateEnvironmentName

> ValidateEnvironmentName(ctx).NameSchema(nameSchema).Execute()

Validates if an environment name exists



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
    nameSchema := *openapiclient.NewNameSchema("betaUser") // NameSchema | nameSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentsApi.ValidateEnvironmentName(context.Background()).NameSchema(nameSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ValidateEnvironmentName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEnvironmentNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameSchema** | [**NameSchema**](NameSchema.md) | nameSchema | 

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

