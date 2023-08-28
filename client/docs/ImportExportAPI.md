# \ImportExportAPI

All URIs are relative to *https://us.app.unleash-hosted.com/ushosted*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallImport**](ImportExportAPI.md#CallImport) | **Post** /api/admin/state/import | Import state (deprecated)
[**Export**](ImportExportAPI.md#Export) | **Get** /api/admin/state/export | Export state (deprecated)
[**ImportToggles**](ImportExportAPI.md#ImportToggles) | **Post** /api/admin/features-batch/import | Import feature toggles
[**ValidateImport**](ImportExportAPI.md#ValidateImport) | **Post** /api/admin/features-batch/validate | Validate feature import data



## CallImport

> CallImport(ctx).StateSchema(stateSchema).Execute()

Import state (deprecated)



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
    stateSchema := *openapiclient.NewStateSchema(int32(1)) // StateSchema | stateSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportExportAPI.CallImport(context.Background()).StateSchema(stateSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportAPI.CallImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCallImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateSchema** | [**StateSchema**](StateSchema.md) | stateSchema | 

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


## Export

> StateSchema Export(ctx).Format(format).Download(download).Strategies(strategies).FeatureToggles(featureToggles).Projects(projects).Tags(tags).Environments(environments).Execute()

Export state (deprecated)



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
    format := "format_example" // string | Desired export format. Must be either `json` or `yaml`. (optional) (default to "json")
    download := *openapiclient.NewExportDownloadParameter() // ExportDownloadParameter | Whether exported data should be downloaded as a file. (optional) (default to false)
    strategies := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether strategies should be included in the exported data. (optional) (default to true)
    featureToggles := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether feature toggles should be included in the exported data. (optional) (default to true)
    projects := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether projects should be included in the exported data. (optional) (default to true)
    tags := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether tag types, tags, and feature_tags should be included in the exported data. (optional) (default to true)
    environments := *openapiclient.NewExportStrategiesParameter() // ExportStrategiesParameter | Whether environments should be included in the exported data. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportExportAPI.Export(context.Background()).Format(format).Download(download).Strategies(strategies).FeatureToggles(featureToggles).Projects(projects).Tags(tags).Environments(environments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportAPI.Export``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Export`: StateSchema
    fmt.Fprintf(os.Stdout, "Response from `ImportExportAPI.Export`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | Desired export format. Must be either &#x60;json&#x60; or &#x60;yaml&#x60;. | [default to &quot;json&quot;]
 **download** | [**ExportDownloadParameter**](ExportDownloadParameter.md) | Whether exported data should be downloaded as a file. | [default to false]
 **strategies** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether strategies should be included in the exported data. | [default to true]
 **featureToggles** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether feature toggles should be included in the exported data. | [default to true]
 **projects** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether projects should be included in the exported data. | [default to true]
 **tags** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether tag types, tags, and feature_tags should be included in the exported data. | [default to true]
 **environments** | [**ExportStrategiesParameter**](ExportStrategiesParameter.md) | Whether environments should be included in the exported data. | [default to true]

### Return type

[**StateSchema**](StateSchema.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportToggles

> ImportToggles(ctx).ImportTogglesSchema(importTogglesSchema).Execute()

Import feature toggles



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
    importTogglesSchema := *openapiclient.NewImportTogglesSchema("My awesome project", "development", *openapiclient.NewExportResultSchema([]openapiclient.FeatureSchema{*openapiclient.NewFeatureSchema("disable-comments")}, []openapiclient.FeatureStrategySchema{*openapiclient.NewFeatureStrategySchema("flexibleRollout")}, []openapiclient.TagTypeSchema{*openapiclient.NewTagTypeSchema("color")})) // ImportTogglesSchema | importTogglesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImportExportAPI.ImportToggles(context.Background()).ImportTogglesSchema(importTogglesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportAPI.ImportToggles``: %v\n", err)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateImport

> ImportTogglesValidateSchema ValidateImport(ctx).ImportTogglesSchema(importTogglesSchema).Execute()

Validate feature import data



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
    importTogglesSchema := *openapiclient.NewImportTogglesSchema("My awesome project", "development", *openapiclient.NewExportResultSchema([]openapiclient.FeatureSchema{*openapiclient.NewFeatureSchema("disable-comments")}, []openapiclient.FeatureStrategySchema{*openapiclient.NewFeatureStrategySchema("flexibleRollout")}, []openapiclient.TagTypeSchema{*openapiclient.NewTagTypeSchema("color")})) // ImportTogglesSchema | importTogglesSchema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportExportAPI.ValidateImport(context.Background()).ImportTogglesSchema(importTogglesSchema).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportExportAPI.ValidateImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateImport`: ImportTogglesValidateSchema
    fmt.Fprintf(os.Stdout, "Response from `ImportExportAPI.ValidateImport`: %v\n", resp)
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

