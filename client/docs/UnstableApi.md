# \UnstableApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportFeatures**](UnstableApi.md#ExportFeatures) | **Post** /api/admin/features-batch/export | 
[**GetAdvancedPlayground**](UnstableApi.md#GetAdvancedPlayground) | **Post** /api/admin/playground/advanced | Batch evaluate an Unleash context against a set of environments and projects.
[**ImportToggles**](UnstableApi.md#ImportToggles) | **Post** /api/admin/features-batch/import | Import feature toggles for an environment in the project
[**ValidateImport**](UnstableApi.md#ValidateImport) | **Post** /api/admin/features-batch/validate | Validate import of feature toggles for an environment in the project



## ExportFeatures

> ExportResultSchema ExportFeatures(ctx).RequestBody(requestBody).Execute()



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | exportQuerySchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.ExportFeatures(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.ExportFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFeatures`: ExportResultSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.ExportFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** | exportQuerySchema | 

### Return type

[**ExportResultSchema**](ExportResultSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdvancedPlayground

> AdvancedPlaygroundResponseSchema GetAdvancedPlayground(ctx).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()

Batch evaluate an Unleash context against a set of environments and projects.



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
    advancedPlaygroundRequestSchema := *openapiclient.NewAdvancedPlaygroundRequestSchema([]string{"Environments_example"}, *openapiclient.NewSdkContextSchema("My cool application.")) // AdvancedPlaygroundRequestSchema | advancedPlaygroundRequestSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.GetAdvancedPlayground(context.Background()).AdvancedPlaygroundRequestSchema(advancedPlaygroundRequestSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.GetAdvancedPlayground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdvancedPlayground`: AdvancedPlaygroundResponseSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.GetAdvancedPlayground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedPlaygroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **advancedPlaygroundRequestSchema** | [**AdvancedPlaygroundRequestSchema**](AdvancedPlaygroundRequestSchema.md) | advancedPlaygroundRequestSchema | 

### Return type

[**AdvancedPlaygroundResponseSchema**](AdvancedPlaygroundResponseSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportToggles

> ImportToggles(ctx).ImportTogglesSchema(importTogglesSchema).Execute()

Import feature toggles for an environment in the project



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
    importTogglesSchema := *openapiclient.NewImportTogglesSchema("Project_example", "Environment_example", *openapiclient.NewExportResultSchema([]openapiclient.FeatureSchema{*openapiclient.NewFeatureSchema("disable-comments")}, []openapiclient.FeatureStrategySchema{*openapiclient.NewFeatureStrategySchema("flexibleRollout")}, []openapiclient.TagTypeSchema{*openapiclient.NewTagTypeSchema("Name_example")})) // ImportTogglesSchema | importTogglesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UnstableApi.ImportToggles(context.Background()).ImportTogglesSchema(importTogglesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.ImportToggles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportTogglesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importTogglesSchema** | [**ImportTogglesSchema**](ImportTogglesSchema.md) | importTogglesSchema | 

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


## ValidateImport

> ImportTogglesValidateSchema ValidateImport(ctx).ImportTogglesSchema(importTogglesSchema).Execute()

Validate import of feature toggles for an environment in the project



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
    importTogglesSchema := *openapiclient.NewImportTogglesSchema("Project_example", "Environment_example", *openapiclient.NewExportResultSchema([]openapiclient.FeatureSchema{*openapiclient.NewFeatureSchema("disable-comments")}, []openapiclient.FeatureStrategySchema{*openapiclient.NewFeatureStrategySchema("flexibleRollout")}, []openapiclient.TagTypeSchema{*openapiclient.NewTagTypeSchema("Name_example")})) // ImportTogglesSchema | importTogglesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnstableApi.ValidateImport(context.Background()).ImportTogglesSchema(importTogglesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnstableApi.ValidateImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateImport`: ImportTogglesValidateSchema
    fmt.Fprintf(os.Stdout, "Response from `UnstableApi.ValidateImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importTogglesSchema** | [**ImportTogglesSchema**](ImportTogglesSchema.md) | importTogglesSchema | 

### Return type

[**ImportTogglesValidateSchema**](ImportTogglesValidateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

