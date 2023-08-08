# \AddonsApi

All URIs are relative to *https://app.unleash-hosted.com/hosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddon**](AddonsApi.md#CreateAddon) | **Post** /api/admin/addons | Create a new addon
[**DeleteAddon**](AddonsApi.md#DeleteAddon) | **Delete** /api/admin/addons/{id} | Delete an addon
[**GetAddon**](AddonsApi.md#GetAddon) | **Get** /api/admin/addons/{id} | Get a specific addon
[**GetAddons**](AddonsApi.md#GetAddons) | **Get** /api/admin/addons | Get all addons and providers
[**UpdateAddon**](AddonsApi.md#UpdateAddon) | **Put** /api/admin/addons/{id} | Update an addon



## CreateAddon

> AddonSchema CreateAddon(ctx).AddonCreateUpdateSchema(addonCreateUpdateSchema).Execute()

Create a new addon



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
    addonCreateUpdateSchema := *openapiclient.NewAddonCreateUpdateSchema("webhook", false, map[string]interface{}{"key": interface{}(123)}, []string{"Events_example"}) // AddonCreateUpdateSchema | addonCreateUpdateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddonsApi.CreateAddon(context.Background()).AddonCreateUpdateSchema(addonCreateUpdateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsApi.CreateAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAddon`: AddonSchema
    fmt.Fprintf(os.Stdout, "Response from `AddonsApi.CreateAddon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addonCreateUpdateSchema** | [**AddonCreateUpdateSchema**](AddonCreateUpdateSchema.md) | addonCreateUpdateSchema | 

### Return type

[**AddonSchema**](AddonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddon

> DeleteAddon(ctx, id).Execute()

Delete an addon



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
    r, err := apiClient.AddonsApi.DeleteAddon(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsApi.DeleteAddon``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAddonRequest struct via the builder pattern


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


## GetAddon

> AddonSchema GetAddon(ctx, id).Execute()

Get a specific addon



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
    resp, r, err := apiClient.AddonsApi.GetAddon(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsApi.GetAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddon`: AddonSchema
    fmt.Fprintf(os.Stdout, "Response from `AddonsApi.GetAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddonSchema**](AddonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddons

> AddonsSchema GetAddons(ctx).Execute()

Get all addons and providers



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
    resp, r, err := apiClient.AddonsApi.GetAddons(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsApi.GetAddons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddons`: AddonsSchema
    fmt.Fprintf(os.Stdout, "Response from `AddonsApi.GetAddons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddonsRequest struct via the builder pattern


### Return type

[**AddonsSchema**](AddonsSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddon

> AddonSchema UpdateAddon(ctx, id).AddonCreateUpdateSchema(addonCreateUpdateSchema).Execute()

Update an addon



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
    addonCreateUpdateSchema := *openapiclient.NewAddonCreateUpdateSchema("webhook", false, map[string]interface{}{"key": interface{}(123)}, []string{"Events_example"}) // AddonCreateUpdateSchema | addonCreateUpdateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddonsApi.UpdateAddon(context.Background(), id).AddonCreateUpdateSchema(addonCreateUpdateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddonsApi.UpdateAddon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddon`: AddonSchema
    fmt.Fprintf(os.Stdout, "Response from `AddonsApi.UpdateAddon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAddonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addonCreateUpdateSchema** | [**AddonCreateUpdateSchema**](AddonCreateUpdateSchema.md) | addonCreateUpdateSchema | 

### Return type

[**AddonSchema**](AddonSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

